# goscan

goscan libray for command parameters scan usefully.

#import package

```shell

go get github.com/lingdor/goscan
```

#demo

```go
scanner,err:=goscan.NewScanStd()
```
or
```go
words,err:=scanner.ScanWords()
```

input: set key "content\"good"

output: []{"set","key","content\"good"}

or
```go
word,end,err:=scanner.Scan()
if word == "get" {
     keys,err:=scanner.ScanWords()
    ...
}
```
