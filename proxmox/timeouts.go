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
	defaultTimeout:=time.Duration(300)
	if v, ok := os.LookupEnv("PM_TIMEOUT"); ok {
		defaultTimeout, _ = time.Duration((strconv.Atoi(v))
        }

	return &schema.ResourceTimeout{
		Create:  schema.DefaultTimeout(defaultTimeout),
		Read:    schema.DefaultTimeout(defaultTimeout),
		Update:  schema.DefaultTimeout(defaultTimeout),
		Delete:  schema.DefaultTimeout(defaultTimeout),
		Default: schema.DefaultTimeout(defaultTimeout),
	}
}
