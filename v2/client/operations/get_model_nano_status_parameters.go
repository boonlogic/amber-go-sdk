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

// NewGetModelNanoStatusParams creates a new GetModelNanoStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetModelNanoStatusParams() *GetModelNanoStatusParams {
	return &GetModelNanoStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetModelNanoStatusParamsWithTimeout creates a new GetModelNanoStatusParams object
// with the ability to set a timeout on a request.
func NewGetModelNanoStatusParamsWithTimeout(timeout time.Duration) *GetModelNanoStatusParams {
	return &GetModelNanoStatusParams{
		timeout: timeout,
	}
}

// NewGetModelNanoStatusParamsWithContext creates a new GetModelNanoStatusParams object
// with the ability to set a context for a request.
func NewGetModelNanoStatusParamsWithContext(ctx context.Context) *GetModelNanoStatusParams {
	return &GetModelNanoStatusParams{
		Context: ctx,
	}
}

// NewGetModelNanoStatusParamsWithHTTPClient creates a new GetModelNanoStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetModelNanoStatusParamsWithHTTPClient(client *http.Client) *GetModelNanoStatusParams {
	return &GetModelNanoStatusParams{
		HTTPClient: client,
	}
}

/*
GetModelNanoStatusParams contains all the parameters to send to the API endpoint

	for the get model nano status operation.

	Typically these are written to a http.Request.
*/
type GetModelNanoStatusParams struct {

	// ModelID.
	ModelID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get model nano status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetModelNanoStatusParams) WithDefaults() *GetModelNanoStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get model nano status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetModelNanoStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get model nano status params
func (o *GetModelNanoStatusParams) WithTimeout(timeout time.Duration) *GetModelNanoStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get model nano status params
func (o *GetModelNanoStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get model nano status params
func (o *GetModelNanoStatusParams) WithContext(ctx context.Context) *GetModelNanoStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get model nano status params
func (o *GetModelNanoStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get model nano status params
func (o *GetModelNanoStatusParams) WithHTTPClient(client *http.Client) *GetModelNanoStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get model nano status params
func (o *GetModelNanoStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModelID adds the modelID to the get model nano status params
func (o *GetModelNanoStatusParams) WithModelID(modelID string) *GetModelNanoStatusParams {
	o.SetModelID(modelID)
	return o
}

// SetModelID adds the modelId to the get model nano status params
func (o *GetModelNanoStatusParams) SetModelID(modelID string) {
	o.ModelID = modelID
}

// WriteToRequest writes these params to a swagger request
func (o *GetModelNanoStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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