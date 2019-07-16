package config

import "github.com/pubgo/vapper/vapper"

type Config struct {
	UrlPrefix string
}

func init() {
	vapper.Config(&Config{
		UrlPrefix: "hello",
	})
}
