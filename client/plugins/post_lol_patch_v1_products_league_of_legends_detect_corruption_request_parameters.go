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

// NewPostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams creates a new PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams object
// with the default values initialized.
func NewPostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams() *PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams {

	return &PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParamsWithTimeout creates a new PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParamsWithTimeout(timeout time.Duration) *PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams {

	return &PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams{

		timeout: timeout,
	}
}

// NewPostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParamsWithContext creates a new PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParamsWithContext(ctx context.Context) *PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams {

	return &PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams{

		Context: ctx,
	}
}

// NewPostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParamsWithHTTPClient creates a new PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParamsWithHTTPClient(client *http.Client) *PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams {

	return &PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams{
		HTTPClient: client,
	}
}

/*PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams contains all the parameters to send to the API endpoint
for the post lol patch v1 products league of legends detect corruption request operation typically these are written to a http.Request
*/
type PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol patch v1 products league of legends detect corruption request params
func (o *PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams) WithTimeout(timeout time.Duration) *PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol patch v1 products league of legends detect corruption request params
func (o *PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol patch v1 products league of legends detect corruption request params
func (o *PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams) WithContext(ctx context.Context) *PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol patch v1 products league of legends detect corruption request params
func (o *PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol patch v1 products league of legends detect corruption request params
func (o *PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams) WithHTTPClient(client *http.Client) *PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol patch v1 products league of legends detect corruption request params
func (o *PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
