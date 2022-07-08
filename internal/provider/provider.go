package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a *schema.Provider.
func Provider() *schema.Provider {
	// The actual provider
	provider := &schema.Provider{}

	return provider
}
