package main

import (
	_ "github.com/artemsalnikovtulun-dot/TEST/hiddify_extension"

	"github.com/hiddify/hiddify-core/extension/server"
)

func main() {
	server.StartTestExtensionServer()
}
