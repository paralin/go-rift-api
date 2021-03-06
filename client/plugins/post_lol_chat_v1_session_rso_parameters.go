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

// NewPostLolChatV1SessionRsoParams creates a new PostLolChatV1SessionRsoParams object
// with the default values initialized.
func NewPostLolChatV1SessionRsoParams() *PostLolChatV1SessionRsoParams {
	var ()
	return &PostLolChatV1SessionRsoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolChatV1SessionRsoParamsWithTimeout creates a new PostLolChatV1SessionRsoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolChatV1SessionRsoParamsWithTimeout(timeout time.Duration) *PostLolChatV1SessionRsoParams {
	var ()
	return &PostLolChatV1SessionRsoParams{

		timeout: timeout,
	}
}

// NewPostLolChatV1SessionRsoParamsWithContext creates a new PostLolChatV1SessionRsoParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolChatV1SessionRsoParamsWithContext(ctx context.Context) *PostLolChatV1SessionRsoParams {
	var ()
	return &PostLolChatV1SessionRsoParams{

		Context: ctx,
	}
}

// NewPostLolChatV1SessionRsoParamsWithHTTPClient creates a new PostLolChatV1SessionRsoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolChatV1SessionRsoParamsWithHTTPClient(client *http.Client) *PostLolChatV1SessionRsoParams {
	var ()
	return &PostLolChatV1SessionRsoParams{
		HTTPClient: client,
	}
}

/*PostLolChatV1SessionRsoParams contains all the parameters to send to the API endpoint
for the post lol chat v1 session rso operation typically these are written to a http.Request
*/
type PostLolChatV1SessionRsoParams struct {

	/*Auth*/
	Auth *models.LolChatAuthResourceRsoAccessToken

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol chat v1 session rso params
func (o *PostLolChatV1SessionRsoParams) WithTimeout(timeout time.Duration) *PostLolChatV1SessionRsoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol chat v1 session rso params
func (o *PostLolChatV1SessionRsoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol chat v1 session rso params
func (o *PostLolChatV1SessionRsoParams) WithContext(ctx context.Context) *PostLolChatV1SessionRsoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol chat v1 session rso params
func (o *PostLolChatV1SessionRsoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol chat v1 session rso params
func (o *PostLolChatV1SessionRsoParams) WithHTTPClient(client *http.Client) *PostLolChatV1SessionRsoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol chat v1 session rso params
func (o *PostLolChatV1SessionRsoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuth adds the auth to the post lol chat v1 session rso params
func (o *PostLolChatV1SessionRsoParams) WithAuth(auth *models.LolChatAuthResourceRsoAccessToken) *PostLolChatV1SessionRsoParams {
	o.SetAuth(auth)
	return o
}

// SetAuth adds the auth to the post lol chat v1 session rso params
func (o *PostLolChatV1SessionRsoParams) SetAuth(auth *models.LolChatAuthResourceRsoAccessToken) {
	o.Auth = auth
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolChatV1SessionRsoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Auth != nil {
		if err := r.SetBodyParam(o.Auth); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
