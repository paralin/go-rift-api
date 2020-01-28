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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/paralin/go-rift-api/models"
)

// NewPostLolClashV1TournamentByTournamentIDCreateRosterParams creates a new PostLolClashV1TournamentByTournamentIDCreateRosterParams object
// with the default values initialized.
func NewPostLolClashV1TournamentByTournamentIDCreateRosterParams() *PostLolClashV1TournamentByTournamentIDCreateRosterParams {
	var ()
	return &PostLolClashV1TournamentByTournamentIDCreateRosterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolClashV1TournamentByTournamentIDCreateRosterParamsWithTimeout creates a new PostLolClashV1TournamentByTournamentIDCreateRosterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolClashV1TournamentByTournamentIDCreateRosterParamsWithTimeout(timeout time.Duration) *PostLolClashV1TournamentByTournamentIDCreateRosterParams {
	var ()
	return &PostLolClashV1TournamentByTournamentIDCreateRosterParams{

		timeout: timeout,
	}
}

// NewPostLolClashV1TournamentByTournamentIDCreateRosterParamsWithContext creates a new PostLolClashV1TournamentByTournamentIDCreateRosterParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolClashV1TournamentByTournamentIDCreateRosterParamsWithContext(ctx context.Context) *PostLolClashV1TournamentByTournamentIDCreateRosterParams {
	var ()
	return &PostLolClashV1TournamentByTournamentIDCreateRosterParams{

		Context: ctx,
	}
}

// NewPostLolClashV1TournamentByTournamentIDCreateRosterParamsWithHTTPClient creates a new PostLolClashV1TournamentByTournamentIDCreateRosterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolClashV1TournamentByTournamentIDCreateRosterParamsWithHTTPClient(client *http.Client) *PostLolClashV1TournamentByTournamentIDCreateRosterParams {
	var ()
	return &PostLolClashV1TournamentByTournamentIDCreateRosterParams{
		HTTPClient: client,
	}
}

/*PostLolClashV1TournamentByTournamentIDCreateRosterParams contains all the parameters to send to the API endpoint
for the post lol clash v1 tournament by tournament Id create roster operation typically these are written to a http.Request
*/
type PostLolClashV1TournamentByTournamentIDCreateRosterParams struct {

	/*RosterDetails*/
	RosterDetails *models.LolClashRosterDetails
	/*TournamentID*/
	TournamentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol clash v1 tournament by tournament Id create roster params
func (o *PostLolClashV1TournamentByTournamentIDCreateRosterParams) WithTimeout(timeout time.Duration) *PostLolClashV1TournamentByTournamentIDCreateRosterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol clash v1 tournament by tournament Id create roster params
func (o *PostLolClashV1TournamentByTournamentIDCreateRosterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol clash v1 tournament by tournament Id create roster params
func (o *PostLolClashV1TournamentByTournamentIDCreateRosterParams) WithContext(ctx context.Context) *PostLolClashV1TournamentByTournamentIDCreateRosterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol clash v1 tournament by tournament Id create roster params
func (o *PostLolClashV1TournamentByTournamentIDCreateRosterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol clash v1 tournament by tournament Id create roster params
func (o *PostLolClashV1TournamentByTournamentIDCreateRosterParams) WithHTTPClient(client *http.Client) *PostLolClashV1TournamentByTournamentIDCreateRosterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol clash v1 tournament by tournament Id create roster params
func (o *PostLolClashV1TournamentByTournamentIDCreateRosterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRosterDetails adds the rosterDetails to the post lol clash v1 tournament by tournament Id create roster params
func (o *PostLolClashV1TournamentByTournamentIDCreateRosterParams) WithRosterDetails(rosterDetails *models.LolClashRosterDetails) *PostLolClashV1TournamentByTournamentIDCreateRosterParams {
	o.SetRosterDetails(rosterDetails)
	return o
}

// SetRosterDetails adds the rosterDetails to the post lol clash v1 tournament by tournament Id create roster params
func (o *PostLolClashV1TournamentByTournamentIDCreateRosterParams) SetRosterDetails(rosterDetails *models.LolClashRosterDetails) {
	o.RosterDetails = rosterDetails
}

// WithTournamentID adds the tournamentID to the post lol clash v1 tournament by tournament Id create roster params
func (o *PostLolClashV1TournamentByTournamentIDCreateRosterParams) WithTournamentID(tournamentID int64) *PostLolClashV1TournamentByTournamentIDCreateRosterParams {
	o.SetTournamentID(tournamentID)
	return o
}

// SetTournamentID adds the tournamentId to the post lol clash v1 tournament by tournament Id create roster params
func (o *PostLolClashV1TournamentByTournamentIDCreateRosterParams) SetTournamentID(tournamentID int64) {
	o.TournamentID = tournamentID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolClashV1TournamentByTournamentIDCreateRosterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RosterDetails != nil {
		if err := r.SetBodyParam(o.RosterDetails); err != nil {
			return err
		}
	}

	// path param tournamentId
	if err := r.SetPathParam("tournamentId", swag.FormatInt64(o.TournamentID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
