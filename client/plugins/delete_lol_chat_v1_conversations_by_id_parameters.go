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

// NewDeleteLolChatV1ConversationsByIDParams creates a new DeleteLolChatV1ConversationsByIDParams object
// with the default values initialized.
func NewDeleteLolChatV1ConversationsByIDParams() *DeleteLolChatV1ConversationsByIDParams {
	var ()
	return &DeleteLolChatV1ConversationsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLolChatV1ConversationsByIDParamsWithTimeout creates a new DeleteLolChatV1ConversationsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLolChatV1ConversationsByIDParamsWithTimeout(timeout time.Duration) *DeleteLolChatV1ConversationsByIDParams {
	var ()
	return &DeleteLolChatV1ConversationsByIDParams{

		timeout: timeout,
	}
}

// NewDeleteLolChatV1ConversationsByIDParamsWithContext creates a new DeleteLolChatV1ConversationsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLolChatV1ConversationsByIDParamsWithContext(ctx context.Context) *DeleteLolChatV1ConversationsByIDParams {
	var ()
	return &DeleteLolChatV1ConversationsByIDParams{

		Context: ctx,
	}
}

// NewDeleteLolChatV1ConversationsByIDParamsWithHTTPClient creates a new DeleteLolChatV1ConversationsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLolChatV1ConversationsByIDParamsWithHTTPClient(client *http.Client) *DeleteLolChatV1ConversationsByIDParams {
	var ()
	return &DeleteLolChatV1ConversationsByIDParams{
		HTTPClient: client,
	}
}

/*DeleteLolChatV1ConversationsByIDParams contains all the parameters to send to the API endpoint
for the delete lol chat v1 conversations by Id operation typically these are written to a http.Request
*/
type DeleteLolChatV1ConversationsByIDParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete lol chat v1 conversations by Id params
func (o *DeleteLolChatV1ConversationsByIDParams) WithTimeout(timeout time.Duration) *DeleteLolChatV1ConversationsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete lol chat v1 conversations by Id params
func (o *DeleteLolChatV1ConversationsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete lol chat v1 conversations by Id params
func (o *DeleteLolChatV1ConversationsByIDParams) WithContext(ctx context.Context) *DeleteLolChatV1ConversationsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete lol chat v1 conversations by Id params
func (o *DeleteLolChatV1ConversationsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete lol chat v1 conversations by Id params
func (o *DeleteLolChatV1ConversationsByIDParams) WithHTTPClient(client *http.Client) *DeleteLolChatV1ConversationsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete lol chat v1 conversations by Id params
func (o *DeleteLolChatV1ConversationsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete lol chat v1 conversations by Id params
func (o *DeleteLolChatV1ConversationsByIDParams) WithID(id string) *DeleteLolChatV1ConversationsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete lol chat v1 conversations by Id params
func (o *DeleteLolChatV1ConversationsByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLolChatV1ConversationsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
