# October Linux management utility

Octoberctl is the management utility of October Linux. It allows the user to manage or updates their [October configuration](https://github.com/october-os/october-config).

## Quickstart

You can install the AUR package to quickly download a binary.

[![octoberctl-bin](https://img.shields.io/aur/version/octoberctl-bin?color=1793d1&label=octoberctl-bin&logo=arch-linux&style=for-the-badge)](https://aur.archlinux.org/packages/octoberctl-bin)

To compile it, you need the [Go compiler](https://go.dev/). After that, you can run:
```bash
git clone https://github.com/october-os/octoberctl.git
cd octoberctl
go build -o octoberctl cmd/main.go
```
