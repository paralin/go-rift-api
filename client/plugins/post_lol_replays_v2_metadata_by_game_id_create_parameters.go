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

// NewPostLolReplaysV2MetadataByGameIDCreateParams creates a new PostLolReplaysV2MetadataByGameIDCreateParams object
// with the default values initialized.
func NewPostLolReplaysV2MetadataByGameIDCreateParams() *PostLolReplaysV2MetadataByGameIDCreateParams {
	var ()
	return &PostLolReplaysV2MetadataByGameIDCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolReplaysV2MetadataByGameIDCreateParamsWithTimeout creates a new PostLolReplaysV2MetadataByGameIDCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolReplaysV2MetadataByGameIDCreateParamsWithTimeout(timeout time.Duration) *PostLolReplaysV2MetadataByGameIDCreateParams {
	var ()
	return &PostLolReplaysV2MetadataByGameIDCreateParams{

		timeout: timeout,
	}
}

// NewPostLolReplaysV2MetadataByGameIDCreateParamsWithContext creates a new PostLolReplaysV2MetadataByGameIDCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolReplaysV2MetadataByGameIDCreateParamsWithContext(ctx context.Context) *PostLolReplaysV2MetadataByGameIDCreateParams {
	var ()
	return &PostLolReplaysV2MetadataByGameIDCreateParams{

		Context: ctx,
	}
}

// NewPostLolReplaysV2MetadataByGameIDCreateParamsWithHTTPClient creates a new PostLolReplaysV2MetadataByGameIDCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolReplaysV2MetadataByGameIDCreateParamsWithHTTPClient(client *http.Client) *PostLolReplaysV2MetadataByGameIDCreateParams {
	var ()
	return &PostLolReplaysV2MetadataByGameIDCreateParams{
		HTTPClient: client,
	}
}

/*PostLolReplaysV2MetadataByGameIDCreateParams contains all the parameters to send to the API endpoint
for the post lol replays v2 metadata by game Id create operation typically these are written to a http.Request
*/
type PostLolReplaysV2MetadataByGameIDCreateParams struct {

	/*GameID*/
	GameID int64
	/*Request*/
	Request *models.LolReplaysReplayCreateMetadata

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol replays v2 metadata by game Id create params
func (o *PostLolReplaysV2MetadataByGameIDCreateParams) WithTimeout(timeout time.Duration) *PostLolReplaysV2MetadataByGameIDCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol replays v2 metadata by game Id create params
func (o *PostLolReplaysV2MetadataByGameIDCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol replays v2 metadata by game Id create params
func (o *PostLolReplaysV2MetadataByGameIDCreateParams) WithContext(ctx context.Context) *PostLolReplaysV2MetadataByGameIDCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol replays v2 metadata by game Id create params
func (o *PostLolReplaysV2MetadataByGameIDCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol replays v2 metadata by game Id create params
func (o *PostLolReplaysV2MetadataByGameIDCreateParams) WithHTTPClient(client *http.Client) *PostLolReplaysV2MetadataByGameIDCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol replays v2 metadata by game Id create params
func (o *PostLolReplaysV2MetadataByGameIDCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGameID adds the gameID to the post lol replays v2 metadata by game Id create params
func (o *PostLolReplaysV2MetadataByGameIDCreateParams) WithGameID(gameID int64) *PostLolReplaysV2MetadataByGameIDCreateParams {
	o.SetGameID(gameID)
	return o
}

// SetGameID adds the gameId to the post lol replays v2 metadata by game Id create params
func (o *PostLolReplaysV2MetadataByGameIDCreateParams) SetGameID(gameID int64) {
	o.GameID = gameID
}

// WithRequest adds the request to the post lol replays v2 metadata by game Id create params
func (o *PostLolReplaysV2MetadataByGameIDCreateParams) WithRequest(request *models.LolReplaysReplayCreateMetadata) *PostLolReplaysV2MetadataByGameIDCreateParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post lol replays v2 metadata by game Id create params
func (o *PostLolReplaysV2MetadataByGameIDCreateParams) SetRequest(request *models.LolReplaysReplayCreateMetadata) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolReplaysV2MetadataByGameIDCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gameId
	if err := r.SetPathParam("gameId", swag.FormatInt64(o.GameID)); err != nil {
		return err
	}

	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}