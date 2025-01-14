# biganimal_region (Data Source)
The region data source shows the available regions within a cloud provider.

## Example Usage
```terraform
variable "cloud_provider" {
  type        = string
  description = "Cloud Provider"

  validation {
    condition     = contains(["aws", "azure"], var.cloud_provider)
    error_message = "Please select one of the supported regions: aws, azure."
  }
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"

}

data "biganimal_region" "this" {
  cloud_provider = var.cloud_provider
  project_id     = var.project_id
  // region_id   = "us-west-1" //optional
  // query       = "eu" // optional
}

output "regions" {
  value = data.biganimal_region.this.regions
}

output "cloud_provider_id" {
  value = data.biganimal_region.this.cloud_provider
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cloud_provider` (String) Cloud provider to list the regions. For example, "aws" or "azure".
- `project_id` (String) BigAnimal Project ID.

### Optional

- `query` (String) Query to filter region list.
- `region_id` (String) Unique region ID. For example, "germanywestcentral" in the Azure cloud provider, "eu-west-1" in the AWS cloud provider.

### Read-Only

- `id` (String) The ID of this resource.
- `regions` (List of Object) Region information. (see [below for nested schema](#nestedatt--regions))

<a id="nestedatt--regions"></a>
### Nested Schema for `regions`

Read-Only:

- `continent` (String)
- `name` (String)
- `region_id` (String)
- `status` (String)
