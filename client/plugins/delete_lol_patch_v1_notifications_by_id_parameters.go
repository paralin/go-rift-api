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

// NewDeleteLolPatchV1NotificationsByIDParams creates a new DeleteLolPatchV1NotificationsByIDParams object
// with the default values initialized.
func NewDeleteLolPatchV1NotificationsByIDParams() *DeleteLolPatchV1NotificationsByIDParams {
	var ()
	return &DeleteLolPatchV1NotificationsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLolPatchV1NotificationsByIDParamsWithTimeout creates a new DeleteLolPatchV1NotificationsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLolPatchV1NotificationsByIDParamsWithTimeout(timeout time.Duration) *DeleteLolPatchV1NotificationsByIDParams {
	var ()
	return &DeleteLolPatchV1NotificationsByIDParams{

		timeout: timeout,
	}
}

// NewDeleteLolPatchV1NotificationsByIDParamsWithContext creates a new DeleteLolPatchV1NotificationsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLolPatchV1NotificationsByIDParamsWithContext(ctx context.Context) *DeleteLolPatchV1NotificationsByIDParams {
	var ()
	return &DeleteLolPatchV1NotificationsByIDParams{

		Context: ctx,
	}
}

// NewDeleteLolPatchV1NotificationsByIDParamsWithHTTPClient creates a new DeleteLolPatchV1NotificationsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLolPatchV1NotificationsByIDParamsWithHTTPClient(client *http.Client) *DeleteLolPatchV1NotificationsByIDParams {
	var ()
	return &DeleteLolPatchV1NotificationsByIDParams{
		HTTPClient: client,
	}
}

/*DeleteLolPatchV1NotificationsByIDParams contains all the parameters to send to the API endpoint
for the delete lol patch v1 notifications by Id operation typically these are written to a http.Request
*/
type DeleteLolPatchV1NotificationsByIDParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete lol patch v1 notifications by Id params
func (o *DeleteLolPatchV1NotificationsByIDParams) WithTimeout(timeout time.Duration) *DeleteLolPatchV1NotificationsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete lol patch v1 notifications by Id params
func (o *DeleteLolPatchV1NotificationsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete lol patch v1 notifications by Id params
func (o *DeleteLolPatchV1NotificationsByIDParams) WithContext(ctx context.Context) *DeleteLolPatchV1NotificationsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete lol patch v1 notifications by Id params
func (o *DeleteLolPatchV1NotificationsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete lol patch v1 notifications by Id params
func (o *DeleteLolPatchV1NotificationsByIDParams) WithHTTPClient(client *http.Client) *DeleteLolPatchV1NotificationsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete lol patch v1 notifications by Id params
func (o *DeleteLolPatchV1NotificationsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete lol patch v1 notifications by Id params
func (o *DeleteLolPatchV1NotificationsByIDParams) WithID(id string) *DeleteLolPatchV1NotificationsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete lol patch v1 notifications by Id params
func (o *DeleteLolPatchV1NotificationsByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLolPatchV1NotificationsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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