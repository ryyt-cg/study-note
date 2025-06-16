# Go Setup

## Download and Install

1. Download Go SDK here at this URL https://go.dev/doc/install
2. Install Go quickly and follow the steps

I would recommend to install Go Version Manager to manage your Go SDK. It's easy and very useful to manage multiple
versions.

---

## Go Version Manager

* Simple go version manager<p/>
  Go to this [link](https://github.com/stefanmaric/g "Go Version Manager link") and follow the instructions for
  installation.

## Single-line Installation

Install by curl command

```shell
curl -sSL https://git.io/g-install | sh -s
```

Validate installation

```
g --help

  Usage: g [COMMAND] [options] [args]

  Commands:

    g                         Open interactive UI with downloaded versions
    g install latest          Download and set the latest go release
    g install <version>       Download and set go <version>
    g download <version>      Download go <version>
    g set <version>           Switch to go <version>
    g run <version>           Run a given version of go
    g which <version>         Output bin path for <version>
    g remove <version ...>    Remove the given version(s)
    g prune                   Remove all versions except the current version
    g list                    Output downloaded go versions
    g list-all                Output all available, remote go versions
    g self-upgrade            Upgrades g to the latest version
    g help                    Display help information, same as g --help

  Options:

    -h, --help                Display help information and exit
    -v, --version             Output current version of g and exit
    -q, --quiet               Suppress almost all output
    -c, --no-color            Force disabled color output
    -y, --non-interactive     Prevent prompts
    -o, --os                  Override operating system
    -a, --arch                Override system architecture
    -u, --unstable            Include unstable versions in list
```

