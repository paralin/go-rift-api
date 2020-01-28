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

// NewPostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams creates a new PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams object
// with the default values initialized.
func NewPostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams() *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams {
	var ()
	return &PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParamsWithTimeout creates a new PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParamsWithTimeout(timeout time.Duration) *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams {
	var ()
	return &PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams{

		timeout: timeout,
	}
}

// NewPostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParamsWithContext creates a new PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParamsWithContext(ctx context.Context) *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams {
	var ()
	return &PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams{

		Context: ctx,
	}
}

// NewPostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParamsWithHTTPClient creates a new PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParamsWithHTTPClient(client *http.Client) *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams {
	var ()
	return &PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams{
		HTTPClient: client,
	}
}

/*PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams contains all the parameters to send to the API endpoint
for the post lol clubs v1 clubs by club key promotions by summoner Id operation typically these are written to a http.Request
*/
type PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams struct {

	/*ClubKey*/
	ClubKey string
	/*SummonerID*/
	SummonerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol clubs v1 clubs by club key promotions by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams) WithTimeout(timeout time.Duration) *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol clubs v1 clubs by club key promotions by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol clubs v1 clubs by club key promotions by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams) WithContext(ctx context.Context) *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol clubs v1 clubs by club key promotions by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol clubs v1 clubs by club key promotions by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams) WithHTTPClient(client *http.Client) *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol clubs v1 clubs by club key promotions by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClubKey adds the clubKey to the post lol clubs v1 clubs by club key promotions by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams) WithClubKey(clubKey string) *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams {
	o.SetClubKey(clubKey)
	return o
}

// SetClubKey adds the clubKey to the post lol clubs v1 clubs by club key promotions by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams) SetClubKey(clubKey string) {
	o.ClubKey = clubKey
}

// WithSummonerID adds the summonerID to the post lol clubs v1 clubs by club key promotions by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams) WithSummonerID(summonerID int64) *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams {
	o.SetSummonerID(summonerID)
	return o
}

// SetSummonerID adds the summonerId to the post lol clubs v1 clubs by club key promotions by summoner Id params
func (o *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams) SetSummonerID(summonerID int64) {
	o.SummonerID = summonerID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolClubsV1ClubsByClubKeyPromotionsBySummonerIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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