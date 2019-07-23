package config

import "github.com/pubgo/vapper/jsvapper"

type Config struct {
	UrlPrefix string
}

func init() {
	app := jsvapper.Default()
	app.Config(&Config{
		UrlPrefix: "hello",
	})
}
