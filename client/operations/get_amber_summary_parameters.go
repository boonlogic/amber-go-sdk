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

// NewGetAmberSummaryParams creates a new GetAmberSummaryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAmberSummaryParams() *GetAmberSummaryParams {
	return &GetAmberSummaryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAmberSummaryParamsWithTimeout creates a new GetAmberSummaryParams object
// with the ability to set a timeout on a request.
func NewGetAmberSummaryParamsWithTimeout(timeout time.Duration) *GetAmberSummaryParams {
	return &GetAmberSummaryParams{
		timeout: timeout,
	}
}

// NewGetAmberSummaryParamsWithContext creates a new GetAmberSummaryParams object
// with the ability to set a context for a request.
func NewGetAmberSummaryParamsWithContext(ctx context.Context) *GetAmberSummaryParams {
	return &GetAmberSummaryParams{
		Context: ctx,
	}
}

// NewGetAmberSummaryParamsWithHTTPClient creates a new GetAmberSummaryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAmberSummaryParamsWithHTTPClient(client *http.Client) *GetAmberSummaryParams {
	return &GetAmberSummaryParams{
		HTTPClient: client,
	}
}

/*
GetAmberSummaryParams contains all the parameters to send to the API endpoint

	for the get amber summary operation.

	Typically these are written to a http.Request.
*/
type GetAmberSummaryParams struct {

	/* SensorID.

	   Unique identifier for sensor
	*/
	SensorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get amber summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAmberSummaryParams) WithDefaults() *GetAmberSummaryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get amber summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAmberSummaryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get amber summary params
func (o *GetAmberSummaryParams) WithTimeout(timeout time.Duration) *GetAmberSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get amber summary params
func (o *GetAmberSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get amber summary params
func (o *GetAmberSummaryParams) WithContext(ctx context.Context) *GetAmberSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get amber summary params
func (o *GetAmberSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get amber summary params
func (o *GetAmberSummaryParams) WithHTTPClient(client *http.Client) *GetAmberSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get amber summary params
func (o *GetAmberSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSensorID adds the sensorID to the get amber summary params
func (o *GetAmberSummaryParams) WithSensorID(sensorID string) *GetAmberSummaryParams {
	o.SetSensorID(sensorID)
	return o
}

// SetSensorID adds the sensorId to the get amber summary params
func (o *GetAmberSummaryParams) SetSensorID(sensorID string) {
	o.SensorID = sensorID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAmberSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param sensorId
	if err := r.SetHeaderParam("sensorId", o.SensorID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
