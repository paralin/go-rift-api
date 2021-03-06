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

// NewDeleteLolChatV1BlockedPlayersByIDParams creates a new DeleteLolChatV1BlockedPlayersByIDParams object
// with the default values initialized.
func NewDeleteLolChatV1BlockedPlayersByIDParams() *DeleteLolChatV1BlockedPlayersByIDParams {
	var ()
	return &DeleteLolChatV1BlockedPlayersByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLolChatV1BlockedPlayersByIDParamsWithTimeout creates a new DeleteLolChatV1BlockedPlayersByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLolChatV1BlockedPlayersByIDParamsWithTimeout(timeout time.Duration) *DeleteLolChatV1BlockedPlayersByIDParams {
	var ()
	return &DeleteLolChatV1BlockedPlayersByIDParams{

		timeout: timeout,
	}
}

// NewDeleteLolChatV1BlockedPlayersByIDParamsWithContext creates a new DeleteLolChatV1BlockedPlayersByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLolChatV1BlockedPlayersByIDParamsWithContext(ctx context.Context) *DeleteLolChatV1BlockedPlayersByIDParams {
	var ()
	return &DeleteLolChatV1BlockedPlayersByIDParams{

		Context: ctx,
	}
}

// NewDeleteLolChatV1BlockedPlayersByIDParamsWithHTTPClient creates a new DeleteLolChatV1BlockedPlayersByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLolChatV1BlockedPlayersByIDParamsWithHTTPClient(client *http.Client) *DeleteLolChatV1BlockedPlayersByIDParams {
	var ()
	return &DeleteLolChatV1BlockedPlayersByIDParams{
		HTTPClient: client,
	}
}

/*DeleteLolChatV1BlockedPlayersByIDParams contains all the parameters to send to the API endpoint
for the delete lol chat v1 blocked players by Id operation typically these are written to a http.Request
*/
type DeleteLolChatV1BlockedPlayersByIDParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete lol chat v1 blocked players by Id params
func (o *DeleteLolChatV1BlockedPlayersByIDParams) WithTimeout(timeout time.Duration) *DeleteLolChatV1BlockedPlayersByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete lol chat v1 blocked players by Id params
func (o *DeleteLolChatV1BlockedPlayersByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete lol chat v1 blocked players by Id params
func (o *DeleteLolChatV1BlockedPlayersByIDParams) WithContext(ctx context.Context) *DeleteLolChatV1BlockedPlayersByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete lol chat v1 blocked players by Id params
func (o *DeleteLolChatV1BlockedPlayersByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete lol chat v1 blocked players by Id params
func (o *DeleteLolChatV1BlockedPlayersByIDParams) WithHTTPClient(client *http.Client) *DeleteLolChatV1BlockedPlayersByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete lol chat v1 blocked players by Id params
func (o *DeleteLolChatV1BlockedPlayersByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete lol chat v1 blocked players by Id params
func (o *DeleteLolChatV1BlockedPlayersByIDParams) WithID(id string) *DeleteLolChatV1BlockedPlayersByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete lol chat v1 blocked players by Id params
func (o *DeleteLolChatV1BlockedPlayersByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLolChatV1BlockedPlayersByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
