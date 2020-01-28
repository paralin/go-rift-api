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

// NewGetLolChatV1FriendsParams creates a new GetLolChatV1FriendsParams object
// with the default values initialized.
func NewGetLolChatV1FriendsParams() *GetLolChatV1FriendsParams {

	return &GetLolChatV1FriendsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolChatV1FriendsParamsWithTimeout creates a new GetLolChatV1FriendsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolChatV1FriendsParamsWithTimeout(timeout time.Duration) *GetLolChatV1FriendsParams {

	return &GetLolChatV1FriendsParams{

		timeout: timeout,
	}
}

// NewGetLolChatV1FriendsParamsWithContext creates a new GetLolChatV1FriendsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolChatV1FriendsParamsWithContext(ctx context.Context) *GetLolChatV1FriendsParams {

	return &GetLolChatV1FriendsParams{

		Context: ctx,
	}
}

// NewGetLolChatV1FriendsParamsWithHTTPClient creates a new GetLolChatV1FriendsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolChatV1FriendsParamsWithHTTPClient(client *http.Client) *GetLolChatV1FriendsParams {

	return &GetLolChatV1FriendsParams{
		HTTPClient: client,
	}
}

/*GetLolChatV1FriendsParams contains all the parameters to send to the API endpoint
for the get lol chat v1 friends operation typically these are written to a http.Request
*/
type GetLolChatV1FriendsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol chat v1 friends params
func (o *GetLolChatV1FriendsParams) WithTimeout(timeout time.Duration) *GetLolChatV1FriendsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol chat v1 friends params
func (o *GetLolChatV1FriendsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol chat v1 friends params
func (o *GetLolChatV1FriendsParams) WithContext(ctx context.Context) *GetLolChatV1FriendsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol chat v1 friends params
func (o *GetLolChatV1FriendsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol chat v1 friends params
func (o *GetLolChatV1FriendsParams) WithHTTPClient(client *http.Client) *GetLolChatV1FriendsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol chat v1 friends params
func (o *GetLolChatV1FriendsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolChatV1FriendsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
