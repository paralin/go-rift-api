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

	models "github.com/paralin/go-rift-api/models"
)

// NewPutLolItemSetsV1ItemSetsBySummonerIDSetsParams creates a new PutLolItemSetsV1ItemSetsBySummonerIDSetsParams object
// with the default values initialized.
func NewPutLolItemSetsV1ItemSetsBySummonerIDSetsParams() *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams {
	var ()
	return &PutLolItemSetsV1ItemSetsBySummonerIDSetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLolItemSetsV1ItemSetsBySummonerIDSetsParamsWithTimeout creates a new PutLolItemSetsV1ItemSetsBySummonerIDSetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLolItemSetsV1ItemSetsBySummonerIDSetsParamsWithTimeout(timeout time.Duration) *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams {
	var ()
	return &PutLolItemSetsV1ItemSetsBySummonerIDSetsParams{

		timeout: timeout,
	}
}

// NewPutLolItemSetsV1ItemSetsBySummonerIDSetsParamsWithContext creates a new PutLolItemSetsV1ItemSetsBySummonerIDSetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLolItemSetsV1ItemSetsBySummonerIDSetsParamsWithContext(ctx context.Context) *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams {
	var ()
	return &PutLolItemSetsV1ItemSetsBySummonerIDSetsParams{

		Context: ctx,
	}
}

// NewPutLolItemSetsV1ItemSetsBySummonerIDSetsParamsWithHTTPClient creates a new PutLolItemSetsV1ItemSetsBySummonerIDSetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLolItemSetsV1ItemSetsBySummonerIDSetsParamsWithHTTPClient(client *http.Client) *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams {
	var ()
	return &PutLolItemSetsV1ItemSetsBySummonerIDSetsParams{
		HTTPClient: client,
	}
}

/*PutLolItemSetsV1ItemSetsBySummonerIDSetsParams contains all the parameters to send to the API endpoint
for the put lol item sets v1 item sets by summoner Id sets operation typically these are written to a http.Request
*/
type PutLolItemSetsV1ItemSetsBySummonerIDSetsParams struct {

	/*ItemSets*/
	ItemSets *models.LolItemSetsItemSets
	/*SummonerID*/
	SummonerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put lol item sets v1 item sets by summoner Id sets params
func (o *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams) WithTimeout(timeout time.Duration) *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put lol item sets v1 item sets by summoner Id sets params
func (o *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put lol item sets v1 item sets by summoner Id sets params
func (o *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams) WithContext(ctx context.Context) *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put lol item sets v1 item sets by summoner Id sets params
func (o *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put lol item sets v1 item sets by summoner Id sets params
func (o *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams) WithHTTPClient(client *http.Client) *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put lol item sets v1 item sets by summoner Id sets params
func (o *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithItemSets adds the itemSets to the put lol item sets v1 item sets by summoner Id sets params
func (o *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams) WithItemSets(itemSets *models.LolItemSetsItemSets) *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams {
	o.SetItemSets(itemSets)
	return o
}

// SetItemSets adds the itemSets to the put lol item sets v1 item sets by summoner Id sets params
func (o *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams) SetItemSets(itemSets *models.LolItemSetsItemSets) {
	o.ItemSets = itemSets
}

// WithSummonerID adds the summonerID to the put lol item sets v1 item sets by summoner Id sets params
func (o *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams) WithSummonerID(summonerID int64) *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams {
	o.SetSummonerID(summonerID)
	return o
}

// SetSummonerID adds the summonerId to the put lol item sets v1 item sets by summoner Id sets params
func (o *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams) SetSummonerID(summonerID int64) {
	o.SummonerID = summonerID
}

// WriteToRequest writes these params to a swagger request
func (o *PutLolItemSetsV1ItemSetsBySummonerIDSetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ItemSets != nil {
		if err := r.SetBodyParam(o.ItemSets); err != nil {
			return err
		}
	}

	// path param summonerId
	if err := r.SetPathParam("summonerId", swag.FormatInt64(o.SummonerID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}