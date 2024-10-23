// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/boonlogic/amber-go-sdk/v2/models"
)

// NewPostModelPretrainParams creates a new PostModelPretrainParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostModelPretrainParams() *PostModelPretrainParams {
	return &PostModelPretrainParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostModelPretrainParamsWithTimeout creates a new PostModelPretrainParams object
// with the ability to set a timeout on a request.
func NewPostModelPretrainParamsWithTimeout(timeout time.Duration) *PostModelPretrainParams {
	return &PostModelPretrainParams{
		timeout: timeout,
	}
}

// NewPostModelPretrainParamsWithContext creates a new PostModelPretrainParams object
// with the ability to set a context for a request.
func NewPostModelPretrainParamsWithContext(ctx context.Context) *PostModelPretrainParams {
	return &PostModelPretrainParams{
		Context: ctx,
	}
}

// NewPostModelPretrainParamsWithHTTPClient creates a new PostModelPretrainParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostModelPretrainParamsWithHTTPClient(client *http.Client) *PostModelPretrainParams {
	return &PostModelPretrainParams{
		HTTPClient: client,
	}
}

/*
PostModelPretrainParams contains all the parameters to send to the API endpoint

	for the post model pretrain operation.

	Typically these are written to a http.Request.
*/
type PostModelPretrainParams struct {

	/* Chunkspec.

	   Chunk specifier for chunked uploads. In a chunked upload, each request is sent with a `chunkspec` of the form `1:3`, `2:10`, `7:7`, etc. where the first number is the index for the chunk being uploaded (1-based) and the second number is the total number of chunks in the transaction. The chunks may be uploaded in any order. Pretraining starts once all chunks have been received.
	*/
	Chunkspec *string

	// ModelID.
	ModelID string

	/* PostPretrainRequest.

	   Data to use for pretraining.
	*/
	PostPretrainRequest *models.PostPretrainRequest

	/* TxnID.

	   Transaction id for chunked uploads. The response body for the first request in a chunked upload will contain a `txnId` which uniquely identifies the chunking transaction across multiple requests. That `txnId` must be included in the header of all remaining chunks uploaded in the transaction.
	*/
	TxnID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post model pretrain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostModelPretrainParams) WithDefaults() *PostModelPretrainParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post model pretrain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostModelPretrainParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post model pretrain params
func (o *PostModelPretrainParams) WithTimeout(timeout time.Duration) *PostModelPretrainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post model pretrain params
func (o *PostModelPretrainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post model pretrain params
func (o *PostModelPretrainParams) WithContext(ctx context.Context) *PostModelPretrainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post model pretrain params
func (o *PostModelPretrainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post model pretrain params
func (o *PostModelPretrainParams) WithHTTPClient(client *http.Client) *PostModelPretrainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post model pretrain params
func (o *PostModelPretrainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChunkspec adds the chunkspec to the post model pretrain params
func (o *PostModelPretrainParams) WithChunkspec(chunkspec *string) *PostModelPretrainParams {
	o.SetChunkspec(chunkspec)
	return o
}

// SetChunkspec adds the chunkspec to the post model pretrain params
func (o *PostModelPretrainParams) SetChunkspec(chunkspec *string) {
	o.Chunkspec = chunkspec
}

// WithModelID adds the modelID to the post model pretrain params
func (o *PostModelPretrainParams) WithModelID(modelID string) *PostModelPretrainParams {
	o.SetModelID(modelID)
	return o
}

// SetModelID adds the modelId to the post model pretrain params
func (o *PostModelPretrainParams) SetModelID(modelID string) {
	o.ModelID = modelID
}

// WithPostPretrainRequest adds the postPretrainRequest to the post model pretrain params
func (o *PostModelPretrainParams) WithPostPretrainRequest(postPretrainRequest *models.PostPretrainRequest) *PostModelPretrainParams {
	o.SetPostPretrainRequest(postPretrainRequest)
	return o
}

// SetPostPretrainRequest adds the postPretrainRequest to the post model pretrain params
func (o *PostModelPretrainParams) SetPostPretrainRequest(postPretrainRequest *models.PostPretrainRequest) {
	o.PostPretrainRequest = postPretrainRequest
}

// WithTxnID adds the txnID to the post model pretrain params
func (o *PostModelPretrainParams) WithTxnID(txnID *string) *PostModelPretrainParams {
	o.SetTxnID(txnID)
	return o
}

// SetTxnID adds the txnId to the post model pretrain params
func (o *PostModelPretrainParams) SetTxnID(txnID *string) {
	o.TxnID = txnID
}

// WriteToRequest writes these params to a swagger request
func (o *PostModelPretrainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Chunkspec != nil {

		// header param chunkspec
		if err := r.SetHeaderParam("chunkspec", *o.Chunkspec); err != nil {
			return err
		}
	}

	// path param modelId
	if err := r.SetPathParam("modelId", o.ModelID); err != nil {
		return err
	}
	if o.PostPretrainRequest != nil {
		if err := r.SetBodyParam(o.PostPretrainRequest); err != nil {
			return err
		}
	}

	if o.TxnID != nil {

		// header param txnId
		if err := r.SetHeaderParam("txnId", *o.TxnID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
