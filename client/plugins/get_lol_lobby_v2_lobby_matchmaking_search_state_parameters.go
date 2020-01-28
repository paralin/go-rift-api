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

// NewGetLolLobbyV2LobbyMatchmakingSearchStateParams creates a new GetLolLobbyV2LobbyMatchmakingSearchStateParams object
// with the default values initialized.
func NewGetLolLobbyV2LobbyMatchmakingSearchStateParams() *GetLolLobbyV2LobbyMatchmakingSearchStateParams {

	return &GetLolLobbyV2LobbyMatchmakingSearchStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolLobbyV2LobbyMatchmakingSearchStateParamsWithTimeout creates a new GetLolLobbyV2LobbyMatchmakingSearchStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolLobbyV2LobbyMatchmakingSearchStateParamsWithTimeout(timeout time.Duration) *GetLolLobbyV2LobbyMatchmakingSearchStateParams {

	return &GetLolLobbyV2LobbyMatchmakingSearchStateParams{

		timeout: timeout,
	}
}

// NewGetLolLobbyV2LobbyMatchmakingSearchStateParamsWithContext creates a new GetLolLobbyV2LobbyMatchmakingSearchStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolLobbyV2LobbyMatchmakingSearchStateParamsWithContext(ctx context.Context) *GetLolLobbyV2LobbyMatchmakingSearchStateParams {

	return &GetLolLobbyV2LobbyMatchmakingSearchStateParams{

		Context: ctx,
	}
}

// NewGetLolLobbyV2LobbyMatchmakingSearchStateParamsWithHTTPClient creates a new GetLolLobbyV2LobbyMatchmakingSearchStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolLobbyV2LobbyMatchmakingSearchStateParamsWithHTTPClient(client *http.Client) *GetLolLobbyV2LobbyMatchmakingSearchStateParams {

	return &GetLolLobbyV2LobbyMatchmakingSearchStateParams{
		HTTPClient: client,
	}
}

/*GetLolLobbyV2LobbyMatchmakingSearchStateParams contains all the parameters to send to the API endpoint
for the get lol lobby v2 lobby matchmaking search state operation typically these are written to a http.Request
*/
type GetLolLobbyV2LobbyMatchmakingSearchStateParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol lobby v2 lobby matchmaking search state params
func (o *GetLolLobbyV2LobbyMatchmakingSearchStateParams) WithTimeout(timeout time.Duration) *GetLolLobbyV2LobbyMatchmakingSearchStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol lobby v2 lobby matchmaking search state params
func (o *GetLolLobbyV2LobbyMatchmakingSearchStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol lobby v2 lobby matchmaking search state params
func (o *GetLolLobbyV2LobbyMatchmakingSearchStateParams) WithContext(ctx context.Context) *GetLolLobbyV2LobbyMatchmakingSearchStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol lobby v2 lobby matchmaking search state params
func (o *GetLolLobbyV2LobbyMatchmakingSearchStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol lobby v2 lobby matchmaking search state params
func (o *GetLolLobbyV2LobbyMatchmakingSearchStateParams) WithHTTPClient(client *http.Client) *GetLolLobbyV2LobbyMatchmakingSearchStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol lobby v2 lobby matchmaking search state params
func (o *GetLolLobbyV2LobbyMatchmakingSearchStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolLobbyV2LobbyMatchmakingSearchStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
