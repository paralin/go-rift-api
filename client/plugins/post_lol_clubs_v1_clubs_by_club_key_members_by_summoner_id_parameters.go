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

// NewPostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams creates a new PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams object
// with the default values initialized.
func NewPostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams() *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams {
	var ()
	return &PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolClubsV1ClubsByClubKeyMembersBySummonerIDParamsWithTimeout creates a new PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolClubsV1ClubsByClubKeyMembersBySummonerIDParamsWithTimeout(timeout time.Duration) *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams {
	var ()
	return &PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams{

		timeout: timeout,
	}
}

// NewPostLolClubsV1ClubsByClubKeyMembersBySummonerIDParamsWithContext creates a new PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolClubsV1ClubsByClubKeyMembersBySummonerIDParamsWithContext(ctx context.Context) *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams {
	var ()
	return &PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams{

		Context: ctx,
	}
}

// NewPostLolClubsV1ClubsByClubKeyMembersBySummonerIDParamsWithHTTPClient creates a new PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolClubsV1ClubsByClubKeyMembersBySummonerIDParamsWithHTTPClient(client *http.Client) *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams {
	var ()
	return &PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams{
		HTTPClient: client,
	}
}

/*PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams contains all the parameters to send to the API endpoint
for the post lol clubs v1 clubs by club key members by summoner Id operation typically these are written to a http.Request
*/
type PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams struct {

	/*ClubKey*/
	ClubKey string
	/*SummonerID*/
	SummonerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol clubs v1 clubs by club key members by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams) WithTimeout(timeout time.Duration) *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol clubs v1 clubs by club key members by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol clubs v1 clubs by club key members by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams) WithContext(ctx context.Context) *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol clubs v1 clubs by club key members by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol clubs v1 clubs by club key members by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams) WithHTTPClient(client *http.Client) *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol clubs v1 clubs by club key members by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClubKey adds the clubKey to the post lol clubs v1 clubs by club key members by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams) WithClubKey(clubKey string) *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams {
	o.SetClubKey(clubKey)
	return o
}

// SetClubKey adds the clubKey to the post lol clubs v1 clubs by club key members by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams) SetClubKey(clubKey string) {
	o.ClubKey = clubKey
}

// WithSummonerID adds the summonerID to the post lol clubs v1 clubs by club key members by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams) WithSummonerID(summonerID int64) *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams {
	o.SetSummonerID(summonerID)
	return o
}

// SetSummonerID adds the summonerId to the post lol clubs v1 clubs by club key members by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams) SetSummonerID(summonerID int64) {
	o.SummonerID = summonerID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolClubsV1ClubsByClubKeyMembersBySummonerIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
