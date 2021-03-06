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

// NewPutLolRegaliaV2CurrentSummonerRegaliaParams creates a new PutLolRegaliaV2CurrentSummonerRegaliaParams object
// with the default values initialized.
func NewPutLolRegaliaV2CurrentSummonerRegaliaParams() *PutLolRegaliaV2CurrentSummonerRegaliaParams {
	var ()
	return &PutLolRegaliaV2CurrentSummonerRegaliaParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLolRegaliaV2CurrentSummonerRegaliaParamsWithTimeout creates a new PutLolRegaliaV2CurrentSummonerRegaliaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLolRegaliaV2CurrentSummonerRegaliaParamsWithTimeout(timeout time.Duration) *PutLolRegaliaV2CurrentSummonerRegaliaParams {
	var ()
	return &PutLolRegaliaV2CurrentSummonerRegaliaParams{

		timeout: timeout,
	}
}

// NewPutLolRegaliaV2CurrentSummonerRegaliaParamsWithContext creates a new PutLolRegaliaV2CurrentSummonerRegaliaParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLolRegaliaV2CurrentSummonerRegaliaParamsWithContext(ctx context.Context) *PutLolRegaliaV2CurrentSummonerRegaliaParams {
	var ()
	return &PutLolRegaliaV2CurrentSummonerRegaliaParams{

		Context: ctx,
	}
}

// NewPutLolRegaliaV2CurrentSummonerRegaliaParamsWithHTTPClient creates a new PutLolRegaliaV2CurrentSummonerRegaliaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLolRegaliaV2CurrentSummonerRegaliaParamsWithHTTPClient(client *http.Client) *PutLolRegaliaV2CurrentSummonerRegaliaParams {
	var ()
	return &PutLolRegaliaV2CurrentSummonerRegaliaParams{
		HTTPClient: client,
	}
}

/*PutLolRegaliaV2CurrentSummonerRegaliaParams contains all the parameters to send to the API endpoint
for the put lol regalia v2 current summoner regalia operation typically these are written to a http.Request
*/
type PutLolRegaliaV2CurrentSummonerRegaliaParams struct {

	/*Regalia*/
	Regalia *models.LolRegaliaRegaliaPreferences

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put lol regalia v2 current summoner regalia params
func (o *PutLolRegaliaV2CurrentSummonerRegaliaParams) WithTimeout(timeout time.Duration) *PutLolRegaliaV2CurrentSummonerRegaliaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put lol regalia v2 current summoner regalia params
func (o *PutLolRegaliaV2CurrentSummonerRegaliaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put lol regalia v2 current summoner regalia params
func (o *PutLolRegaliaV2CurrentSummonerRegaliaParams) WithContext(ctx context.Context) *PutLolRegaliaV2CurrentSummonerRegaliaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put lol regalia v2 current summoner regalia params
func (o *PutLolRegaliaV2CurrentSummonerRegaliaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put lol regalia v2 current summoner regalia params
func (o *PutLolRegaliaV2CurrentSummonerRegaliaParams) WithHTTPClient(client *http.Client) *PutLolRegaliaV2CurrentSummonerRegaliaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put lol regalia v2 current summoner regalia params
func (o *PutLolRegaliaV2CurrentSummonerRegaliaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegalia adds the regalia to the put lol regalia v2 current summoner regalia params
func (o *PutLolRegaliaV2CurrentSummonerRegaliaParams) WithRegalia(regalia *models.LolRegaliaRegaliaPreferences) *PutLolRegaliaV2CurrentSummonerRegaliaParams {
	o.SetRegalia(regalia)
	return o
}

// SetRegalia adds the regalia to the put lol regalia v2 current summoner regalia params
func (o *PutLolRegaliaV2CurrentSummonerRegaliaParams) SetRegalia(regalia *models.LolRegaliaRegaliaPreferences) {
	o.Regalia = regalia
}

// WriteToRequest writes these params to a swagger request
func (o *PutLolRegaliaV2CurrentSummonerRegaliaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Regalia != nil {
		if err := r.SetBodyParam(o.Regalia); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
