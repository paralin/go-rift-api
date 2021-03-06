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

// NewPostLolClubsV1ClubsMembershipParams creates a new PostLolClubsV1ClubsMembershipParams object
// with the default values initialized.
func NewPostLolClubsV1ClubsMembershipParams() *PostLolClubsV1ClubsMembershipParams {
	var ()
	return &PostLolClubsV1ClubsMembershipParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolClubsV1ClubsMembershipParamsWithTimeout creates a new PostLolClubsV1ClubsMembershipParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolClubsV1ClubsMembershipParamsWithTimeout(timeout time.Duration) *PostLolClubsV1ClubsMembershipParams {
	var ()
	return &PostLolClubsV1ClubsMembershipParams{

		timeout: timeout,
	}
}

// NewPostLolClubsV1ClubsMembershipParamsWithContext creates a new PostLolClubsV1ClubsMembershipParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolClubsV1ClubsMembershipParamsWithContext(ctx context.Context) *PostLolClubsV1ClubsMembershipParams {
	var ()
	return &PostLolClubsV1ClubsMembershipParams{

		Context: ctx,
	}
}

// NewPostLolClubsV1ClubsMembershipParamsWithHTTPClient creates a new PostLolClubsV1ClubsMembershipParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolClubsV1ClubsMembershipParamsWithHTTPClient(client *http.Client) *PostLolClubsV1ClubsMembershipParams {
	var ()
	return &PostLolClubsV1ClubsMembershipParams{
		HTTPClient: client,
	}
}

/*PostLolClubsV1ClubsMembershipParams contains all the parameters to send to the API endpoint
for the post lol clubs v1 clubs membership operation typically these are written to a http.Request
*/
type PostLolClubsV1ClubsMembershipParams struct {

	/*Name*/
	Name *models.LolClubsClubName

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol clubs v1 clubs membership params
func (o *PostLolClubsV1ClubsMembershipParams) WithTimeout(timeout time.Duration) *PostLolClubsV1ClubsMembershipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol clubs v1 clubs membership params
func (o *PostLolClubsV1ClubsMembershipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol clubs v1 clubs membership params
func (o *PostLolClubsV1ClubsMembershipParams) WithContext(ctx context.Context) *PostLolClubsV1ClubsMembershipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol clubs v1 clubs membership params
func (o *PostLolClubsV1ClubsMembershipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol clubs v1 clubs membership params
func (o *PostLolClubsV1ClubsMembershipParams) WithHTTPClient(client *http.Client) *PostLolClubsV1ClubsMembershipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol clubs v1 clubs membership params
func (o *PostLolClubsV1ClubsMembershipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the post lol clubs v1 clubs membership params
func (o *PostLolClubsV1ClubsMembershipParams) WithName(name *models.LolClubsClubName) *PostLolClubsV1ClubsMembershipParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post lol clubs v1 clubs membership params
func (o *PostLolClubsV1ClubsMembershipParams) SetName(name *models.LolClubsClubName) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolClubsV1ClubsMembershipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {
		if err := r.SetBodyParam(o.Name); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
