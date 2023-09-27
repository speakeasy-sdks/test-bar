# Orders
(*Orders*)

## Overview

The orders endpoints.

### Available Operations

* [CreateOrder](#createorder) - Create an order.

## CreateOrder

Create an order for a drink.

### Example Usage

```go
package main

import(
	"context"
	"log"
	testbar "github.com/speakeasy-sdks/test-bar"
	"github.com/speakeasy-sdks/test-bar/pkg/models/shared"
	"github.com/speakeasy-sdks/test-bar/pkg/models/operations"
	"github.com/speakeasy-sdks/test-bar/pkg/models/callbacks"
	"net/http"
)

func main() {
    s := testbar.New(
        testbar.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Orders.CreateOrder(ctx, operations.CreateOrderRequest{
        RequestBody: []shared.OrderInput{
            shared.OrderInput{
                ProductCode: "NAC-3F2D1",
                Quantity: 392785,
                Type: shared.OrderTypeIngredient,
            },
        },
        CallbackURL: testbar.String("temporibus"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.CreateOrderRequest](../../models/operations/createorderrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CreateOrderResponse](../../models/operations/createorderresponse.md), error**

