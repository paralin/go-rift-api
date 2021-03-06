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

// NewPutLolPremadeVoiceV1SelfActivationSensitivityParams creates a new PutLolPremadeVoiceV1SelfActivationSensitivityParams object
// with the default values initialized.
func NewPutLolPremadeVoiceV1SelfActivationSensitivityParams() *PutLolPremadeVoiceV1SelfActivationSensitivityParams {
	var ()
	return &PutLolPremadeVoiceV1SelfActivationSensitivityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLolPremadeVoiceV1SelfActivationSensitivityParamsWithTimeout creates a new PutLolPremadeVoiceV1SelfActivationSensitivityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLolPremadeVoiceV1SelfActivationSensitivityParamsWithTimeout(timeout time.Duration) *PutLolPremadeVoiceV1SelfActivationSensitivityParams {
	var ()
	return &PutLolPremadeVoiceV1SelfActivationSensitivityParams{

		timeout: timeout,
	}
}

// NewPutLolPremadeVoiceV1SelfActivationSensitivityParamsWithContext creates a new PutLolPremadeVoiceV1SelfActivationSensitivityParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLolPremadeVoiceV1SelfActivationSensitivityParamsWithContext(ctx context.Context) *PutLolPremadeVoiceV1SelfActivationSensitivityParams {
	var ()
	return &PutLolPremadeVoiceV1SelfActivationSensitivityParams{

		Context: ctx,
	}
}

// NewPutLolPremadeVoiceV1SelfActivationSensitivityParamsWithHTTPClient creates a new PutLolPremadeVoiceV1SelfActivationSensitivityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLolPremadeVoiceV1SelfActivationSensitivityParamsWithHTTPClient(client *http.Client) *PutLolPremadeVoiceV1SelfActivationSensitivityParams {
	var ()
	return &PutLolPremadeVoiceV1SelfActivationSensitivityParams{
		HTTPClient: client,
	}
}

/*PutLolPremadeVoiceV1SelfActivationSensitivityParams contains all the parameters to send to the API endpoint
for the put lol premade voice v1 self activation sensitivity operation typically these are written to a http.Request
*/
type PutLolPremadeVoiceV1SelfActivationSensitivityParams struct {

	/*Sensitivity*/
	Sensitivity int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put lol premade voice v1 self activation sensitivity params
func (o *PutLolPremadeVoiceV1SelfActivationSensitivityParams) WithTimeout(timeout time.Duration) *PutLolPremadeVoiceV1SelfActivationSensitivityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put lol premade voice v1 self activation sensitivity params
func (o *PutLolPremadeVoiceV1SelfActivationSensitivityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put lol premade voice v1 self activation sensitivity params
func (o *PutLolPremadeVoiceV1SelfActivationSensitivityParams) WithContext(ctx context.Context) *PutLolPremadeVoiceV1SelfActivationSensitivityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put lol premade voice v1 self activation sensitivity params
func (o *PutLolPremadeVoiceV1SelfActivationSensitivityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put lol premade voice v1 self activation sensitivity params
func (o *PutLolPremadeVoiceV1SelfActivationSensitivityParams) WithHTTPClient(client *http.Client) *PutLolPremadeVoiceV1SelfActivationSensitivityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put lol premade voice v1 self activation sensitivity params
func (o *PutLolPremadeVoiceV1SelfActivationSensitivityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSensitivity adds the sensitivity to the put lol premade voice v1 self activation sensitivity params
func (o *PutLolPremadeVoiceV1SelfActivationSensitivityParams) WithSensitivity(sensitivity int32) *PutLolPremadeVoiceV1SelfActivationSensitivityParams {
	o.SetSensitivity(sensitivity)
	return o
}

// SetSensitivity adds the sensitivity to the put lol premade voice v1 self activation sensitivity params
func (o *PutLolPremadeVoiceV1SelfActivationSensitivityParams) SetSensitivity(sensitivity int32) {
	o.Sensitivity = sensitivity
}

// WriteToRequest writes these params to a swagger request
func (o *PutLolPremadeVoiceV1SelfActivationSensitivityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Sensitivity); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
