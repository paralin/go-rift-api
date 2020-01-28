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

// NewPostLolLoginV1LeagueSessionStatusParams creates a new PostLolLoginV1LeagueSessionStatusParams object
// with the default values initialized.
func NewPostLolLoginV1LeagueSessionStatusParams() *PostLolLoginV1LeagueSessionStatusParams {
	var ()
	return &PostLolLoginV1LeagueSessionStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolLoginV1LeagueSessionStatusParamsWithTimeout creates a new PostLolLoginV1LeagueSessionStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolLoginV1LeagueSessionStatusParamsWithTimeout(timeout time.Duration) *PostLolLoginV1LeagueSessionStatusParams {
	var ()
	return &PostLolLoginV1LeagueSessionStatusParams{

		timeout: timeout,
	}
}

// NewPostLolLoginV1LeagueSessionStatusParamsWithContext creates a new PostLolLoginV1LeagueSessionStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolLoginV1LeagueSessionStatusParamsWithContext(ctx context.Context) *PostLolLoginV1LeagueSessionStatusParams {
	var ()
	return &PostLolLoginV1LeagueSessionStatusParams{

		Context: ctx,
	}
}

// NewPostLolLoginV1LeagueSessionStatusParamsWithHTTPClient creates a new PostLolLoginV1LeagueSessionStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolLoginV1LeagueSessionStatusParamsWithHTTPClient(client *http.Client) *PostLolLoginV1LeagueSessionStatusParams {
	var ()
	return &PostLolLoginV1LeagueSessionStatusParams{
		HTTPClient: client,
	}
}

/*PostLolLoginV1LeagueSessionStatusParams contains all the parameters to send to the API endpoint
for the post lol login v1 league session status operation typically these are written to a http.Request
*/
type PostLolLoginV1LeagueSessionStatusParams struct {

	/*LeagueSession*/
	LeagueSession models.LolLoginLeagueSessionStatus

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol login v1 league session status params
func (o *PostLolLoginV1LeagueSessionStatusParams) WithTimeout(timeout time.Duration) *PostLolLoginV1LeagueSessionStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol login v1 league session status params
func (o *PostLolLoginV1LeagueSessionStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol login v1 league session status params
func (o *PostLolLoginV1LeagueSessionStatusParams) WithContext(ctx context.Context) *PostLolLoginV1LeagueSessionStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol login v1 league session status params
func (o *PostLolLoginV1LeagueSessionStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol login v1 league session status params
func (o *PostLolLoginV1LeagueSessionStatusParams) WithHTTPClient(client *http.Client) *PostLolLoginV1LeagueSessionStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol login v1 league session status params
func (o *PostLolLoginV1LeagueSessionStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLeagueSession adds the leagueSession to the post lol login v1 league session status params
func (o *PostLolLoginV1LeagueSessionStatusParams) WithLeagueSession(leagueSession models.LolLoginLeagueSessionStatus) *PostLolLoginV1LeagueSessionStatusParams {
	o.SetLeagueSession(leagueSession)
	return o
}

// SetLeagueSession adds the leagueSession to the post lol login v1 league session status params
func (o *PostLolLoginV1LeagueSessionStatusParams) SetLeagueSession(leagueSession models.LolLoginLeagueSessionStatus) {
	o.LeagueSession = leagueSession
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolLoginV1LeagueSessionStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.LeagueSession); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
