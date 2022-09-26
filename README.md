# ggen

## Overview

'ggen' is a simple CLI tool via Golang.

You can select some options for generating passwords.

The options are like:

- specifing the number of characters
- containing lowercase / uppercase characters or not
- containing numeric characters or not
- containing symbol characters or not

## Install

Download the latest version of the binary that matches your local environment from the [releases](https://github.com/Nabby12/ggen/releases).

or

```shell
brew tap Nabby12/ggen
brew install Nabby12/ggen/ggen
```

## Usage

Run one simple command below:

```shell
ggen # default setting is '30 characters containing lowercase, uppercase, numeric and symbol'.
```

### options

- `-c`, `--count [int]`    # specify the number of characters (default 30)
- `""`, `--no-lowercase`   # not contain lowercase character
- `""`, `--no-numeric`     # not contain numeric character
- `""`, `--no-symbol`      # not contain symbol character
- `""`, `--no-uppercase`   # nto contain uppercase character

### example

```shell
ggen                   # 30 characters containing lowercase, uppercase, numeric and symbol.
ggen -c 15 --no-symbol # 15 characters not containing symbol.
```

## License

ggen is released under the MIT license. See [LICENSE](./LICENSE).
