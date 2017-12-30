package main


import (
	"os"
)

import (
	"github.com/kflange/gocli/cmd"
)


func main() {
	os.Exit(int(cmd.Execute()))
}
