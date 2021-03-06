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

// NewPostLolPerksV1UpdatePageOrderParams creates a new PostLolPerksV1UpdatePageOrderParams object
// with the default values initialized.
func NewPostLolPerksV1UpdatePageOrderParams() *PostLolPerksV1UpdatePageOrderParams {
	var ()
	return &PostLolPerksV1UpdatePageOrderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolPerksV1UpdatePageOrderParamsWithTimeout creates a new PostLolPerksV1UpdatePageOrderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolPerksV1UpdatePageOrderParamsWithTimeout(timeout time.Duration) *PostLolPerksV1UpdatePageOrderParams {
	var ()
	return &PostLolPerksV1UpdatePageOrderParams{

		timeout: timeout,
	}
}

// NewPostLolPerksV1UpdatePageOrderParamsWithContext creates a new PostLolPerksV1UpdatePageOrderParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolPerksV1UpdatePageOrderParamsWithContext(ctx context.Context) *PostLolPerksV1UpdatePageOrderParams {
	var ()
	return &PostLolPerksV1UpdatePageOrderParams{

		Context: ctx,
	}
}

// NewPostLolPerksV1UpdatePageOrderParamsWithHTTPClient creates a new PostLolPerksV1UpdatePageOrderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolPerksV1UpdatePageOrderParamsWithHTTPClient(client *http.Client) *PostLolPerksV1UpdatePageOrderParams {
	var ()
	return &PostLolPerksV1UpdatePageOrderParams{
		HTTPClient: client,
	}
}

/*PostLolPerksV1UpdatePageOrderParams contains all the parameters to send to the API endpoint
for the post lol perks v1 update page order operation typically these are written to a http.Request
*/
type PostLolPerksV1UpdatePageOrderParams struct {

	/*Request*/
	Request *models.LolPerksUpdatePageOrderRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol perks v1 update page order params
func (o *PostLolPerksV1UpdatePageOrderParams) WithTimeout(timeout time.Duration) *PostLolPerksV1UpdatePageOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol perks v1 update page order params
func (o *PostLolPerksV1UpdatePageOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol perks v1 update page order params
func (o *PostLolPerksV1UpdatePageOrderParams) WithContext(ctx context.Context) *PostLolPerksV1UpdatePageOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol perks v1 update page order params
func (o *PostLolPerksV1UpdatePageOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol perks v1 update page order params
func (o *PostLolPerksV1UpdatePageOrderParams) WithHTTPClient(client *http.Client) *PostLolPerksV1UpdatePageOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol perks v1 update page order params
func (o *PostLolPerksV1UpdatePageOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post lol perks v1 update page order params
func (o *PostLolPerksV1UpdatePageOrderParams) WithRequest(request *models.LolPerksUpdatePageOrderRequest) *PostLolPerksV1UpdatePageOrderParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post lol perks v1 update page order params
func (o *PostLolPerksV1UpdatePageOrderParams) SetRequest(request *models.LolPerksUpdatePageOrderRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolPerksV1UpdatePageOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
