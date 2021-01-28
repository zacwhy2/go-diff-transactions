# Getting started

```sh
$ go build

$ ./go-diff-transactions
Usage: ./go-diff-transactions FILE1 FILE2

$ ./go-diff-transactions diff/testdata/po/local.csv diff/testdata/po/remote.csv
local
< [2020-01-02 100][0] [2020-01-02 100 local only]
po
> [2020-01-03 100][0] [03 Jan 2020 po only S$1.00]

# or

$ go run main.go diff/testdata/po/local.csv diff/testdata/po/remote.csv
local
< [2020-01-02 100][0] [2020-01-02 100 local only]
po
> [2020-01-03 100][0] [03 Jan 2020 po only S$1.00]
```
