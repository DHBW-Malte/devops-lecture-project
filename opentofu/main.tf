terraform {
  required_providers {
    azurerm = {
      source = "hashicorp/azurerm"
      version = "~> 4.0"
    }
  }
  required_version = ">= 1.1.0"
}

provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "rg" {
  name = "rg-kubernetes"
  location = var.region
}

resource "azurerm_kubernetes_cluster" "aks" {
  name = "aks-cluster"
  location = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  dns_prefix = "aks-cluster"

  default_node_pool {
    name = "default"
    node_count = 1
    vm_size = "Standard_D4pds_v6"
  }

  identity {
    type = "SystemAssigned"
  }
}
