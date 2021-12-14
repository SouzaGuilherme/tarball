variable "gcp_credentials" {
    type = string
    description = "Location of service account for GCP"
}

variable "gcp_project_id" {
    type = string
    description = "GCP project id"
}

variable "gcp_region" {
    type = string
    description = "GKE region selected"
}

variable "gcp_zone" {
    type = string
    description = "GCP zonoe select for project"
}

variable "gke_regional" {
    type = bool
    description = "GKE regional selected"
}

variable "gcp_cluster_name" {
    type = string
    description = "GCP name for cluster k8s" 
}

variable "gke_machine_type" {
    type = string
    description = "Machine type selected for GKE project"
}

variable "gke_zone" {
    type = list(string)
    description = "GKE zone selected for project"
}

variable "initial_node_count" {
    description = "Number initital of node"
}

variable "max_node_count"{
    description = "Number max of node"
}

variable "gke_network" {
    type = string
    description = "VPC network name"
}

variable "gke_subnetwork" {
    type = string
    description = "VPC subnetwork name"
}

variable "gke_service_account" {
    type = string
    description = "GKE service name account for project"
}

variable "gke_node_name_pool" {
    type = string
    description = "GKE node name pool"
}

variable "disk_size_gb" {
    description = "Disc size in gb for node pool"
}