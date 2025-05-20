package main

import (
	"github.com/caddyserver/caddy/v2"
	_ "github.com/silentloop/testcaddy/handler"
)

func main() {
	cfg := &caddy.Config{}
	caddy.Run(cfg)
}
