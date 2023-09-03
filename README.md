# goscan

goscan libray for command parameters scan usefully.

#demo

```go
scanner,err:=goscan.NewScanStd()
```
or
```go
words,err:=scanner.ScanWords()
```
or
```go
word,end,err:=scanner.Scan()
if word == "get" {
     keys,err:=scanner.ScanWords()
    ...
}
```