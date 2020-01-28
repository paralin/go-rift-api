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

// NewPutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams creates a new PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams object
// with the default values initialized.
func NewPutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams() *PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams {
	var ()
	return &PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParamsWithTimeout creates a new PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParamsWithTimeout(timeout time.Duration) *PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams {
	var ()
	return &PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams{

		timeout: timeout,
	}
}

// NewPutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParamsWithContext creates a new PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParamsWithContext(ctx context.Context) *PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams {
	var ()
	return &PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams{

		Context: ctx,
	}
}

// NewPutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParamsWithHTTPClient creates a new PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParamsWithHTTPClient(client *http.Client) *PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams {
	var ()
	return &PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams{
		HTTPClient: client,
	}
}

/*PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams contains all the parameters to send to the API endpoint
for the put lol lobby v2 lobby members local member position preferences operation typically these are written to a http.Request
*/
type PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams struct {

	/*PositionPreferences*/
	PositionPreferences *models.LolLobbyLobbyPositionPreferences

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put lol lobby v2 lobby members local member position preferences params
func (o *PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams) WithTimeout(timeout time.Duration) *PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put lol lobby v2 lobby members local member position preferences params
func (o *PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put lol lobby v2 lobby members local member position preferences params
func (o *PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams) WithContext(ctx context.Context) *PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put lol lobby v2 lobby members local member position preferences params
func (o *PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put lol lobby v2 lobby members local member position preferences params
func (o *PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams) WithHTTPClient(client *http.Client) *PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put lol lobby v2 lobby members local member position preferences params
func (o *PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPositionPreferences adds the positionPreferences to the put lol lobby v2 lobby members local member position preferences params
func (o *PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams) WithPositionPreferences(positionPreferences *models.LolLobbyLobbyPositionPreferences) *PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams {
	o.SetPositionPreferences(positionPreferences)
	return o
}

// SetPositionPreferences adds the positionPreferences to the put lol lobby v2 lobby members local member position preferences params
func (o *PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams) SetPositionPreferences(positionPreferences *models.LolLobbyLobbyPositionPreferences) {
	o.PositionPreferences = positionPreferences
}

// WriteToRequest writes these params to a swagger request
func (o *PutLolLobbyV2LobbyMembersLocalMemberPositionPreferencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PositionPreferences != nil {
		if err := r.SetBodyParam(o.PositionPreferences); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
