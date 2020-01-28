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

// NewPutLolChatV1SettingsParams creates a new PutLolChatV1SettingsParams object
// with the default values initialized.
func NewPutLolChatV1SettingsParams() *PutLolChatV1SettingsParams {
	var ()
	return &PutLolChatV1SettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLolChatV1SettingsParamsWithTimeout creates a new PutLolChatV1SettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLolChatV1SettingsParamsWithTimeout(timeout time.Duration) *PutLolChatV1SettingsParams {
	var ()
	return &PutLolChatV1SettingsParams{

		timeout: timeout,
	}
}

// NewPutLolChatV1SettingsParamsWithContext creates a new PutLolChatV1SettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLolChatV1SettingsParamsWithContext(ctx context.Context) *PutLolChatV1SettingsParams {
	var ()
	return &PutLolChatV1SettingsParams{

		Context: ctx,
	}
}

// NewPutLolChatV1SettingsParamsWithHTTPClient creates a new PutLolChatV1SettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLolChatV1SettingsParamsWithHTTPClient(client *http.Client) *PutLolChatV1SettingsParams {
	var ()
	return &PutLolChatV1SettingsParams{
		HTTPClient: client,
	}
}

/*PutLolChatV1SettingsParams contains all the parameters to send to the API endpoint
for the put lol chat v1 settings operation typically these are written to a http.Request
*/
type PutLolChatV1SettingsParams struct {

	/*Data*/
	Data interface{}
	/*DoAsync*/
	DoAsync *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put lol chat v1 settings params
func (o *PutLolChatV1SettingsParams) WithTimeout(timeout time.Duration) *PutLolChatV1SettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put lol chat v1 settings params
func (o *PutLolChatV1SettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put lol chat v1 settings params
func (o *PutLolChatV1SettingsParams) WithContext(ctx context.Context) *PutLolChatV1SettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put lol chat v1 settings params
func (o *PutLolChatV1SettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put lol chat v1 settings params
func (o *PutLolChatV1SettingsParams) WithHTTPClient(client *http.Client) *PutLolChatV1SettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put lol chat v1 settings params
func (o *PutLolChatV1SettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the put lol chat v1 settings params
func (o *PutLolChatV1SettingsParams) WithData(data interface{}) *PutLolChatV1SettingsParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the put lol chat v1 settings params
func (o *PutLolChatV1SettingsParams) SetData(data interface{}) {
	o.Data = data
}

// WithDoAsync adds the doAsync to the put lol chat v1 settings params
func (o *PutLolChatV1SettingsParams) WithDoAsync(doAsync *bool) *PutLolChatV1SettingsParams {
	o.SetDoAsync(doAsync)
	return o
}

// SetDoAsync adds the doAsync to the put lol chat v1 settings params
func (o *PutLolChatV1SettingsParams) SetDoAsync(doAsync *bool) {
	o.DoAsync = doAsync
}

// WriteToRequest writes these params to a swagger request
func (o *PutLolChatV1SettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if o.DoAsync != nil {

		// query param doAsync
		var qrDoAsync bool
		if o.DoAsync != nil {
			qrDoAsync = *o.DoAsync
		}
		qDoAsync := swag.FormatBool(qrDoAsync)
		if qDoAsync != "" {
			if err := r.SetQueryParam("doAsync", qDoAsync); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}