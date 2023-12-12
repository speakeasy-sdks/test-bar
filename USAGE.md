<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	testbar "github.com/speakeasy-sdks/test-bar"
	"github.com/speakeasy-sdks/test-bar/pkg/models/operations"
	"github.com/speakeasy-sdks/test-bar/pkg/models/shared"
	"log"
)

func main() {
	s := testbar.New(
		testbar.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Drinks.GetDrink(ctx, operations.GetDrinkRequest{
		Name: "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Drink != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->