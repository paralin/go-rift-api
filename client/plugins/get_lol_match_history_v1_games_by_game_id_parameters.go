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

// NewGetLolMatchHistoryV1GamesByGameIDParams creates a new GetLolMatchHistoryV1GamesByGameIDParams object
// with the default values initialized.
func NewGetLolMatchHistoryV1GamesByGameIDParams() *GetLolMatchHistoryV1GamesByGameIDParams {
	var ()
	return &GetLolMatchHistoryV1GamesByGameIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolMatchHistoryV1GamesByGameIDParamsWithTimeout creates a new GetLolMatchHistoryV1GamesByGameIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolMatchHistoryV1GamesByGameIDParamsWithTimeout(timeout time.Duration) *GetLolMatchHistoryV1GamesByGameIDParams {
	var ()
	return &GetLolMatchHistoryV1GamesByGameIDParams{

		timeout: timeout,
	}
}

// NewGetLolMatchHistoryV1GamesByGameIDParamsWithContext creates a new GetLolMatchHistoryV1GamesByGameIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolMatchHistoryV1GamesByGameIDParamsWithContext(ctx context.Context) *GetLolMatchHistoryV1GamesByGameIDParams {
	var ()
	return &GetLolMatchHistoryV1GamesByGameIDParams{

		Context: ctx,
	}
}

// NewGetLolMatchHistoryV1GamesByGameIDParamsWithHTTPClient creates a new GetLolMatchHistoryV1GamesByGameIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolMatchHistoryV1GamesByGameIDParamsWithHTTPClient(client *http.Client) *GetLolMatchHistoryV1GamesByGameIDParams {
	var ()
	return &GetLolMatchHistoryV1GamesByGameIDParams{
		HTTPClient: client,
	}
}

/*GetLolMatchHistoryV1GamesByGameIDParams contains all the parameters to send to the API endpoint
for the get lol match history v1 games by game Id operation typically these are written to a http.Request
*/
type GetLolMatchHistoryV1GamesByGameIDParams struct {

	/*GameID*/
	GameID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol match history v1 games by game Id params
func (o *GetLolMatchHistoryV1GamesByGameIDParams) WithTimeout(timeout time.Duration) *GetLolMatchHistoryV1GamesByGameIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol match history v1 games by game Id params
func (o *GetLolMatchHistoryV1GamesByGameIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol match history v1 games by game Id params
func (o *GetLolMatchHistoryV1GamesByGameIDParams) WithContext(ctx context.Context) *GetLolMatchHistoryV1GamesByGameIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol match history v1 games by game Id params
func (o *GetLolMatchHistoryV1GamesByGameIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol match history v1 games by game Id params
func (o *GetLolMatchHistoryV1GamesByGameIDParams) WithHTTPClient(client *http.Client) *GetLolMatchHistoryV1GamesByGameIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol match history v1 games by game Id params
func (o *GetLolMatchHistoryV1GamesByGameIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGameID adds the gameID to the get lol match history v1 games by game Id params
func (o *GetLolMatchHistoryV1GamesByGameIDParams) WithGameID(gameID int64) *GetLolMatchHistoryV1GamesByGameIDParams {
	o.SetGameID(gameID)
	return o
}

// SetGameID adds the gameId to the get lol match history v1 games by game Id params
func (o *GetLolMatchHistoryV1GamesByGameIDParams) SetGameID(gameID int64) {
	o.GameID = gameID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolMatchHistoryV1GamesByGameIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gameId
	if err := r.SetPathParam("gameId", swag.FormatInt64(o.GameID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
