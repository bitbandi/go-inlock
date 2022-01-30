# go-inlock
INLOCK.io API library in Golang

# Example

```php
package main

import (
	"context"
	"fmt"
	"github.com/bitbandi/go-inlock"
)

func main() {
	auth := context.WithValue(context.Background(), inlock.ContextAPIKey, inlock.APIKey{
		Key:    "apikey",
		Secret: "apisecret",
	})
	cfg := inlock.NewConfiguration()
	client := inlock.NewAPIClient(cfg)
	balance, resp, err := client.OtherApi.GetBalance(auth, &inlock.OtherApiGetBalanceOpts{})
	fmt.Println(balance, resp, err)
}
```
