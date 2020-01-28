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

// NewPostLolLobbyV2EligibilityPartyParams creates a new PostLolLobbyV2EligibilityPartyParams object
// with the default values initialized.
func NewPostLolLobbyV2EligibilityPartyParams() *PostLolLobbyV2EligibilityPartyParams {

	return &PostLolLobbyV2EligibilityPartyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolLobbyV2EligibilityPartyParamsWithTimeout creates a new PostLolLobbyV2EligibilityPartyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolLobbyV2EligibilityPartyParamsWithTimeout(timeout time.Duration) *PostLolLobbyV2EligibilityPartyParams {

	return &PostLolLobbyV2EligibilityPartyParams{

		timeout: timeout,
	}
}

// NewPostLolLobbyV2EligibilityPartyParamsWithContext creates a new PostLolLobbyV2EligibilityPartyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolLobbyV2EligibilityPartyParamsWithContext(ctx context.Context) *PostLolLobbyV2EligibilityPartyParams {

	return &PostLolLobbyV2EligibilityPartyParams{

		Context: ctx,
	}
}

// NewPostLolLobbyV2EligibilityPartyParamsWithHTTPClient creates a new PostLolLobbyV2EligibilityPartyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolLobbyV2EligibilityPartyParamsWithHTTPClient(client *http.Client) *PostLolLobbyV2EligibilityPartyParams {

	return &PostLolLobbyV2EligibilityPartyParams{
		HTTPClient: client,
	}
}

/*PostLolLobbyV2EligibilityPartyParams contains all the parameters to send to the API endpoint
for the post lol lobby v2 eligibility party operation typically these are written to a http.Request
*/
type PostLolLobbyV2EligibilityPartyParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol lobby v2 eligibility party params
func (o *PostLolLobbyV2EligibilityPartyParams) WithTimeout(timeout time.Duration) *PostLolLobbyV2EligibilityPartyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol lobby v2 eligibility party params
func (o *PostLolLobbyV2EligibilityPartyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol lobby v2 eligibility party params
func (o *PostLolLobbyV2EligibilityPartyParams) WithContext(ctx context.Context) *PostLolLobbyV2EligibilityPartyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol lobby v2 eligibility party params
func (o *PostLolLobbyV2EligibilityPartyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol lobby v2 eligibility party params
func (o *PostLolLobbyV2EligibilityPartyParams) WithHTTPClient(client *http.Client) *PostLolLobbyV2EligibilityPartyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol lobby v2 eligibility party params
func (o *PostLolLobbyV2EligibilityPartyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolLobbyV2EligibilityPartyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
