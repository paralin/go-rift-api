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

// NewPostLolStoreV1NotificationsAcknowledgeParams creates a new PostLolStoreV1NotificationsAcknowledgeParams object
// with the default values initialized.
func NewPostLolStoreV1NotificationsAcknowledgeParams() *PostLolStoreV1NotificationsAcknowledgeParams {
	var ()
	return &PostLolStoreV1NotificationsAcknowledgeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolStoreV1NotificationsAcknowledgeParamsWithTimeout creates a new PostLolStoreV1NotificationsAcknowledgeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolStoreV1NotificationsAcknowledgeParamsWithTimeout(timeout time.Duration) *PostLolStoreV1NotificationsAcknowledgeParams {
	var ()
	return &PostLolStoreV1NotificationsAcknowledgeParams{

		timeout: timeout,
	}
}

// NewPostLolStoreV1NotificationsAcknowledgeParamsWithContext creates a new PostLolStoreV1NotificationsAcknowledgeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolStoreV1NotificationsAcknowledgeParamsWithContext(ctx context.Context) *PostLolStoreV1NotificationsAcknowledgeParams {
	var ()
	return &PostLolStoreV1NotificationsAcknowledgeParams{

		Context: ctx,
	}
}

// NewPostLolStoreV1NotificationsAcknowledgeParamsWithHTTPClient creates a new PostLolStoreV1NotificationsAcknowledgeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolStoreV1NotificationsAcknowledgeParamsWithHTTPClient(client *http.Client) *PostLolStoreV1NotificationsAcknowledgeParams {
	var ()
	return &PostLolStoreV1NotificationsAcknowledgeParams{
		HTTPClient: client,
	}
}

/*PostLolStoreV1NotificationsAcknowledgeParams contains all the parameters to send to the API endpoint
for the post lol store v1 notifications acknowledge operation typically these are written to a http.Request
*/
type PostLolStoreV1NotificationsAcknowledgeParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol store v1 notifications acknowledge params
func (o *PostLolStoreV1NotificationsAcknowledgeParams) WithTimeout(timeout time.Duration) *PostLolStoreV1NotificationsAcknowledgeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol store v1 notifications acknowledge params
func (o *PostLolStoreV1NotificationsAcknowledgeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol store v1 notifications acknowledge params
func (o *PostLolStoreV1NotificationsAcknowledgeParams) WithContext(ctx context.Context) *PostLolStoreV1NotificationsAcknowledgeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol store v1 notifications acknowledge params
func (o *PostLolStoreV1NotificationsAcknowledgeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol store v1 notifications acknowledge params
func (o *PostLolStoreV1NotificationsAcknowledgeParams) WithHTTPClient(client *http.Client) *PostLolStoreV1NotificationsAcknowledgeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol store v1 notifications acknowledge params
func (o *PostLolStoreV1NotificationsAcknowledgeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post lol store v1 notifications acknowledge params
func (o *PostLolStoreV1NotificationsAcknowledgeParams) WithID(id string) *PostLolStoreV1NotificationsAcknowledgeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post lol store v1 notifications acknowledge params
func (o *PostLolStoreV1NotificationsAcknowledgeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolStoreV1NotificationsAcknowledgeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
