package allow

import (
	"fmt"
	"net/url"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/juliengk/go-docker/image"
	"github.com/juliengk/go-log"
	"github.com/juliengk/go-log/driver"
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/allow/types"
	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/version"
)

var (
	// AllImages allow all images to be pulled
	AllImages bool
)

// ImageCreate called from plugin
func ImageCreate(req authorization.Request, config *types.Config) *types.AllowResult {
	u, err := url.ParseRequestURI(req.RequestURI)
	if err != nil {
		return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Could not parse URL query")}
	}

	params := u.Query()

	if !allowImage(params["fromImage"][0], config) {
		return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Image %s is not allowed to be pulled", params["fromImage"][0])}
	}

	return &types.AllowResult{Allow: true}
}

func allowImage(img string, config *types.Config) bool {
	defer utils.RecoverFunc()

	l, _ := log.NewDriver("standard", nil)

	s, err := storage.NewDriver("sqlite", config.AppPath)
	if err != nil {
		l.WithFields(driver.Fields{
			"storagedriver": "sqlite",
			"logdriver":     "standard",
			"version":       version.Version,
		}).Fatal(err)
	}
	defer s.End()

	i := image.NewImage(img)

	if i.Official {
		if s.ValidatePolicy(config.Username, "config", "image_create_official", "") {
			return true
		}
	}

	if s.ValidatePolicy(config.Username, "registry", i.Registry, "") {
		return true
	}

	if s.ValidatePolicy(config.Username, "image", i.String(), "") {
		return true
	}

	return false
}
