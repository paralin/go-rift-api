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

// NewPostLolClashV1RosterByRosterIDChangeIconParams creates a new PostLolClashV1RosterByRosterIDChangeIconParams object
// with the default values initialized.
func NewPostLolClashV1RosterByRosterIDChangeIconParams() *PostLolClashV1RosterByRosterIDChangeIconParams {
	var ()
	return &PostLolClashV1RosterByRosterIDChangeIconParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolClashV1RosterByRosterIDChangeIconParamsWithTimeout creates a new PostLolClashV1RosterByRosterIDChangeIconParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolClashV1RosterByRosterIDChangeIconParamsWithTimeout(timeout time.Duration) *PostLolClashV1RosterByRosterIDChangeIconParams {
	var ()
	return &PostLolClashV1RosterByRosterIDChangeIconParams{

		timeout: timeout,
	}
}

// NewPostLolClashV1RosterByRosterIDChangeIconParamsWithContext creates a new PostLolClashV1RosterByRosterIDChangeIconParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolClashV1RosterByRosterIDChangeIconParamsWithContext(ctx context.Context) *PostLolClashV1RosterByRosterIDChangeIconParams {
	var ()
	return &PostLolClashV1RosterByRosterIDChangeIconParams{

		Context: ctx,
	}
}

// NewPostLolClashV1RosterByRosterIDChangeIconParamsWithHTTPClient creates a new PostLolClashV1RosterByRosterIDChangeIconParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolClashV1RosterByRosterIDChangeIconParamsWithHTTPClient(client *http.Client) *PostLolClashV1RosterByRosterIDChangeIconParams {
	var ()
	return &PostLolClashV1RosterByRosterIDChangeIconParams{
		HTTPClient: client,
	}
}

/*PostLolClashV1RosterByRosterIDChangeIconParams contains all the parameters to send to the API endpoint
for the post lol clash v1 roster by roster Id change icon operation typically these are written to a http.Request
*/
type PostLolClashV1RosterByRosterIDChangeIconParams struct {

	/*ChangeIconRequest*/
	ChangeIconRequest *models.LolClashChangeIconRequest
	/*RosterID*/
	RosterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol clash v1 roster by roster Id change icon params
func (o *PostLolClashV1RosterByRosterIDChangeIconParams) WithTimeout(timeout time.Duration) *PostLolClashV1RosterByRosterIDChangeIconParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol clash v1 roster by roster Id change icon params
func (o *PostLolClashV1RosterByRosterIDChangeIconParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol clash v1 roster by roster Id change icon params
func (o *PostLolClashV1RosterByRosterIDChangeIconParams) WithContext(ctx context.Context) *PostLolClashV1RosterByRosterIDChangeIconParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol clash v1 roster by roster Id change icon params
func (o *PostLolClashV1RosterByRosterIDChangeIconParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol clash v1 roster by roster Id change icon params
func (o *PostLolClashV1RosterByRosterIDChangeIconParams) WithHTTPClient(client *http.Client) *PostLolClashV1RosterByRosterIDChangeIconParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol clash v1 roster by roster Id change icon params
func (o *PostLolClashV1RosterByRosterIDChangeIconParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChangeIconRequest adds the changeIconRequest to the post lol clash v1 roster by roster Id change icon params
func (o *PostLolClashV1RosterByRosterIDChangeIconParams) WithChangeIconRequest(changeIconRequest *models.LolClashChangeIconRequest) *PostLolClashV1RosterByRosterIDChangeIconParams {
	o.SetChangeIconRequest(changeIconRequest)
	return o
}

// SetChangeIconRequest adds the changeIconRequest to the post lol clash v1 roster by roster Id change icon params
func (o *PostLolClashV1RosterByRosterIDChangeIconParams) SetChangeIconRequest(changeIconRequest *models.LolClashChangeIconRequest) {
	o.ChangeIconRequest = changeIconRequest
}

// WithRosterID adds the rosterID to the post lol clash v1 roster by roster Id change icon params
func (o *PostLolClashV1RosterByRosterIDChangeIconParams) WithRosterID(rosterID string) *PostLolClashV1RosterByRosterIDChangeIconParams {
	o.SetRosterID(rosterID)
	return o
}

// SetRosterID adds the rosterId to the post lol clash v1 roster by roster Id change icon params
func (o *PostLolClashV1RosterByRosterIDChangeIconParams) SetRosterID(rosterID string) {
	o.RosterID = rosterID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolClashV1RosterByRosterIDChangeIconParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ChangeIconRequest != nil {
		if err := r.SetBodyParam(o.ChangeIconRequest); err != nil {
			return err
		}
	}

	// path param rosterId
	if err := r.SetPathParam("rosterId", o.RosterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}