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

// NewDeleteLolChatV1SettingsByKeyParams creates a new DeleteLolChatV1SettingsByKeyParams object
// with the default values initialized.
func NewDeleteLolChatV1SettingsByKeyParams() *DeleteLolChatV1SettingsByKeyParams {
	var ()
	return &DeleteLolChatV1SettingsByKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLolChatV1SettingsByKeyParamsWithTimeout creates a new DeleteLolChatV1SettingsByKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLolChatV1SettingsByKeyParamsWithTimeout(timeout time.Duration) *DeleteLolChatV1SettingsByKeyParams {
	var ()
	return &DeleteLolChatV1SettingsByKeyParams{

		timeout: timeout,
	}
}

// NewDeleteLolChatV1SettingsByKeyParamsWithContext creates a new DeleteLolChatV1SettingsByKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLolChatV1SettingsByKeyParamsWithContext(ctx context.Context) *DeleteLolChatV1SettingsByKeyParams {
	var ()
	return &DeleteLolChatV1SettingsByKeyParams{

		Context: ctx,
	}
}

// NewDeleteLolChatV1SettingsByKeyParamsWithHTTPClient creates a new DeleteLolChatV1SettingsByKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLolChatV1SettingsByKeyParamsWithHTTPClient(client *http.Client) *DeleteLolChatV1SettingsByKeyParams {
	var ()
	return &DeleteLolChatV1SettingsByKeyParams{
		HTTPClient: client,
	}
}

/*DeleteLolChatV1SettingsByKeyParams contains all the parameters to send to the API endpoint
for the delete lol chat v1 settings by key operation typically these are written to a http.Request
*/
type DeleteLolChatV1SettingsByKeyParams struct {

	/*DoAsync*/
	DoAsync *bool
	/*Key*/
	Key string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete lol chat v1 settings by key params
func (o *DeleteLolChatV1SettingsByKeyParams) WithTimeout(timeout time.Duration) *DeleteLolChatV1SettingsByKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete lol chat v1 settings by key params
func (o *DeleteLolChatV1SettingsByKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete lol chat v1 settings by key params
func (o *DeleteLolChatV1SettingsByKeyParams) WithContext(ctx context.Context) *DeleteLolChatV1SettingsByKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete lol chat v1 settings by key params
func (o *DeleteLolChatV1SettingsByKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete lol chat v1 settings by key params
func (o *DeleteLolChatV1SettingsByKeyParams) WithHTTPClient(client *http.Client) *DeleteLolChatV1SettingsByKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete lol chat v1 settings by key params
func (o *DeleteLolChatV1SettingsByKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDoAsync adds the doAsync to the delete lol chat v1 settings by key params
func (o *DeleteLolChatV1SettingsByKeyParams) WithDoAsync(doAsync *bool) *DeleteLolChatV1SettingsByKeyParams {
	o.SetDoAsync(doAsync)
	return o
}

// SetDoAsync adds the doAsync to the delete lol chat v1 settings by key params
func (o *DeleteLolChatV1SettingsByKeyParams) SetDoAsync(doAsync *bool) {
	o.DoAsync = doAsync
}

// WithKey adds the key to the delete lol chat v1 settings by key params
func (o *DeleteLolChatV1SettingsByKeyParams) WithKey(key string) *DeleteLolChatV1SettingsByKeyParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the delete lol chat v1 settings by key params
func (o *DeleteLolChatV1SettingsByKeyParams) SetKey(key string) {
	o.Key = key
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLolChatV1SettingsByKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}