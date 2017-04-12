package main

import "github.com/smithatlanta/goshortener/client/cmd"

var version string
var buildDate string

func main() {
	cmd.Execute(version, buildDate)
}
