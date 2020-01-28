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

// NewGetLolChatV1ResourcesParams creates a new GetLolChatV1ResourcesParams object
// with the default values initialized.
func NewGetLolChatV1ResourcesParams() *GetLolChatV1ResourcesParams {

	return &GetLolChatV1ResourcesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolChatV1ResourcesParamsWithTimeout creates a new GetLolChatV1ResourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolChatV1ResourcesParamsWithTimeout(timeout time.Duration) *GetLolChatV1ResourcesParams {

	return &GetLolChatV1ResourcesParams{

		timeout: timeout,
	}
}

// NewGetLolChatV1ResourcesParamsWithContext creates a new GetLolChatV1ResourcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolChatV1ResourcesParamsWithContext(ctx context.Context) *GetLolChatV1ResourcesParams {

	return &GetLolChatV1ResourcesParams{

		Context: ctx,
	}
}

// NewGetLolChatV1ResourcesParamsWithHTTPClient creates a new GetLolChatV1ResourcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolChatV1ResourcesParamsWithHTTPClient(client *http.Client) *GetLolChatV1ResourcesParams {

	return &GetLolChatV1ResourcesParams{
		HTTPClient: client,
	}
}

/*GetLolChatV1ResourcesParams contains all the parameters to send to the API endpoint
for the get lol chat v1 resources operation typically these are written to a http.Request
*/
type GetLolChatV1ResourcesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol chat v1 resources params
func (o *GetLolChatV1ResourcesParams) WithTimeout(timeout time.Duration) *GetLolChatV1ResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol chat v1 resources params
func (o *GetLolChatV1ResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol chat v1 resources params
func (o *GetLolChatV1ResourcesParams) WithContext(ctx context.Context) *GetLolChatV1ResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol chat v1 resources params
func (o *GetLolChatV1ResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol chat v1 resources params
func (o *GetLolChatV1ResourcesParams) WithHTTPClient(client *http.Client) *GetLolChatV1ResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol chat v1 resources params
func (o *GetLolChatV1ResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolChatV1ResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}