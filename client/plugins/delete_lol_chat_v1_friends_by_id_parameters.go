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

// NewDeleteLolChatV1FriendsByIDParams creates a new DeleteLolChatV1FriendsByIDParams object
// with the default values initialized.
func NewDeleteLolChatV1FriendsByIDParams() *DeleteLolChatV1FriendsByIDParams {
	var ()
	return &DeleteLolChatV1FriendsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLolChatV1FriendsByIDParamsWithTimeout creates a new DeleteLolChatV1FriendsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLolChatV1FriendsByIDParamsWithTimeout(timeout time.Duration) *DeleteLolChatV1FriendsByIDParams {
	var ()
	return &DeleteLolChatV1FriendsByIDParams{

		timeout: timeout,
	}
}

// NewDeleteLolChatV1FriendsByIDParamsWithContext creates a new DeleteLolChatV1FriendsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLolChatV1FriendsByIDParamsWithContext(ctx context.Context) *DeleteLolChatV1FriendsByIDParams {
	var ()
	return &DeleteLolChatV1FriendsByIDParams{

		Context: ctx,
	}
}

// NewDeleteLolChatV1FriendsByIDParamsWithHTTPClient creates a new DeleteLolChatV1FriendsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLolChatV1FriendsByIDParamsWithHTTPClient(client *http.Client) *DeleteLolChatV1FriendsByIDParams {
	var ()
	return &DeleteLolChatV1FriendsByIDParams{
		HTTPClient: client,
	}
}

/*DeleteLolChatV1FriendsByIDParams contains all the parameters to send to the API endpoint
for the delete lol chat v1 friends by Id operation typically these are written to a http.Request
*/
type DeleteLolChatV1FriendsByIDParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete lol chat v1 friends by Id params
func (o *DeleteLolChatV1FriendsByIDParams) WithTimeout(timeout time.Duration) *DeleteLolChatV1FriendsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete lol chat v1 friends by Id params
func (o *DeleteLolChatV1FriendsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete lol chat v1 friends by Id params
func (o *DeleteLolChatV1FriendsByIDParams) WithContext(ctx context.Context) *DeleteLolChatV1FriendsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete lol chat v1 friends by Id params
func (o *DeleteLolChatV1FriendsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete lol chat v1 friends by Id params
func (o *DeleteLolChatV1FriendsByIDParams) WithHTTPClient(client *http.Client) *DeleteLolChatV1FriendsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete lol chat v1 friends by Id params
func (o *DeleteLolChatV1FriendsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete lol chat v1 friends by Id params
func (o *DeleteLolChatV1FriendsByIDParams) WithID(id string) *DeleteLolChatV1FriendsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete lol chat v1 friends by Id params
func (o *DeleteLolChatV1FriendsByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLolChatV1FriendsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
