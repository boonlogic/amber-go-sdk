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

// NewPostModelMigrateParams creates a new PostModelMigrateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostModelMigrateParams() *PostModelMigrateParams {
	return &PostModelMigrateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostModelMigrateParamsWithTimeout creates a new PostModelMigrateParams object
// with the ability to set a timeout on a request.
func NewPostModelMigrateParamsWithTimeout(timeout time.Duration) *PostModelMigrateParams {
	return &PostModelMigrateParams{
		timeout: timeout,
	}
}

// NewPostModelMigrateParamsWithContext creates a new PostModelMigrateParams object
// with the ability to set a context for a request.
func NewPostModelMigrateParamsWithContext(ctx context.Context) *PostModelMigrateParams {
	return &PostModelMigrateParams{
		Context: ctx,
	}
}

// NewPostModelMigrateParamsWithHTTPClient creates a new PostModelMigrateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostModelMigrateParamsWithHTTPClient(client *http.Client) *PostModelMigrateParams {
	return &PostModelMigrateParams{
		HTTPClient: client,
	}
}

/*
PostModelMigrateParams contains all the parameters to send to the API endpoint

	for the post model migrate operation.

	Typically these are written to a http.Request.
*/
type PostModelMigrateParams struct {

	/* V1ModelID.

	   version 1 model id
	*/
	V1ModelID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post model migrate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostModelMigrateParams) WithDefaults() *PostModelMigrateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post model migrate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostModelMigrateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post model migrate params
func (o *PostModelMigrateParams) WithTimeout(timeout time.Duration) *PostModelMigrateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post model migrate params
func (o *PostModelMigrateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post model migrate params
func (o *PostModelMigrateParams) WithContext(ctx context.Context) *PostModelMigrateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post model migrate params
func (o *PostModelMigrateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post model migrate params
func (o *PostModelMigrateParams) WithHTTPClient(client *http.Client) *PostModelMigrateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post model migrate params
func (o *PostModelMigrateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1ModelID adds the v1ModelID to the post model migrate params
func (o *PostModelMigrateParams) WithV1ModelID(v1ModelID string) *PostModelMigrateParams {
	o.SetV1ModelID(v1ModelID)
	return o
}

// SetV1ModelID adds the v1ModelId to the post model migrate params
func (o *PostModelMigrateParams) SetV1ModelID(v1ModelID string) {
	o.V1ModelID = v1ModelID
}

// WriteToRequest writes these params to a swagger request
func (o *PostModelMigrateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param v1ModelId
	if err := r.SetPathParam("v1ModelId", o.V1ModelID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
