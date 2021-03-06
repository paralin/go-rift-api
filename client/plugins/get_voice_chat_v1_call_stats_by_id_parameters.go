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

// NewGetVoiceChatV1CallStatsByIDParams creates a new GetVoiceChatV1CallStatsByIDParams object
// with the default values initialized.
func NewGetVoiceChatV1CallStatsByIDParams() *GetVoiceChatV1CallStatsByIDParams {
	var ()
	return &GetVoiceChatV1CallStatsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVoiceChatV1CallStatsByIDParamsWithTimeout creates a new GetVoiceChatV1CallStatsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVoiceChatV1CallStatsByIDParamsWithTimeout(timeout time.Duration) *GetVoiceChatV1CallStatsByIDParams {
	var ()
	return &GetVoiceChatV1CallStatsByIDParams{

		timeout: timeout,
	}
}

// NewGetVoiceChatV1CallStatsByIDParamsWithContext creates a new GetVoiceChatV1CallStatsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVoiceChatV1CallStatsByIDParamsWithContext(ctx context.Context) *GetVoiceChatV1CallStatsByIDParams {
	var ()
	return &GetVoiceChatV1CallStatsByIDParams{

		Context: ctx,
	}
}

// NewGetVoiceChatV1CallStatsByIDParamsWithHTTPClient creates a new GetVoiceChatV1CallStatsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVoiceChatV1CallStatsByIDParamsWithHTTPClient(client *http.Client) *GetVoiceChatV1CallStatsByIDParams {
	var ()
	return &GetVoiceChatV1CallStatsByIDParams{
		HTTPClient: client,
	}
}

/*GetVoiceChatV1CallStatsByIDParams contains all the parameters to send to the API endpoint
for the get voice chat v1 call stats by Id operation typically these are written to a http.Request
*/
type GetVoiceChatV1CallStatsByIDParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get voice chat v1 call stats by Id params
func (o *GetVoiceChatV1CallStatsByIDParams) WithTimeout(timeout time.Duration) *GetVoiceChatV1CallStatsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get voice chat v1 call stats by Id params
func (o *GetVoiceChatV1CallStatsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get voice chat v1 call stats by Id params
func (o *GetVoiceChatV1CallStatsByIDParams) WithContext(ctx context.Context) *GetVoiceChatV1CallStatsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get voice chat v1 call stats by Id params
func (o *GetVoiceChatV1CallStatsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get voice chat v1 call stats by Id params
func (o *GetVoiceChatV1CallStatsByIDParams) WithHTTPClient(client *http.Client) *GetVoiceChatV1CallStatsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get voice chat v1 call stats by Id params
func (o *GetVoiceChatV1CallStatsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get voice chat v1 call stats by Id params
func (o *GetVoiceChatV1CallStatsByIDParams) WithID(id string) *GetVoiceChatV1CallStatsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get voice chat v1 call stats by Id params
func (o *GetVoiceChatV1CallStatsByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetVoiceChatV1CallStatsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
