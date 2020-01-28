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

// NewDeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams creates a new DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams object
// with the default values initialized.
func NewDeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams() *DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams {
	var ()
	return &DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParamsWithTimeout creates a new DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParamsWithTimeout(timeout time.Duration) *DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams {
	var ()
	return &DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams{

		timeout: timeout,
	}
}

// NewDeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParamsWithContext creates a new DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParamsWithContext(ctx context.Context) *DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams {
	var ()
	return &DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams{

		Context: ctx,
	}
}

// NewDeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParamsWithHTTPClient creates a new DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParamsWithHTTPClient(client *http.Client) *DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams {
	var ()
	return &DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams{
		HTTPClient: client,
	}
}

/*DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams contains all the parameters to send to the API endpoint
for the delete lol player messaging v1 notification by Id acknowledge operation typically these are written to a http.Request
*/
type DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete lol player messaging v1 notification by Id acknowledge params
func (o *DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams) WithTimeout(timeout time.Duration) *DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete lol player messaging v1 notification by Id acknowledge params
func (o *DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete lol player messaging v1 notification by Id acknowledge params
func (o *DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams) WithContext(ctx context.Context) *DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete lol player messaging v1 notification by Id acknowledge params
func (o *DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete lol player messaging v1 notification by Id acknowledge params
func (o *DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams) WithHTTPClient(client *http.Client) *DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete lol player messaging v1 notification by Id acknowledge params
func (o *DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete lol player messaging v1 notification by Id acknowledge params
func (o *DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams) WithID(id int32) *DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete lol player messaging v1 notification by Id acknowledge params
func (o *DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLolPlayerMessagingV1NotificationByIDAcknowledgeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
