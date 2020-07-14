[![GoDoc](https://godoc.org/github.com/IamStefin/nsego?status.svg)](https://godoc.org/github.com/IamStefin/visago)

# visago

## Installation

To install Gin package, you need to install Go and set your Go workspace first.

1. The first need [Go](https://golang.org/) installed (**version 1.11+ is required**), then you can use the below Go command to install nsego.

```sh
go get -u github.com/iamstefin/visago
```

2. Import it in your code:

```go
import "github.com/iamstefin/visago"
```

## Usage

```go
package main

import (
  "fmt"
  "github.com/iamstefin/visago"
)

func main()  {

  vis := visago.VisaFetcher("France","Indonesia").VisaReq
  fmt.Println(vis)
}
```
Feel free to make this code better! :)
