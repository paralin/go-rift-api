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

// NewGetLolChatV1FriendsByIDParams creates a new GetLolChatV1FriendsByIDParams object
// with the default values initialized.
func NewGetLolChatV1FriendsByIDParams() *GetLolChatV1FriendsByIDParams {
	var ()
	return &GetLolChatV1FriendsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolChatV1FriendsByIDParamsWithTimeout creates a new GetLolChatV1FriendsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolChatV1FriendsByIDParamsWithTimeout(timeout time.Duration) *GetLolChatV1FriendsByIDParams {
	var ()
	return &GetLolChatV1FriendsByIDParams{

		timeout: timeout,
	}
}

// NewGetLolChatV1FriendsByIDParamsWithContext creates a new GetLolChatV1FriendsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolChatV1FriendsByIDParamsWithContext(ctx context.Context) *GetLolChatV1FriendsByIDParams {
	var ()
	return &GetLolChatV1FriendsByIDParams{

		Context: ctx,
	}
}

// NewGetLolChatV1FriendsByIDParamsWithHTTPClient creates a new GetLolChatV1FriendsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolChatV1FriendsByIDParamsWithHTTPClient(client *http.Client) *GetLolChatV1FriendsByIDParams {
	var ()
	return &GetLolChatV1FriendsByIDParams{
		HTTPClient: client,
	}
}

/*GetLolChatV1FriendsByIDParams contains all the parameters to send to the API endpoint
for the get lol chat v1 friends by Id operation typically these are written to a http.Request
*/
type GetLolChatV1FriendsByIDParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol chat v1 friends by Id params
func (o *GetLolChatV1FriendsByIDParams) WithTimeout(timeout time.Duration) *GetLolChatV1FriendsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol chat v1 friends by Id params
func (o *GetLolChatV1FriendsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol chat v1 friends by Id params
func (o *GetLolChatV1FriendsByIDParams) WithContext(ctx context.Context) *GetLolChatV1FriendsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol chat v1 friends by Id params
func (o *GetLolChatV1FriendsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol chat v1 friends by Id params
func (o *GetLolChatV1FriendsByIDParams) WithHTTPClient(client *http.Client) *GetLolChatV1FriendsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol chat v1 friends by Id params
func (o *GetLolChatV1FriendsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get lol chat v1 friends by Id params
func (o *GetLolChatV1FriendsByIDParams) WithID(id string) *GetLolChatV1FriendsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get lol chat v1 friends by Id params
func (o *GetLolChatV1FriendsByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolChatV1FriendsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
