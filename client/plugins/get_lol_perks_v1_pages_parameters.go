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

// NewGetLolPerksV1PagesParams creates a new GetLolPerksV1PagesParams object
// with the default values initialized.
func NewGetLolPerksV1PagesParams() *GetLolPerksV1PagesParams {

	return &GetLolPerksV1PagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolPerksV1PagesParamsWithTimeout creates a new GetLolPerksV1PagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolPerksV1PagesParamsWithTimeout(timeout time.Duration) *GetLolPerksV1PagesParams {

	return &GetLolPerksV1PagesParams{

		timeout: timeout,
	}
}

// NewGetLolPerksV1PagesParamsWithContext creates a new GetLolPerksV1PagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolPerksV1PagesParamsWithContext(ctx context.Context) *GetLolPerksV1PagesParams {

	return &GetLolPerksV1PagesParams{

		Context: ctx,
	}
}

// NewGetLolPerksV1PagesParamsWithHTTPClient creates a new GetLolPerksV1PagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolPerksV1PagesParamsWithHTTPClient(client *http.Client) *GetLolPerksV1PagesParams {

	return &GetLolPerksV1PagesParams{
		HTTPClient: client,
	}
}

/*GetLolPerksV1PagesParams contains all the parameters to send to the API endpoint
for the get lol perks v1 pages operation typically these are written to a http.Request
*/
type GetLolPerksV1PagesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol perks v1 pages params
func (o *GetLolPerksV1PagesParams) WithTimeout(timeout time.Duration) *GetLolPerksV1PagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol perks v1 pages params
func (o *GetLolPerksV1PagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol perks v1 pages params
func (o *GetLolPerksV1PagesParams) WithContext(ctx context.Context) *GetLolPerksV1PagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol perks v1 pages params
func (o *GetLolPerksV1PagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol perks v1 pages params
func (o *GetLolPerksV1PagesParams) WithHTTPClient(client *http.Client) *GetLolPerksV1PagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol perks v1 pages params
func (o *GetLolPerksV1PagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolPerksV1PagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
