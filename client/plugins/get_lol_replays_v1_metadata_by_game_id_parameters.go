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

// NewGetLolReplaysV1MetadataByGameIDParams creates a new GetLolReplaysV1MetadataByGameIDParams object
// with the default values initialized.
func NewGetLolReplaysV1MetadataByGameIDParams() *GetLolReplaysV1MetadataByGameIDParams {
	var ()
	return &GetLolReplaysV1MetadataByGameIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolReplaysV1MetadataByGameIDParamsWithTimeout creates a new GetLolReplaysV1MetadataByGameIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolReplaysV1MetadataByGameIDParamsWithTimeout(timeout time.Duration) *GetLolReplaysV1MetadataByGameIDParams {
	var ()
	return &GetLolReplaysV1MetadataByGameIDParams{

		timeout: timeout,
	}
}

// NewGetLolReplaysV1MetadataByGameIDParamsWithContext creates a new GetLolReplaysV1MetadataByGameIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolReplaysV1MetadataByGameIDParamsWithContext(ctx context.Context) *GetLolReplaysV1MetadataByGameIDParams {
	var ()
	return &GetLolReplaysV1MetadataByGameIDParams{

		Context: ctx,
	}
}

// NewGetLolReplaysV1MetadataByGameIDParamsWithHTTPClient creates a new GetLolReplaysV1MetadataByGameIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolReplaysV1MetadataByGameIDParamsWithHTTPClient(client *http.Client) *GetLolReplaysV1MetadataByGameIDParams {
	var ()
	return &GetLolReplaysV1MetadataByGameIDParams{
		HTTPClient: client,
	}
}

/*GetLolReplaysV1MetadataByGameIDParams contains all the parameters to send to the API endpoint
for the get lol replays v1 metadata by game Id operation typically these are written to a http.Request
*/
type GetLolReplaysV1MetadataByGameIDParams struct {

	/*GameID*/
	GameID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol replays v1 metadata by game Id params
func (o *GetLolReplaysV1MetadataByGameIDParams) WithTimeout(timeout time.Duration) *GetLolReplaysV1MetadataByGameIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol replays v1 metadata by game Id params
func (o *GetLolReplaysV1MetadataByGameIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol replays v1 metadata by game Id params
func (o *GetLolReplaysV1MetadataByGameIDParams) WithContext(ctx context.Context) *GetLolReplaysV1MetadataByGameIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol replays v1 metadata by game Id params
func (o *GetLolReplaysV1MetadataByGameIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol replays v1 metadata by game Id params
func (o *GetLolReplaysV1MetadataByGameIDParams) WithHTTPClient(client *http.Client) *GetLolReplaysV1MetadataByGameIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol replays v1 metadata by game Id params
func (o *GetLolReplaysV1MetadataByGameIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGameID adds the gameID to the get lol replays v1 metadata by game Id params
func (o *GetLolReplaysV1MetadataByGameIDParams) WithGameID(gameID int64) *GetLolReplaysV1MetadataByGameIDParams {
	o.SetGameID(gameID)
	return o
}

// SetGameID adds the gameId to the get lol replays v1 metadata by game Id params
func (o *GetLolReplaysV1MetadataByGameIDParams) SetGameID(gameID int64) {
	o.GameID = gameID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolReplaysV1MetadataByGameIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
