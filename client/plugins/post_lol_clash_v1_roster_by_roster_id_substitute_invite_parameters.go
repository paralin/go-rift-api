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

// NewPostLolClashV1RosterByRosterIDSubstituteInviteParams creates a new PostLolClashV1RosterByRosterIDSubstituteInviteParams object
// with the default values initialized.
func NewPostLolClashV1RosterByRosterIDSubstituteInviteParams() *PostLolClashV1RosterByRosterIDSubstituteInviteParams {
	var ()
	return &PostLolClashV1RosterByRosterIDSubstituteInviteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolClashV1RosterByRosterIDSubstituteInviteParamsWithTimeout creates a new PostLolClashV1RosterByRosterIDSubstituteInviteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolClashV1RosterByRosterIDSubstituteInviteParamsWithTimeout(timeout time.Duration) *PostLolClashV1RosterByRosterIDSubstituteInviteParams {
	var ()
	return &PostLolClashV1RosterByRosterIDSubstituteInviteParams{

		timeout: timeout,
	}
}

// NewPostLolClashV1RosterByRosterIDSubstituteInviteParamsWithContext creates a new PostLolClashV1RosterByRosterIDSubstituteInviteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolClashV1RosterByRosterIDSubstituteInviteParamsWithContext(ctx context.Context) *PostLolClashV1RosterByRosterIDSubstituteInviteParams {
	var ()
	return &PostLolClashV1RosterByRosterIDSubstituteInviteParams{

		Context: ctx,
	}
}

// NewPostLolClashV1RosterByRosterIDSubstituteInviteParamsWithHTTPClient creates a new PostLolClashV1RosterByRosterIDSubstituteInviteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolClashV1RosterByRosterIDSubstituteInviteParamsWithHTTPClient(client *http.Client) *PostLolClashV1RosterByRosterIDSubstituteInviteParams {
	var ()
	return &PostLolClashV1RosterByRosterIDSubstituteInviteParams{
		HTTPClient: client,
	}
}

/*PostLolClashV1RosterByRosterIDSubstituteInviteParams contains all the parameters to send to the API endpoint
for the post lol clash v1 roster by roster Id substitute invite operation typically these are written to a http.Request
*/
type PostLolClashV1RosterByRosterIDSubstituteInviteParams struct {

	/*InviteSubRequest*/
	InviteSubRequest *models.LolClashInviteSubRequest
	/*RosterID*/
	RosterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol clash v1 roster by roster Id substitute invite params
func (o *PostLolClashV1RosterByRosterIDSubstituteInviteParams) WithTimeout(timeout time.Duration) *PostLolClashV1RosterByRosterIDSubstituteInviteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol clash v1 roster by roster Id substitute invite params
func (o *PostLolClashV1RosterByRosterIDSubstituteInviteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol clash v1 roster by roster Id substitute invite params
func (o *PostLolClashV1RosterByRosterIDSubstituteInviteParams) WithContext(ctx context.Context) *PostLolClashV1RosterByRosterIDSubstituteInviteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol clash v1 roster by roster Id substitute invite params
func (o *PostLolClashV1RosterByRosterIDSubstituteInviteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol clash v1 roster by roster Id substitute invite params
func (o *PostLolClashV1RosterByRosterIDSubstituteInviteParams) WithHTTPClient(client *http.Client) *PostLolClashV1RosterByRosterIDSubstituteInviteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol clash v1 roster by roster Id substitute invite params
func (o *PostLolClashV1RosterByRosterIDSubstituteInviteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInviteSubRequest adds the inviteSubRequest to the post lol clash v1 roster by roster Id substitute invite params
func (o *PostLolClashV1RosterByRosterIDSubstituteInviteParams) WithInviteSubRequest(inviteSubRequest *models.LolClashInviteSubRequest) *PostLolClashV1RosterByRosterIDSubstituteInviteParams {
	o.SetInviteSubRequest(inviteSubRequest)
	return o
}

// SetInviteSubRequest adds the inviteSubRequest to the post lol clash v1 roster by roster Id substitute invite params
func (o *PostLolClashV1RosterByRosterIDSubstituteInviteParams) SetInviteSubRequest(inviteSubRequest *models.LolClashInviteSubRequest) {
	o.InviteSubRequest = inviteSubRequest
}

// WithRosterID adds the rosterID to the post lol clash v1 roster by roster Id substitute invite params
func (o *PostLolClashV1RosterByRosterIDSubstituteInviteParams) WithRosterID(rosterID string) *PostLolClashV1RosterByRosterIDSubstituteInviteParams {
	o.SetRosterID(rosterID)
	return o
}

// SetRosterID adds the rosterId to the post lol clash v1 roster by roster Id substitute invite params
func (o *PostLolClashV1RosterByRosterIDSubstituteInviteParams) SetRosterID(rosterID string) {
	o.RosterID = rosterID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolClashV1RosterByRosterIDSubstituteInviteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InviteSubRequest != nil {
		if err := r.SetBodyParam(o.InviteSubRequest); err != nil {
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