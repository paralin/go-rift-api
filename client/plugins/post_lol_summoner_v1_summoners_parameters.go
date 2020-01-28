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

// NewPostLolSummonerV1SummonersParams creates a new PostLolSummonerV1SummonersParams object
// with the default values initialized.
func NewPostLolSummonerV1SummonersParams() *PostLolSummonerV1SummonersParams {
	var ()
	return &PostLolSummonerV1SummonersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolSummonerV1SummonersParamsWithTimeout creates a new PostLolSummonerV1SummonersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolSummonerV1SummonersParamsWithTimeout(timeout time.Duration) *PostLolSummonerV1SummonersParams {
	var ()
	return &PostLolSummonerV1SummonersParams{

		timeout: timeout,
	}
}

// NewPostLolSummonerV1SummonersParamsWithContext creates a new PostLolSummonerV1SummonersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolSummonerV1SummonersParamsWithContext(ctx context.Context) *PostLolSummonerV1SummonersParams {
	var ()
	return &PostLolSummonerV1SummonersParams{

		Context: ctx,
	}
}

// NewPostLolSummonerV1SummonersParamsWithHTTPClient creates a new PostLolSummonerV1SummonersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolSummonerV1SummonersParamsWithHTTPClient(client *http.Client) *PostLolSummonerV1SummonersParams {
	var ()
	return &PostLolSummonerV1SummonersParams{
		HTTPClient: client,
	}
}

/*PostLolSummonerV1SummonersParams contains all the parameters to send to the API endpoint
for the post lol summoner v1 summoners operation typically these are written to a http.Request
*/
type PostLolSummonerV1SummonersParams struct {

	/*Name*/
	Name *models.LolSummonerSummonerRequestedName

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol summoner v1 summoners params
func (o *PostLolSummonerV1SummonersParams) WithTimeout(timeout time.Duration) *PostLolSummonerV1SummonersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol summoner v1 summoners params
func (o *PostLolSummonerV1SummonersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol summoner v1 summoners params
func (o *PostLolSummonerV1SummonersParams) WithContext(ctx context.Context) *PostLolSummonerV1SummonersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol summoner v1 summoners params
func (o *PostLolSummonerV1SummonersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol summoner v1 summoners params
func (o *PostLolSummonerV1SummonersParams) WithHTTPClient(client *http.Client) *PostLolSummonerV1SummonersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol summoner v1 summoners params
func (o *PostLolSummonerV1SummonersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the post lol summoner v1 summoners params
func (o *PostLolSummonerV1SummonersParams) WithName(name *models.LolSummonerSummonerRequestedName) *PostLolSummonerV1SummonersParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post lol summoner v1 summoners params
func (o *PostLolSummonerV1SummonersParams) SetName(name *models.LolSummonerSummonerRequestedName) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolSummonerV1SummonersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {
		if err := r.SetBodyParam(o.Name); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
