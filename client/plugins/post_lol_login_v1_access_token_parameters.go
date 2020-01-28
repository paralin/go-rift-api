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

	models "github.com/paralin/go-rift-api/models"
)

// NewPostLolLoginV1AccessTokenParams creates a new PostLolLoginV1AccessTokenParams object
// with the default values initialized.
func NewPostLolLoginV1AccessTokenParams() *PostLolLoginV1AccessTokenParams {
	var ()
	return &PostLolLoginV1AccessTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolLoginV1AccessTokenParamsWithTimeout creates a new PostLolLoginV1AccessTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolLoginV1AccessTokenParamsWithTimeout(timeout time.Duration) *PostLolLoginV1AccessTokenParams {
	var ()
	return &PostLolLoginV1AccessTokenParams{

		timeout: timeout,
	}
}

// NewPostLolLoginV1AccessTokenParamsWithContext creates a new PostLolLoginV1AccessTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolLoginV1AccessTokenParamsWithContext(ctx context.Context) *PostLolLoginV1AccessTokenParams {
	var ()
	return &PostLolLoginV1AccessTokenParams{

		Context: ctx,
	}
}

// NewPostLolLoginV1AccessTokenParamsWithHTTPClient creates a new PostLolLoginV1AccessTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolLoginV1AccessTokenParamsWithHTTPClient(client *http.Client) *PostLolLoginV1AccessTokenParams {
	var ()
	return &PostLolLoginV1AccessTokenParams{
		HTTPClient: client,
	}
}

/*PostLolLoginV1AccessTokenParams contains all the parameters to send to the API endpoint
for the post lol login v1 access token operation typically these are written to a http.Request
*/
type PostLolLoginV1AccessTokenParams struct {

	/*AccessToken*/
	AccessToken *models.LolLoginAccessToken

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol login v1 access token params
func (o *PostLolLoginV1AccessTokenParams) WithTimeout(timeout time.Duration) *PostLolLoginV1AccessTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol login v1 access token params
func (o *PostLolLoginV1AccessTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol login v1 access token params
func (o *PostLolLoginV1AccessTokenParams) WithContext(ctx context.Context) *PostLolLoginV1AccessTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol login v1 access token params
func (o *PostLolLoginV1AccessTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol login v1 access token params
func (o *PostLolLoginV1AccessTokenParams) WithHTTPClient(client *http.Client) *PostLolLoginV1AccessTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol login v1 access token params
func (o *PostLolLoginV1AccessTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessToken adds the accessToken to the post lol login v1 access token params
func (o *PostLolLoginV1AccessTokenParams) WithAccessToken(accessToken *models.LolLoginAccessToken) *PostLolLoginV1AccessTokenParams {
	o.SetAccessToken(accessToken)
	return o
}

// SetAccessToken adds the accessToken to the post lol login v1 access token params
func (o *PostLolLoginV1AccessTokenParams) SetAccessToken(accessToken *models.LolLoginAccessToken) {
	o.AccessToken = accessToken
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolLoginV1AccessTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessToken != nil {
		if err := r.SetBodyParam(o.AccessToken); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}