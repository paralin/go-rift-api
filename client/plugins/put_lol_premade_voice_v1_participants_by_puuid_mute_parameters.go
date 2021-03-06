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

// NewPutLolPremadeVoiceV1ParticipantsByPuuidMuteParams creates a new PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams object
// with the default values initialized.
func NewPutLolPremadeVoiceV1ParticipantsByPuuidMuteParams() *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams {
	var ()
	return &PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLolPremadeVoiceV1ParticipantsByPuuidMuteParamsWithTimeout creates a new PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLolPremadeVoiceV1ParticipantsByPuuidMuteParamsWithTimeout(timeout time.Duration) *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams {
	var ()
	return &PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams{

		timeout: timeout,
	}
}

// NewPutLolPremadeVoiceV1ParticipantsByPuuidMuteParamsWithContext creates a new PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLolPremadeVoiceV1ParticipantsByPuuidMuteParamsWithContext(ctx context.Context) *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams {
	var ()
	return &PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams{

		Context: ctx,
	}
}

// NewPutLolPremadeVoiceV1ParticipantsByPuuidMuteParamsWithHTTPClient creates a new PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLolPremadeVoiceV1ParticipantsByPuuidMuteParamsWithHTTPClient(client *http.Client) *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams {
	var ()
	return &PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams{
		HTTPClient: client,
	}
}

/*PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams contains all the parameters to send to the API endpoint
for the put lol premade voice v1 participants by puuid mute operation typically these are written to a http.Request
*/
type PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams struct {

	/*Muted*/
	Muted int32
	/*Puuid*/
	Puuid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put lol premade voice v1 participants by puuid mute params
func (o *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams) WithTimeout(timeout time.Duration) *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put lol premade voice v1 participants by puuid mute params
func (o *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put lol premade voice v1 participants by puuid mute params
func (o *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams) WithContext(ctx context.Context) *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put lol premade voice v1 participants by puuid mute params
func (o *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put lol premade voice v1 participants by puuid mute params
func (o *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams) WithHTTPClient(client *http.Client) *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put lol premade voice v1 participants by puuid mute params
func (o *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMuted adds the muted to the put lol premade voice v1 participants by puuid mute params
func (o *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams) WithMuted(muted int32) *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams {
	o.SetMuted(muted)
	return o
}

// SetMuted adds the muted to the put lol premade voice v1 participants by puuid mute params
func (o *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams) SetMuted(muted int32) {
	o.Muted = muted
}

// WithPuuid adds the puuid to the put lol premade voice v1 participants by puuid mute params
func (o *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams) WithPuuid(puuid string) *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams {
	o.SetPuuid(puuid)
	return o
}

// SetPuuid adds the puuid to the put lol premade voice v1 participants by puuid mute params
func (o *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams) SetPuuid(puuid string) {
	o.Puuid = puuid
}

// WriteToRequest writes these params to a swagger request
func (o *PutLolPremadeVoiceV1ParticipantsByPuuidMuteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Muted); err != nil {
		return err
	}

	// path param puuid
	if err := r.SetPathParam("puuid", o.Puuid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
