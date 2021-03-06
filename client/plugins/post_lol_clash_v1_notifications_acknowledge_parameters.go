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

// NewPostLolClashV1NotificationsAcknowledgeParams creates a new PostLolClashV1NotificationsAcknowledgeParams object
// with the default values initialized.
func NewPostLolClashV1NotificationsAcknowledgeParams() *PostLolClashV1NotificationsAcknowledgeParams {

	return &PostLolClashV1NotificationsAcknowledgeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolClashV1NotificationsAcknowledgeParamsWithTimeout creates a new PostLolClashV1NotificationsAcknowledgeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolClashV1NotificationsAcknowledgeParamsWithTimeout(timeout time.Duration) *PostLolClashV1NotificationsAcknowledgeParams {

	return &PostLolClashV1NotificationsAcknowledgeParams{

		timeout: timeout,
	}
}

// NewPostLolClashV1NotificationsAcknowledgeParamsWithContext creates a new PostLolClashV1NotificationsAcknowledgeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolClashV1NotificationsAcknowledgeParamsWithContext(ctx context.Context) *PostLolClashV1NotificationsAcknowledgeParams {

	return &PostLolClashV1NotificationsAcknowledgeParams{

		Context: ctx,
	}
}

// NewPostLolClashV1NotificationsAcknowledgeParamsWithHTTPClient creates a new PostLolClashV1NotificationsAcknowledgeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolClashV1NotificationsAcknowledgeParamsWithHTTPClient(client *http.Client) *PostLolClashV1NotificationsAcknowledgeParams {

	return &PostLolClashV1NotificationsAcknowledgeParams{
		HTTPClient: client,
	}
}

/*PostLolClashV1NotificationsAcknowledgeParams contains all the parameters to send to the API endpoint
for the post lol clash v1 notifications acknowledge operation typically these are written to a http.Request
*/
type PostLolClashV1NotificationsAcknowledgeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol clash v1 notifications acknowledge params
func (o *PostLolClashV1NotificationsAcknowledgeParams) WithTimeout(timeout time.Duration) *PostLolClashV1NotificationsAcknowledgeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol clash v1 notifications acknowledge params
func (o *PostLolClashV1NotificationsAcknowledgeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol clash v1 notifications acknowledge params
func (o *PostLolClashV1NotificationsAcknowledgeParams) WithContext(ctx context.Context) *PostLolClashV1NotificationsAcknowledgeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol clash v1 notifications acknowledge params
func (o *PostLolClashV1NotificationsAcknowledgeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol clash v1 notifications acknowledge params
func (o *PostLolClashV1NotificationsAcknowledgeParams) WithHTTPClient(client *http.Client) *PostLolClashV1NotificationsAcknowledgeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol clash v1 notifications acknowledge params
func (o *PostLolClashV1NotificationsAcknowledgeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolClashV1NotificationsAcknowledgeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
