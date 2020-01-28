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

// NewGetVoiceChatV1AudioPropertiesParams creates a new GetVoiceChatV1AudioPropertiesParams object
// with the default values initialized.
func NewGetVoiceChatV1AudioPropertiesParams() *GetVoiceChatV1AudioPropertiesParams {

	return &GetVoiceChatV1AudioPropertiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVoiceChatV1AudioPropertiesParamsWithTimeout creates a new GetVoiceChatV1AudioPropertiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVoiceChatV1AudioPropertiesParamsWithTimeout(timeout time.Duration) *GetVoiceChatV1AudioPropertiesParams {

	return &GetVoiceChatV1AudioPropertiesParams{

		timeout: timeout,
	}
}

// NewGetVoiceChatV1AudioPropertiesParamsWithContext creates a new GetVoiceChatV1AudioPropertiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVoiceChatV1AudioPropertiesParamsWithContext(ctx context.Context) *GetVoiceChatV1AudioPropertiesParams {

	return &GetVoiceChatV1AudioPropertiesParams{

		Context: ctx,
	}
}

// NewGetVoiceChatV1AudioPropertiesParamsWithHTTPClient creates a new GetVoiceChatV1AudioPropertiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVoiceChatV1AudioPropertiesParamsWithHTTPClient(client *http.Client) *GetVoiceChatV1AudioPropertiesParams {

	return &GetVoiceChatV1AudioPropertiesParams{
		HTTPClient: client,
	}
}

/*GetVoiceChatV1AudioPropertiesParams contains all the parameters to send to the API endpoint
for the get voice chat v1 audio properties operation typically these are written to a http.Request
*/
type GetVoiceChatV1AudioPropertiesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get voice chat v1 audio properties params
func (o *GetVoiceChatV1AudioPropertiesParams) WithTimeout(timeout time.Duration) *GetVoiceChatV1AudioPropertiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get voice chat v1 audio properties params
func (o *GetVoiceChatV1AudioPropertiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get voice chat v1 audio properties params
func (o *GetVoiceChatV1AudioPropertiesParams) WithContext(ctx context.Context) *GetVoiceChatV1AudioPropertiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get voice chat v1 audio properties params
func (o *GetVoiceChatV1AudioPropertiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get voice chat v1 audio properties params
func (o *GetVoiceChatV1AudioPropertiesParams) WithHTTPClient(client *http.Client) *GetVoiceChatV1AudioPropertiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get voice chat v1 audio properties params
func (o *GetVoiceChatV1AudioPropertiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetVoiceChatV1AudioPropertiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
