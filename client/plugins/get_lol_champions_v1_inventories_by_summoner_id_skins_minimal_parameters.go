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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams creates a new GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams object
// with the default values initialized.
func NewGetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams() *GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams {
	var ()
	return &GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParamsWithTimeout creates a new GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParamsWithTimeout(timeout time.Duration) *GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams {
	var ()
	return &GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams{

		timeout: timeout,
	}
}

// NewGetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParamsWithContext creates a new GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParamsWithContext(ctx context.Context) *GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams {
	var ()
	return &GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams{

		Context: ctx,
	}
}

// NewGetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParamsWithHTTPClient creates a new GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParamsWithHTTPClient(client *http.Client) *GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams {
	var ()
	return &GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams{
		HTTPClient: client,
	}
}

/*GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams contains all the parameters to send to the API endpoint
for the get lol champions v1 inventories by summoner Id skins minimal operation typically these are written to a http.Request
*/
type GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams struct {

	/*SummonerID*/
	SummonerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol champions v1 inventories by summoner Id skins minimal params
func (o *GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams) WithTimeout(timeout time.Duration) *GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol champions v1 inventories by summoner Id skins minimal params
func (o *GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol champions v1 inventories by summoner Id skins minimal params
func (o *GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams) WithContext(ctx context.Context) *GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol champions v1 inventories by summoner Id skins minimal params
func (o *GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol champions v1 inventories by summoner Id skins minimal params
func (o *GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams) WithHTTPClient(client *http.Client) *GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol champions v1 inventories by summoner Id skins minimal params
func (o *GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSummonerID adds the summonerID to the get lol champions v1 inventories by summoner Id skins minimal params
func (o *GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams) WithSummonerID(summonerID int64) *GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams {
	o.SetSummonerID(summonerID)
	return o
}

// SetSummonerID adds the summonerId to the get lol champions v1 inventories by summoner Id skins minimal params
func (o *GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams) SetSummonerID(summonerID int64) {
	o.SummonerID = summonerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolChampionsV1InventoriesBySummonerIDSkinsMinimalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param summonerId
	if err := r.SetPathParam("summonerId", swag.FormatInt64(o.SummonerID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
