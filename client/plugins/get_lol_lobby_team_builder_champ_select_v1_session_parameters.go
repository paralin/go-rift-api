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

// NewGetLolLobbyTeamBuilderChampSelectV1SessionParams creates a new GetLolLobbyTeamBuilderChampSelectV1SessionParams object
// with the default values initialized.
func NewGetLolLobbyTeamBuilderChampSelectV1SessionParams() *GetLolLobbyTeamBuilderChampSelectV1SessionParams {

	return &GetLolLobbyTeamBuilderChampSelectV1SessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolLobbyTeamBuilderChampSelectV1SessionParamsWithTimeout creates a new GetLolLobbyTeamBuilderChampSelectV1SessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolLobbyTeamBuilderChampSelectV1SessionParamsWithTimeout(timeout time.Duration) *GetLolLobbyTeamBuilderChampSelectV1SessionParams {

	return &GetLolLobbyTeamBuilderChampSelectV1SessionParams{

		timeout: timeout,
	}
}

// NewGetLolLobbyTeamBuilderChampSelectV1SessionParamsWithContext creates a new GetLolLobbyTeamBuilderChampSelectV1SessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolLobbyTeamBuilderChampSelectV1SessionParamsWithContext(ctx context.Context) *GetLolLobbyTeamBuilderChampSelectV1SessionParams {

	return &GetLolLobbyTeamBuilderChampSelectV1SessionParams{

		Context: ctx,
	}
}

// NewGetLolLobbyTeamBuilderChampSelectV1SessionParamsWithHTTPClient creates a new GetLolLobbyTeamBuilderChampSelectV1SessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolLobbyTeamBuilderChampSelectV1SessionParamsWithHTTPClient(client *http.Client) *GetLolLobbyTeamBuilderChampSelectV1SessionParams {

	return &GetLolLobbyTeamBuilderChampSelectV1SessionParams{
		HTTPClient: client,
	}
}

/*GetLolLobbyTeamBuilderChampSelectV1SessionParams contains all the parameters to send to the API endpoint
for the get lol lobby team builder champ select v1 session operation typically these are written to a http.Request
*/
type GetLolLobbyTeamBuilderChampSelectV1SessionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol lobby team builder champ select v1 session params
func (o *GetLolLobbyTeamBuilderChampSelectV1SessionParams) WithTimeout(timeout time.Duration) *GetLolLobbyTeamBuilderChampSelectV1SessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol lobby team builder champ select v1 session params
func (o *GetLolLobbyTeamBuilderChampSelectV1SessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol lobby team builder champ select v1 session params
func (o *GetLolLobbyTeamBuilderChampSelectV1SessionParams) WithContext(ctx context.Context) *GetLolLobbyTeamBuilderChampSelectV1SessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol lobby team builder champ select v1 session params
func (o *GetLolLobbyTeamBuilderChampSelectV1SessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol lobby team builder champ select v1 session params
func (o *GetLolLobbyTeamBuilderChampSelectV1SessionParams) WithHTTPClient(client *http.Client) *GetLolLobbyTeamBuilderChampSelectV1SessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol lobby team builder champ select v1 session params
func (o *GetLolLobbyTeamBuilderChampSelectV1SessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolLobbyTeamBuilderChampSelectV1SessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
