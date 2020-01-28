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

// NewGetVoiceChatV1PushToTalkParams creates a new GetVoiceChatV1PushToTalkParams object
// with the default values initialized.
func NewGetVoiceChatV1PushToTalkParams() *GetVoiceChatV1PushToTalkParams {

	return &GetVoiceChatV1PushToTalkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVoiceChatV1PushToTalkParamsWithTimeout creates a new GetVoiceChatV1PushToTalkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVoiceChatV1PushToTalkParamsWithTimeout(timeout time.Duration) *GetVoiceChatV1PushToTalkParams {

	return &GetVoiceChatV1PushToTalkParams{

		timeout: timeout,
	}
}

// NewGetVoiceChatV1PushToTalkParamsWithContext creates a new GetVoiceChatV1PushToTalkParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVoiceChatV1PushToTalkParamsWithContext(ctx context.Context) *GetVoiceChatV1PushToTalkParams {

	return &GetVoiceChatV1PushToTalkParams{

		Context: ctx,
	}
}

// NewGetVoiceChatV1PushToTalkParamsWithHTTPClient creates a new GetVoiceChatV1PushToTalkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVoiceChatV1PushToTalkParamsWithHTTPClient(client *http.Client) *GetVoiceChatV1PushToTalkParams {

	return &GetVoiceChatV1PushToTalkParams{
		HTTPClient: client,
	}
}

/*GetVoiceChatV1PushToTalkParams contains all the parameters to send to the API endpoint
for the get voice chat v1 push to talk operation typically these are written to a http.Request
*/
type GetVoiceChatV1PushToTalkParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get voice chat v1 push to talk params
func (o *GetVoiceChatV1PushToTalkParams) WithTimeout(timeout time.Duration) *GetVoiceChatV1PushToTalkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get voice chat v1 push to talk params
func (o *GetVoiceChatV1PushToTalkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get voice chat v1 push to talk params
func (o *GetVoiceChatV1PushToTalkParams) WithContext(ctx context.Context) *GetVoiceChatV1PushToTalkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get voice chat v1 push to talk params
func (o *GetVoiceChatV1PushToTalkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get voice chat v1 push to talk params
func (o *GetVoiceChatV1PushToTalkParams) WithHTTPClient(client *http.Client) *GetVoiceChatV1PushToTalkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get voice chat v1 push to talk params
func (o *GetVoiceChatV1PushToTalkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetVoiceChatV1PushToTalkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}