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

	"github.com/boonlogic/amber-go-sdk/models"
)

// NewPostSensorParams creates a new PostSensorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSensorParams() *PostSensorParams {
	return &PostSensorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSensorParamsWithTimeout creates a new PostSensorParams object
// with the ability to set a timeout on a request.
func NewPostSensorParamsWithTimeout(timeout time.Duration) *PostSensorParams {
	return &PostSensorParams{
		timeout: timeout,
	}
}

// NewPostSensorParamsWithContext creates a new PostSensorParams object
// with the ability to set a context for a request.
func NewPostSensorParamsWithContext(ctx context.Context) *PostSensorParams {
	return &PostSensorParams{
		Context: ctx,
	}
}

// NewPostSensorParamsWithHTTPClient creates a new PostSensorParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSensorParamsWithHTTPClient(client *http.Client) *PostSensorParams {
	return &PostSensorParams{
		HTTPClient: client,
	}
}

/* PostSensorParams contains all the parameters to send to the API endpoint
   for the post sensor operation.

   Typically these are written to a http.Request.
*/
type PostSensorParams struct {

	/* PostSensorRequest.

	   Label for new sensor instance to be created
	*/
	PostSensorRequest *models.PostSensorRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post sensor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSensorParams) WithDefaults() *PostSensorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post sensor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSensorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post sensor params
func (o *PostSensorParams) WithTimeout(timeout time.Duration) *PostSensorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post sensor params
func (o *PostSensorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post sensor params
func (o *PostSensorParams) WithContext(ctx context.Context) *PostSensorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post sensor params
func (o *PostSensorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post sensor params
func (o *PostSensorParams) WithHTTPClient(client *http.Client) *PostSensorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post sensor params
func (o *PostSensorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPostSensorRequest adds the postSensorRequest to the post sensor params
func (o *PostSensorParams) WithPostSensorRequest(postSensorRequest *models.PostSensorRequest) *PostSensorParams {
	o.SetPostSensorRequest(postSensorRequest)
	return o
}

// SetPostSensorRequest adds the postSensorRequest to the post sensor params
func (o *PostSensorParams) SetPostSensorRequest(postSensorRequest *models.PostSensorRequest) {
	o.PostSensorRequest = postSensorRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostSensorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PostSensorRequest != nil {
		if err := r.SetBodyParam(o.PostSensorRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}