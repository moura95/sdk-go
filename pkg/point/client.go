package point

import (
	"context"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/baseclient"
)

const (
	urlBase                = "https://api.mercadopago.com/point"
	urlDevices             = urlBase + "/integration-api/devices/"
	urlPaymentIntent       = urlDevices + ":device_id/payment-intents"
	urlPaymentIntentGet    = urlBase + "/integration-api/payment-intents/:payment_intent_id"
	urlPaymentIntentCancel = urlDevices + ":device_id/payment-intents/:payment_intent_id"
	urlDevicesWithID       = urlDevices + ":device_id"
)

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// Client contains the methods to interact with the Point API.
type Client interface {

	// Create a point payment intent.
	// It is a post request to the endpoint: https://api.mercadopago.com/point/integration-api/devices/{device_id}/payment-intents
	// Reference: https://www.mercadopago.com/developers/en/reference/integrations_api_paymentintent_mlb/_point_integration-api_devices_deviceid_payment-intents/post
	Create(ctx context.Context, deviceID string, request CreateRequest) (*Response, error)

	// Get a point payment intent.
	// It is a get request to the endpoint: https://api.mercadopago.com/point/integration-api/payment-intents/{payment_intent_id}
	// Reference: https://www.mercadopago.com/developers/en/reference/integrations_api/_point_integration-api_payment-intents_paymentintentid/get
	Get(ctx context.Context, paymentIntentID string) (*Response, error)

	// Cancel a point payment intent.
	// It is a cancel request to the endpoint: https://api.mercadopago.com/point/integration-api/devices/{device_id}/payment-intents/{payment_intent_id}
	// Reference: https://www.mercadopago.com/developers/en/reference/integrations_api/_point_integration-api_devices_deviceid_payment-intents_paymentintentid/delete
	Cancel(ctx context.Context, deviceID string, paymentIntentID string) (*CancelResponse, error)

	// ListDevices retrieve devices.
	// It is a get request to the endpoint: https://api.mercadopago.com/point/integration-api/devices
	// Reference: https://www.mercadopago.com/developers/en/reference/integrations_api/_point_integration-api_devices/get
	ListDevices(ctx context.Context) (*DevicesResponse, error)

	// UpdateDeviceOperatingMode update operating mode from device.
	// It is a patch request to the endpoint: https://api.mercadopago.com/point/integration-api/devices/{device-id}
	// Reference: https://www.mercadopago.com/developers/en/reference/integrations_api/_point_integration-api_devices_device-id/patch
	UpdateDeviceOperatingMode(ctx context.Context, deviceID string, request UpdateDeviceOperatingModeRequest) (*OperatingModeResponse, error)
}

// NewClient returns a new Point Client.
func NewClient(c *config.Config) Client {
	return &client{cfg: c}
}

func (c *client) Create(ctx context.Context, deviceID string, request CreateRequest) (*Response, error) {
	params := map[string]string{
		"device_id": deviceID,
	}

	result, err := baseclient.Post[*Response](ctx, c.cfg, urlPaymentIntent, request, baseclient.WithPathParams(params))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Get(ctx context.Context, paymentIntentID string) (*Response, error) {
	params := map[string]string{
		"payment_intent_id": paymentIntentID,
	}

	result, err := baseclient.Get[*Response](ctx, c.cfg, urlPaymentIntentGet, baseclient.WithPathParams(params))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Cancel(ctx context.Context, deviceID string, paymentIntentID string) (*CancelResponse, error) {
	params := map[string]string{
		"device_id":         deviceID,
		"payment_intent_id": paymentIntentID,
	}

	result, err := baseclient.Delete[*CancelResponse](ctx, c.cfg, urlPaymentIntentCancel, nil, baseclient.WithPathParams(params))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) ListDevices(ctx context.Context) (*DevicesResponse, error) {
	result, err := baseclient.Get[*DevicesResponse](ctx, c.cfg, urlDevices)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) UpdateDeviceOperatingMode(ctx context.Context, deviceID string, request UpdateDeviceOperatingModeRequest) (*OperatingModeResponse, error) {
	params := map[string]string{
		"device_id": deviceID,
	}

	result, err := baseclient.Patch[*OperatingModeResponse](ctx, c.cfg, urlDevicesWithID, request, baseclient.WithPathParams(params))
	if err != nil {
		return nil, err
	}

	return result, nil
}
