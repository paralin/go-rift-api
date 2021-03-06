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

// NewPostLolPersonalizedOffersV1OffersByIDPurchaseParams creates a new PostLolPersonalizedOffersV1OffersByIDPurchaseParams object
// with the default values initialized.
func NewPostLolPersonalizedOffersV1OffersByIDPurchaseParams() *PostLolPersonalizedOffersV1OffersByIDPurchaseParams {
	var ()
	return &PostLolPersonalizedOffersV1OffersByIDPurchaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolPersonalizedOffersV1OffersByIDPurchaseParamsWithTimeout creates a new PostLolPersonalizedOffersV1OffersByIDPurchaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolPersonalizedOffersV1OffersByIDPurchaseParamsWithTimeout(timeout time.Duration) *PostLolPersonalizedOffersV1OffersByIDPurchaseParams {
	var ()
	return &PostLolPersonalizedOffersV1OffersByIDPurchaseParams{

		timeout: timeout,
	}
}

// NewPostLolPersonalizedOffersV1OffersByIDPurchaseParamsWithContext creates a new PostLolPersonalizedOffersV1OffersByIDPurchaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolPersonalizedOffersV1OffersByIDPurchaseParamsWithContext(ctx context.Context) *PostLolPersonalizedOffersV1OffersByIDPurchaseParams {
	var ()
	return &PostLolPersonalizedOffersV1OffersByIDPurchaseParams{

		Context: ctx,
	}
}

// NewPostLolPersonalizedOffersV1OffersByIDPurchaseParamsWithHTTPClient creates a new PostLolPersonalizedOffersV1OffersByIDPurchaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolPersonalizedOffersV1OffersByIDPurchaseParamsWithHTTPClient(client *http.Client) *PostLolPersonalizedOffersV1OffersByIDPurchaseParams {
	var ()
	return &PostLolPersonalizedOffersV1OffersByIDPurchaseParams{
		HTTPClient: client,
	}
}

/*PostLolPersonalizedOffersV1OffersByIDPurchaseParams contains all the parameters to send to the API endpoint
for the post lol personalized offers v1 offers by Id purchase operation typically these are written to a http.Request
*/
type PostLolPersonalizedOffersV1OffersByIDPurchaseParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol personalized offers v1 offers by Id purchase params
func (o *PostLolPersonalizedOffersV1OffersByIDPurchaseParams) WithTimeout(timeout time.Duration) *PostLolPersonalizedOffersV1OffersByIDPurchaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol personalized offers v1 offers by Id purchase params
func (o *PostLolPersonalizedOffersV1OffersByIDPurchaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol personalized offers v1 offers by Id purchase params
func (o *PostLolPersonalizedOffersV1OffersByIDPurchaseParams) WithContext(ctx context.Context) *PostLolPersonalizedOffersV1OffersByIDPurchaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol personalized offers v1 offers by Id purchase params
func (o *PostLolPersonalizedOffersV1OffersByIDPurchaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol personalized offers v1 offers by Id purchase params
func (o *PostLolPersonalizedOffersV1OffersByIDPurchaseParams) WithHTTPClient(client *http.Client) *PostLolPersonalizedOffersV1OffersByIDPurchaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol personalized offers v1 offers by Id purchase params
func (o *PostLolPersonalizedOffersV1OffersByIDPurchaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post lol personalized offers v1 offers by Id purchase params
func (o *PostLolPersonalizedOffersV1OffersByIDPurchaseParams) WithID(id string) *PostLolPersonalizedOffersV1OffersByIDPurchaseParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post lol personalized offers v1 offers by Id purchase params
func (o *PostLolPersonalizedOffersV1OffersByIDPurchaseParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolPersonalizedOffersV1OffersByIDPurchaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
