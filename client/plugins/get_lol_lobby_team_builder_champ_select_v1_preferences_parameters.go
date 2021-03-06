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

// NewGetLolLobbyTeamBuilderChampSelectV1PreferencesParams creates a new GetLolLobbyTeamBuilderChampSelectV1PreferencesParams object
// with the default values initialized.
func NewGetLolLobbyTeamBuilderChampSelectV1PreferencesParams() *GetLolLobbyTeamBuilderChampSelectV1PreferencesParams {

	return &GetLolLobbyTeamBuilderChampSelectV1PreferencesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolLobbyTeamBuilderChampSelectV1PreferencesParamsWithTimeout creates a new GetLolLobbyTeamBuilderChampSelectV1PreferencesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolLobbyTeamBuilderChampSelectV1PreferencesParamsWithTimeout(timeout time.Duration) *GetLolLobbyTeamBuilderChampSelectV1PreferencesParams {

	return &GetLolLobbyTeamBuilderChampSelectV1PreferencesParams{

		timeout: timeout,
	}
}

// NewGetLolLobbyTeamBuilderChampSelectV1PreferencesParamsWithContext creates a new GetLolLobbyTeamBuilderChampSelectV1PreferencesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolLobbyTeamBuilderChampSelectV1PreferencesParamsWithContext(ctx context.Context) *GetLolLobbyTeamBuilderChampSelectV1PreferencesParams {

	return &GetLolLobbyTeamBuilderChampSelectV1PreferencesParams{

		Context: ctx,
	}
}

// NewGetLolLobbyTeamBuilderChampSelectV1PreferencesParamsWithHTTPClient creates a new GetLolLobbyTeamBuilderChampSelectV1PreferencesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolLobbyTeamBuilderChampSelectV1PreferencesParamsWithHTTPClient(client *http.Client) *GetLolLobbyTeamBuilderChampSelectV1PreferencesParams {

	return &GetLolLobbyTeamBuilderChampSelectV1PreferencesParams{
		HTTPClient: client,
	}
}

/*GetLolLobbyTeamBuilderChampSelectV1PreferencesParams contains all the parameters to send to the API endpoint
for the get lol lobby team builder champ select v1 preferences operation typically these are written to a http.Request
*/
type GetLolLobbyTeamBuilderChampSelectV1PreferencesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol lobby team builder champ select v1 preferences params
func (o *GetLolLobbyTeamBuilderChampSelectV1PreferencesParams) WithTimeout(timeout time.Duration) *GetLolLobbyTeamBuilderChampSelectV1PreferencesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol lobby team builder champ select v1 preferences params
func (o *GetLolLobbyTeamBuilderChampSelectV1PreferencesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol lobby team builder champ select v1 preferences params
func (o *GetLolLobbyTeamBuilderChampSelectV1PreferencesParams) WithContext(ctx context.Context) *GetLolLobbyTeamBuilderChampSelectV1PreferencesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol lobby team builder champ select v1 preferences params
func (o *GetLolLobbyTeamBuilderChampSelectV1PreferencesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol lobby team builder champ select v1 preferences params
func (o *GetLolLobbyTeamBuilderChampSelectV1PreferencesParams) WithHTTPClient(client *http.Client) *GetLolLobbyTeamBuilderChampSelectV1PreferencesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol lobby team builder champ select v1 preferences params
func (o *GetLolLobbyTeamBuilderChampSelectV1PreferencesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolLobbyTeamBuilderChampSelectV1PreferencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
