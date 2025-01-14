package provider

import (
	"errors"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func fromBigAnimalErr(err error) diag.Diagnostics {
	if err == nil {
		return nil
	}

	var baError *api.BigAnimalError
	if errors.As(err, &baError) {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  baError.Error(),
				Detail:   baError.GetDetails(),
			},
		}
	}
	return diag.FromErr(err)
}

func unsupportedWarning(message string) diag.Diagnostics {
	return diag.Diagnostics{
		diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Unsupported",
			Detail:   message,
		},
	}
}
