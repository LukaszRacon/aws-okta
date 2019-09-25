package main

import (
	"github.com/LukaszRacon/aws-okta/cmd"
)

// These are set via linker flags
var (
	Version           = "dev-keyring-key"
	AnalyticsWriteKey = ""
)

func main() {
	// vars set by linker flags must be strings...
	cmd.Execute(Version, AnalyticsWriteKey)
}
