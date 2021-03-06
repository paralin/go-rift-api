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

// NewGetLolChatV1SettingsByKeyParams creates a new GetLolChatV1SettingsByKeyParams object
// with the default values initialized.
func NewGetLolChatV1SettingsByKeyParams() *GetLolChatV1SettingsByKeyParams {
	var ()
	return &GetLolChatV1SettingsByKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolChatV1SettingsByKeyParamsWithTimeout creates a new GetLolChatV1SettingsByKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolChatV1SettingsByKeyParamsWithTimeout(timeout time.Duration) *GetLolChatV1SettingsByKeyParams {
	var ()
	return &GetLolChatV1SettingsByKeyParams{

		timeout: timeout,
	}
}

// NewGetLolChatV1SettingsByKeyParamsWithContext creates a new GetLolChatV1SettingsByKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolChatV1SettingsByKeyParamsWithContext(ctx context.Context) *GetLolChatV1SettingsByKeyParams {
	var ()
	return &GetLolChatV1SettingsByKeyParams{

		Context: ctx,
	}
}

// NewGetLolChatV1SettingsByKeyParamsWithHTTPClient creates a new GetLolChatV1SettingsByKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolChatV1SettingsByKeyParamsWithHTTPClient(client *http.Client) *GetLolChatV1SettingsByKeyParams {
	var ()
	return &GetLolChatV1SettingsByKeyParams{
		HTTPClient: client,
	}
}

/*GetLolChatV1SettingsByKeyParams contains all the parameters to send to the API endpoint
for the get lol chat v1 settings by key operation typically these are written to a http.Request
*/
type GetLolChatV1SettingsByKeyParams struct {

	/*Key*/
	Key string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol chat v1 settings by key params
func (o *GetLolChatV1SettingsByKeyParams) WithTimeout(timeout time.Duration) *GetLolChatV1SettingsByKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol chat v1 settings by key params
func (o *GetLolChatV1SettingsByKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol chat v1 settings by key params
func (o *GetLolChatV1SettingsByKeyParams) WithContext(ctx context.Context) *GetLolChatV1SettingsByKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol chat v1 settings by key params
func (o *GetLolChatV1SettingsByKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol chat v1 settings by key params
func (o *GetLolChatV1SettingsByKeyParams) WithHTTPClient(client *http.Client) *GetLolChatV1SettingsByKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol chat v1 settings by key params
func (o *GetLolChatV1SettingsByKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the get lol chat v1 settings by key params
func (o *GetLolChatV1SettingsByKeyParams) WithKey(key string) *GetLolChatV1SettingsByKeyParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the get lol chat v1 settings by key params
func (o *GetLolChatV1SettingsByKeyParams) SetKey(key string) {
	o.Key = key
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolChatV1SettingsByKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
