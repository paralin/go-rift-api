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

// NewPostLolHighlightsV1FileBrowserByHighlightIDParams creates a new PostLolHighlightsV1FileBrowserByHighlightIDParams object
// with the default values initialized.
func NewPostLolHighlightsV1FileBrowserByHighlightIDParams() *PostLolHighlightsV1FileBrowserByHighlightIDParams {
	var ()
	return &PostLolHighlightsV1FileBrowserByHighlightIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolHighlightsV1FileBrowserByHighlightIDParamsWithTimeout creates a new PostLolHighlightsV1FileBrowserByHighlightIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolHighlightsV1FileBrowserByHighlightIDParamsWithTimeout(timeout time.Duration) *PostLolHighlightsV1FileBrowserByHighlightIDParams {
	var ()
	return &PostLolHighlightsV1FileBrowserByHighlightIDParams{

		timeout: timeout,
	}
}

// NewPostLolHighlightsV1FileBrowserByHighlightIDParamsWithContext creates a new PostLolHighlightsV1FileBrowserByHighlightIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolHighlightsV1FileBrowserByHighlightIDParamsWithContext(ctx context.Context) *PostLolHighlightsV1FileBrowserByHighlightIDParams {
	var ()
	return &PostLolHighlightsV1FileBrowserByHighlightIDParams{

		Context: ctx,
	}
}

// NewPostLolHighlightsV1FileBrowserByHighlightIDParamsWithHTTPClient creates a new PostLolHighlightsV1FileBrowserByHighlightIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolHighlightsV1FileBrowserByHighlightIDParamsWithHTTPClient(client *http.Client) *PostLolHighlightsV1FileBrowserByHighlightIDParams {
	var ()
	return &PostLolHighlightsV1FileBrowserByHighlightIDParams{
		HTTPClient: client,
	}
}

/*PostLolHighlightsV1FileBrowserByHighlightIDParams contains all the parameters to send to the API endpoint
for the post lol highlights v1 file browser by highlight Id operation typically these are written to a http.Request
*/
type PostLolHighlightsV1FileBrowserByHighlightIDParams struct {

	/*HighlightID*/
	HighlightID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol highlights v1 file browser by highlight Id params
func (o *PostLolHighlightsV1FileBrowserByHighlightIDParams) WithTimeout(timeout time.Duration) *PostLolHighlightsV1FileBrowserByHighlightIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol highlights v1 file browser by highlight Id params
func (o *PostLolHighlightsV1FileBrowserByHighlightIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol highlights v1 file browser by highlight Id params
func (o *PostLolHighlightsV1FileBrowserByHighlightIDParams) WithContext(ctx context.Context) *PostLolHighlightsV1FileBrowserByHighlightIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol highlights v1 file browser by highlight Id params
func (o *PostLolHighlightsV1FileBrowserByHighlightIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol highlights v1 file browser by highlight Id params
func (o *PostLolHighlightsV1FileBrowserByHighlightIDParams) WithHTTPClient(client *http.Client) *PostLolHighlightsV1FileBrowserByHighlightIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol highlights v1 file browser by highlight Id params
func (o *PostLolHighlightsV1FileBrowserByHighlightIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHighlightID adds the highlightID to the post lol highlights v1 file browser by highlight Id params
func (o *PostLolHighlightsV1FileBrowserByHighlightIDParams) WithHighlightID(highlightID int64) *PostLolHighlightsV1FileBrowserByHighlightIDParams {
	o.SetHighlightID(highlightID)
	return o
}

// SetHighlightID adds the highlightId to the post lol highlights v1 file browser by highlight Id params
func (o *PostLolHighlightsV1FileBrowserByHighlightIDParams) SetHighlightID(highlightID int64) {
	o.HighlightID = highlightID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolHighlightsV1FileBrowserByHighlightIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param highlightId
	if err := r.SetPathParam("highlightId", swag.FormatInt64(o.HighlightID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
