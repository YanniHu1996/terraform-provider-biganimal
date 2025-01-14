# biganimal_cluster (Resource)
The cluster resource is used to manage BigAnimal clusters. See [Creating a cluster](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/creating_a_cluster/) for more details.



## Single Node Cluster Example
```terraform
terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "0.3.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.3.1"
    }
  }
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

variable "cluster_name" {
  type        = string
  description = "The name of the cluster."
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

resource "biganimal_cluster" "single_node_cluster" {
  cluster_name = var.cluster_name
  project_id   = var.project_id

  allowed_ip_ranges {
    cidr_block  = "127.0.0.1/32"
    description = "localhost"
  }

  allowed_ip_ranges {
    cidr_block  = "192.168.0.1/32"
    description = "description!"
  }

  backup_retention_period = "6d"
  cluster_architecture {
    id    = "single"
    nodes = 1
  }
  csp_auth = false

  instance_type = "azure:Standard_D2s_v3"
  password      = resource.random_password.password.result
  pg_config {
    name  = "application_name"
    value = "created through terraform"
  }

  pg_config {
    name  = "array_nulls"
    value = "off"
  }

  storage {
    volume_type       = "azurepremiumstorage"
    volume_properties = "P1"
    size              = "4 Gi"
  }

  pg_type               = "epas"
  pg_version            = "14"
  private_networking    = false
  cloud_provider        = "azure"
  read_only_connections = false
  region                = "eastus2"
}

output "password" {
  sensitive = true
  value     = resource.biganimal_cluster.single_node_cluster.password
}

output "faraway_replica_ids" {
  value = biganimal_cluster.single_node_cluster.faraway_replica_ids
}
```

## High Availability Cluster Example
```terraform
terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "0.3.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.3.1"
    }
  }
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

variable "cluster_name" {
  type        = string
  description = "The name of the cluster."
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

resource "biganimal_cluster" "ha_cluster" {
  cluster_name = var.cluster_name
  project_id   = var.project_id

  allowed_ip_ranges {
    cidr_block  = "127.0.0.1/32"
    description = "localhost"
  }

  allowed_ip_ranges {
    cidr_block  = "192.168.0.1/32"
    description = "description!"
  }

  backup_retention_period = "6d"
  cluster_architecture {
    id    = "ha"
    nodes = 3
  }

  instance_type = "aws:c5.large"
  password      = resource.random_password.password.result
  pg_config {
    name  = "application_name"
    value = "created through terraform"
  }

  pg_config {
    name  = "array_nulls"
    value = "off"
  }

  storage {
    volume_type       = "gp3"
    volume_properties = "gp3"
    size              = "4 Gi"
  }

  pg_type               = "epas"
  pg_version            = "14"
  private_networking    = false
  cloud_provider        = "aws"
  read_only_connections = true
  region                = "us-east-1"
}

output "password" {
  sensitive = true
  value     = resource.biganimal_cluster.ha_cluster.password
}

output "ro_connection_uri" {
  value = resource.biganimal_cluster.ha_cluster.ro_connection_uri
}

output "faraway_replica_ids" {
  value = resource.biganimal_cluster.ha_cluster.faraway_replica_ids
}
```

## Extreme High Availability Cluster Example
```terraform
terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "0.3.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.3.1"
    }
  }
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

variable "cluster_name" {
  type        = string
  description = "The name of the cluster."
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

resource "biganimal_cluster" "eha_cluster" {
  cluster_name = var.cluster_name
  project_id   = var.project_id

  allowed_ip_ranges {
    cidr_block  = "127.0.0.1/32"
    description = "localhost"
  }

  allowed_ip_ranges {
    cidr_block  = "192.168.0.1/32"
    description = "description!"
  }

  backup_retention_period = "6d"
  cluster_architecture {
    id    = "eha"
    nodes = 5
  }

  instance_type = "aws:c5.large"
  password      = resource.random_password.password.result
  pg_config {
    name  = "application_name"
    value = "created through terraform"
  }

  pg_config {
    name  = "array_nulls"
    value = "off"
  }

  storage {
    volume_type       = "gp3"
    volume_properties = "gp3"
    size              = "4 Gi"
  }

  pg_type            = "epas"
  pg_version         = "14"
  private_networking = false
  cloud_provider     = "aws"
  region             = "us-east-1"
}

output "password" {
  sensitive = true
  value     = resource.biganimal_cluster.eha_cluster.password
}

output "connection_uri" {
  value = resource.biganimal_cluster.eha_cluster.connection_uri
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cloud_provider` (String) Cloud provider. For example, "aws" or "azure".
- `cluster_architecture` (Block List, Min: 1) Cluster architecture. See [Supported cluster types](https://www.enterprisedb.com/docs/biganimal/latest/overview/02_high_availability/) for details. (see [below for nested schema](#nestedblock--cluster_architecture))
- `cluster_name` (String) Name of the cluster.
- `instance_type` (String) Instance type. For example, "azure:Standard_D2s_v3" or "aws:c5.large".
- `password` (String, Sensitive) Password for the user edb_admin. It must be 12 characters or more.
- `pg_type` (String) Postgres type. For example, "epas", "pgextended", or "postgres".
- `pg_version` (String) Postgres version. See [Supported Postgres types and versions](https://www.enterprisedb.com/docs/biganimal/latest/overview/05_database_version_policy/#supported-postgres-types-and-versions) for supported Postgres types and versions.
- `project_id` (String) BigAnimal Project ID.
- `region` (String) Region to deploy the cluster. See [Supported regions](https://www.enterprisedb.com/docs/biganimal/latest/overview/03a_region_support/) for supported regions.
- `storage` (Block List, Min: 1) Storage. (see [below for nested schema](#nestedblock--storage))

### Optional

- `allowed_ip_ranges` (Block List) Allowed IP ranges. (see [below for nested schema](#nestedblock--allowed_ip_ranges))
- `backup_retention_period` (String) Backup retention period. For example, "7d", "2w", or "3m".
- `csp_auth` (Boolean) Is authentication handled by the cloud service provider. Available for AWS only, See [Authentication](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/creating_a_cluster/#authentication) for details.
- `pg_config` (Block List) Database configuration parameters. See [Modifying database configuration parameters](https://www.enterprisedb.com/docs/biganimal/latest/using_cluster/03_modifying_your_cluster/05_db_configuration_parameters/) for details. (see [below for nested schema](#nestedblock--pg_config))
- `private_networking` (Boolean) Is private networking enabled.
- `read_only_connections` (Boolean) Is read only connection enabled.
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `cluster_id` (String) Cluster ID.
- `connection_uri` (String) Cluster connection URI.
- `created_at` (String) Cluster creation time.
- `deleted_at` (String) Cluster deletion time.
- `expired_at` (String) Cluster expiry time.
- `faraway_replica_ids` (Set of String)
- `first_recoverability_point_at` (String) Earliest backup recover time.
- `id` (String) The ID of this resource.
- `logs_url` (String) The URL to find the logs of this cluster.
- `metrics_url` (String) The URL to find the metrics of this cluster.
- `phase` (String) Current phase of the cluster.
- `resizing_pvc` (List of String) Resizing PVC.
- `ro_connection_uri` (String) Cluster read-only connection URI. Only available for high availability clusters.

<a id="nestedblock--cluster_architecture"></a>
### Nested Schema for `cluster_architecture`

Required:

- `id` (String) Cluster architecture ID. For example, "single", "ha", or "eha".
- `nodes` (Number) Node count.

Read-Only:

- `name` (String) Name.


<a id="nestedblock--storage"></a>
### Nested Schema for `storage`

Required:

- `size` (String) Size of the volume. It can be set to different values depending on your volume type and properties.
- `volume_properties` (String) Volume properties in accordance with the selected volume type.
- `volume_type` (String) Volume type. For Azure: "azurepremiumstorage" or "ultradisk". For AWS: "gp3", "io2", or "io2-block-express".

Optional:

- `iops` (String) IOPS for the selected volume. It can be set to different values depending on your volume type and properties.

Read-Only:

- `throughput` (String) Throughput is automatically calculated by BigAnimal based on the IOPS input.


<a id="nestedblock--allowed_ip_ranges"></a>
### Nested Schema for `allowed_ip_ranges`

Required:

- `cidr_block` (String) CIDR block.

Optional:

- `description` (String) CIDR block description.


<a id="nestedblock--pg_config"></a>
### Nested Schema for `pg_config`

Required:

- `name` (String) GUC name.
- `value` (String) GUC value.


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
