# ggen

## Overview

'ggen' is a simple CLI tool via Golang.

You can select some options for generating passwords.

The options are like:

- specifing the count
- containing lowercase / uppercase characters or not
- containing numeric characters or not
- containing symbol or not

## Usage

Run one simple command below:

```shell
ggen # default setting is '10 lowercase characters'.
```

### options

- `-c`, `--count [int]`    # specify the number of characters (default 30)
- `""`, `--no-lowercase`   # not contain lowercase character
- `""`, `--no-numeric`     # not contain numeric character
- `""`, `--no-symbol`      # not contain symbol character
- `""`, `--no-uppercase`   # nto contain uppercase character

example:

```shell
ggen                   # 30 characters containing lowercase, uppercase, numeric and symbol.
ggen -c 15 --no-symbol # 15 characters not containing symbol.
```

## License

ggen is released under the MIT license. See [LICENSE](./LICENSE).
