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

// NewGetLolMapsV2MapsParams creates a new GetLolMapsV2MapsParams object
// with the default values initialized.
func NewGetLolMapsV2MapsParams() *GetLolMapsV2MapsParams {

	return &GetLolMapsV2MapsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolMapsV2MapsParamsWithTimeout creates a new GetLolMapsV2MapsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolMapsV2MapsParamsWithTimeout(timeout time.Duration) *GetLolMapsV2MapsParams {

	return &GetLolMapsV2MapsParams{

		timeout: timeout,
	}
}

// NewGetLolMapsV2MapsParamsWithContext creates a new GetLolMapsV2MapsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolMapsV2MapsParamsWithContext(ctx context.Context) *GetLolMapsV2MapsParams {

	return &GetLolMapsV2MapsParams{

		Context: ctx,
	}
}

// NewGetLolMapsV2MapsParamsWithHTTPClient creates a new GetLolMapsV2MapsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolMapsV2MapsParamsWithHTTPClient(client *http.Client) *GetLolMapsV2MapsParams {

	return &GetLolMapsV2MapsParams{
		HTTPClient: client,
	}
}

/*GetLolMapsV2MapsParams contains all the parameters to send to the API endpoint
for the get lol maps v2 maps operation typically these are written to a http.Request
*/
type GetLolMapsV2MapsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol maps v2 maps params
func (o *GetLolMapsV2MapsParams) WithTimeout(timeout time.Duration) *GetLolMapsV2MapsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol maps v2 maps params
func (o *GetLolMapsV2MapsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol maps v2 maps params
func (o *GetLolMapsV2MapsParams) WithContext(ctx context.Context) *GetLolMapsV2MapsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol maps v2 maps params
func (o *GetLolMapsV2MapsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol maps v2 maps params
func (o *GetLolMapsV2MapsParams) WithHTTPClient(client *http.Client) *GetLolMapsV2MapsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol maps v2 maps params
func (o *GetLolMapsV2MapsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolMapsV2MapsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
