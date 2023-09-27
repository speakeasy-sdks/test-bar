# Ingredients
(*Ingredients*)

## Overview

The ingredients endpoints.

### Available Operations

* [ListIngredients](#listingredients) - Get a list of ingredients.

## ListIngredients

Get a list of ingredients, if authenticated this will include stock levels and product codes otherwise it will only include public information.

### Example Usage

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
    res, err := s.Ingredients.ListIngredients(ctx, operations.ListIngredientsRequest{
        Ingredients: []string{
            "iusto",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Ingredients != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListIngredientsRequest](../../models/operations/listingredientsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ListIngredientsResponse](../../models/operations/listingredientsresponse.md), error**
