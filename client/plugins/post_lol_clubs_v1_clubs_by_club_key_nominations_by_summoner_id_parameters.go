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

// NewPostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams creates a new PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams object
// with the default values initialized.
func NewPostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams() *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams {
	var ()
	return &PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParamsWithTimeout creates a new PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParamsWithTimeout(timeout time.Duration) *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams {
	var ()
	return &PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams{

		timeout: timeout,
	}
}

// NewPostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParamsWithContext creates a new PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParamsWithContext(ctx context.Context) *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams {
	var ()
	return &PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams{

		Context: ctx,
	}
}

// NewPostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParamsWithHTTPClient creates a new PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParamsWithHTTPClient(client *http.Client) *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams {
	var ()
	return &PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams{
		HTTPClient: client,
	}
}

/*PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams contains all the parameters to send to the API endpoint
for the post lol clubs v1 clubs by club key nominations by summoner Id operation typically these are written to a http.Request
*/
type PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams struct {

	/*ClubKey*/
	ClubKey string
	/*SummonerID*/
	SummonerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol clubs v1 clubs by club key nominations by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams) WithTimeout(timeout time.Duration) *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol clubs v1 clubs by club key nominations by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol clubs v1 clubs by club key nominations by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams) WithContext(ctx context.Context) *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol clubs v1 clubs by club key nominations by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol clubs v1 clubs by club key nominations by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams) WithHTTPClient(client *http.Client) *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol clubs v1 clubs by club key nominations by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClubKey adds the clubKey to the post lol clubs v1 clubs by club key nominations by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams) WithClubKey(clubKey string) *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams {
	o.SetClubKey(clubKey)
	return o
}

// SetClubKey adds the clubKey to the post lol clubs v1 clubs by club key nominations by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams) SetClubKey(clubKey string) {
	o.ClubKey = clubKey
}

// WithSummonerID adds the summonerID to the post lol clubs v1 clubs by club key nominations by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams) WithSummonerID(summonerID int64) *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams {
	o.SetSummonerID(summonerID)
	return o
}

// SetSummonerID adds the summonerId to the post lol clubs v1 clubs by club key nominations by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams) SetSummonerID(summonerID int64) {
	o.SummonerID = summonerID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolClubsV1ClubsByClubKeyNominationsBySummonerIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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