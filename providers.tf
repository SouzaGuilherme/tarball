terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "~>3.88.0"
    }
  }
}

provider "google" {
  credentials = file(var.gcp_credentials)
  project = var.gcp_project_id
  region = var.gcp_region
  zone = var.gcp_zone
}
provider "google-beta" {
  credentials = file(var.gcp_credentials)
  project = var.gcp_project_id
  region = var.gcp_region
  zone = var.gcp_zone
}

# google_client_config and kubernetes provider must be explicitly specified like the following.
data "google_client_config" "default" {
}

provider "kubernetes" {
  host  = "https://${module.kubernetes-engine.endpoint}"
  token = data.google_client_config.default.access_token
  cluster_ca_certificate = base64decode(module.kubernetes-engine.ca_certificate)
}

provider "helm" {
    kubernetes {
      host  = "https://${module.kubernetes-engine.endpoint}"
      token = data.google_client_config.default.access_token
      cluster_ca_certificate = base64decode(module.kubernetes-engine.ca_certificate)
  }
}