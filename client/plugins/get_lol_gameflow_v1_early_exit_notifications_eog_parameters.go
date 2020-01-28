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

// NewGetLolGameflowV1EarlyExitNotificationsEogParams creates a new GetLolGameflowV1EarlyExitNotificationsEogParams object
// with the default values initialized.
func NewGetLolGameflowV1EarlyExitNotificationsEogParams() *GetLolGameflowV1EarlyExitNotificationsEogParams {

	return &GetLolGameflowV1EarlyExitNotificationsEogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolGameflowV1EarlyExitNotificationsEogParamsWithTimeout creates a new GetLolGameflowV1EarlyExitNotificationsEogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolGameflowV1EarlyExitNotificationsEogParamsWithTimeout(timeout time.Duration) *GetLolGameflowV1EarlyExitNotificationsEogParams {

	return &GetLolGameflowV1EarlyExitNotificationsEogParams{

		timeout: timeout,
	}
}

// NewGetLolGameflowV1EarlyExitNotificationsEogParamsWithContext creates a new GetLolGameflowV1EarlyExitNotificationsEogParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolGameflowV1EarlyExitNotificationsEogParamsWithContext(ctx context.Context) *GetLolGameflowV1EarlyExitNotificationsEogParams {

	return &GetLolGameflowV1EarlyExitNotificationsEogParams{

		Context: ctx,
	}
}

// NewGetLolGameflowV1EarlyExitNotificationsEogParamsWithHTTPClient creates a new GetLolGameflowV1EarlyExitNotificationsEogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolGameflowV1EarlyExitNotificationsEogParamsWithHTTPClient(client *http.Client) *GetLolGameflowV1EarlyExitNotificationsEogParams {

	return &GetLolGameflowV1EarlyExitNotificationsEogParams{
		HTTPClient: client,
	}
}

/*GetLolGameflowV1EarlyExitNotificationsEogParams contains all the parameters to send to the API endpoint
for the get lol gameflow v1 early exit notifications eog operation typically these are written to a http.Request
*/
type GetLolGameflowV1EarlyExitNotificationsEogParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol gameflow v1 early exit notifications eog params
func (o *GetLolGameflowV1EarlyExitNotificationsEogParams) WithTimeout(timeout time.Duration) *GetLolGameflowV1EarlyExitNotificationsEogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol gameflow v1 early exit notifications eog params
func (o *GetLolGameflowV1EarlyExitNotificationsEogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol gameflow v1 early exit notifications eog params
func (o *GetLolGameflowV1EarlyExitNotificationsEogParams) WithContext(ctx context.Context) *GetLolGameflowV1EarlyExitNotificationsEogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol gameflow v1 early exit notifications eog params
func (o *GetLolGameflowV1EarlyExitNotificationsEogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol gameflow v1 early exit notifications eog params
func (o *GetLolGameflowV1EarlyExitNotificationsEogParams) WithHTTPClient(client *http.Client) *GetLolGameflowV1EarlyExitNotificationsEogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol gameflow v1 early exit notifications eog params
func (o *GetLolGameflowV1EarlyExitNotificationsEogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolGameflowV1EarlyExitNotificationsEogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
