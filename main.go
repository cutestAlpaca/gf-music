package main

import (
	_ "gf-music/boot"
	_ "gf-music/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
