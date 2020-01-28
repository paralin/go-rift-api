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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteLolStatstonesV1VignetteNotificationsByKeyParams creates a new DeleteLolStatstonesV1VignetteNotificationsByKeyParams object
// with the default values initialized.
func NewDeleteLolStatstonesV1VignetteNotificationsByKeyParams() *DeleteLolStatstonesV1VignetteNotificationsByKeyParams {
	var ()
	return &DeleteLolStatstonesV1VignetteNotificationsByKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLolStatstonesV1VignetteNotificationsByKeyParamsWithTimeout creates a new DeleteLolStatstonesV1VignetteNotificationsByKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLolStatstonesV1VignetteNotificationsByKeyParamsWithTimeout(timeout time.Duration) *DeleteLolStatstonesV1VignetteNotificationsByKeyParams {
	var ()
	return &DeleteLolStatstonesV1VignetteNotificationsByKeyParams{

		timeout: timeout,
	}
}

// NewDeleteLolStatstonesV1VignetteNotificationsByKeyParamsWithContext creates a new DeleteLolStatstonesV1VignetteNotificationsByKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLolStatstonesV1VignetteNotificationsByKeyParamsWithContext(ctx context.Context) *DeleteLolStatstonesV1VignetteNotificationsByKeyParams {
	var ()
	return &DeleteLolStatstonesV1VignetteNotificationsByKeyParams{

		Context: ctx,
	}
}

// NewDeleteLolStatstonesV1VignetteNotificationsByKeyParamsWithHTTPClient creates a new DeleteLolStatstonesV1VignetteNotificationsByKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLolStatstonesV1VignetteNotificationsByKeyParamsWithHTTPClient(client *http.Client) *DeleteLolStatstonesV1VignetteNotificationsByKeyParams {
	var ()
	return &DeleteLolStatstonesV1VignetteNotificationsByKeyParams{
		HTTPClient: client,
	}
}

/*DeleteLolStatstonesV1VignetteNotificationsByKeyParams contains all the parameters to send to the API endpoint
for the delete lol statstones v1 vignette notifications by key operation typically these are written to a http.Request
*/
type DeleteLolStatstonesV1VignetteNotificationsByKeyParams struct {

	/*Key*/
	Key int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete lol statstones v1 vignette notifications by key params
func (o *DeleteLolStatstonesV1VignetteNotificationsByKeyParams) WithTimeout(timeout time.Duration) *DeleteLolStatstonesV1VignetteNotificationsByKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete lol statstones v1 vignette notifications by key params
func (o *DeleteLolStatstonesV1VignetteNotificationsByKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete lol statstones v1 vignette notifications by key params
func (o *DeleteLolStatstonesV1VignetteNotificationsByKeyParams) WithContext(ctx context.Context) *DeleteLolStatstonesV1VignetteNotificationsByKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete lol statstones v1 vignette notifications by key params
func (o *DeleteLolStatstonesV1VignetteNotificationsByKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete lol statstones v1 vignette notifications by key params
func (o *DeleteLolStatstonesV1VignetteNotificationsByKeyParams) WithHTTPClient(client *http.Client) *DeleteLolStatstonesV1VignetteNotificationsByKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete lol statstones v1 vignette notifications by key params
func (o *DeleteLolStatstonesV1VignetteNotificationsByKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the delete lol statstones v1 vignette notifications by key params
func (o *DeleteLolStatstonesV1VignetteNotificationsByKeyParams) WithKey(key int32) *DeleteLolStatstonesV1VignetteNotificationsByKeyParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the delete lol statstones v1 vignette notifications by key params
func (o *DeleteLolStatstonesV1VignetteNotificationsByKeyParams) SetKey(key int32) {
	o.Key = key
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLolStatstonesV1VignetteNotificationsByKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param key
	if err := r.SetPathParam("key", swag.FormatInt32(o.Key)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
