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

// NewPostLolSimpleDialogMessagesV1MessagesParams creates a new PostLolSimpleDialogMessagesV1MessagesParams object
// with the default values initialized.
func NewPostLolSimpleDialogMessagesV1MessagesParams() *PostLolSimpleDialogMessagesV1MessagesParams {
	var ()
	return &PostLolSimpleDialogMessagesV1MessagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolSimpleDialogMessagesV1MessagesParamsWithTimeout creates a new PostLolSimpleDialogMessagesV1MessagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolSimpleDialogMessagesV1MessagesParamsWithTimeout(timeout time.Duration) *PostLolSimpleDialogMessagesV1MessagesParams {
	var ()
	return &PostLolSimpleDialogMessagesV1MessagesParams{

		timeout: timeout,
	}
}

// NewPostLolSimpleDialogMessagesV1MessagesParamsWithContext creates a new PostLolSimpleDialogMessagesV1MessagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolSimpleDialogMessagesV1MessagesParamsWithContext(ctx context.Context) *PostLolSimpleDialogMessagesV1MessagesParams {
	var ()
	return &PostLolSimpleDialogMessagesV1MessagesParams{

		Context: ctx,
	}
}

// NewPostLolSimpleDialogMessagesV1MessagesParamsWithHTTPClient creates a new PostLolSimpleDialogMessagesV1MessagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolSimpleDialogMessagesV1MessagesParamsWithHTTPClient(client *http.Client) *PostLolSimpleDialogMessagesV1MessagesParams {
	var ()
	return &PostLolSimpleDialogMessagesV1MessagesParams{
		HTTPClient: client,
	}
}

/*PostLolSimpleDialogMessagesV1MessagesParams contains all the parameters to send to the API endpoint
for the post lol simple dialog messages v1 messages operation typically these are written to a http.Request
*/
type PostLolSimpleDialogMessagesV1MessagesParams struct {

	/*MessageRequest*/
	MessageRequest *models.LolSimpleDialogMessagesLocalMessageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol simple dialog messages v1 messages params
func (o *PostLolSimpleDialogMessagesV1MessagesParams) WithTimeout(timeout time.Duration) *PostLolSimpleDialogMessagesV1MessagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol simple dialog messages v1 messages params
func (o *PostLolSimpleDialogMessagesV1MessagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol simple dialog messages v1 messages params
func (o *PostLolSimpleDialogMessagesV1MessagesParams) WithContext(ctx context.Context) *PostLolSimpleDialogMessagesV1MessagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol simple dialog messages v1 messages params
func (o *PostLolSimpleDialogMessagesV1MessagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol simple dialog messages v1 messages params
func (o *PostLolSimpleDialogMessagesV1MessagesParams) WithHTTPClient(client *http.Client) *PostLolSimpleDialogMessagesV1MessagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol simple dialog messages v1 messages params
func (o *PostLolSimpleDialogMessagesV1MessagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMessageRequest adds the messageRequest to the post lol simple dialog messages v1 messages params
func (o *PostLolSimpleDialogMessagesV1MessagesParams) WithMessageRequest(messageRequest *models.LolSimpleDialogMessagesLocalMessageRequest) *PostLolSimpleDialogMessagesV1MessagesParams {
	o.SetMessageRequest(messageRequest)
	return o
}

// SetMessageRequest adds the messageRequest to the post lol simple dialog messages v1 messages params
func (o *PostLolSimpleDialogMessagesV1MessagesParams) SetMessageRequest(messageRequest *models.LolSimpleDialogMessagesLocalMessageRequest) {
	o.MessageRequest = messageRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolSimpleDialogMessagesV1MessagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MessageRequest != nil {
		if err := r.SetBodyParam(o.MessageRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
