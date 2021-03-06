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

// NewGetLolSettingsV2AccountByPpTypeByCategoryParams creates a new GetLolSettingsV2AccountByPpTypeByCategoryParams object
// with the default values initialized.
func NewGetLolSettingsV2AccountByPpTypeByCategoryParams() *GetLolSettingsV2AccountByPpTypeByCategoryParams {
	var ()
	return &GetLolSettingsV2AccountByPpTypeByCategoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolSettingsV2AccountByPpTypeByCategoryParamsWithTimeout creates a new GetLolSettingsV2AccountByPpTypeByCategoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolSettingsV2AccountByPpTypeByCategoryParamsWithTimeout(timeout time.Duration) *GetLolSettingsV2AccountByPpTypeByCategoryParams {
	var ()
	return &GetLolSettingsV2AccountByPpTypeByCategoryParams{

		timeout: timeout,
	}
}

// NewGetLolSettingsV2AccountByPpTypeByCategoryParamsWithContext creates a new GetLolSettingsV2AccountByPpTypeByCategoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolSettingsV2AccountByPpTypeByCategoryParamsWithContext(ctx context.Context) *GetLolSettingsV2AccountByPpTypeByCategoryParams {
	var ()
	return &GetLolSettingsV2AccountByPpTypeByCategoryParams{

		Context: ctx,
	}
}

// NewGetLolSettingsV2AccountByPpTypeByCategoryParamsWithHTTPClient creates a new GetLolSettingsV2AccountByPpTypeByCategoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolSettingsV2AccountByPpTypeByCategoryParamsWithHTTPClient(client *http.Client) *GetLolSettingsV2AccountByPpTypeByCategoryParams {
	var ()
	return &GetLolSettingsV2AccountByPpTypeByCategoryParams{
		HTTPClient: client,
	}
}

/*GetLolSettingsV2AccountByPpTypeByCategoryParams contains all the parameters to send to the API endpoint
for the get lol settings v2 account by pp type by category operation typically these are written to a http.Request
*/
type GetLolSettingsV2AccountByPpTypeByCategoryParams struct {

	/*Category*/
	Category string
	/*PpType*/
	PpType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol settings v2 account by pp type by category params
func (o *GetLolSettingsV2AccountByPpTypeByCategoryParams) WithTimeout(timeout time.Duration) *GetLolSettingsV2AccountByPpTypeByCategoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol settings v2 account by pp type by category params
func (o *GetLolSettingsV2AccountByPpTypeByCategoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol settings v2 account by pp type by category params
func (o *GetLolSettingsV2AccountByPpTypeByCategoryParams) WithContext(ctx context.Context) *GetLolSettingsV2AccountByPpTypeByCategoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol settings v2 account by pp type by category params
func (o *GetLolSettingsV2AccountByPpTypeByCategoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol settings v2 account by pp type by category params
func (o *GetLolSettingsV2AccountByPpTypeByCategoryParams) WithHTTPClient(client *http.Client) *GetLolSettingsV2AccountByPpTypeByCategoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol settings v2 account by pp type by category params
func (o *GetLolSettingsV2AccountByPpTypeByCategoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategory adds the category to the get lol settings v2 account by pp type by category params
func (o *GetLolSettingsV2AccountByPpTypeByCategoryParams) WithCategory(category string) *GetLolSettingsV2AccountByPpTypeByCategoryParams {
	o.SetCategory(category)
	return o
}

// SetCategory adds the category to the get lol settings v2 account by pp type by category params
func (o *GetLolSettingsV2AccountByPpTypeByCategoryParams) SetCategory(category string) {
	o.Category = category
}

// WithPpType adds the ppType to the get lol settings v2 account by pp type by category params
func (o *GetLolSettingsV2AccountByPpTypeByCategoryParams) WithPpType(ppType string) *GetLolSettingsV2AccountByPpTypeByCategoryParams {
	o.SetPpType(ppType)
	return o
}

// SetPpType adds the ppType to the get lol settings v2 account by pp type by category params
func (o *GetLolSettingsV2AccountByPpTypeByCategoryParams) SetPpType(ppType string) {
	o.PpType = ppType
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolSettingsV2AccountByPpTypeByCategoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param category
	if err := r.SetPathParam("category", o.Category); err != nil {
		return err
	}

	// path param ppType
	if err := r.SetPathParam("ppType", o.PpType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
