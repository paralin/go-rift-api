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

// NewPostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams creates a new PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams object
// with the default values initialized.
func NewPostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams() *PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams {
	var ()
	return &PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParamsWithTimeout creates a new PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParamsWithTimeout(timeout time.Duration) *PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams {
	var ()
	return &PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams{

		timeout: timeout,
	}
}

// NewPostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParamsWithContext creates a new PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParamsWithContext(ctx context.Context) *PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams {
	var ()
	return &PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams{

		Context: ctx,
	}
}

// NewPostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParamsWithHTTPClient creates a new PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParamsWithHTTPClient(client *http.Client) *PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams {
	var ()
	return &PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams{
		HTTPClient: client,
	}
}

/*PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams contains all the parameters to send to the API endpoint
for the post lol lobby v2 received invitations by invitation Id accept operation typically these are written to a http.Request
*/
type PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams struct {

	/*InvitationID*/
	InvitationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol lobby v2 received invitations by invitation Id accept params
func (o *PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams) WithTimeout(timeout time.Duration) *PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol lobby v2 received invitations by invitation Id accept params
func (o *PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol lobby v2 received invitations by invitation Id accept params
func (o *PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams) WithContext(ctx context.Context) *PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol lobby v2 received invitations by invitation Id accept params
func (o *PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol lobby v2 received invitations by invitation Id accept params
func (o *PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams) WithHTTPClient(client *http.Client) *PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol lobby v2 received invitations by invitation Id accept params
func (o *PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvitationID adds the invitationID to the post lol lobby v2 received invitations by invitation Id accept params
func (o *PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams) WithInvitationID(invitationID string) *PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams {
	o.SetInvitationID(invitationID)
	return o
}

// SetInvitationID adds the invitationId to the post lol lobby v2 received invitations by invitation Id accept params
func (o *PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams) SetInvitationID(invitationID string) {
	o.InvitationID = invitationID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolLobbyV2ReceivedInvitationsByInvitationIDAcceptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param invitationId
	if err := r.SetPathParam("invitationId", o.InvitationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
