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

// NewPostLolRsoAuthV2ConfigParams creates a new PostLolRsoAuthV2ConfigParams object
// with the default values initialized.
func NewPostLolRsoAuthV2ConfigParams() *PostLolRsoAuthV2ConfigParams {
	var ()
	return &PostLolRsoAuthV2ConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolRsoAuthV2ConfigParamsWithTimeout creates a new PostLolRsoAuthV2ConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolRsoAuthV2ConfigParamsWithTimeout(timeout time.Duration) *PostLolRsoAuthV2ConfigParams {
	var ()
	return &PostLolRsoAuthV2ConfigParams{

		timeout: timeout,
	}
}

// NewPostLolRsoAuthV2ConfigParamsWithContext creates a new PostLolRsoAuthV2ConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolRsoAuthV2ConfigParamsWithContext(ctx context.Context) *PostLolRsoAuthV2ConfigParams {
	var ()
	return &PostLolRsoAuthV2ConfigParams{

		Context: ctx,
	}
}

// NewPostLolRsoAuthV2ConfigParamsWithHTTPClient creates a new PostLolRsoAuthV2ConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolRsoAuthV2ConfigParamsWithHTTPClient(client *http.Client) *PostLolRsoAuthV2ConfigParams {
	var ()
	return &PostLolRsoAuthV2ConfigParams{
		HTTPClient: client,
	}
}

/*PostLolRsoAuthV2ConfigParams contains all the parameters to send to the API endpoint
for the post lol rso auth v2 config operation typically these are written to a http.Request
*/
type PostLolRsoAuthV2ConfigParams struct {

	/*Config*/
	Config *models.LolRsoAuthPublicClientConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol rso auth v2 config params
func (o *PostLolRsoAuthV2ConfigParams) WithTimeout(timeout time.Duration) *PostLolRsoAuthV2ConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol rso auth v2 config params
func (o *PostLolRsoAuthV2ConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol rso auth v2 config params
func (o *PostLolRsoAuthV2ConfigParams) WithContext(ctx context.Context) *PostLolRsoAuthV2ConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol rso auth v2 config params
func (o *PostLolRsoAuthV2ConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol rso auth v2 config params
func (o *PostLolRsoAuthV2ConfigParams) WithHTTPClient(client *http.Client) *PostLolRsoAuthV2ConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol rso auth v2 config params
func (o *PostLolRsoAuthV2ConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the post lol rso auth v2 config params
func (o *PostLolRsoAuthV2ConfigParams) WithConfig(config *models.LolRsoAuthPublicClientConfig) *PostLolRsoAuthV2ConfigParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the post lol rso auth v2 config params
func (o *PostLolRsoAuthV2ConfigParams) SetConfig(config *models.LolRsoAuthPublicClientConfig) {
	o.Config = config
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolRsoAuthV2ConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Config != nil {
		if err := r.SetBodyParam(o.Config); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}