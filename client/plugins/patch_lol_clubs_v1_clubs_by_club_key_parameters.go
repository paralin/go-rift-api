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

// NewPatchLolClubsV1ClubsByClubKeyParams creates a new PatchLolClubsV1ClubsByClubKeyParams object
// with the default values initialized.
func NewPatchLolClubsV1ClubsByClubKeyParams() *PatchLolClubsV1ClubsByClubKeyParams {
	var ()
	return &PatchLolClubsV1ClubsByClubKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchLolClubsV1ClubsByClubKeyParamsWithTimeout creates a new PatchLolClubsV1ClubsByClubKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchLolClubsV1ClubsByClubKeyParamsWithTimeout(timeout time.Duration) *PatchLolClubsV1ClubsByClubKeyParams {
	var ()
	return &PatchLolClubsV1ClubsByClubKeyParams{

		timeout: timeout,
	}
}

// NewPatchLolClubsV1ClubsByClubKeyParamsWithContext creates a new PatchLolClubsV1ClubsByClubKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchLolClubsV1ClubsByClubKeyParamsWithContext(ctx context.Context) *PatchLolClubsV1ClubsByClubKeyParams {
	var ()
	return &PatchLolClubsV1ClubsByClubKeyParams{

		Context: ctx,
	}
}

// NewPatchLolClubsV1ClubsByClubKeyParamsWithHTTPClient creates a new PatchLolClubsV1ClubsByClubKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchLolClubsV1ClubsByClubKeyParamsWithHTTPClient(client *http.Client) *PatchLolClubsV1ClubsByClubKeyParams {
	var ()
	return &PatchLolClubsV1ClubsByClubKeyParams{
		HTTPClient: client,
	}
}

/*PatchLolClubsV1ClubsByClubKeyParams contains all the parameters to send to the API endpoint
for the patch lol clubs v1 clubs by club key operation typically these are written to a http.Request
*/
type PatchLolClubsV1ClubsByClubKeyParams struct {

	/*ClubKey*/
	ClubKey string
	/*Tag*/
	Tag *models.LolClubsClubTag

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch lol clubs v1 clubs by club key params
func (o *PatchLolClubsV1ClubsByClubKeyParams) WithTimeout(timeout time.Duration) *PatchLolClubsV1ClubsByClubKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch lol clubs v1 clubs by club key params
func (o *PatchLolClubsV1ClubsByClubKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch lol clubs v1 clubs by club key params
func (o *PatchLolClubsV1ClubsByClubKeyParams) WithContext(ctx context.Context) *PatchLolClubsV1ClubsByClubKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch lol clubs v1 clubs by club key params
func (o *PatchLolClubsV1ClubsByClubKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch lol clubs v1 clubs by club key params
func (o *PatchLolClubsV1ClubsByClubKeyParams) WithHTTPClient(client *http.Client) *PatchLolClubsV1ClubsByClubKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch lol clubs v1 clubs by club key params
func (o *PatchLolClubsV1ClubsByClubKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClubKey adds the clubKey to the patch lol clubs v1 clubs by club key params
func (o *PatchLolClubsV1ClubsByClubKeyParams) WithClubKey(clubKey string) *PatchLolClubsV1ClubsByClubKeyParams {
	o.SetClubKey(clubKey)
	return o
}

// SetClubKey adds the clubKey to the patch lol clubs v1 clubs by club key params
func (o *PatchLolClubsV1ClubsByClubKeyParams) SetClubKey(clubKey string) {
	o.ClubKey = clubKey
}

// WithTag adds the tag to the patch lol clubs v1 clubs by club key params
func (o *PatchLolClubsV1ClubsByClubKeyParams) WithTag(tag *models.LolClubsClubTag) *PatchLolClubsV1ClubsByClubKeyParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the patch lol clubs v1 clubs by club key params
func (o *PatchLolClubsV1ClubsByClubKeyParams) SetTag(tag *models.LolClubsClubTag) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *PatchLolClubsV1ClubsByClubKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clubKey
	if err := r.SetPathParam("clubKey", o.ClubKey); err != nil {
		return err
	}

	if o.Tag != nil {
		if err := r.SetBodyParam(o.Tag); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
