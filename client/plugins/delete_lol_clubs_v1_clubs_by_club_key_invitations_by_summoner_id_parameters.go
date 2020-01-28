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

// NewDeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams creates a new DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams object
// with the default values initialized.
func NewDeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams() *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams {
	var ()
	return &DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParamsWithTimeout creates a new DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParamsWithTimeout(timeout time.Duration) *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams {
	var ()
	return &DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams{

		timeout: timeout,
	}
}

// NewDeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParamsWithContext creates a new DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParamsWithContext(ctx context.Context) *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams {
	var ()
	return &DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams{

		Context: ctx,
	}
}

// NewDeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParamsWithHTTPClient creates a new DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParamsWithHTTPClient(client *http.Client) *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams {
	var ()
	return &DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams{
		HTTPClient: client,
	}
}

/*DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams contains all the parameters to send to the API endpoint
for the delete lol clubs v1 clubs by club key invitations by summoner Id operation typically these are written to a http.Request
*/
type DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams struct {

	/*ClubKey*/
	ClubKey string
	/*SummonerID*/
	SummonerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete lol clubs v1 clubs by club key invitations by summoner Id params
func (o *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams) WithTimeout(timeout time.Duration) *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete lol clubs v1 clubs by club key invitations by summoner Id params
func (o *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete lol clubs v1 clubs by club key invitations by summoner Id params
func (o *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams) WithContext(ctx context.Context) *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete lol clubs v1 clubs by club key invitations by summoner Id params
func (o *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete lol clubs v1 clubs by club key invitations by summoner Id params
func (o *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams) WithHTTPClient(client *http.Client) *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete lol clubs v1 clubs by club key invitations by summoner Id params
func (o *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClubKey adds the clubKey to the delete lol clubs v1 clubs by club key invitations by summoner Id params
func (o *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams) WithClubKey(clubKey string) *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams {
	o.SetClubKey(clubKey)
	return o
}

// SetClubKey adds the clubKey to the delete lol clubs v1 clubs by club key invitations by summoner Id params
func (o *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams) SetClubKey(clubKey string) {
	o.ClubKey = clubKey
}

// WithSummonerID adds the summonerID to the delete lol clubs v1 clubs by club key invitations by summoner Id params
func (o *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams) WithSummonerID(summonerID int64) *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams {
	o.SetSummonerID(summonerID)
	return o
}

// SetSummonerID adds the summonerId to the delete lol clubs v1 clubs by club key invitations by summoner Id params
func (o *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams) SetSummonerID(summonerID int64) {
	o.SummonerID = summonerID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLolClubsV1ClubsByClubKeyInvitationsBySummonerIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clubKey
	if err := r.SetPathParam("clubKey", o.ClubKey); err != nil {
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
