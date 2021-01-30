# Getting started

```sh
$ go build

$ ./go-diff-transactions
Usage: ./go-diff-transactions FILE1 FILE2

$ ./go-diff-transactions diff/testdata/po.csv diff/testdata/local.csv
po
< [2020-01-02 100][0] [02 Jan 2020 po only S$1.00]
local
> [2020-01-03 100][0] [2020-01-03 100 local only]

# or

$ go run main.go diff/testdata/po.csv diff/testdata/local.csv
po
< [2020-01-02 100][0] [02 Jan 2020 po only S$1.00]
local
> [2020-01-03 100][0] [2020-01-03 100 local only]
```
