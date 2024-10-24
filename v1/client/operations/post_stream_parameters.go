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

	"github.com/boonlogic/amber-go-sdk/v1/models"
)

// NewPostStreamParams creates a new PostStreamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStreamParams() *PostStreamParams {
	return &PostStreamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStreamParamsWithTimeout creates a new PostStreamParams object
// with the ability to set a timeout on a request.
func NewPostStreamParamsWithTimeout(timeout time.Duration) *PostStreamParams {
	return &PostStreamParams{
		timeout: timeout,
	}
}

// NewPostStreamParamsWithContext creates a new PostStreamParams object
// with the ability to set a context for a request.
func NewPostStreamParamsWithContext(ctx context.Context) *PostStreamParams {
	return &PostStreamParams{
		Context: ctx,
	}
}

// NewPostStreamParamsWithHTTPClient creates a new PostStreamParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStreamParamsWithHTTPClient(client *http.Client) *PostStreamParams {
	return &PostStreamParams{
		HTTPClient: client,
	}
}

/*
PostStreamParams contains all the parameters to send to the API endpoint

	for the post stream operation.

	Typically these are written to a http.Request.
*/
type PostStreamParams struct {

	/* PostStreamRequest.

	   Data to be streamed to sensor. Should be formatted as a simple string of comma-separated numbers with no spaces (e.g. "0,0.5,1,1.5,2").
	*/
	PostStreamRequest *models.PostStreamRequest

	/* SensorID.

	   Unique identifier for sensor
	*/
	SensorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post stream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStreamParams) WithDefaults() *PostStreamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post stream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStreamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post stream params
func (o *PostStreamParams) WithTimeout(timeout time.Duration) *PostStreamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post stream params
func (o *PostStreamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post stream params
func (o *PostStreamParams) WithContext(ctx context.Context) *PostStreamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post stream params
func (o *PostStreamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post stream params
func (o *PostStreamParams) WithHTTPClient(client *http.Client) *PostStreamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post stream params
func (o *PostStreamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPostStreamRequest adds the postStreamRequest to the post stream params
func (o *PostStreamParams) WithPostStreamRequest(postStreamRequest *models.PostStreamRequest) *PostStreamParams {
	o.SetPostStreamRequest(postStreamRequest)
	return o
}

// SetPostStreamRequest adds the postStreamRequest to the post stream params
func (o *PostStreamParams) SetPostStreamRequest(postStreamRequest *models.PostStreamRequest) {
	o.PostStreamRequest = postStreamRequest
}

// WithSensorID adds the sensorID to the post stream params
func (o *PostStreamParams) WithSensorID(sensorID string) *PostStreamParams {
	o.SetSensorID(sensorID)
	return o
}

// SetSensorID adds the sensorId to the post stream params
func (o *PostStreamParams) SetSensorID(sensorID string) {
	o.SensorID = sensorID
}

// WriteToRequest writes these params to a swagger request
func (o *PostStreamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PostStreamRequest != nil {
		if err := r.SetBodyParam(o.PostStreamRequest); err != nil {
			return err
		}
	}

	// header param sensorId
	if err := r.SetHeaderParam("sensorId", o.SensorID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
