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

// NewPostLolLootV1RefreshParams creates a new PostLolLootV1RefreshParams object
// with the default values initialized.
func NewPostLolLootV1RefreshParams() *PostLolLootV1RefreshParams {
	var ()
	return &PostLolLootV1RefreshParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolLootV1RefreshParamsWithTimeout creates a new PostLolLootV1RefreshParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolLootV1RefreshParamsWithTimeout(timeout time.Duration) *PostLolLootV1RefreshParams {
	var ()
	return &PostLolLootV1RefreshParams{

		timeout: timeout,
	}
}

// NewPostLolLootV1RefreshParamsWithContext creates a new PostLolLootV1RefreshParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolLootV1RefreshParamsWithContext(ctx context.Context) *PostLolLootV1RefreshParams {
	var ()
	return &PostLolLootV1RefreshParams{

		Context: ctx,
	}
}

// NewPostLolLootV1RefreshParamsWithHTTPClient creates a new PostLolLootV1RefreshParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolLootV1RefreshParamsWithHTTPClient(client *http.Client) *PostLolLootV1RefreshParams {
	var ()
	return &PostLolLootV1RefreshParams{
		HTTPClient: client,
	}
}

/*PostLolLootV1RefreshParams contains all the parameters to send to the API endpoint
for the post lol loot v1 refresh operation typically these are written to a http.Request
*/
type PostLolLootV1RefreshParams struct {

	/*Force*/
	Force bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol loot v1 refresh params
func (o *PostLolLootV1RefreshParams) WithTimeout(timeout time.Duration) *PostLolLootV1RefreshParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol loot v1 refresh params
func (o *PostLolLootV1RefreshParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol loot v1 refresh params
func (o *PostLolLootV1RefreshParams) WithContext(ctx context.Context) *PostLolLootV1RefreshParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol loot v1 refresh params
func (o *PostLolLootV1RefreshParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol loot v1 refresh params
func (o *PostLolLootV1RefreshParams) WithHTTPClient(client *http.Client) *PostLolLootV1RefreshParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol loot v1 refresh params
func (o *PostLolLootV1RefreshParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the post lol loot v1 refresh params
func (o *PostLolLootV1RefreshParams) WithForce(force bool) *PostLolLootV1RefreshParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the post lol loot v1 refresh params
func (o *PostLolLootV1RefreshParams) SetForce(force bool) {
	o.Force = force
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolLootV1RefreshParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param force
	qrForce := o.Force
	qForce := swag.FormatBool(qrForce)
	if qForce != "" {
		if err := r.SetQueryParam("force", qForce); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}