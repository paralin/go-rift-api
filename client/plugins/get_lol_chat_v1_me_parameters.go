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

// NewGetLolChatV1MeParams creates a new GetLolChatV1MeParams object
// with the default values initialized.
func NewGetLolChatV1MeParams() *GetLolChatV1MeParams {

	return &GetLolChatV1MeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolChatV1MeParamsWithTimeout creates a new GetLolChatV1MeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolChatV1MeParamsWithTimeout(timeout time.Duration) *GetLolChatV1MeParams {

	return &GetLolChatV1MeParams{

		timeout: timeout,
	}
}

// NewGetLolChatV1MeParamsWithContext creates a new GetLolChatV1MeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolChatV1MeParamsWithContext(ctx context.Context) *GetLolChatV1MeParams {

	return &GetLolChatV1MeParams{

		Context: ctx,
	}
}

// NewGetLolChatV1MeParamsWithHTTPClient creates a new GetLolChatV1MeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolChatV1MeParamsWithHTTPClient(client *http.Client) *GetLolChatV1MeParams {

	return &GetLolChatV1MeParams{
		HTTPClient: client,
	}
}

/*GetLolChatV1MeParams contains all the parameters to send to the API endpoint
for the get lol chat v1 me operation typically these are written to a http.Request
*/
type GetLolChatV1MeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol chat v1 me params
func (o *GetLolChatV1MeParams) WithTimeout(timeout time.Duration) *GetLolChatV1MeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol chat v1 me params
func (o *GetLolChatV1MeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol chat v1 me params
func (o *GetLolChatV1MeParams) WithContext(ctx context.Context) *GetLolChatV1MeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol chat v1 me params
func (o *GetLolChatV1MeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol chat v1 me params
func (o *GetLolChatV1MeParams) WithHTTPClient(client *http.Client) *GetLolChatV1MeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol chat v1 me params
func (o *GetLolChatV1MeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolChatV1MeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}