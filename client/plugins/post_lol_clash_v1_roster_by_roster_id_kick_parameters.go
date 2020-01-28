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

// NewPostLolClashV1RosterByRosterIDKickParams creates a new PostLolClashV1RosterByRosterIDKickParams object
// with the default values initialized.
func NewPostLolClashV1RosterByRosterIDKickParams() *PostLolClashV1RosterByRosterIDKickParams {
	var ()
	return &PostLolClashV1RosterByRosterIDKickParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolClashV1RosterByRosterIDKickParamsWithTimeout creates a new PostLolClashV1RosterByRosterIDKickParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolClashV1RosterByRosterIDKickParamsWithTimeout(timeout time.Duration) *PostLolClashV1RosterByRosterIDKickParams {
	var ()
	return &PostLolClashV1RosterByRosterIDKickParams{

		timeout: timeout,
	}
}

// NewPostLolClashV1RosterByRosterIDKickParamsWithContext creates a new PostLolClashV1RosterByRosterIDKickParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolClashV1RosterByRosterIDKickParamsWithContext(ctx context.Context) *PostLolClashV1RosterByRosterIDKickParams {
	var ()
	return &PostLolClashV1RosterByRosterIDKickParams{

		Context: ctx,
	}
}

// NewPostLolClashV1RosterByRosterIDKickParamsWithHTTPClient creates a new PostLolClashV1RosterByRosterIDKickParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolClashV1RosterByRosterIDKickParamsWithHTTPClient(client *http.Client) *PostLolClashV1RosterByRosterIDKickParams {
	var ()
	return &PostLolClashV1RosterByRosterIDKickParams{
		HTTPClient: client,
	}
}

/*PostLolClashV1RosterByRosterIDKickParams contains all the parameters to send to the API endpoint
for the post lol clash v1 roster by roster Id kick operation typically these are written to a http.Request
*/
type PostLolClashV1RosterByRosterIDKickParams struct {

	/*KickRequest*/
	KickRequest *models.LolClashKickRequest
	/*RosterID*/
	RosterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol clash v1 roster by roster Id kick params
func (o *PostLolClashV1RosterByRosterIDKickParams) WithTimeout(timeout time.Duration) *PostLolClashV1RosterByRosterIDKickParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol clash v1 roster by roster Id kick params
func (o *PostLolClashV1RosterByRosterIDKickParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol clash v1 roster by roster Id kick params
func (o *PostLolClashV1RosterByRosterIDKickParams) WithContext(ctx context.Context) *PostLolClashV1RosterByRosterIDKickParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol clash v1 roster by roster Id kick params
func (o *PostLolClashV1RosterByRosterIDKickParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol clash v1 roster by roster Id kick params
func (o *PostLolClashV1RosterByRosterIDKickParams) WithHTTPClient(client *http.Client) *PostLolClashV1RosterByRosterIDKickParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol clash v1 roster by roster Id kick params
func (o *PostLolClashV1RosterByRosterIDKickParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKickRequest adds the kickRequest to the post lol clash v1 roster by roster Id kick params
func (o *PostLolClashV1RosterByRosterIDKickParams) WithKickRequest(kickRequest *models.LolClashKickRequest) *PostLolClashV1RosterByRosterIDKickParams {
	o.SetKickRequest(kickRequest)
	return o
}

// SetKickRequest adds the kickRequest to the post lol clash v1 roster by roster Id kick params
func (o *PostLolClashV1RosterByRosterIDKickParams) SetKickRequest(kickRequest *models.LolClashKickRequest) {
	o.KickRequest = kickRequest
}

// WithRosterID adds the rosterID to the post lol clash v1 roster by roster Id kick params
func (o *PostLolClashV1RosterByRosterIDKickParams) WithRosterID(rosterID string) *PostLolClashV1RosterByRosterIDKickParams {
	o.SetRosterID(rosterID)
	return o
}

// SetRosterID adds the rosterId to the post lol clash v1 roster by roster Id kick params
func (o *PostLolClashV1RosterByRosterIDKickParams) SetRosterID(rosterID string) {
	o.RosterID = rosterID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolClashV1RosterByRosterIDKickParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.KickRequest != nil {
		if err := r.SetBodyParam(o.KickRequest); err != nil {
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
