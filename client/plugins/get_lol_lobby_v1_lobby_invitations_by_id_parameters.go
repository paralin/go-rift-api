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
)

// NewGetLolLobbyV1LobbyInvitationsByIDParams creates a new GetLolLobbyV1LobbyInvitationsByIDParams object
// with the default values initialized.
func NewGetLolLobbyV1LobbyInvitationsByIDParams() *GetLolLobbyV1LobbyInvitationsByIDParams {
	var ()
	return &GetLolLobbyV1LobbyInvitationsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolLobbyV1LobbyInvitationsByIDParamsWithTimeout creates a new GetLolLobbyV1LobbyInvitationsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolLobbyV1LobbyInvitationsByIDParamsWithTimeout(timeout time.Duration) *GetLolLobbyV1LobbyInvitationsByIDParams {
	var ()
	return &GetLolLobbyV1LobbyInvitationsByIDParams{

		timeout: timeout,
	}
}

// NewGetLolLobbyV1LobbyInvitationsByIDParamsWithContext creates a new GetLolLobbyV1LobbyInvitationsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolLobbyV1LobbyInvitationsByIDParamsWithContext(ctx context.Context) *GetLolLobbyV1LobbyInvitationsByIDParams {
	var ()
	return &GetLolLobbyV1LobbyInvitationsByIDParams{

		Context: ctx,
	}
}

// NewGetLolLobbyV1LobbyInvitationsByIDParamsWithHTTPClient creates a new GetLolLobbyV1LobbyInvitationsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolLobbyV1LobbyInvitationsByIDParamsWithHTTPClient(client *http.Client) *GetLolLobbyV1LobbyInvitationsByIDParams {
	var ()
	return &GetLolLobbyV1LobbyInvitationsByIDParams{
		HTTPClient: client,
	}
}

/*GetLolLobbyV1LobbyInvitationsByIDParams contains all the parameters to send to the API endpoint
for the get lol lobby v1 lobby invitations by Id operation typically these are written to a http.Request
*/
type GetLolLobbyV1LobbyInvitationsByIDParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol lobby v1 lobby invitations by Id params
func (o *GetLolLobbyV1LobbyInvitationsByIDParams) WithTimeout(timeout time.Duration) *GetLolLobbyV1LobbyInvitationsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol lobby v1 lobby invitations by Id params
func (o *GetLolLobbyV1LobbyInvitationsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol lobby v1 lobby invitations by Id params
func (o *GetLolLobbyV1LobbyInvitationsByIDParams) WithContext(ctx context.Context) *GetLolLobbyV1LobbyInvitationsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol lobby v1 lobby invitations by Id params
func (o *GetLolLobbyV1LobbyInvitationsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol lobby v1 lobby invitations by Id params
func (o *GetLolLobbyV1LobbyInvitationsByIDParams) WithHTTPClient(client *http.Client) *GetLolLobbyV1LobbyInvitationsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol lobby v1 lobby invitations by Id params
func (o *GetLolLobbyV1LobbyInvitationsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get lol lobby v1 lobby invitations by Id params
func (o *GetLolLobbyV1LobbyInvitationsByIDParams) WithID(id string) *GetLolLobbyV1LobbyInvitationsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get lol lobby v1 lobby invitations by Id params
func (o *GetLolLobbyV1LobbyInvitationsByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolLobbyV1LobbyInvitationsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
