# Overview

The project described here provides the necessary files to upload a Kubernetes cluster on the Google Cloud Platform using Terraform.

GKE has a toy service installed via helm chart, it also has prometheus and grafana to perform cluster monitoring.

The chart was created locally and has a mariadb database to carry out the consumption. The Dockerfile for creating the image used to create the service and database used here is also provided.

## Toy Service

<img src="img_readme/gopher.svg" width="48">

First we have a simple API written in golang. Its function is nothing more to be a crud of cities in the world, providing the id, city and country. Allowing the main functions:
- [x] /GET
- [x] /POST
- [x] /PUT
- [x] /DELETE

The API can be viewed at:
[IMAGE AQUI]

Inside the folder contains the Dockerfile used to generate the image that is available in my dockerhub. To do this, just run the command:
~~~
$ docker built -t name_dockerhub/name_image:tag .
~~~
Note that the images will be downloaded directly from dockerhub as they were created outside the custer, in a development environment, that is, they do not yet exist in the cluster.

<img src="img_readme/mariadb-logo.png" width="80">

The database chosen for the project was Mariadb because it is opensource. Inside the mariadb folder there is also the dockerfile for creating the image and the worldcities.sql file that is used to populate our database.

## Helm Chart

Inside the deployment folder, the templates used for testing were created before creating the chart. So just create the chart in a simple way to be deployed via terraform in our GKE.

## Terraform
For the project, a GKE was raised using terraform, for the infra, two points were considered:
>region = southamerica-east1-a
>type-machine = n2-standard-2

This concern comes from saving my credits at GCP and being closer to my location :information_desk_person:
to run the project you need to enter your credentials, in the file:
>variable.auto.tfvars

change the variable:
~~~
gcp_credentials = "your_credentials.json"
~~~

Right. To make it easier, we check if the tools needed for execution are already installed in the current execution environment.

In this way, just run the file:

~~~
./prepare_environment.sh