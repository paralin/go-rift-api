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

// NewPostLolLobbyV2MatchmakingQuickSearchParams creates a new PostLolLobbyV2MatchmakingQuickSearchParams object
// with the default values initialized.
func NewPostLolLobbyV2MatchmakingQuickSearchParams() *PostLolLobbyV2MatchmakingQuickSearchParams {
	var ()
	return &PostLolLobbyV2MatchmakingQuickSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolLobbyV2MatchmakingQuickSearchParamsWithTimeout creates a new PostLolLobbyV2MatchmakingQuickSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolLobbyV2MatchmakingQuickSearchParamsWithTimeout(timeout time.Duration) *PostLolLobbyV2MatchmakingQuickSearchParams {
	var ()
	return &PostLolLobbyV2MatchmakingQuickSearchParams{

		timeout: timeout,
	}
}

// NewPostLolLobbyV2MatchmakingQuickSearchParamsWithContext creates a new PostLolLobbyV2MatchmakingQuickSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolLobbyV2MatchmakingQuickSearchParamsWithContext(ctx context.Context) *PostLolLobbyV2MatchmakingQuickSearchParams {
	var ()
	return &PostLolLobbyV2MatchmakingQuickSearchParams{

		Context: ctx,
	}
}

// NewPostLolLobbyV2MatchmakingQuickSearchParamsWithHTTPClient creates a new PostLolLobbyV2MatchmakingQuickSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolLobbyV2MatchmakingQuickSearchParamsWithHTTPClient(client *http.Client) *PostLolLobbyV2MatchmakingQuickSearchParams {
	var ()
	return &PostLolLobbyV2MatchmakingQuickSearchParams{
		HTTPClient: client,
	}
}

/*PostLolLobbyV2MatchmakingQuickSearchParams contains all the parameters to send to the API endpoint
for the post lol lobby v2 matchmaking quick search operation typically these are written to a http.Request
*/
type PostLolLobbyV2MatchmakingQuickSearchParams struct {

	/*LobbyChange*/
	LobbyChange *models.LolLobbyLobbyChangeGameDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol lobby v2 matchmaking quick search params
func (o *PostLolLobbyV2MatchmakingQuickSearchParams) WithTimeout(timeout time.Duration) *PostLolLobbyV2MatchmakingQuickSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol lobby v2 matchmaking quick search params
func (o *PostLolLobbyV2MatchmakingQuickSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol lobby v2 matchmaking quick search params
func (o *PostLolLobbyV2MatchmakingQuickSearchParams) WithContext(ctx context.Context) *PostLolLobbyV2MatchmakingQuickSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol lobby v2 matchmaking quick search params
func (o *PostLolLobbyV2MatchmakingQuickSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol lobby v2 matchmaking quick search params
func (o *PostLolLobbyV2MatchmakingQuickSearchParams) WithHTTPClient(client *http.Client) *PostLolLobbyV2MatchmakingQuickSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol lobby v2 matchmaking quick search params
func (o *PostLolLobbyV2MatchmakingQuickSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLobbyChange adds the lobbyChange to the post lol lobby v2 matchmaking quick search params
func (o *PostLolLobbyV2MatchmakingQuickSearchParams) WithLobbyChange(lobbyChange *models.LolLobbyLobbyChangeGameDto) *PostLolLobbyV2MatchmakingQuickSearchParams {
	o.SetLobbyChange(lobbyChange)
	return o
}

// SetLobbyChange adds the lobbyChange to the post lol lobby v2 matchmaking quick search params
func (o *PostLolLobbyV2MatchmakingQuickSearchParams) SetLobbyChange(lobbyChange *models.LolLobbyLobbyChangeGameDto) {
	o.LobbyChange = lobbyChange
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolLobbyV2MatchmakingQuickSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LobbyChange != nil {
		if err := r.SetBodyParam(o.LobbyChange); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}