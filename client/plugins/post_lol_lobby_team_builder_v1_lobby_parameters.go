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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/paralin/go-rift-api/models"
)

// NewPostLolLobbyTeamBuilderV1LobbyParams creates a new PostLolLobbyTeamBuilderV1LobbyParams object
// with the default values initialized.
func NewPostLolLobbyTeamBuilderV1LobbyParams() *PostLolLobbyTeamBuilderV1LobbyParams {
	var ()
	return &PostLolLobbyTeamBuilderV1LobbyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolLobbyTeamBuilderV1LobbyParamsWithTimeout creates a new PostLolLobbyTeamBuilderV1LobbyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolLobbyTeamBuilderV1LobbyParamsWithTimeout(timeout time.Duration) *PostLolLobbyTeamBuilderV1LobbyParams {
	var ()
	return &PostLolLobbyTeamBuilderV1LobbyParams{

		timeout: timeout,
	}
}

// NewPostLolLobbyTeamBuilderV1LobbyParamsWithContext creates a new PostLolLobbyTeamBuilderV1LobbyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolLobbyTeamBuilderV1LobbyParamsWithContext(ctx context.Context) *PostLolLobbyTeamBuilderV1LobbyParams {
	var ()
	return &PostLolLobbyTeamBuilderV1LobbyParams{

		Context: ctx,
	}
}

// NewPostLolLobbyTeamBuilderV1LobbyParamsWithHTTPClient creates a new PostLolLobbyTeamBuilderV1LobbyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolLobbyTeamBuilderV1LobbyParamsWithHTTPClient(client *http.Client) *PostLolLobbyTeamBuilderV1LobbyParams {
	var ()
	return &PostLolLobbyTeamBuilderV1LobbyParams{
		HTTPClient: client,
	}
}

/*PostLolLobbyTeamBuilderV1LobbyParams contains all the parameters to send to the API endpoint
for the post lol lobby team builder v1 lobby operation typically these are written to a http.Request
*/
type PostLolLobbyTeamBuilderV1LobbyParams struct {

	/*Lobby*/
	Lobby *models.LolLobbyTeamBuilderLobby

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol lobby team builder v1 lobby params
func (o *PostLolLobbyTeamBuilderV1LobbyParams) WithTimeout(timeout time.Duration) *PostLolLobbyTeamBuilderV1LobbyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol lobby team builder v1 lobby params
func (o *PostLolLobbyTeamBuilderV1LobbyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol lobby team builder v1 lobby params
func (o *PostLolLobbyTeamBuilderV1LobbyParams) WithContext(ctx context.Context) *PostLolLobbyTeamBuilderV1LobbyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol lobby team builder v1 lobby params
func (o *PostLolLobbyTeamBuilderV1LobbyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol lobby team builder v1 lobby params
func (o *PostLolLobbyTeamBuilderV1LobbyParams) WithHTTPClient(client *http.Client) *PostLolLobbyTeamBuilderV1LobbyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol lobby team builder v1 lobby params
func (o *PostLolLobbyTeamBuilderV1LobbyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLobby adds the lobby to the post lol lobby team builder v1 lobby params
func (o *PostLolLobbyTeamBuilderV1LobbyParams) WithLobby(lobby *models.LolLobbyTeamBuilderLobby) *PostLolLobbyTeamBuilderV1LobbyParams {
	o.SetLobby(lobby)
	return o
}

// SetLobby adds the lobby to the post lol lobby team builder v1 lobby params
func (o *PostLolLobbyTeamBuilderV1LobbyParams) SetLobby(lobby *models.LolLobbyTeamBuilderLobby) {
	o.Lobby = lobby
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolLobbyTeamBuilderV1LobbyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Lobby != nil {
		if err := r.SetBodyParam(o.Lobby); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
