# typeassert

[![godoc.org][godoc-badge]][godoc]

`typeassert` finds unreceived type assert result.

```go
package main

import "fmt"

func main() {
    var n interface{} = 123
    v := n.(int)
    fmt.Println(v)
}
```

```sh
$ go vet -vettool=`which typeassert` main.go
./main.go:7:2: type assertion must be checked
```

<!-- links -->
[godoc]: https://godoc.org/github.com/gostaticanalysis/typeassert
[godoc-badge]: https://img.shields.io/badge/godoc-reference-4F73B3.svg?style=flat-square&label=%20godoc.org
