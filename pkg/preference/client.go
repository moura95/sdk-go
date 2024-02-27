package preference

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/baseclient"
)

const (
	urlBase   = "https://api.mercadopago.com/checkout/preferences"
	urlSearch = urlBase + "/search"
	urlWithID = urlBase + "/:id"
)

// Client contains the methods to interact with the Preference API.
type Client interface {
	// Create a preference with information about a product or service and obtain the URL needed to start the payment flow.
	// It is a post request to the endpoint: https://api.mercadopago.com/checkout/preferences
	// Reference: https://www.mercadopago.com/developers/en/reference/preferences/_checkout_preferences/post
	Create(ctx context.Context, request Request) (*Response, error)

	// Get finds a preference by ID.
	// It is a get request to the endpoint: https://api.mercadopago.com/checkout/preferences/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/preferences/_checkout_preferences_id/get
	Get(ctx context.Context, id string) (*Response, error)

	// Update updates details for a payment preference.
	// It is a put request to the endpoint: https://api.mercadopago.com/checkout/preferences/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/preferences/_checkout_preferences_id/put
	Update(ctx context.Context, request Request, id string) (*Response, error)

	// Search finds all preference information generated through specific filters
	// It is a get request to the endpoint: https://api.mercadopago.com/checkout/preferences/search
	// Reference: https://www.mercadopago.com/developers/en/reference/preferences/_checkout_preferences_search/get
	Search(ctx context.Context, request SearchRequest) (*SearchResponsePage, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// NewClient returns a new Preference API Client.
func NewClient(c *config.Config) Client {
	return &client{
		cfg: c,
	}
}

func (c *client) Create(ctx context.Context, request Request) (*Response, error) {
	result, err := baseclient.Post[*Response](ctx, c.cfg, urlBase, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Get(ctx context.Context, id string) (*Response, error) {
	params := map[string]string{
		"id": id,
	}

	result, err := baseclient.Get[*Response](ctx, c.cfg, urlWithID, baseclient.WithPathParams(params))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Update(ctx context.Context, request Request, id string) (*Response, error) {
	params := map[string]string{
		"id": id,
	}

	result, err := baseclient.Put[*Response](ctx, c.cfg, urlWithID, request, baseclient.WithPathParams(params))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Search(ctx context.Context, request SearchRequest) (*SearchResponsePage, error) {
	params := request.Parameters()

	url, err := url.Parse(urlSearch)
	if err != nil {
		return nil, fmt.Errorf("error parsing url: %w", err)
	}
	url.RawQuery = params

	result, err := baseclient.Get[*SearchResponsePage](ctx, c.cfg, url.String())
	if err != nil {
		return nil, err
	}

	return result, nil
}