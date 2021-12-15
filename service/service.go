package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//db, err := sql.Open("mysql", "MainUser:MainPassword@tcp(192.168.8.132:3306)/Worldcities")
	// Acess DB in cluster
	db, err := sql.Open("mysql", "root:RootPassword@tcp(localhost:3306)/Worldcities")

	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	type Worldcities struct {
		Id         int
		City string
		Country  string
	}
	router := gin.Default()

	// GET a customer detail based on the id
	router.GET("/Worldcities/:id", func(c *gin.Context) {
		var (
			worldcities Worldcities
			result   gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("select id, city, country from mytable where id = ?;", id)
		err = row.Scan(&worldcities.Id, &worldcities.City, &worldcities.Country)
		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": worldcities,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	// GET all Worldcities
	router.GET("/Worldcities", func(c *gin.Context) {
		var (
			worldcities  Worldcities
			worldcities_all []Worldcities
		)
		rows, err := db.Query("select id, city, country from mytable;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&worldcities.Id, &worldcities.City, &worldcities.Country)
			worldcities_all = append(worldcities_all, worldcities)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": worldcities_all,
			"count":  len(worldcities_all),
		})
	})

	// POST new customer details
	router.POST("/Worldcities", func(c *gin.Context) {
		var buffer bytes.Buffer
		city := c.PostForm("city")
		country := c.PostForm("county")
		stmt, err := db.Prepare("insert into mytable (city, country) values(?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(city, country)

		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(city)
		buffer.WriteString(" ")
		buffer.WriteString(country)
		defer stmt.Close()
		city_add := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %s successfully created", city_add),
		})
	})

	// PUT - update a customer details
	router.PUT("/Worldcities", func(c *gin.Context) {
		var buffer bytes.Buffer
		id := c.Query("id")
		city := c.PostForm("city")
		country := c.PostForm("country")
		stmt, err := db.Prepare("update mytable set city= ?, country= ? where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(city, country, id)
		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(city)
		buffer.WriteString(" ")
		buffer.WriteString(country)
		defer stmt.Close()
		city_update := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Worldcities Successfully updated to %s", city_update),
		})
	})

	// Delete customer
	router.DELETE("/Wolrdcities", func(c *gin.Context) {
		id := c.Query("id")
		stmt, err := db.Prepare("delete from mytable where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(id)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted Worldcities: %s", id),
		})
	})
	router.Run(":3000")
}