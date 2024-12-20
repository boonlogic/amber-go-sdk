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

// NewPostConfigParams creates a new PostConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostConfigParams() *PostConfigParams {
	return &PostConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostConfigParamsWithTimeout creates a new PostConfigParams object
// with the ability to set a timeout on a request.
func NewPostConfigParamsWithTimeout(timeout time.Duration) *PostConfigParams {
	return &PostConfigParams{
		timeout: timeout,
	}
}

// NewPostConfigParamsWithContext creates a new PostConfigParams object
// with the ability to set a context for a request.
func NewPostConfigParamsWithContext(ctx context.Context) *PostConfigParams {
	return &PostConfigParams{
		Context: ctx,
	}
}

// NewPostConfigParamsWithHTTPClient creates a new PostConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostConfigParamsWithHTTPClient(client *http.Client) *PostConfigParams {
	return &PostConfigParams{
		HTTPClient: client,
	}
}

/*
PostConfigParams contains all the parameters to send to the API endpoint

	for the post config operation.

	Typically these are written to a http.Request.
*/
type PostConfigParams struct {

	/* PostConfigRequest.

	   Sensor configuration to be applied
	*/
	PostConfigRequest *models.PostConfigRequest

	/* SensorID.

	   Unique identifier for sensor
	*/
	SensorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostConfigParams) WithDefaults() *PostConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post config params
func (o *PostConfigParams) WithTimeout(timeout time.Duration) *PostConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post config params
func (o *PostConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post config params
func (o *PostConfigParams) WithContext(ctx context.Context) *PostConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post config params
func (o *PostConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post config params
func (o *PostConfigParams) WithHTTPClient(client *http.Client) *PostConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post config params
func (o *PostConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPostConfigRequest adds the postConfigRequest to the post config params
func (o *PostConfigParams) WithPostConfigRequest(postConfigRequest *models.PostConfigRequest) *PostConfigParams {
	o.SetPostConfigRequest(postConfigRequest)
	return o
}

// SetPostConfigRequest adds the postConfigRequest to the post config params
func (o *PostConfigParams) SetPostConfigRequest(postConfigRequest *models.PostConfigRequest) {
	o.PostConfigRequest = postConfigRequest
}

// WithSensorID adds the sensorID to the post config params
func (o *PostConfigParams) WithSensorID(sensorID string) *PostConfigParams {
	o.SetSensorID(sensorID)
	return o
}

// SetSensorID adds the sensorId to the post config params
func (o *PostConfigParams) SetSensorID(sensorID string) {
	o.SensorID = sensorID
}

// WriteToRequest writes these params to a swagger request
func (o *PostConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PostConfigRequest != nil {
		if err := r.SetBodyParam(o.PostConfigRequest); err != nil {
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
