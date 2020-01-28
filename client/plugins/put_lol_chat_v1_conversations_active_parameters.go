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

	models "github.com/paralin/go-rift-api/models"
)

// NewPutLolChatV1ConversationsActiveParams creates a new PutLolChatV1ConversationsActiveParams object
// with the default values initialized.
func NewPutLolChatV1ConversationsActiveParams() *PutLolChatV1ConversationsActiveParams {
	var ()
	return &PutLolChatV1ConversationsActiveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLolChatV1ConversationsActiveParamsWithTimeout creates a new PutLolChatV1ConversationsActiveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLolChatV1ConversationsActiveParamsWithTimeout(timeout time.Duration) *PutLolChatV1ConversationsActiveParams {
	var ()
	return &PutLolChatV1ConversationsActiveParams{

		timeout: timeout,
	}
}

// NewPutLolChatV1ConversationsActiveParamsWithContext creates a new PutLolChatV1ConversationsActiveParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLolChatV1ConversationsActiveParamsWithContext(ctx context.Context) *PutLolChatV1ConversationsActiveParams {
	var ()
	return &PutLolChatV1ConversationsActiveParams{

		Context: ctx,
	}
}

// NewPutLolChatV1ConversationsActiveParamsWithHTTPClient creates a new PutLolChatV1ConversationsActiveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLolChatV1ConversationsActiveParamsWithHTTPClient(client *http.Client) *PutLolChatV1ConversationsActiveParams {
	var ()
	return &PutLolChatV1ConversationsActiveParams{
		HTTPClient: client,
	}
}

/*PutLolChatV1ConversationsActiveParams contains all the parameters to send to the API endpoint
for the put lol chat v1 conversations active operation typically these are written to a http.Request
*/
type PutLolChatV1ConversationsActiveParams struct {

	/*ActiveConversation*/
	ActiveConversation *models.LolChatActiveConversationResource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put lol chat v1 conversations active params
func (o *PutLolChatV1ConversationsActiveParams) WithTimeout(timeout time.Duration) *PutLolChatV1ConversationsActiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put lol chat v1 conversations active params
func (o *PutLolChatV1ConversationsActiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put lol chat v1 conversations active params
func (o *PutLolChatV1ConversationsActiveParams) WithContext(ctx context.Context) *PutLolChatV1ConversationsActiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put lol chat v1 conversations active params
func (o *PutLolChatV1ConversationsActiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put lol chat v1 conversations active params
func (o *PutLolChatV1ConversationsActiveParams) WithHTTPClient(client *http.Client) *PutLolChatV1ConversationsActiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put lol chat v1 conversations active params
func (o *PutLolChatV1ConversationsActiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActiveConversation adds the activeConversation to the put lol chat v1 conversations active params
func (o *PutLolChatV1ConversationsActiveParams) WithActiveConversation(activeConversation *models.LolChatActiveConversationResource) *PutLolChatV1ConversationsActiveParams {
	o.SetActiveConversation(activeConversation)
	return o
}

// SetActiveConversation adds the activeConversation to the put lol chat v1 conversations active params
func (o *PutLolChatV1ConversationsActiveParams) SetActiveConversation(activeConversation *models.LolChatActiveConversationResource) {
	o.ActiveConversation = activeConversation
}

// WriteToRequest writes these params to a swagger request
func (o *PutLolChatV1ConversationsActiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActiveConversation != nil {
		if err := r.SetBodyParam(o.ActiveConversation); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
