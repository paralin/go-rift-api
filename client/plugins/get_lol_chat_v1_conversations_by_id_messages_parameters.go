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

// NewGetLolChatV1ConversationsByIDMessagesParams creates a new GetLolChatV1ConversationsByIDMessagesParams object
// with the default values initialized.
func NewGetLolChatV1ConversationsByIDMessagesParams() *GetLolChatV1ConversationsByIDMessagesParams {
	var ()
	return &GetLolChatV1ConversationsByIDMessagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolChatV1ConversationsByIDMessagesParamsWithTimeout creates a new GetLolChatV1ConversationsByIDMessagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolChatV1ConversationsByIDMessagesParamsWithTimeout(timeout time.Duration) *GetLolChatV1ConversationsByIDMessagesParams {
	var ()
	return &GetLolChatV1ConversationsByIDMessagesParams{

		timeout: timeout,
	}
}

// NewGetLolChatV1ConversationsByIDMessagesParamsWithContext creates a new GetLolChatV1ConversationsByIDMessagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolChatV1ConversationsByIDMessagesParamsWithContext(ctx context.Context) *GetLolChatV1ConversationsByIDMessagesParams {
	var ()
	return &GetLolChatV1ConversationsByIDMessagesParams{

		Context: ctx,
	}
}

// NewGetLolChatV1ConversationsByIDMessagesParamsWithHTTPClient creates a new GetLolChatV1ConversationsByIDMessagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolChatV1ConversationsByIDMessagesParamsWithHTTPClient(client *http.Client) *GetLolChatV1ConversationsByIDMessagesParams {
	var ()
	return &GetLolChatV1ConversationsByIDMessagesParams{
		HTTPClient: client,
	}
}

/*GetLolChatV1ConversationsByIDMessagesParams contains all the parameters to send to the API endpoint
for the get lol chat v1 conversations by Id messages operation typically these are written to a http.Request
*/
type GetLolChatV1ConversationsByIDMessagesParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol chat v1 conversations by Id messages params
func (o *GetLolChatV1ConversationsByIDMessagesParams) WithTimeout(timeout time.Duration) *GetLolChatV1ConversationsByIDMessagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol chat v1 conversations by Id messages params
func (o *GetLolChatV1ConversationsByIDMessagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol chat v1 conversations by Id messages params
func (o *GetLolChatV1ConversationsByIDMessagesParams) WithContext(ctx context.Context) *GetLolChatV1ConversationsByIDMessagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol chat v1 conversations by Id messages params
func (o *GetLolChatV1ConversationsByIDMessagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol chat v1 conversations by Id messages params
func (o *GetLolChatV1ConversationsByIDMessagesParams) WithHTTPClient(client *http.Client) *GetLolChatV1ConversationsByIDMessagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol chat v1 conversations by Id messages params
func (o *GetLolChatV1ConversationsByIDMessagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get lol chat v1 conversations by Id messages params
func (o *GetLolChatV1ConversationsByIDMessagesParams) WithID(id string) *GetLolChatV1ConversationsByIDMessagesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get lol chat v1 conversations by Id messages params
func (o *GetLolChatV1ConversationsByIDMessagesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolChatV1ConversationsByIDMessagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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