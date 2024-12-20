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

// NewGetModelDiagnosticParams creates a new GetModelDiagnosticParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetModelDiagnosticParams() *GetModelDiagnosticParams {
	return &GetModelDiagnosticParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetModelDiagnosticParamsWithTimeout creates a new GetModelDiagnosticParams object
// with the ability to set a timeout on a request.
func NewGetModelDiagnosticParamsWithTimeout(timeout time.Duration) *GetModelDiagnosticParams {
	return &GetModelDiagnosticParams{
		timeout: timeout,
	}
}

// NewGetModelDiagnosticParamsWithContext creates a new GetModelDiagnosticParams object
// with the ability to set a context for a request.
func NewGetModelDiagnosticParamsWithContext(ctx context.Context) *GetModelDiagnosticParams {
	return &GetModelDiagnosticParams{
		Context: ctx,
	}
}

// NewGetModelDiagnosticParamsWithHTTPClient creates a new GetModelDiagnosticParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetModelDiagnosticParamsWithHTTPClient(client *http.Client) *GetModelDiagnosticParams {
	return &GetModelDiagnosticParams{
		HTTPClient: client,
	}
}

/*
GetModelDiagnosticParams contains all the parameters to send to the API endpoint

	for the get model diagnostic operation.

	Typically these are written to a http.Request.
*/
type GetModelDiagnosticParams struct {

	// ModelID.
	ModelID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get model diagnostic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetModelDiagnosticParams) WithDefaults() *GetModelDiagnosticParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get model diagnostic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetModelDiagnosticParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get model diagnostic params
func (o *GetModelDiagnosticParams) WithTimeout(timeout time.Duration) *GetModelDiagnosticParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get model diagnostic params
func (o *GetModelDiagnosticParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get model diagnostic params
func (o *GetModelDiagnosticParams) WithContext(ctx context.Context) *GetModelDiagnosticParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get model diagnostic params
func (o *GetModelDiagnosticParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get model diagnostic params
func (o *GetModelDiagnosticParams) WithHTTPClient(client *http.Client) *GetModelDiagnosticParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get model diagnostic params
func (o *GetModelDiagnosticParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModelID adds the modelID to the get model diagnostic params
func (o *GetModelDiagnosticParams) WithModelID(modelID string) *GetModelDiagnosticParams {
	o.SetModelID(modelID)
	return o
}

// SetModelID adds the modelId to the get model diagnostic params
func (o *GetModelDiagnosticParams) SetModelID(modelID string) {
	o.ModelID = modelID
}

// WriteToRequest writes these params to a swagger request
func (o *GetModelDiagnosticParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
