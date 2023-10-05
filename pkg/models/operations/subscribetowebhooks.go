// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/test-bar/pkg/models/shared"
	"net/http"
)

type SubscribeToWebhooksRequestBodyWebhook string

const (
	SubscribeToWebhooksRequestBodyWebhookStockUpdate SubscribeToWebhooksRequestBodyWebhook = "stockUpdate"
)

func (e SubscribeToWebhooksRequestBodyWebhook) ToPointer() *SubscribeToWebhooksRequestBodyWebhook {
	return &e
}

func (e *SubscribeToWebhooksRequestBodyWebhook) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "stockUpdate":
		*e = SubscribeToWebhooksRequestBodyWebhook(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SubscribeToWebhooksRequestBodyWebhook: %v", v)
	}
}

type SubscribeToWebhooksRequestBody struct {
	URL     *string                                `json:"url,omitempty"`
	Webhook *SubscribeToWebhooksRequestBodyWebhook `json:"webhook,omitempty"`
}

func (o *SubscribeToWebhooksRequestBody) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *SubscribeToWebhooksRequestBody) GetWebhook() *SubscribeToWebhooksRequestBodyWebhook {
	if o == nil {
		return nil
	}
	return o.Webhook
}

type SubscribeToWebhooksResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// An unknown error occurred interacting with the API.
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SubscribeToWebhooksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SubscribeToWebhooksResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *SubscribeToWebhooksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SubscribeToWebhooksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
