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

// NewPostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams creates a new PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams object
// with the default values initialized.
func NewPostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams() *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams {
	var ()
	return &PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolLootV1PlayerLootNotificationsByIDAcknowledgeParamsWithTimeout creates a new PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolLootV1PlayerLootNotificationsByIDAcknowledgeParamsWithTimeout(timeout time.Duration) *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams {
	var ()
	return &PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams{

		timeout: timeout,
	}
}

// NewPostLolLootV1PlayerLootNotificationsByIDAcknowledgeParamsWithContext creates a new PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolLootV1PlayerLootNotificationsByIDAcknowledgeParamsWithContext(ctx context.Context) *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams {
	var ()
	return &PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams{

		Context: ctx,
	}
}

// NewPostLolLootV1PlayerLootNotificationsByIDAcknowledgeParamsWithHTTPClient creates a new PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolLootV1PlayerLootNotificationsByIDAcknowledgeParamsWithHTTPClient(client *http.Client) *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams {
	var ()
	return &PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams{
		HTTPClient: client,
	}
}

/*PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams contains all the parameters to send to the API endpoint
for the post lol loot v1 player loot notifications by Id acknowledge operation typically these are written to a http.Request
*/
type PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol loot v1 player loot notifications by Id acknowledge params
func (o *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams) WithTimeout(timeout time.Duration) *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol loot v1 player loot notifications by Id acknowledge params
func (o *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol loot v1 player loot notifications by Id acknowledge params
func (o *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams) WithContext(ctx context.Context) *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol loot v1 player loot notifications by Id acknowledge params
func (o *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol loot v1 player loot notifications by Id acknowledge params
func (o *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams) WithHTTPClient(client *http.Client) *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol loot v1 player loot notifications by Id acknowledge params
func (o *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post lol loot v1 player loot notifications by Id acknowledge params
func (o *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams) WithID(id string) *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post lol loot v1 player loot notifications by Id acknowledge params
func (o *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
