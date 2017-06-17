package dcb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/docker/engine-api/types"
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

// VolumeList function
func VolumeList(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("volume")
	cmd.Add("ls")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		if _, ok := cmd.Params["filters"]; ok {
			var v map[string]map[string]bool

			err := json.Unmarshal([]byte(cmd.Params["filters"][0]), &v)
			if err != nil {
				panic(err)
			}

			var r []string

			for k, val := range v {
				r = append(r, k)

				for ka := range val {
					r = append(r, ka)
				}
			}

			cmd.Add(fmt.Sprintf("--filter \"%s=%s\"", r[0], r[1]))
		}
	}

	return cmd.String()
}

// VolumeCreate function
func VolumeCreate(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("volume")
	cmd.Add("create")

	vol := &types.VolumeCreateRequest{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(vol); err != nil {
			panic(err)
		}
	}

	if len(vol.Driver) > 0 {
		cmd.Add(fmt.Sprintf("--driver %s", vol.Driver))
	}

	if len(vol.DriverOpts) > 0 {
		for k, v := range vol.DriverOpts {
			cmd.Add(fmt.Sprintf("--opt %s=%s", k, v))
		}
	}

	if len(vol.Labels) > 0 {
		for k, v := range vol.Labels {
			cmd.Add(fmt.Sprintf("--label %s=%s", k, v))
		}
	}

	if len(vol.Name) > 0 {
		cmd.Add(fmt.Sprintf("--name %s", vol.Name))
	}

	return cmd.String()
}

// VolumeInspect function
func VolumeInspect(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("volume")
	cmd.Add("inspect")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

// VolumeRemove function
func VolumeRemove(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("volume")
	cmd.Add("rm")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}
