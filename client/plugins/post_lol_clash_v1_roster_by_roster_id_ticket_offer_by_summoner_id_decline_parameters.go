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
)

// NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams creates a new PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams object
// with the default values initialized.
func NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams() *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams {
	var ()
	return &PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParamsWithTimeout creates a new PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParamsWithTimeout(timeout time.Duration) *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams {
	var ()
	return &PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams{

		timeout: timeout,
	}
}

// NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParamsWithContext creates a new PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParamsWithContext(ctx context.Context) *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams {
	var ()
	return &PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams{

		Context: ctx,
	}
}

// NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParamsWithHTTPClient creates a new PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParamsWithHTTPClient(client *http.Client) *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams {
	var ()
	return &PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams{
		HTTPClient: client,
	}
}

/*PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams contains all the parameters to send to the API endpoint
for the post lol clash v1 roster by roster Id ticket offer by summoner Id decline operation typically these are written to a http.Request
*/
type PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams struct {

	/*RosterID*/
	RosterID string
	/*SummonerID*/
	SummonerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol clash v1 roster by roster Id ticket offer by summoner Id decline params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams) WithTimeout(timeout time.Duration) *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol clash v1 roster by roster Id ticket offer by summoner Id decline params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol clash v1 roster by roster Id ticket offer by summoner Id decline params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams) WithContext(ctx context.Context) *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol clash v1 roster by roster Id ticket offer by summoner Id decline params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol clash v1 roster by roster Id ticket offer by summoner Id decline params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams) WithHTTPClient(client *http.Client) *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol clash v1 roster by roster Id ticket offer by summoner Id decline params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRosterID adds the rosterID to the post lol clash v1 roster by roster Id ticket offer by summoner Id decline params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams) WithRosterID(rosterID string) *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams {
	o.SetRosterID(rosterID)
	return o
}

// SetRosterID adds the rosterId to the post lol clash v1 roster by roster Id ticket offer by summoner Id decline params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams) SetRosterID(rosterID string) {
	o.RosterID = rosterID
}

// WithSummonerID adds the summonerID to the post lol clash v1 roster by roster Id ticket offer by summoner Id decline params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams) WithSummonerID(summonerID int64) *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams {
	o.SetSummonerID(summonerID)
	return o
}

// SetSummonerID adds the summonerId to the post lol clash v1 roster by roster Id ticket offer by summoner Id decline params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams) SetSummonerID(summonerID int64) {
	o.SummonerID = summonerID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDDeclineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param rosterId
	if err := r.SetPathParam("rosterId", o.RosterID); err != nil {
		return err
	}

	// path param summonerId
	if err := r.SetPathParam("summonerId", swag.FormatInt64(o.SummonerID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
