# Random user agent

This is a simple package for random browser user agent.

## Installation
```bash
$ go get github.com/mark1002/randomuseragent
```

## Usage
```go
package main

import (
    "fmt"

    "github.com/mark1002/randomuseragent"
)

func main() {
    fmt.Println(randomuseragent.GetRandomUserAgent())
    // "Mozilla/5.0 (X11; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0"
}

```