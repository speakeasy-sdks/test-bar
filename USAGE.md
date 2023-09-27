<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	testbar "github.com/speakeasy-sdks/test-bar"
	"github.com/speakeasy-sdks/test-bar/pkg/models/shared"
	"github.com/speakeasy-sdks/test-bar/pkg/models/operations"
)

func main() {
    s := testbar.New(
        testbar.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Drinks.GetDrink(ctx, operations.GetDrinkRequest{
        Name: "Terrence Rau",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Drink != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->