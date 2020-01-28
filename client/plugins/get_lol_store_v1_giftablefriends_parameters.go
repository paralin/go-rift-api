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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLolStoreV1GiftablefriendsParams creates a new GetLolStoreV1GiftablefriendsParams object
// with the default values initialized.
func NewGetLolStoreV1GiftablefriendsParams() *GetLolStoreV1GiftablefriendsParams {

	return &GetLolStoreV1GiftablefriendsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolStoreV1GiftablefriendsParamsWithTimeout creates a new GetLolStoreV1GiftablefriendsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolStoreV1GiftablefriendsParamsWithTimeout(timeout time.Duration) *GetLolStoreV1GiftablefriendsParams {

	return &GetLolStoreV1GiftablefriendsParams{

		timeout: timeout,
	}
}

// NewGetLolStoreV1GiftablefriendsParamsWithContext creates a new GetLolStoreV1GiftablefriendsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolStoreV1GiftablefriendsParamsWithContext(ctx context.Context) *GetLolStoreV1GiftablefriendsParams {

	return &GetLolStoreV1GiftablefriendsParams{

		Context: ctx,
	}
}

// NewGetLolStoreV1GiftablefriendsParamsWithHTTPClient creates a new GetLolStoreV1GiftablefriendsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolStoreV1GiftablefriendsParamsWithHTTPClient(client *http.Client) *GetLolStoreV1GiftablefriendsParams {

	return &GetLolStoreV1GiftablefriendsParams{
		HTTPClient: client,
	}
}

/*GetLolStoreV1GiftablefriendsParams contains all the parameters to send to the API endpoint
for the get lol store v1 giftablefriends operation typically these are written to a http.Request
*/
type GetLolStoreV1GiftablefriendsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol store v1 giftablefriends params
func (o *GetLolStoreV1GiftablefriendsParams) WithTimeout(timeout time.Duration) *GetLolStoreV1GiftablefriendsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol store v1 giftablefriends params
func (o *GetLolStoreV1GiftablefriendsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol store v1 giftablefriends params
func (o *GetLolStoreV1GiftablefriendsParams) WithContext(ctx context.Context) *GetLolStoreV1GiftablefriendsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol store v1 giftablefriends params
func (o *GetLolStoreV1GiftablefriendsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol store v1 giftablefriends params
func (o *GetLolStoreV1GiftablefriendsParams) WithHTTPClient(client *http.Client) *GetLolStoreV1GiftablefriendsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol store v1 giftablefriends params
func (o *GetLolStoreV1GiftablefriendsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolStoreV1GiftablefriendsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
