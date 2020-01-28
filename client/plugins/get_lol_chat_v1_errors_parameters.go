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

// NewGetLolChatV1ErrorsParams creates a new GetLolChatV1ErrorsParams object
// with the default values initialized.
func NewGetLolChatV1ErrorsParams() *GetLolChatV1ErrorsParams {

	return &GetLolChatV1ErrorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolChatV1ErrorsParamsWithTimeout creates a new GetLolChatV1ErrorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolChatV1ErrorsParamsWithTimeout(timeout time.Duration) *GetLolChatV1ErrorsParams {

	return &GetLolChatV1ErrorsParams{

		timeout: timeout,
	}
}

// NewGetLolChatV1ErrorsParamsWithContext creates a new GetLolChatV1ErrorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolChatV1ErrorsParamsWithContext(ctx context.Context) *GetLolChatV1ErrorsParams {

	return &GetLolChatV1ErrorsParams{

		Context: ctx,
	}
}

// NewGetLolChatV1ErrorsParamsWithHTTPClient creates a new GetLolChatV1ErrorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolChatV1ErrorsParamsWithHTTPClient(client *http.Client) *GetLolChatV1ErrorsParams {

	return &GetLolChatV1ErrorsParams{
		HTTPClient: client,
	}
}

/*GetLolChatV1ErrorsParams contains all the parameters to send to the API endpoint
for the get lol chat v1 errors operation typically these are written to a http.Request
*/
type GetLolChatV1ErrorsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol chat v1 errors params
func (o *GetLolChatV1ErrorsParams) WithTimeout(timeout time.Duration) *GetLolChatV1ErrorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol chat v1 errors params
func (o *GetLolChatV1ErrorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol chat v1 errors params
func (o *GetLolChatV1ErrorsParams) WithContext(ctx context.Context) *GetLolChatV1ErrorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol chat v1 errors params
func (o *GetLolChatV1ErrorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol chat v1 errors params
func (o *GetLolChatV1ErrorsParams) WithHTTPClient(client *http.Client) *GetLolChatV1ErrorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol chat v1 errors params
func (o *GetLolChatV1ErrorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolChatV1ErrorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
