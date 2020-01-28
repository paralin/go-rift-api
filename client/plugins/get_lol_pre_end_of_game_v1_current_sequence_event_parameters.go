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

// NewGetLolPreEndOfGameV1CurrentSequenceEventParams creates a new GetLolPreEndOfGameV1CurrentSequenceEventParams object
// with the default values initialized.
func NewGetLolPreEndOfGameV1CurrentSequenceEventParams() *GetLolPreEndOfGameV1CurrentSequenceEventParams {

	return &GetLolPreEndOfGameV1CurrentSequenceEventParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolPreEndOfGameV1CurrentSequenceEventParamsWithTimeout creates a new GetLolPreEndOfGameV1CurrentSequenceEventParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolPreEndOfGameV1CurrentSequenceEventParamsWithTimeout(timeout time.Duration) *GetLolPreEndOfGameV1CurrentSequenceEventParams {

	return &GetLolPreEndOfGameV1CurrentSequenceEventParams{

		timeout: timeout,
	}
}

// NewGetLolPreEndOfGameV1CurrentSequenceEventParamsWithContext creates a new GetLolPreEndOfGameV1CurrentSequenceEventParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolPreEndOfGameV1CurrentSequenceEventParamsWithContext(ctx context.Context) *GetLolPreEndOfGameV1CurrentSequenceEventParams {

	return &GetLolPreEndOfGameV1CurrentSequenceEventParams{

		Context: ctx,
	}
}

// NewGetLolPreEndOfGameV1CurrentSequenceEventParamsWithHTTPClient creates a new GetLolPreEndOfGameV1CurrentSequenceEventParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolPreEndOfGameV1CurrentSequenceEventParamsWithHTTPClient(client *http.Client) *GetLolPreEndOfGameV1CurrentSequenceEventParams {

	return &GetLolPreEndOfGameV1CurrentSequenceEventParams{
		HTTPClient: client,
	}
}

/*GetLolPreEndOfGameV1CurrentSequenceEventParams contains all the parameters to send to the API endpoint
for the get lol pre end of game v1 current sequence event operation typically these are written to a http.Request
*/
type GetLolPreEndOfGameV1CurrentSequenceEventParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol pre end of game v1 current sequence event params
func (o *GetLolPreEndOfGameV1CurrentSequenceEventParams) WithTimeout(timeout time.Duration) *GetLolPreEndOfGameV1CurrentSequenceEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol pre end of game v1 current sequence event params
func (o *GetLolPreEndOfGameV1CurrentSequenceEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol pre end of game v1 current sequence event params
func (o *GetLolPreEndOfGameV1CurrentSequenceEventParams) WithContext(ctx context.Context) *GetLolPreEndOfGameV1CurrentSequenceEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol pre end of game v1 current sequence event params
func (o *GetLolPreEndOfGameV1CurrentSequenceEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol pre end of game v1 current sequence event params
func (o *GetLolPreEndOfGameV1CurrentSequenceEventParams) WithHTTPClient(client *http.Client) *GetLolPreEndOfGameV1CurrentSequenceEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol pre end of game v1 current sequence event params
func (o *GetLolPreEndOfGameV1CurrentSequenceEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolPreEndOfGameV1CurrentSequenceEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}