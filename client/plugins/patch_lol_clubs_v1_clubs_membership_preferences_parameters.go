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

// NewPatchLolClubsV1ClubsMembershipPreferencesParams creates a new PatchLolClubsV1ClubsMembershipPreferencesParams object
// with the default values initialized.
func NewPatchLolClubsV1ClubsMembershipPreferencesParams() *PatchLolClubsV1ClubsMembershipPreferencesParams {
	var ()
	return &PatchLolClubsV1ClubsMembershipPreferencesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchLolClubsV1ClubsMembershipPreferencesParamsWithTimeout creates a new PatchLolClubsV1ClubsMembershipPreferencesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchLolClubsV1ClubsMembershipPreferencesParamsWithTimeout(timeout time.Duration) *PatchLolClubsV1ClubsMembershipPreferencesParams {
	var ()
	return &PatchLolClubsV1ClubsMembershipPreferencesParams{

		timeout: timeout,
	}
}

// NewPatchLolClubsV1ClubsMembershipPreferencesParamsWithContext creates a new PatchLolClubsV1ClubsMembershipPreferencesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchLolClubsV1ClubsMembershipPreferencesParamsWithContext(ctx context.Context) *PatchLolClubsV1ClubsMembershipPreferencesParams {
	var ()
	return &PatchLolClubsV1ClubsMembershipPreferencesParams{

		Context: ctx,
	}
}

// NewPatchLolClubsV1ClubsMembershipPreferencesParamsWithHTTPClient creates a new PatchLolClubsV1ClubsMembershipPreferencesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchLolClubsV1ClubsMembershipPreferencesParamsWithHTTPClient(client *http.Client) *PatchLolClubsV1ClubsMembershipPreferencesParams {
	var ()
	return &PatchLolClubsV1ClubsMembershipPreferencesParams{
		HTTPClient: client,
	}
}

/*PatchLolClubsV1ClubsMembershipPreferencesParams contains all the parameters to send to the API endpoint
for the patch lol clubs v1 clubs membership preferences operation typically these are written to a http.Request
*/
type PatchLolClubsV1ClubsMembershipPreferencesParams struct {

	/*Preferences*/
	Preferences *models.LolClubsClubPreferences

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch lol clubs v1 clubs membership preferences params
func (o *PatchLolClubsV1ClubsMembershipPreferencesParams) WithTimeout(timeout time.Duration) *PatchLolClubsV1ClubsMembershipPreferencesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch lol clubs v1 clubs membership preferences params
func (o *PatchLolClubsV1ClubsMembershipPreferencesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch lol clubs v1 clubs membership preferences params
func (o *PatchLolClubsV1ClubsMembershipPreferencesParams) WithContext(ctx context.Context) *PatchLolClubsV1ClubsMembershipPreferencesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch lol clubs v1 clubs membership preferences params
func (o *PatchLolClubsV1ClubsMembershipPreferencesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch lol clubs v1 clubs membership preferences params
func (o *PatchLolClubsV1ClubsMembershipPreferencesParams) WithHTTPClient(client *http.Client) *PatchLolClubsV1ClubsMembershipPreferencesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch lol clubs v1 clubs membership preferences params
func (o *PatchLolClubsV1ClubsMembershipPreferencesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPreferences adds the preferences to the patch lol clubs v1 clubs membership preferences params
func (o *PatchLolClubsV1ClubsMembershipPreferencesParams) WithPreferences(preferences *models.LolClubsClubPreferences) *PatchLolClubsV1ClubsMembershipPreferencesParams {
	o.SetPreferences(preferences)
	return o
}

// SetPreferences adds the preferences to the patch lol clubs v1 clubs membership preferences params
func (o *PatchLolClubsV1ClubsMembershipPreferencesParams) SetPreferences(preferences *models.LolClubsClubPreferences) {
	o.Preferences = preferences
}

// WriteToRequest writes these params to a swagger request
func (o *PatchLolClubsV1ClubsMembershipPreferencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Preferences != nil {
		if err := r.SetBodyParam(o.Preferences); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}