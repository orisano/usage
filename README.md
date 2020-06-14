# usage
## Installation
```
go get github.com/orisano/usage
```

## How to use
```go
package main

import (
    "flag"
    "fmt"
    
    "github.com/orisano/usage"
)

func main() {
    flag.Usage = usage.Ordered("n", "N")
    n := flag.Int("n", 0, "number n")
    N := flag.Int("N", 1, "number N")
    flag.Parse()
    
    fmt.Println(*n, *N)
}
```

## Author
Nao Yonashiro (@orisano)

## License
MIT
