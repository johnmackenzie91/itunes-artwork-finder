package tests

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ory/dockertest/v3"
	dc "github.com/ory/dockertest/v3/docker"
)

type container struct {
	url     string
	options *dockertest.RunOptions
}

func (c container) HealthCheck() error {
	resp, err := http.Get(c.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code not OK")
	}
	return nil
}

func (c container) URL(artist, title string) string {
	artist = strings.Replace(strings.ToLower(artist), " ", "+", -1)
	title = strings.Replace(strings.ToLower(title), " ", "+", -1)
	return fmt.Sprintf("http://0.0.0.0:5678/v1/artist/%s/album/%s?size=500", artist, title)
}

var sutContainer = container{
	url: "http://0.0.0.0:5678/v1/status",
	options: &dockertest.RunOptions{
		Repository:   "itunes-proxy",
		ExposedPorts: []string{"80"},
		PortBindings: map[dc.Port][]dc.PortBinding{
			"80": {
				{HostIP: "0.0.0.0", HostPort: "5678"},
			},
		},
		Tag: "latest",
		Env: []string{"LOG_LEVEL=debug"},
	},
}
