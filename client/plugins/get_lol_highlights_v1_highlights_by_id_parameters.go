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

// NewGetLolHighlightsV1HighlightsByIDParams creates a new GetLolHighlightsV1HighlightsByIDParams object
// with the default values initialized.
func NewGetLolHighlightsV1HighlightsByIDParams() *GetLolHighlightsV1HighlightsByIDParams {
	var ()
	return &GetLolHighlightsV1HighlightsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolHighlightsV1HighlightsByIDParamsWithTimeout creates a new GetLolHighlightsV1HighlightsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolHighlightsV1HighlightsByIDParamsWithTimeout(timeout time.Duration) *GetLolHighlightsV1HighlightsByIDParams {
	var ()
	return &GetLolHighlightsV1HighlightsByIDParams{

		timeout: timeout,
	}
}

// NewGetLolHighlightsV1HighlightsByIDParamsWithContext creates a new GetLolHighlightsV1HighlightsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolHighlightsV1HighlightsByIDParamsWithContext(ctx context.Context) *GetLolHighlightsV1HighlightsByIDParams {
	var ()
	return &GetLolHighlightsV1HighlightsByIDParams{

		Context: ctx,
	}
}

// NewGetLolHighlightsV1HighlightsByIDParamsWithHTTPClient creates a new GetLolHighlightsV1HighlightsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolHighlightsV1HighlightsByIDParamsWithHTTPClient(client *http.Client) *GetLolHighlightsV1HighlightsByIDParams {
	var ()
	return &GetLolHighlightsV1HighlightsByIDParams{
		HTTPClient: client,
	}
}

/*GetLolHighlightsV1HighlightsByIDParams contains all the parameters to send to the API endpoint
for the get lol highlights v1 highlights by Id operation typically these are written to a http.Request
*/
type GetLolHighlightsV1HighlightsByIDParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol highlights v1 highlights by Id params
func (o *GetLolHighlightsV1HighlightsByIDParams) WithTimeout(timeout time.Duration) *GetLolHighlightsV1HighlightsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol highlights v1 highlights by Id params
func (o *GetLolHighlightsV1HighlightsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol highlights v1 highlights by Id params
func (o *GetLolHighlightsV1HighlightsByIDParams) WithContext(ctx context.Context) *GetLolHighlightsV1HighlightsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol highlights v1 highlights by Id params
func (o *GetLolHighlightsV1HighlightsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol highlights v1 highlights by Id params
func (o *GetLolHighlightsV1HighlightsByIDParams) WithHTTPClient(client *http.Client) *GetLolHighlightsV1HighlightsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol highlights v1 highlights by Id params
func (o *GetLolHighlightsV1HighlightsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get lol highlights v1 highlights by Id params
func (o *GetLolHighlightsV1HighlightsByIDParams) WithID(id int64) *GetLolHighlightsV1HighlightsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get lol highlights v1 highlights by Id params
func (o *GetLolHighlightsV1HighlightsByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolHighlightsV1HighlightsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
