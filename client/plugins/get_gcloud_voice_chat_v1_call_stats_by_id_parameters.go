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

// NewGetGcloudVoiceChatV1CallStatsByIDParams creates a new GetGcloudVoiceChatV1CallStatsByIDParams object
// with the default values initialized.
func NewGetGcloudVoiceChatV1CallStatsByIDParams() *GetGcloudVoiceChatV1CallStatsByIDParams {
	var ()
	return &GetGcloudVoiceChatV1CallStatsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGcloudVoiceChatV1CallStatsByIDParamsWithTimeout creates a new GetGcloudVoiceChatV1CallStatsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGcloudVoiceChatV1CallStatsByIDParamsWithTimeout(timeout time.Duration) *GetGcloudVoiceChatV1CallStatsByIDParams {
	var ()
	return &GetGcloudVoiceChatV1CallStatsByIDParams{

		timeout: timeout,
	}
}

// NewGetGcloudVoiceChatV1CallStatsByIDParamsWithContext creates a new GetGcloudVoiceChatV1CallStatsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGcloudVoiceChatV1CallStatsByIDParamsWithContext(ctx context.Context) *GetGcloudVoiceChatV1CallStatsByIDParams {
	var ()
	return &GetGcloudVoiceChatV1CallStatsByIDParams{

		Context: ctx,
	}
}

// NewGetGcloudVoiceChatV1CallStatsByIDParamsWithHTTPClient creates a new GetGcloudVoiceChatV1CallStatsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGcloudVoiceChatV1CallStatsByIDParamsWithHTTPClient(client *http.Client) *GetGcloudVoiceChatV1CallStatsByIDParams {
	var ()
	return &GetGcloudVoiceChatV1CallStatsByIDParams{
		HTTPClient: client,
	}
}

/*GetGcloudVoiceChatV1CallStatsByIDParams contains all the parameters to send to the API endpoint
for the get gcloud voice chat v1 call stats by Id operation typically these are written to a http.Request
*/
type GetGcloudVoiceChatV1CallStatsByIDParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get gcloud voice chat v1 call stats by Id params
func (o *GetGcloudVoiceChatV1CallStatsByIDParams) WithTimeout(timeout time.Duration) *GetGcloudVoiceChatV1CallStatsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gcloud voice chat v1 call stats by Id params
func (o *GetGcloudVoiceChatV1CallStatsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gcloud voice chat v1 call stats by Id params
func (o *GetGcloudVoiceChatV1CallStatsByIDParams) WithContext(ctx context.Context) *GetGcloudVoiceChatV1CallStatsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gcloud voice chat v1 call stats by Id params
func (o *GetGcloudVoiceChatV1CallStatsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gcloud voice chat v1 call stats by Id params
func (o *GetGcloudVoiceChatV1CallStatsByIDParams) WithHTTPClient(client *http.Client) *GetGcloudVoiceChatV1CallStatsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gcloud voice chat v1 call stats by Id params
func (o *GetGcloudVoiceChatV1CallStatsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get gcloud voice chat v1 call stats by Id params
func (o *GetGcloudVoiceChatV1CallStatsByIDParams) WithID(id string) *GetGcloudVoiceChatV1CallStatsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get gcloud voice chat v1 call stats by Id params
func (o *GetGcloudVoiceChatV1CallStatsByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetGcloudVoiceChatV1CallStatsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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