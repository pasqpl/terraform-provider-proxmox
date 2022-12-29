package proxmox

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTimeouts() *schema.ResourceTimeout {
	// resourceCreateTimeout := defaultTimeout
	// resourceReadTimeout := 600
	// resourceUpdateTimeout := 600
	// resourceDeleteTimeout := 1200
	defaultTimeout:=300
	if v, ok := os.LookupEnv("PM_TIMEOUT"); ok {
	 	defaultTimeout, _ = strconv.Atoi(v)
        }

	return &schema.ResourceTimeout{
		Create:  schema.DefaultTimeout(defaultTimeout * time.Minute),
		Read:    schema.DefaultTimeout(defaultTimeout * time.Minute),
		Update:  schema.DefaultTimeout(defaultTimeout * time.Minute),
		Delete:  schema.DefaultTimeout(defaultTimeout * time.Minute),
		Default: schema.DefaultTimeout(defaultTimeout * time.Minute),
	}
}
