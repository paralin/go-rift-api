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

// NewPostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams creates a new PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams object
// with the default values initialized.
func NewPostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams() *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams {
	var ()
	return &PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParamsWithTimeout creates a new PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParamsWithTimeout(timeout time.Duration) *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams {
	var ()
	return &PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams{

		timeout: timeout,
	}
}

// NewPostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParamsWithContext creates a new PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParamsWithContext(ctx context.Context) *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams {
	var ()
	return &PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams{

		Context: ctx,
	}
}

// NewPostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParamsWithHTTPClient creates a new PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParamsWithHTTPClient(client *http.Client) *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams {
	var ()
	return &PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams{
		HTTPClient: client,
	}
}

/*PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams contains all the parameters to send to the API endpoint
for the post lol statstones v1 featured champion statstones by champion item Id by statstone Id operation typically these are written to a http.Request
*/
type PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams struct {

	/*ChampionItemID*/
	ChampionItemID int32
	/*FeaturedInfo*/
	FeaturedInfo *models.LolStatstonesStatstoneFeaturedRequest
	/*StatstoneID*/
	StatstoneID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol statstones v1 featured champion statstones by champion item Id by statstone Id params
func (o *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams) WithTimeout(timeout time.Duration) *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol statstones v1 featured champion statstones by champion item Id by statstone Id params
func (o *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol statstones v1 featured champion statstones by champion item Id by statstone Id params
func (o *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams) WithContext(ctx context.Context) *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol statstones v1 featured champion statstones by champion item Id by statstone Id params
func (o *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol statstones v1 featured champion statstones by champion item Id by statstone Id params
func (o *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams) WithHTTPClient(client *http.Client) *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol statstones v1 featured champion statstones by champion item Id by statstone Id params
func (o *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChampionItemID adds the championItemID to the post lol statstones v1 featured champion statstones by champion item Id by statstone Id params
func (o *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams) WithChampionItemID(championItemID int32) *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams {
	o.SetChampionItemID(championItemID)
	return o
}

// SetChampionItemID adds the championItemId to the post lol statstones v1 featured champion statstones by champion item Id by statstone Id params
func (o *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams) SetChampionItemID(championItemID int32) {
	o.ChampionItemID = championItemID
}

// WithFeaturedInfo adds the featuredInfo to the post lol statstones v1 featured champion statstones by champion item Id by statstone Id params
func (o *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams) WithFeaturedInfo(featuredInfo *models.LolStatstonesStatstoneFeaturedRequest) *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams {
	o.SetFeaturedInfo(featuredInfo)
	return o
}

// SetFeaturedInfo adds the featuredInfo to the post lol statstones v1 featured champion statstones by champion item Id by statstone Id params
func (o *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams) SetFeaturedInfo(featuredInfo *models.LolStatstonesStatstoneFeaturedRequest) {
	o.FeaturedInfo = featuredInfo
}

// WithStatstoneID adds the statstoneID to the post lol statstones v1 featured champion statstones by champion item Id by statstone Id params
func (o *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams) WithStatstoneID(statstoneID string) *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams {
	o.SetStatstoneID(statstoneID)
	return o
}

// SetStatstoneID adds the statstoneId to the post lol statstones v1 featured champion statstones by champion item Id by statstone Id params
func (o *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams) SetStatstoneID(statstoneID string) {
	o.StatstoneID = statstoneID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolStatstonesV1FeaturedChampionStatstonesByChampionItemIDByStatstoneIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param championItemId
	if err := r.SetPathParam("championItemId", swag.FormatInt32(o.ChampionItemID)); err != nil {
		return err
	}

	if o.FeaturedInfo != nil {
		if err := r.SetBodyParam(o.FeaturedInfo); err != nil {
			return err
		}
	}

	// path param statstoneId
	if err := r.SetPathParam("statstoneId", o.StatstoneID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
