// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLolInventoryV1WalletParams creates a new GetLolInventoryV1WalletParams object
// with the default values initialized.
func NewGetLolInventoryV1WalletParams() *GetLolInventoryV1WalletParams {
	var ()
	return &GetLolInventoryV1WalletParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolInventoryV1WalletParamsWithTimeout creates a new GetLolInventoryV1WalletParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolInventoryV1WalletParamsWithTimeout(timeout time.Duration) *GetLolInventoryV1WalletParams {
	var ()
	return &GetLolInventoryV1WalletParams{

		timeout: timeout,
	}
}

// NewGetLolInventoryV1WalletParamsWithContext creates a new GetLolInventoryV1WalletParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolInventoryV1WalletParamsWithContext(ctx context.Context) *GetLolInventoryV1WalletParams {
	var ()
	return &GetLolInventoryV1WalletParams{

		Context: ctx,
	}
}

// NewGetLolInventoryV1WalletParamsWithHTTPClient creates a new GetLolInventoryV1WalletParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolInventoryV1WalletParamsWithHTTPClient(client *http.Client) *GetLolInventoryV1WalletParams {
	var ()
	return &GetLolInventoryV1WalletParams{
		HTTPClient: client,
	}
}

/*GetLolInventoryV1WalletParams contains all the parameters to send to the API endpoint
for the get lol inventory v1 wallet operation typically these are written to a http.Request
*/
type GetLolInventoryV1WalletParams struct {

	/*CurrencyTypes*/
	CurrencyTypes []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol inventory v1 wallet params
func (o *GetLolInventoryV1WalletParams) WithTimeout(timeout time.Duration) *GetLolInventoryV1WalletParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol inventory v1 wallet params
func (o *GetLolInventoryV1WalletParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol inventory v1 wallet params
func (o *GetLolInventoryV1WalletParams) WithContext(ctx context.Context) *GetLolInventoryV1WalletParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol inventory v1 wallet params
func (o *GetLolInventoryV1WalletParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol inventory v1 wallet params
func (o *GetLolInventoryV1WalletParams) WithHTTPClient(client *http.Client) *GetLolInventoryV1WalletParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol inventory v1 wallet params
func (o *GetLolInventoryV1WalletParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurrencyTypes adds the currencyTypes to the get lol inventory v1 wallet params
func (o *GetLolInventoryV1WalletParams) WithCurrencyTypes(currencyTypes []string) *GetLolInventoryV1WalletParams {
	o.SetCurrencyTypes(currencyTypes)
	return o
}

// SetCurrencyTypes adds the currencyTypes to the get lol inventory v1 wallet params
func (o *GetLolInventoryV1WalletParams) SetCurrencyTypes(currencyTypes []string) {
	o.CurrencyTypes = currencyTypes
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolInventoryV1WalletParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesCurrencyTypes := o.CurrencyTypes

	joinedCurrencyTypes := swag.JoinByFormat(valuesCurrencyTypes, "")
	// query array param currencyTypes
	if err := r.SetQueryParam("currencyTypes", joinedCurrencyTypes...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}