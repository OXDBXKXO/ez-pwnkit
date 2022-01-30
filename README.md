# go-PwnKit

A pure-Go implementation of the **CVE-2021-4034 PwnKit** exploit.



## Installation

```bash
git clone git@github.com:OXDBXKXO/go-PwnKit.git
cd go-PwnKit
make
```



As the exploit relies on a malicious shared library, a **PWN.so** file is generated from ***payload.go*** and embed in the resulting `exploit` executable.

The ***Makefile*** uses `sed` to temporarily change the package name of the ***payload.go*** file to `main`, hence making this ***Makefile*** Linux-only.



## Usage

```
Usage of ./exploit:
  -c string
        Command to execute as root (default "/bin/sh")
  -g    Append id and greetings to your command
  -r string
        Optionally open a reverse-shell instead. Format: host:port
```

The exploit can either be used with a command (flag `-c`) or as a reverse-shell (flag `-r`).



## Mitigation

Patch `pkexec` if possible, other disable the ***setuid*** bit on the `pkexec` binary.

```bash
chmod 0755 /usr/bin/pkexec
```



## Credits

This project is inspired by several other PoCs of the **PwnKit** exploit.



Thanks to [An00bRektn](https://github.com/An00bRektn/CVE-2021-4034) for the straight-forward exploit setup.

Thanks to [PaterGottesman](https://github.com/PeterGottesman/pwnkit-exploit) and [berdav](https://github.com/berdav/CVE-2021-4034) for the clarity of the exploit explanation.

Thanks to [dzonerzy](https://github.com/dzonerzy/poc-cve-2021-4034) for the *GIO_USE_VFS* trick.
