# Runtime

A helper package which initializes the Go runtime.

## Installation

``` bash
$ go get -u github.com/vitalie/runtime/setup
```

## Usage

``` go
package main

import (
    "fmt"

    _ "github.com/vitalie/runtime/setup"

    )

func main() {
    fmt.Println("Hello, world!")

}
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
