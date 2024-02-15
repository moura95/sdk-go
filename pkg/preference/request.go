package preference

import (
	"time"
)

// PreferenceRequest contains parameters to create/update a preference.
type Request struct {
	AdditionalInfo      string                 `json:"additional_info,omitempty"`
	AutoReturn          string                 `json:"auto_return,omitempty"`
	ExternalReference   string                 `json:"external_reference,omitempty"`
	Marketplace         string                 `json:"marketplace,omitempty"`
	OperationType       string                 `json:"operation_type,omitempty"`
	NotificationUrl     string                 `json:"notification_url,omitempty"`
	Purpose             string                 `json:"purpose,omitempty"`
	StatementDescriptor string                 `json:"statement_descriptor,omitempty"`
	BinaryMode          *bool                  `json:"binary_mode,omitempty"`
	Expires             *bool                  `json:"expires,omitempty"`
	MarketplaceFee      float64                `json:"marketplace_fee,omitempty"`
	ProcessingModes     []string               `json:"processing_modes,omitempty"`
	Metadata            map[string]interface{} `json:"metadata,omitempty"`

	DateOfExpiration    *time.Time                            `json:"date_of_expiration,omitempty"`
	ExpirationDateFrom  *time.Time                            `json:"expiration_date_from,omitempty"`
	ExpirationDateTo    *time.Time                            `json:"expiration_date_to,omitempty"`
	BackUrls            *PreferenceBackUrlsRequest            `json:"back_urls,omitempty"`
	DifferentialPricing *PreferenceDifferentialPricingRequest `json:"differential_pricing,omitempty"`
	Payer               *PreferencePayerRequest               `json:"payer,omitempty"`
	PaymentMethods      *PreferencePaymentMethodsRequest      `json:"payment_methods,omitempty"`
	Shipments           *PreferenceShipmentsRequest           `json:"shipments,omitempty"`
	Items               []PreferenceItemRequest               `json:"items,omitempty"`
	Taxes               []PreferenceTaxRequest                `json:"taxes,omitempty"`
	Tracks              []PreferenceTrackRequest              `json:"tracks,omitempty"`
}

// PreferenceBackUrlsRequest contains callback URLs.
type PreferenceBackUrlsRequest struct {
	Success string `json:"success,omitempty"`
	Pending string `json:"pending,omitempty"`
	Failure string `json:"failure,omitempty"`
}

// PreferenceDifferentialPricingRequest contains information about differential pricing configuration.
type PreferenceDifferentialPricingRequest struct {
	ID int64 `json:"id,omitempty"`
}

// PreferenceItemRequest represents a purchased item.
type PreferenceItemRequest struct {
	ID          string  `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	PictureURL  string  `json:"picture_url,omitempty"`
	CategoryID  string  `json:"category_id,omitempty"`
	CurrencyID  string  `json:"currency_id,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	UnitPrice   float64 `json:"unit_price,omitempty"`
}

// PreferencePayerRequest contains payer information in the preference.
type PreferencePayerRequest struct {
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	Email   string `json:"email,omitempty"`

	DateCreated    *time.Time             `json:"date_created,omitempty"`
	Phone          *PhoneRequest          `json:"phone,omitempty"`
	Identification *IdentificationRequest `json:"identification,omitempty"`
	Address        *AddressRequest        `json:"address,omitempty"`
}

// PhoneRequest represents a telephone number.
type PhoneRequest struct {
	AreaCode string `json:"area_code,omitempty"`
	Number   string `json:"number,omitempty"`
}

// IdentificationRequest is a base type that represents identifications, such as customer identification.
type IdentificationRequest struct {
	Type   string `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
}

// AddressRequest represents an address.
type AddressRequest struct {
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
}

// PreferencePaymentMethodsRequest contains information about payment methods in the preference.
type PreferencePaymentMethodsRequest struct {
	DefaultPaymentMethodID string `json:"default_payment_method_id,omitempty"`
	Installments           int    `json:"installments,omitempty"`
	DefaultInstallments    int    `json:"default_installments,omitempty"`

	ExcludedPaymentMethods []PreferencePaymentMethodRequest `json:"excluded_payment_methods,omitempty"`
	ExcludedPaymentTypes   []PreferencePaymentTypeRequest   `json:"excluded_payment_types,omitempty"`
}

// PreferencePaymentMethodRequest contains information about the payment method in the preference.
type PreferencePaymentMethodRequest struct {
	ID string `json:"id,omitempty"`
}

// PreferencePaymentTypeRequest contains information about the payment type in the preference.
type PreferencePaymentTypeRequest struct {
	ID string `json:"id,omitempty"`
}

// PreferenceShipmentsRequest contains information about shipments in the preference.
type PreferenceShipmentsRequest struct {
	Mode                  string  `json:"mode,omitempty"`
	Dimensions            string  `json:"dimensions,omitempty"`
	DefaultShippingMethod string  `json:"default_shipping_method,omitempty"`
	LocalPickup           bool    `json:"local_pickup,omitempty"`
	FreeShipping          bool    `json:"free_shipping,omitempty"`
	ExpressShipment       bool    `json:"express_shipment,omitempty"`
	Cost                  float64 `json:"cost,omitempty"`

	ReceiverAddress *PreferenceReceiverAddressRequest `json:"receiver_address,omitempty"`
	FreeMethods     []PreferenceFreeMethodRequest     `json:"free_methods,omitempty"`
}

// PreferenceFreeMethodRequest contains information about free shipping methods in the preference.
type PreferenceFreeMethodRequest struct {
	ID int64 `json:"id,omitempty"`
}

// PreferenceReceiverAddressRequest contains information about the send address in the preference.
type PreferenceReceiverAddressRequest struct {
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
	CountryName  string `json:"country_name,omitempty"`
	StateName    string `json:"state_name,omitempty"`
	Floor        string `json:"floor,omitempty"`
	Apartment    string `json:"apartment,omitempty"`
	CityName     string `json:"city_name,omitempty"`
}

// PreferenceTaxRequest contains information about taxes in the preference.
type PreferenceTaxRequest struct {
	Type  string  `json:"type,omitempty"`
	Value float64 `json:"value,omitempty"`
}

// PreferenceTrackRequest contains information about the tracking to be performed during user interaction in the Checkout flow.
type PreferenceTrackRequest struct {
	Type string `json:"type,omitempty"`

	Values *PreferenceTrackValuesRequest `json:"values,omitempty"`
}

// PreferenceTrackValuesRequest contains the values ​​of the tracks to be executed during user interaction in the Checkout flow.
type PreferenceTrackValuesRequest struct {
	ConversionID    string `json:"conversion_id,omitempty"`
	ConversionLabel string `json:"conversion_label,omitempty"`
	PixelID         string `json:"pixel_id,omitempty"`
}
