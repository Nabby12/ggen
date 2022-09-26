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

- `-c`, `--count [int]`    # generated password word count (default 20)
- `""`, `--no-lowercase`   # not contain lowercase letters
- `-n`, `--numeric`        # contain numbers
- `-s`, `--symbol`         # contain symbols
- `-u`, `--uppercase`      # contain uppercase letters

example:

```shell
ggen -c 30 -n # 30 characters containing lowercase and numeric.
```

## License

ggen is released under the MIT license. See [LICENSE](./LICENSE).
