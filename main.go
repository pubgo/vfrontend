package main

import (
	"github.com/pubgo/vapper/vapper"
	_ "github.com/pubgo/vfrontend/config"
	_ "github.com/pubgo/vfrontend/routes"
)

func main() {
	vapper.Start()
}
