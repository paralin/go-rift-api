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

// NewGetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams creates a new GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams object
// with the default values initialized.
func NewGetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams() *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParamsWithTimeout creates a new GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParamsWithTimeout(timeout time.Duration) *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams{

		timeout: timeout,
	}
}

// NewGetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParamsWithContext creates a new GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParamsWithContext(ctx context.Context) *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams{

		Context: ctx,
	}
}

// NewGetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParamsWithHTTPClient creates a new GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParamsWithHTTPClient(client *http.Client) *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams{
		HTTPClient: client,
	}
}

/*GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams contains all the parameters to send to the API endpoint
for the get lol collections v1 inventories by summoner Id ward skins by ward skin Id operation typically these are written to a http.Request
*/
type GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams struct {

	/*SummonerID*/
	SummonerID int64
	/*WardSkinID*/
	WardSkinID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol collections v1 inventories by summoner Id ward skins by ward skin Id params
func (o *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams) WithTimeout(timeout time.Duration) *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol collections v1 inventories by summoner Id ward skins by ward skin Id params
func (o *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol collections v1 inventories by summoner Id ward skins by ward skin Id params
func (o *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams) WithContext(ctx context.Context) *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol collections v1 inventories by summoner Id ward skins by ward skin Id params
func (o *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol collections v1 inventories by summoner Id ward skins by ward skin Id params
func (o *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams) WithHTTPClient(client *http.Client) *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol collections v1 inventories by summoner Id ward skins by ward skin Id params
func (o *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSummonerID adds the summonerID to the get lol collections v1 inventories by summoner Id ward skins by ward skin Id params
func (o *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams) WithSummonerID(summonerID int64) *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams {
	o.SetSummonerID(summonerID)
	return o
}

// SetSummonerID adds the summonerId to the get lol collections v1 inventories by summoner Id ward skins by ward skin Id params
func (o *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams) SetSummonerID(summonerID int64) {
	o.SummonerID = summonerID
}

// WithWardSkinID adds the wardSkinID to the get lol collections v1 inventories by summoner Id ward skins by ward skin Id params
func (o *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams) WithWardSkinID(wardSkinID int64) *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams {
	o.SetWardSkinID(wardSkinID)
	return o
}

// SetWardSkinID adds the wardSkinId to the get lol collections v1 inventories by summoner Id ward skins by ward skin Id params
func (o *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams) SetWardSkinID(wardSkinID int64) {
	o.WardSkinID = wardSkinID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolCollectionsV1InventoriesBySummonerIDWardSkinsByWardSkinIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param summonerId
	if err := r.SetPathParam("summonerId", swag.FormatInt64(o.SummonerID)); err != nil {
		return err
	}

	// path param wardSkinId
	if err := r.SetPathParam("wardSkinId", swag.FormatInt64(o.WardSkinID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
