# Config
(*Config*)

### Available Operations

* [SubscribeToWebhooks](#subscribetowebhooks) - Subscribe to webhooks.

## SubscribeToWebhooks

Subscribe to webhooks.

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
    res, err := s.Config.SubscribeToWebhooks(ctx, []SubscribeToWebhooksRequestBody{
        operations.SubscribeToWebhooksRequestBody{
            URL: testbar.String("ipsa"),
            Webhook: testbar.String("delectus"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [[]operations.SubscribeToWebhooksRequestBody](../../models//.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |


### Response

**[*operations.SubscribeToWebhooksResponse](../../models/operations/subscribetowebhooksresponse.md), error**

