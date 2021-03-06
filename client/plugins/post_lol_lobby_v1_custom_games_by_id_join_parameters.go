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

// NewPostLolLobbyV1CustomGamesByIDJoinParams creates a new PostLolLobbyV1CustomGamesByIDJoinParams object
// with the default values initialized.
func NewPostLolLobbyV1CustomGamesByIDJoinParams() *PostLolLobbyV1CustomGamesByIDJoinParams {
	var ()
	return &PostLolLobbyV1CustomGamesByIDJoinParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolLobbyV1CustomGamesByIDJoinParamsWithTimeout creates a new PostLolLobbyV1CustomGamesByIDJoinParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolLobbyV1CustomGamesByIDJoinParamsWithTimeout(timeout time.Duration) *PostLolLobbyV1CustomGamesByIDJoinParams {
	var ()
	return &PostLolLobbyV1CustomGamesByIDJoinParams{

		timeout: timeout,
	}
}

// NewPostLolLobbyV1CustomGamesByIDJoinParamsWithContext creates a new PostLolLobbyV1CustomGamesByIDJoinParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolLobbyV1CustomGamesByIDJoinParamsWithContext(ctx context.Context) *PostLolLobbyV1CustomGamesByIDJoinParams {
	var ()
	return &PostLolLobbyV1CustomGamesByIDJoinParams{

		Context: ctx,
	}
}

// NewPostLolLobbyV1CustomGamesByIDJoinParamsWithHTTPClient creates a new PostLolLobbyV1CustomGamesByIDJoinParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolLobbyV1CustomGamesByIDJoinParamsWithHTTPClient(client *http.Client) *PostLolLobbyV1CustomGamesByIDJoinParams {
	var ()
	return &PostLolLobbyV1CustomGamesByIDJoinParams{
		HTTPClient: client,
	}
}

/*PostLolLobbyV1CustomGamesByIDJoinParams contains all the parameters to send to the API endpoint
for the post lol lobby v1 custom games by Id join operation typically these are written to a http.Request
*/
type PostLolLobbyV1CustomGamesByIDJoinParams struct {

	/*ID*/
	ID int64
	/*Parameters*/
	Parameters *models.LolLobbyLobbyCustomJoinParameters

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol lobby v1 custom games by Id join params
func (o *PostLolLobbyV1CustomGamesByIDJoinParams) WithTimeout(timeout time.Duration) *PostLolLobbyV1CustomGamesByIDJoinParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol lobby v1 custom games by Id join params
func (o *PostLolLobbyV1CustomGamesByIDJoinParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol lobby v1 custom games by Id join params
func (o *PostLolLobbyV1CustomGamesByIDJoinParams) WithContext(ctx context.Context) *PostLolLobbyV1CustomGamesByIDJoinParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol lobby v1 custom games by Id join params
func (o *PostLolLobbyV1CustomGamesByIDJoinParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol lobby v1 custom games by Id join params
func (o *PostLolLobbyV1CustomGamesByIDJoinParams) WithHTTPClient(client *http.Client) *PostLolLobbyV1CustomGamesByIDJoinParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol lobby v1 custom games by Id join params
func (o *PostLolLobbyV1CustomGamesByIDJoinParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post lol lobby v1 custom games by Id join params
func (o *PostLolLobbyV1CustomGamesByIDJoinParams) WithID(id int64) *PostLolLobbyV1CustomGamesByIDJoinParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post lol lobby v1 custom games by Id join params
func (o *PostLolLobbyV1CustomGamesByIDJoinParams) SetID(id int64) {
	o.ID = id
}

// WithParameters adds the parameters to the post lol lobby v1 custom games by Id join params
func (o *PostLolLobbyV1CustomGamesByIDJoinParams) WithParameters(parameters *models.LolLobbyLobbyCustomJoinParameters) *PostLolLobbyV1CustomGamesByIDJoinParams {
	o.SetParameters(parameters)
	return o
}

// SetParameters adds the parameters to the post lol lobby v1 custom games by Id join params
func (o *PostLolLobbyV1CustomGamesByIDJoinParams) SetParameters(parameters *models.LolLobbyLobbyCustomJoinParameters) {
	o.Parameters = parameters
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolLobbyV1CustomGamesByIDJoinParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.Parameters != nil {
		if err := r.SetBodyParam(o.Parameters); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
