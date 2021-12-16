variable "gcp_credentials" {
    type = string
    default = ""
    description = "Location of service account for GCP"
}

variable "gcp_project_id" {
    type = string
    default = "maxxer-test"
    description = "GCP project id"
}

variable "gcp_region" {
    type = string
    default = "southamerica-east1"
    description = "GKE region selected"
}

variable "gcp_zone" {
    type = string
    default = "southamerica-east1-a"
    description = "GCP zonoe select for project"
}

variable "gke_regional" {
    type = bool
    default = false
    description = "GKE regional selected"
}

variable "gcp_cluster_name" {
    type = string
    default = "cluster-k8s-maxxer"
    description = "GCP name for cluster k8s" 
}

variable "gke_machine_type" {
    type = string
    default = "n2-standard-2"
    description = "Machine type selected for GKE project"
}

variable "gke_zone" {
    type = list(string)
    default = ["southamerica-east1-c"]
    description = "GKE zone selected for project"
}

variable "initial_node_count" {
    default = 2
    description = "Number initital of node"
}

variable "max_node_count"{
    default = 5
    description = "Number max of node"
}

variable "gke_network" {
    type = string
    default = "default"
    description = "VPC network name"
}

variable "gke_subnetwork" {
    type = string
    default = "default"
    description = "VPC subnetwork name"
}

variable "gke_service_account" {
    type = string
    default = "maxxer-test@maxxer-test.iam.gserviceaccount.com"
    description = "GKE service name account for project"
}

variable "gke_node_name_pool" {
    type = string
    default = "node-name-default"
    description = "GKE node name pool"
}

variable "disk_size_gb" {
    default = 30
    description = "Disc size in gb for node pool"
}