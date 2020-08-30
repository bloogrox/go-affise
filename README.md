## Installation

```
go get github.com/bloogrox/go-affise
```

## Getting Started

### Init client

```go
import (
	"encoding/json"
	"fmt"
	"log"

	affise "github.com/bloogrox/go-affise"
)
```

```go
network := "mynetwork"
token := "c6d5b6ad56b5..."

client := affise.New(network, token)
```

### Get Offer

```go
offer, err := client.OfferGet(42)

if err != nil {
    log.Fatalf(err.Error())
}

fmt.Printf("%+v \n", offer)
```

### Edit Offer

```go
import (
    // ...
    "net/url"
    // ...
)
```

```go
data := url.Values{}
data.Set("title", "New title for this offer")

err := client.OfferEdit(42, &data)

if err != nil {
    log.Fatalf(err.Error())
}
```