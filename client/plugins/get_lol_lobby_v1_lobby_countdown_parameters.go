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

// NewGetLolLobbyV1LobbyCountdownParams creates a new GetLolLobbyV1LobbyCountdownParams object
// with the default values initialized.
func NewGetLolLobbyV1LobbyCountdownParams() *GetLolLobbyV1LobbyCountdownParams {

	return &GetLolLobbyV1LobbyCountdownParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolLobbyV1LobbyCountdownParamsWithTimeout creates a new GetLolLobbyV1LobbyCountdownParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolLobbyV1LobbyCountdownParamsWithTimeout(timeout time.Duration) *GetLolLobbyV1LobbyCountdownParams {

	return &GetLolLobbyV1LobbyCountdownParams{

		timeout: timeout,
	}
}

// NewGetLolLobbyV1LobbyCountdownParamsWithContext creates a new GetLolLobbyV1LobbyCountdownParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolLobbyV1LobbyCountdownParamsWithContext(ctx context.Context) *GetLolLobbyV1LobbyCountdownParams {

	return &GetLolLobbyV1LobbyCountdownParams{

		Context: ctx,
	}
}

// NewGetLolLobbyV1LobbyCountdownParamsWithHTTPClient creates a new GetLolLobbyV1LobbyCountdownParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolLobbyV1LobbyCountdownParamsWithHTTPClient(client *http.Client) *GetLolLobbyV1LobbyCountdownParams {

	return &GetLolLobbyV1LobbyCountdownParams{
		HTTPClient: client,
	}
}

/*GetLolLobbyV1LobbyCountdownParams contains all the parameters to send to the API endpoint
for the get lol lobby v1 lobby countdown operation typically these are written to a http.Request
*/
type GetLolLobbyV1LobbyCountdownParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol lobby v1 lobby countdown params
func (o *GetLolLobbyV1LobbyCountdownParams) WithTimeout(timeout time.Duration) *GetLolLobbyV1LobbyCountdownParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol lobby v1 lobby countdown params
func (o *GetLolLobbyV1LobbyCountdownParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol lobby v1 lobby countdown params
func (o *GetLolLobbyV1LobbyCountdownParams) WithContext(ctx context.Context) *GetLolLobbyV1LobbyCountdownParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol lobby v1 lobby countdown params
func (o *GetLolLobbyV1LobbyCountdownParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol lobby v1 lobby countdown params
func (o *GetLolLobbyV1LobbyCountdownParams) WithHTTPClient(client *http.Client) *GetLolLobbyV1LobbyCountdownParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol lobby v1 lobby countdown params
func (o *GetLolLobbyV1LobbyCountdownParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolLobbyV1LobbyCountdownParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
