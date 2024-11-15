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

// NewPostModelOutageParams creates a new PostModelOutageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostModelOutageParams() *PostModelOutageParams {
	return &PostModelOutageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostModelOutageParamsWithTimeout creates a new PostModelOutageParams object
// with the ability to set a timeout on a request.
func NewPostModelOutageParamsWithTimeout(timeout time.Duration) *PostModelOutageParams {
	return &PostModelOutageParams{
		timeout: timeout,
	}
}

// NewPostModelOutageParamsWithContext creates a new PostModelOutageParams object
// with the ability to set a context for a request.
func NewPostModelOutageParamsWithContext(ctx context.Context) *PostModelOutageParams {
	return &PostModelOutageParams{
		Context: ctx,
	}
}

// NewPostModelOutageParamsWithHTTPClient creates a new PostModelOutageParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostModelOutageParamsWithHTTPClient(client *http.Client) *PostModelOutageParams {
	return &PostModelOutageParams{
		HTTPClient: client,
	}
}

/*
PostModelOutageParams contains all the parameters to send to the API endpoint

	for the post model outage operation.

	Typically these are written to a http.Request.
*/
type PostModelOutageParams struct {

	// ModelID.
	ModelID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post model outage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostModelOutageParams) WithDefaults() *PostModelOutageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post model outage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostModelOutageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post model outage params
func (o *PostModelOutageParams) WithTimeout(timeout time.Duration) *PostModelOutageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post model outage params
func (o *PostModelOutageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post model outage params
func (o *PostModelOutageParams) WithContext(ctx context.Context) *PostModelOutageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post model outage params
func (o *PostModelOutageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post model outage params
func (o *PostModelOutageParams) WithHTTPClient(client *http.Client) *PostModelOutageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post model outage params
func (o *PostModelOutageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModelID adds the modelID to the post model outage params
func (o *PostModelOutageParams) WithModelID(modelID string) *PostModelOutageParams {
	o.SetModelID(modelID)
	return o
}

// SetModelID adds the modelId to the post model outage params
func (o *PostModelOutageParams) SetModelID(modelID string) {
	o.ModelID = modelID
}

// WriteToRequest writes these params to a swagger request
func (o *PostModelOutageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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