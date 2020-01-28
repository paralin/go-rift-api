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

// NewGetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams creates a new GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams object
// with the default values initialized.
func NewGetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams() *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams {
	var ()
	return &GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParamsWithTimeout creates a new GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParamsWithTimeout(timeout time.Duration) *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams {
	var ()
	return &GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams{

		timeout: timeout,
	}
}

// NewGetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParamsWithContext creates a new GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParamsWithContext(ctx context.Context) *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams {
	var ()
	return &GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams{

		Context: ctx,
	}
}

// NewGetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParamsWithHTTPClient creates a new GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParamsWithHTTPClient(client *http.Client) *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams {
	var ()
	return &GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams{
		HTTPClient: client,
	}
}

/*GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams contains all the parameters to send to the API endpoint
for the get lol career stats v1 champion averages season by season by champion Id by position by tier by queue operation typically these are written to a http.Request
*/
type GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams struct {

	/*ChampionID*/
	ChampionID int32
	/*Position*/
	Position string
	/*Queue*/
	Queue string
	/*Season*/
	Season int32
	/*Tier*/
	Tier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol career stats v1 champion averages season by season by champion Id by position by tier by queue params
func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams) WithTimeout(timeout time.Duration) *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol career stats v1 champion averages season by season by champion Id by position by tier by queue params
func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol career stats v1 champion averages season by season by champion Id by position by tier by queue params
func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams) WithContext(ctx context.Context) *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol career stats v1 champion averages season by season by champion Id by position by tier by queue params
func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol career stats v1 champion averages season by season by champion Id by position by tier by queue params
func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams) WithHTTPClient(client *http.Client) *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol career stats v1 champion averages season by season by champion Id by position by tier by queue params
func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChampionID adds the championID to the get lol career stats v1 champion averages season by season by champion Id by position by tier by queue params
func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams) WithChampionID(championID int32) *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams {
	o.SetChampionID(championID)
	return o
}

// SetChampionID adds the championId to the get lol career stats v1 champion averages season by season by champion Id by position by tier by queue params
func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams) SetChampionID(championID int32) {
	o.ChampionID = championID
}

// WithPosition adds the position to the get lol career stats v1 champion averages season by season by champion Id by position by tier by queue params
func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams) WithPosition(position string) *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams {
	o.SetPosition(position)
	return o
}

// SetPosition adds the position to the get lol career stats v1 champion averages season by season by champion Id by position by tier by queue params
func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams) SetPosition(position string) {
	o.Position = position
}

// WithQueue adds the queue to the get lol career stats v1 champion averages season by season by champion Id by position by tier by queue params
func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams) WithQueue(queue string) *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams {
	o.SetQueue(queue)
	return o
}

// SetQueue adds the queue to the get lol career stats v1 champion averages season by season by champion Id by position by tier by queue params
func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams) SetQueue(queue string) {
	o.Queue = queue
}

// WithSeason adds the season to the get lol career stats v1 champion averages season by season by champion Id by position by tier by queue params
func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams) WithSeason(season int32) *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams {
	o.SetSeason(season)
	return o
}

// SetSeason adds the season to the get lol career stats v1 champion averages season by season by champion Id by position by tier by queue params
func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams) SetSeason(season int32) {
	o.Season = season
}

// WithTier adds the tier to the get lol career stats v1 champion averages season by season by champion Id by position by tier by queue params
func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams) WithTier(tier string) *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams {
	o.SetTier(tier)
	return o
}

// SetTier adds the tier to the get lol career stats v1 champion averages season by season by champion Id by position by tier by queue params
func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams) SetTier(tier string) {
	o.Tier = tier
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param championId
	if err := r.SetPathParam("championId", swag.FormatInt32(o.ChampionID)); err != nil {
		return err
	}

	// path param position
	if err := r.SetPathParam("position", o.Position); err != nil {
		return err
	}

	// path param queue
	if err := r.SetPathParam("queue", o.Queue); err != nil {
		return err
	}

	// path param season
	if err := r.SetPathParam("season", swag.FormatInt32(o.Season)); err != nil {
		return err
	}

	// path param tier
	if err := r.SetPathParam("tier", o.Tier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}