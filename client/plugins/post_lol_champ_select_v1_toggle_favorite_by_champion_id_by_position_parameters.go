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

// NewPostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams creates a new PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams object
// with the default values initialized.
func NewPostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams() *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams {
	var ()
	return &PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParamsWithTimeout creates a new PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParamsWithTimeout(timeout time.Duration) *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams {
	var ()
	return &PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams{

		timeout: timeout,
	}
}

// NewPostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParamsWithContext creates a new PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParamsWithContext(ctx context.Context) *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams {
	var ()
	return &PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams{

		Context: ctx,
	}
}

// NewPostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParamsWithHTTPClient creates a new PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParamsWithHTTPClient(client *http.Client) *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams {
	var ()
	return &PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams{
		HTTPClient: client,
	}
}

/*PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams contains all the parameters to send to the API endpoint
for the post lol champ select v1 toggle favorite by champion Id by position operation typically these are written to a http.Request
*/
type PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams struct {

	/*ChampionID*/
	ChampionID int64
	/*Position*/
	Position string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol champ select v1 toggle favorite by champion Id by position params
func (o *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams) WithTimeout(timeout time.Duration) *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol champ select v1 toggle favorite by champion Id by position params
func (o *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol champ select v1 toggle favorite by champion Id by position params
func (o *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams) WithContext(ctx context.Context) *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol champ select v1 toggle favorite by champion Id by position params
func (o *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol champ select v1 toggle favorite by champion Id by position params
func (o *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams) WithHTTPClient(client *http.Client) *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol champ select v1 toggle favorite by champion Id by position params
func (o *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChampionID adds the championID to the post lol champ select v1 toggle favorite by champion Id by position params
func (o *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams) WithChampionID(championID int64) *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams {
	o.SetChampionID(championID)
	return o
}

// SetChampionID adds the championId to the post lol champ select v1 toggle favorite by champion Id by position params
func (o *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams) SetChampionID(championID int64) {
	o.ChampionID = championID
}

// WithPosition adds the position to the post lol champ select v1 toggle favorite by champion Id by position params
func (o *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams) WithPosition(position string) *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams {
	o.SetPosition(position)
	return o
}

// SetPosition adds the position to the post lol champ select v1 toggle favorite by champion Id by position params
func (o *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams) SetPosition(position string) {
	o.Position = position
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolChampSelectV1ToggleFavoriteByChampionIDByPositionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param championId
	if err := r.SetPathParam("championId", swag.FormatInt64(o.ChampionID)); err != nil {
		return err
	}

	// path param position
	if err := r.SetPathParam("position", o.Position); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
