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

// NewPostLolRsoAuthV1AuthorizationRefreshParams creates a new PostLolRsoAuthV1AuthorizationRefreshParams object
// with the default values initialized.
func NewPostLolRsoAuthV1AuthorizationRefreshParams() *PostLolRsoAuthV1AuthorizationRefreshParams {

	return &PostLolRsoAuthV1AuthorizationRefreshParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolRsoAuthV1AuthorizationRefreshParamsWithTimeout creates a new PostLolRsoAuthV1AuthorizationRefreshParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolRsoAuthV1AuthorizationRefreshParamsWithTimeout(timeout time.Duration) *PostLolRsoAuthV1AuthorizationRefreshParams {

	return &PostLolRsoAuthV1AuthorizationRefreshParams{

		timeout: timeout,
	}
}

// NewPostLolRsoAuthV1AuthorizationRefreshParamsWithContext creates a new PostLolRsoAuthV1AuthorizationRefreshParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolRsoAuthV1AuthorizationRefreshParamsWithContext(ctx context.Context) *PostLolRsoAuthV1AuthorizationRefreshParams {

	return &PostLolRsoAuthV1AuthorizationRefreshParams{

		Context: ctx,
	}
}

// NewPostLolRsoAuthV1AuthorizationRefreshParamsWithHTTPClient creates a new PostLolRsoAuthV1AuthorizationRefreshParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolRsoAuthV1AuthorizationRefreshParamsWithHTTPClient(client *http.Client) *PostLolRsoAuthV1AuthorizationRefreshParams {

	return &PostLolRsoAuthV1AuthorizationRefreshParams{
		HTTPClient: client,
	}
}

/*PostLolRsoAuthV1AuthorizationRefreshParams contains all the parameters to send to the API endpoint
for the post lol rso auth v1 authorization refresh operation typically these are written to a http.Request
*/
type PostLolRsoAuthV1AuthorizationRefreshParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol rso auth v1 authorization refresh params
func (o *PostLolRsoAuthV1AuthorizationRefreshParams) WithTimeout(timeout time.Duration) *PostLolRsoAuthV1AuthorizationRefreshParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol rso auth v1 authorization refresh params
func (o *PostLolRsoAuthV1AuthorizationRefreshParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol rso auth v1 authorization refresh params
func (o *PostLolRsoAuthV1AuthorizationRefreshParams) WithContext(ctx context.Context) *PostLolRsoAuthV1AuthorizationRefreshParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol rso auth v1 authorization refresh params
func (o *PostLolRsoAuthV1AuthorizationRefreshParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol rso auth v1 authorization refresh params
func (o *PostLolRsoAuthV1AuthorizationRefreshParams) WithHTTPClient(client *http.Client) *PostLolRsoAuthV1AuthorizationRefreshParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol rso auth v1 authorization refresh params
func (o *PostLolRsoAuthV1AuthorizationRefreshParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolRsoAuthV1AuthorizationRefreshParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}