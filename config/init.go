package config

import "github.com/pubgo/vapper/vapper"

type Config struct {
	UrlPrefix string
}

func init() {
	app := vapper.Default()
	app.Config(&Config{
		UrlPrefix: "hello",
	})
}
