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

// NewGetLolPremadeVoiceV1CapturedevicesParams creates a new GetLolPremadeVoiceV1CapturedevicesParams object
// with the default values initialized.
func NewGetLolPremadeVoiceV1CapturedevicesParams() *GetLolPremadeVoiceV1CapturedevicesParams {

	return &GetLolPremadeVoiceV1CapturedevicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolPremadeVoiceV1CapturedevicesParamsWithTimeout creates a new GetLolPremadeVoiceV1CapturedevicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolPremadeVoiceV1CapturedevicesParamsWithTimeout(timeout time.Duration) *GetLolPremadeVoiceV1CapturedevicesParams {

	return &GetLolPremadeVoiceV1CapturedevicesParams{

		timeout: timeout,
	}
}

// NewGetLolPremadeVoiceV1CapturedevicesParamsWithContext creates a new GetLolPremadeVoiceV1CapturedevicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolPremadeVoiceV1CapturedevicesParamsWithContext(ctx context.Context) *GetLolPremadeVoiceV1CapturedevicesParams {

	return &GetLolPremadeVoiceV1CapturedevicesParams{

		Context: ctx,
	}
}

// NewGetLolPremadeVoiceV1CapturedevicesParamsWithHTTPClient creates a new GetLolPremadeVoiceV1CapturedevicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolPremadeVoiceV1CapturedevicesParamsWithHTTPClient(client *http.Client) *GetLolPremadeVoiceV1CapturedevicesParams {

	return &GetLolPremadeVoiceV1CapturedevicesParams{
		HTTPClient: client,
	}
}

/*GetLolPremadeVoiceV1CapturedevicesParams contains all the parameters to send to the API endpoint
for the get lol premade voice v1 capturedevices operation typically these are written to a http.Request
*/
type GetLolPremadeVoiceV1CapturedevicesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol premade voice v1 capturedevices params
func (o *GetLolPremadeVoiceV1CapturedevicesParams) WithTimeout(timeout time.Duration) *GetLolPremadeVoiceV1CapturedevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol premade voice v1 capturedevices params
func (o *GetLolPremadeVoiceV1CapturedevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol premade voice v1 capturedevices params
func (o *GetLolPremadeVoiceV1CapturedevicesParams) WithContext(ctx context.Context) *GetLolPremadeVoiceV1CapturedevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol premade voice v1 capturedevices params
func (o *GetLolPremadeVoiceV1CapturedevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol premade voice v1 capturedevices params
func (o *GetLolPremadeVoiceV1CapturedevicesParams) WithHTTPClient(client *http.Client) *GetLolPremadeVoiceV1CapturedevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol premade voice v1 capturedevices params
func (o *GetLolPremadeVoiceV1CapturedevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolPremadeVoiceV1CapturedevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}