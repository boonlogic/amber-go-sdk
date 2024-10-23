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
)

// NewGetModelRootCauseParams creates a new GetModelRootCauseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetModelRootCauseParams() *GetModelRootCauseParams {
	return &GetModelRootCauseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetModelRootCauseParamsWithTimeout creates a new GetModelRootCauseParams object
// with the ability to set a timeout on a request.
func NewGetModelRootCauseParamsWithTimeout(timeout time.Duration) *GetModelRootCauseParams {
	return &GetModelRootCauseParams{
		timeout: timeout,
	}
}

// NewGetModelRootCauseParamsWithContext creates a new GetModelRootCauseParams object
// with the ability to set a context for a request.
func NewGetModelRootCauseParamsWithContext(ctx context.Context) *GetModelRootCauseParams {
	return &GetModelRootCauseParams{
		Context: ctx,
	}
}

// NewGetModelRootCauseParamsWithHTTPClient creates a new GetModelRootCauseParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetModelRootCauseParamsWithHTTPClient(client *http.Client) *GetModelRootCauseParams {
	return &GetModelRootCauseParams{
		HTTPClient: client,
	}
}

/*
GetModelRootCauseParams contains all the parameters to send to the API endpoint

	for the get model root cause operation.

	Typically these are written to a http.Request.
*/
type GetModelRootCauseParams struct {

	/* Clusters.

	   Clusters to analyze (list of comma-separated integers).
	*/
	Clusters *string

	// ModelID.
	ModelID string

	/* Vectors.

	   Vectors to analyze, as a flat list of comma-separated floats. Number of values must be a multiple of the configured number of features.
	*/
	Vectors *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get model root cause params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetModelRootCauseParams) WithDefaults() *GetModelRootCauseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get model root cause params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetModelRootCauseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get model root cause params
func (o *GetModelRootCauseParams) WithTimeout(timeout time.Duration) *GetModelRootCauseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get model root cause params
func (o *GetModelRootCauseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get model root cause params
func (o *GetModelRootCauseParams) WithContext(ctx context.Context) *GetModelRootCauseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get model root cause params
func (o *GetModelRootCauseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get model root cause params
func (o *GetModelRootCauseParams) WithHTTPClient(client *http.Client) *GetModelRootCauseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get model root cause params
func (o *GetModelRootCauseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusters adds the clusters to the get model root cause params
func (o *GetModelRootCauseParams) WithClusters(clusters *string) *GetModelRootCauseParams {
	o.SetClusters(clusters)
	return o
}

// SetClusters adds the clusters to the get model root cause params
func (o *GetModelRootCauseParams) SetClusters(clusters *string) {
	o.Clusters = clusters
}

// WithModelID adds the modelID to the get model root cause params
func (o *GetModelRootCauseParams) WithModelID(modelID string) *GetModelRootCauseParams {
	o.SetModelID(modelID)
	return o
}

// SetModelID adds the modelId to the get model root cause params
func (o *GetModelRootCauseParams) SetModelID(modelID string) {
	o.ModelID = modelID
}

// WithVectors adds the vectors to the get model root cause params
func (o *GetModelRootCauseParams) WithVectors(vectors *string) *GetModelRootCauseParams {
	o.SetVectors(vectors)
	return o
}

// SetVectors adds the vectors to the get model root cause params
func (o *GetModelRootCauseParams) SetVectors(vectors *string) {
	o.Vectors = vectors
}

// WriteToRequest writes these params to a swagger request
func (o *GetModelRootCauseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Clusters != nil {

		// query param clusters
		var qrClusters string

		if o.Clusters != nil {
			qrClusters = *o.Clusters
		}
		qClusters := qrClusters
		if qClusters != "" {

			if err := r.SetQueryParam("clusters", qClusters); err != nil {
				return err
			}
		}
	}

	// path param modelId
	if err := r.SetPathParam("modelId", o.ModelID); err != nil {
		return err
	}

	if o.Vectors != nil {

		// query param vectors
		var qrVectors string

		if o.Vectors != nil {
			qrVectors = *o.Vectors
		}
		qVectors := qrVectors
		if qVectors != "" {

			if err := r.SetQueryParam("vectors", qVectors); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
