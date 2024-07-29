package config

import (
	"net/http"

	"github.com/terraform-providers/terraform-provider-fmgdevice/sdk/auth"
)

// Config provides configuration to a FMG client instance
// It saves authentication information and a http connection
// for FMG Client instance to create New connction to FMG
// and Send data to FMG,  etc. (needs to be extended later.)
type Config struct {
	Auth     *auth.Auth
	HTTPCon  *http.Client
	FwTarget string
}
