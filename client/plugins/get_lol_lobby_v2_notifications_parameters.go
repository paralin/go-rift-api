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

// NewGetLolLobbyV2NotificationsParams creates a new GetLolLobbyV2NotificationsParams object
// with the default values initialized.
func NewGetLolLobbyV2NotificationsParams() *GetLolLobbyV2NotificationsParams {

	return &GetLolLobbyV2NotificationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolLobbyV2NotificationsParamsWithTimeout creates a new GetLolLobbyV2NotificationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolLobbyV2NotificationsParamsWithTimeout(timeout time.Duration) *GetLolLobbyV2NotificationsParams {

	return &GetLolLobbyV2NotificationsParams{

		timeout: timeout,
	}
}

// NewGetLolLobbyV2NotificationsParamsWithContext creates a new GetLolLobbyV2NotificationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolLobbyV2NotificationsParamsWithContext(ctx context.Context) *GetLolLobbyV2NotificationsParams {

	return &GetLolLobbyV2NotificationsParams{

		Context: ctx,
	}
}

// NewGetLolLobbyV2NotificationsParamsWithHTTPClient creates a new GetLolLobbyV2NotificationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolLobbyV2NotificationsParamsWithHTTPClient(client *http.Client) *GetLolLobbyV2NotificationsParams {

	return &GetLolLobbyV2NotificationsParams{
		HTTPClient: client,
	}
}

/*GetLolLobbyV2NotificationsParams contains all the parameters to send to the API endpoint
for the get lol lobby v2 notifications operation typically these are written to a http.Request
*/
type GetLolLobbyV2NotificationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol lobby v2 notifications params
func (o *GetLolLobbyV2NotificationsParams) WithTimeout(timeout time.Duration) *GetLolLobbyV2NotificationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol lobby v2 notifications params
func (o *GetLolLobbyV2NotificationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol lobby v2 notifications params
func (o *GetLolLobbyV2NotificationsParams) WithContext(ctx context.Context) *GetLolLobbyV2NotificationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol lobby v2 notifications params
func (o *GetLolLobbyV2NotificationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol lobby v2 notifications params
func (o *GetLolLobbyV2NotificationsParams) WithHTTPClient(client *http.Client) *GetLolLobbyV2NotificationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol lobby v2 notifications params
func (o *GetLolLobbyV2NotificationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolLobbyV2NotificationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
