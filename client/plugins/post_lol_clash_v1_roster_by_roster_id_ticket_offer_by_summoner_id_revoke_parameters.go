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

// NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams creates a new PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams object
// with the default values initialized.
func NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams() *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams {
	var ()
	return &PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParamsWithTimeout creates a new PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParamsWithTimeout(timeout time.Duration) *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams {
	var ()
	return &PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams{

		timeout: timeout,
	}
}

// NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParamsWithContext creates a new PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParamsWithContext(ctx context.Context) *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams {
	var ()
	return &PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams{

		Context: ctx,
	}
}

// NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParamsWithHTTPClient creates a new PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParamsWithHTTPClient(client *http.Client) *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams {
	var ()
	return &PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams{
		HTTPClient: client,
	}
}

/*PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams contains all the parameters to send to the API endpoint
for the post lol clash v1 roster by roster Id ticket offer by summoner Id revoke operation typically these are written to a http.Request
*/
type PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams struct {

	/*RosterID*/
	RosterID string
	/*SummonerID*/
	SummonerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol clash v1 roster by roster Id ticket offer by summoner Id revoke params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams) WithTimeout(timeout time.Duration) *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol clash v1 roster by roster Id ticket offer by summoner Id revoke params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol clash v1 roster by roster Id ticket offer by summoner Id revoke params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams) WithContext(ctx context.Context) *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol clash v1 roster by roster Id ticket offer by summoner Id revoke params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol clash v1 roster by roster Id ticket offer by summoner Id revoke params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams) WithHTTPClient(client *http.Client) *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol clash v1 roster by roster Id ticket offer by summoner Id revoke params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRosterID adds the rosterID to the post lol clash v1 roster by roster Id ticket offer by summoner Id revoke params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams) WithRosterID(rosterID string) *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams {
	o.SetRosterID(rosterID)
	return o
}

// SetRosterID adds the rosterId to the post lol clash v1 roster by roster Id ticket offer by summoner Id revoke params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams) SetRosterID(rosterID string) {
	o.RosterID = rosterID
}

// WithSummonerID adds the summonerID to the post lol clash v1 roster by roster Id ticket offer by summoner Id revoke params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams) WithSummonerID(summonerID int64) *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams {
	o.SetSummonerID(summonerID)
	return o
}

// SetSummonerID adds the summonerId to the post lol clash v1 roster by roster Id ticket offer by summoner Id revoke params
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams) SetSummonerID(summonerID int64) {
	o.SummonerID = summonerID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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