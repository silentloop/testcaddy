package main

import (
	"fmt"

	caddycmd "github.com/caddyserver/caddy/v2/cmd"

	_ "github.com/caddyserver/caddy/v2/modules/standard"
	_ "github.com/silentloop/testcaddy/handler"
)

func main() {

	fmt.Println("🟢 Plugin module loading: encrypt")
	caddycmd.Main()
}
