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

// NewPostLolStoreV1LastPageParams creates a new PostLolStoreV1LastPageParams object
// with the default values initialized.
func NewPostLolStoreV1LastPageParams() *PostLolStoreV1LastPageParams {
	var ()
	return &PostLolStoreV1LastPageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolStoreV1LastPageParamsWithTimeout creates a new PostLolStoreV1LastPageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolStoreV1LastPageParamsWithTimeout(timeout time.Duration) *PostLolStoreV1LastPageParams {
	var ()
	return &PostLolStoreV1LastPageParams{

		timeout: timeout,
	}
}

// NewPostLolStoreV1LastPageParamsWithContext creates a new PostLolStoreV1LastPageParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolStoreV1LastPageParamsWithContext(ctx context.Context) *PostLolStoreV1LastPageParams {
	var ()
	return &PostLolStoreV1LastPageParams{

		Context: ctx,
	}
}

// NewPostLolStoreV1LastPageParamsWithHTTPClient creates a new PostLolStoreV1LastPageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolStoreV1LastPageParamsWithHTTPClient(client *http.Client) *PostLolStoreV1LastPageParams {
	var ()
	return &PostLolStoreV1LastPageParams{
		HTTPClient: client,
	}
}

/*PostLolStoreV1LastPageParams contains all the parameters to send to the API endpoint
for the post lol store v1 last page operation typically these are written to a http.Request
*/
type PostLolStoreV1LastPageParams struct {

	/*PageType*/
	PageType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol store v1 last page params
func (o *PostLolStoreV1LastPageParams) WithTimeout(timeout time.Duration) *PostLolStoreV1LastPageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol store v1 last page params
func (o *PostLolStoreV1LastPageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol store v1 last page params
func (o *PostLolStoreV1LastPageParams) WithContext(ctx context.Context) *PostLolStoreV1LastPageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol store v1 last page params
func (o *PostLolStoreV1LastPageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol store v1 last page params
func (o *PostLolStoreV1LastPageParams) WithHTTPClient(client *http.Client) *PostLolStoreV1LastPageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol store v1 last page params
func (o *PostLolStoreV1LastPageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageType adds the pageType to the post lol store v1 last page params
func (o *PostLolStoreV1LastPageParams) WithPageType(pageType string) *PostLolStoreV1LastPageParams {
	o.SetPageType(pageType)
	return o
}

// SetPageType adds the pageType to the post lol store v1 last page params
func (o *PostLolStoreV1LastPageParams) SetPageType(pageType string) {
	o.PageType = pageType
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolStoreV1LastPageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.PageType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
