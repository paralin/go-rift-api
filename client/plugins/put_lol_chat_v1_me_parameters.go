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

// NewPutLolChatV1MeParams creates a new PutLolChatV1MeParams object
// with the default values initialized.
func NewPutLolChatV1MeParams() *PutLolChatV1MeParams {
	var ()
	return &PutLolChatV1MeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLolChatV1MeParamsWithTimeout creates a new PutLolChatV1MeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLolChatV1MeParamsWithTimeout(timeout time.Duration) *PutLolChatV1MeParams {
	var ()
	return &PutLolChatV1MeParams{

		timeout: timeout,
	}
}

// NewPutLolChatV1MeParamsWithContext creates a new PutLolChatV1MeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLolChatV1MeParamsWithContext(ctx context.Context) *PutLolChatV1MeParams {
	var ()
	return &PutLolChatV1MeParams{

		Context: ctx,
	}
}

// NewPutLolChatV1MeParamsWithHTTPClient creates a new PutLolChatV1MeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLolChatV1MeParamsWithHTTPClient(client *http.Client) *PutLolChatV1MeParams {
	var ()
	return &PutLolChatV1MeParams{
		HTTPClient: client,
	}
}

/*PutLolChatV1MeParams contains all the parameters to send to the API endpoint
for the put lol chat v1 me operation typically these are written to a http.Request
*/
type PutLolChatV1MeParams struct {

	/*Me*/
	Me *models.LolChatUserResource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put lol chat v1 me params
func (o *PutLolChatV1MeParams) WithTimeout(timeout time.Duration) *PutLolChatV1MeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put lol chat v1 me params
func (o *PutLolChatV1MeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put lol chat v1 me params
func (o *PutLolChatV1MeParams) WithContext(ctx context.Context) *PutLolChatV1MeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put lol chat v1 me params
func (o *PutLolChatV1MeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put lol chat v1 me params
func (o *PutLolChatV1MeParams) WithHTTPClient(client *http.Client) *PutLolChatV1MeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put lol chat v1 me params
func (o *PutLolChatV1MeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMe adds the me to the put lol chat v1 me params
func (o *PutLolChatV1MeParams) WithMe(me *models.LolChatUserResource) *PutLolChatV1MeParams {
	o.SetMe(me)
	return o
}

// SetMe adds the me to the put lol chat v1 me params
func (o *PutLolChatV1MeParams) SetMe(me *models.LolChatUserResource) {
	o.Me = me
}

// WriteToRequest writes these params to a swagger request
func (o *PutLolChatV1MeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Me != nil {
		if err := r.SetBodyParam(o.Me); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}