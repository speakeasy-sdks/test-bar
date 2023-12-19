# Drinks
(*Drinks*)

## Overview

The drinks endpoints.

### Available Operations

* [GetDrink](#getdrink) - Get a drink.
* [ListDrinks](#listdrinks) - Get a list of drinks.

## GetDrink

Get a drink by name, if authenticated this will include stock levels and product codes otherwise it will only include public information.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/test-bar/pkg/models/shared"
	testbar "github.com/speakeasy-sdks/test-bar"
	"context"
	"github.com/speakeasy-sdks/test-bar/pkg/models/operations"
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

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetDrinkRequest](../../pkg/models/operations/getdrinkrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetDrinkResponse](../../pkg/models/operations/getdrinkresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 5XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListDrinks

Get a list of drinks, if authenticated this will include stock levels and product codes otherwise it will only include public information.

### Example Usage

```go
package main

import(
	testbar "github.com/speakeasy-sdks/test-bar"
	"context"
	"github.com/speakeasy-sdks/test-bar/pkg/models/operations"
	"log"
)

func main() {
    s := testbar.New()

    ctx := context.Background()
    res, err := s.Drinks.ListDrinks(ctx, operations.ListDrinksRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListDrinksRequest](../../pkg/models/operations/listdrinksrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListDrinksResponse](../../pkg/models/operations/listdrinksresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 5XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
