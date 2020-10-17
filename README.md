# control-go-sdk

Golang client library for the [API of the VSHN control panel][api-docs].

## Installation

    go get github.com/vshn/control-go-sdk

## Usage

```go
import "github.com/vshn/control-go-sdk"
```

### Authentication

In order to access the API, you need an access token. See ["API: Basics" in the VSHN Knowledge Base](https://kb.vshn.ch/kb/api_basics.html#_authentication) how to obtain a token.

### Example

```go
package main

import (
	"github.com/vshn/control-go-sdk"
)

func main() {
	c := control.NewClientFromToken("my-vshn-portal-token")
	fqdns, _, _ := c.Servers.ListFQDNs("")
	print(fqdns)
}
```

## Documentation

For more information about the VSHN Control APIs, check out the [API documentation][api-docs]

For details about this library, see the [GoDoc documentation](http://pkg.go.dev/github.com/vshn/control-go-sdk).

## Contributing

Pull requests welcome! Please see the [contribution guidelines](CONTRIBUTING.md).

[api-docs]: https://kb.vshn.ch/kb/api_basics.html
