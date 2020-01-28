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

// NewGetLolPremadeVoiceV1ParticipantsParams creates a new GetLolPremadeVoiceV1ParticipantsParams object
// with the default values initialized.
func NewGetLolPremadeVoiceV1ParticipantsParams() *GetLolPremadeVoiceV1ParticipantsParams {

	return &GetLolPremadeVoiceV1ParticipantsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolPremadeVoiceV1ParticipantsParamsWithTimeout creates a new GetLolPremadeVoiceV1ParticipantsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolPremadeVoiceV1ParticipantsParamsWithTimeout(timeout time.Duration) *GetLolPremadeVoiceV1ParticipantsParams {

	return &GetLolPremadeVoiceV1ParticipantsParams{

		timeout: timeout,
	}
}

// NewGetLolPremadeVoiceV1ParticipantsParamsWithContext creates a new GetLolPremadeVoiceV1ParticipantsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolPremadeVoiceV1ParticipantsParamsWithContext(ctx context.Context) *GetLolPremadeVoiceV1ParticipantsParams {

	return &GetLolPremadeVoiceV1ParticipantsParams{

		Context: ctx,
	}
}

// NewGetLolPremadeVoiceV1ParticipantsParamsWithHTTPClient creates a new GetLolPremadeVoiceV1ParticipantsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolPremadeVoiceV1ParticipantsParamsWithHTTPClient(client *http.Client) *GetLolPremadeVoiceV1ParticipantsParams {

	return &GetLolPremadeVoiceV1ParticipantsParams{
		HTTPClient: client,
	}
}

/*GetLolPremadeVoiceV1ParticipantsParams contains all the parameters to send to the API endpoint
for the get lol premade voice v1 participants operation typically these are written to a http.Request
*/
type GetLolPremadeVoiceV1ParticipantsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol premade voice v1 participants params
func (o *GetLolPremadeVoiceV1ParticipantsParams) WithTimeout(timeout time.Duration) *GetLolPremadeVoiceV1ParticipantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol premade voice v1 participants params
func (o *GetLolPremadeVoiceV1ParticipantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol premade voice v1 participants params
func (o *GetLolPremadeVoiceV1ParticipantsParams) WithContext(ctx context.Context) *GetLolPremadeVoiceV1ParticipantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol premade voice v1 participants params
func (o *GetLolPremadeVoiceV1ParticipantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol premade voice v1 participants params
func (o *GetLolPremadeVoiceV1ParticipantsParams) WithHTTPClient(client *http.Client) *GetLolPremadeVoiceV1ParticipantsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol premade voice v1 participants params
func (o *GetLolPremadeVoiceV1ParticipantsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolPremadeVoiceV1ParticipantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
