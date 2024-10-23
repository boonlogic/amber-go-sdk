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

// NewDeleteModelParams creates a new DeleteModelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteModelParams() *DeleteModelParams {
	return &DeleteModelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteModelParamsWithTimeout creates a new DeleteModelParams object
// with the ability to set a timeout on a request.
func NewDeleteModelParamsWithTimeout(timeout time.Duration) *DeleteModelParams {
	return &DeleteModelParams{
		timeout: timeout,
	}
}

// NewDeleteModelParamsWithContext creates a new DeleteModelParams object
// with the ability to set a context for a request.
func NewDeleteModelParamsWithContext(ctx context.Context) *DeleteModelParams {
	return &DeleteModelParams{
		Context: ctx,
	}
}

// NewDeleteModelParamsWithHTTPClient creates a new DeleteModelParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteModelParamsWithHTTPClient(client *http.Client) *DeleteModelParams {
	return &DeleteModelParams{
		HTTPClient: client,
	}
}

/*
DeleteModelParams contains all the parameters to send to the API endpoint

	for the delete model operation.

	Typically these are written to a http.Request.
*/
type DeleteModelParams struct {

	// ModelID.
	ModelID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete model params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteModelParams) WithDefaults() *DeleteModelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete model params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteModelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete model params
func (o *DeleteModelParams) WithTimeout(timeout time.Duration) *DeleteModelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete model params
func (o *DeleteModelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete model params
func (o *DeleteModelParams) WithContext(ctx context.Context) *DeleteModelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete model params
func (o *DeleteModelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete model params
func (o *DeleteModelParams) WithHTTPClient(client *http.Client) *DeleteModelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete model params
func (o *DeleteModelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModelID adds the modelID to the delete model params
func (o *DeleteModelParams) WithModelID(modelID string) *DeleteModelParams {
	o.SetModelID(modelID)
	return o
}

// SetModelID adds the modelId to the delete model params
func (o *DeleteModelParams) SetModelID(modelID string) {
	o.ModelID = modelID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteModelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param modelId
	if err := r.SetPathParam("modelId", o.ModelID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
