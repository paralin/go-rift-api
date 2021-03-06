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

// NewPostGcloudVoiceChatV2SessionsParams creates a new PostGcloudVoiceChatV2SessionsParams object
// with the default values initialized.
func NewPostGcloudVoiceChatV2SessionsParams() *PostGcloudVoiceChatV2SessionsParams {
	var ()
	return &PostGcloudVoiceChatV2SessionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostGcloudVoiceChatV2SessionsParamsWithTimeout creates a new PostGcloudVoiceChatV2SessionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostGcloudVoiceChatV2SessionsParamsWithTimeout(timeout time.Duration) *PostGcloudVoiceChatV2SessionsParams {
	var ()
	return &PostGcloudVoiceChatV2SessionsParams{

		timeout: timeout,
	}
}

// NewPostGcloudVoiceChatV2SessionsParamsWithContext creates a new PostGcloudVoiceChatV2SessionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostGcloudVoiceChatV2SessionsParamsWithContext(ctx context.Context) *PostGcloudVoiceChatV2SessionsParams {
	var ()
	return &PostGcloudVoiceChatV2SessionsParams{

		Context: ctx,
	}
}

// NewPostGcloudVoiceChatV2SessionsParamsWithHTTPClient creates a new PostGcloudVoiceChatV2SessionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostGcloudVoiceChatV2SessionsParamsWithHTTPClient(client *http.Client) *PostGcloudVoiceChatV2SessionsParams {
	var ()
	return &PostGcloudVoiceChatV2SessionsParams{
		HTTPClient: client,
	}
}

/*PostGcloudVoiceChatV2SessionsParams contains all the parameters to send to the API endpoint
for the post gcloud voice chat v2 sessions operation typically these are written to a http.Request
*/
type PostGcloudVoiceChatV2SessionsParams struct {

	/*JWT*/
	JWT string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post gcloud voice chat v2 sessions params
func (o *PostGcloudVoiceChatV2SessionsParams) WithTimeout(timeout time.Duration) *PostGcloudVoiceChatV2SessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post gcloud voice chat v2 sessions params
func (o *PostGcloudVoiceChatV2SessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post gcloud voice chat v2 sessions params
func (o *PostGcloudVoiceChatV2SessionsParams) WithContext(ctx context.Context) *PostGcloudVoiceChatV2SessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post gcloud voice chat v2 sessions params
func (o *PostGcloudVoiceChatV2SessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post gcloud voice chat v2 sessions params
func (o *PostGcloudVoiceChatV2SessionsParams) WithHTTPClient(client *http.Client) *PostGcloudVoiceChatV2SessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post gcloud voice chat v2 sessions params
func (o *PostGcloudVoiceChatV2SessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJWT adds the jWT to the post gcloud voice chat v2 sessions params
func (o *PostGcloudVoiceChatV2SessionsParams) WithJWT(jWT string) *PostGcloudVoiceChatV2SessionsParams {
	o.SetJWT(jWT)
	return o
}

// SetJWT adds the jWT to the post gcloud voice chat v2 sessions params
func (o *PostGcloudVoiceChatV2SessionsParams) SetJWT(jWT string) {
	o.JWT = jWT
}

// WriteToRequest writes these params to a swagger request
func (o *PostGcloudVoiceChatV2SessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param JWT
	if err := r.SetHeaderParam("JWT", o.JWT); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
