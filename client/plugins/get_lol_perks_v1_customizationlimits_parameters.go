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

// NewGetLolPerksV1CustomizationlimitsParams creates a new GetLolPerksV1CustomizationlimitsParams object
// with the default values initialized.
func NewGetLolPerksV1CustomizationlimitsParams() *GetLolPerksV1CustomizationlimitsParams {

	return &GetLolPerksV1CustomizationlimitsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolPerksV1CustomizationlimitsParamsWithTimeout creates a new GetLolPerksV1CustomizationlimitsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolPerksV1CustomizationlimitsParamsWithTimeout(timeout time.Duration) *GetLolPerksV1CustomizationlimitsParams {

	return &GetLolPerksV1CustomizationlimitsParams{

		timeout: timeout,
	}
}

// NewGetLolPerksV1CustomizationlimitsParamsWithContext creates a new GetLolPerksV1CustomizationlimitsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolPerksV1CustomizationlimitsParamsWithContext(ctx context.Context) *GetLolPerksV1CustomizationlimitsParams {

	return &GetLolPerksV1CustomizationlimitsParams{

		Context: ctx,
	}
}

// NewGetLolPerksV1CustomizationlimitsParamsWithHTTPClient creates a new GetLolPerksV1CustomizationlimitsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolPerksV1CustomizationlimitsParamsWithHTTPClient(client *http.Client) *GetLolPerksV1CustomizationlimitsParams {

	return &GetLolPerksV1CustomizationlimitsParams{
		HTTPClient: client,
	}
}

/*GetLolPerksV1CustomizationlimitsParams contains all the parameters to send to the API endpoint
for the get lol perks v1 customizationlimits operation typically these are written to a http.Request
*/
type GetLolPerksV1CustomizationlimitsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol perks v1 customizationlimits params
func (o *GetLolPerksV1CustomizationlimitsParams) WithTimeout(timeout time.Duration) *GetLolPerksV1CustomizationlimitsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol perks v1 customizationlimits params
func (o *GetLolPerksV1CustomizationlimitsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol perks v1 customizationlimits params
func (o *GetLolPerksV1CustomizationlimitsParams) WithContext(ctx context.Context) *GetLolPerksV1CustomizationlimitsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol perks v1 customizationlimits params
func (o *GetLolPerksV1CustomizationlimitsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol perks v1 customizationlimits params
func (o *GetLolPerksV1CustomizationlimitsParams) WithHTTPClient(client *http.Client) *GetLolPerksV1CustomizationlimitsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol perks v1 customizationlimits params
func (o *GetLolPerksV1CustomizationlimitsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolPerksV1CustomizationlimitsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
