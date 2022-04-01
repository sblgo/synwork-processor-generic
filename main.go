package main

import (
	"sbl.system/synwork/generic/generic"
	"sbl.systems/go/synwork/plugin-sdk/plugin"
)

func main() {
	plugin.Serve(generic.Opts)
}
