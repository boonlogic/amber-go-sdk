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

	"github.com/boonlogic/amber-go-sdk/v2/models"
)

// NewPostModelParams creates a new PostModelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostModelParams() *PostModelParams {
	return &PostModelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostModelParamsWithTimeout creates a new PostModelParams object
// with the ability to set a timeout on a request.
func NewPostModelParamsWithTimeout(timeout time.Duration) *PostModelParams {
	return &PostModelParams{
		timeout: timeout,
	}
}

// NewPostModelParamsWithContext creates a new PostModelParams object
// with the ability to set a context for a request.
func NewPostModelParamsWithContext(ctx context.Context) *PostModelParams {
	return &PostModelParams{
		Context: ctx,
	}
}

// NewPostModelParamsWithHTTPClient creates a new PostModelParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostModelParamsWithHTTPClient(client *http.Client) *PostModelParams {
	return &PostModelParams{
		HTTPClient: client,
	}
}

/*
PostModelParams contains all the parameters to send to the API endpoint

	for the post model operation.

	Typically these are written to a http.Request.
*/
type PostModelParams struct {

	/* PostModelRequest.

	   initial metadata for new model
	*/
	PostModelRequest *models.PostModelRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post model params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostModelParams) WithDefaults() *PostModelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post model params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostModelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post model params
func (o *PostModelParams) WithTimeout(timeout time.Duration) *PostModelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post model params
func (o *PostModelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post model params
func (o *PostModelParams) WithContext(ctx context.Context) *PostModelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post model params
func (o *PostModelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post model params
func (o *PostModelParams) WithHTTPClient(client *http.Client) *PostModelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post model params
func (o *PostModelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPostModelRequest adds the postModelRequest to the post model params
func (o *PostModelParams) WithPostModelRequest(postModelRequest *models.PostModelRequest) *PostModelParams {
	o.SetPostModelRequest(postModelRequest)
	return o
}

// SetPostModelRequest adds the postModelRequest to the post model params
func (o *PostModelParams) SetPostModelRequest(postModelRequest *models.PostModelRequest) {
	o.PostModelRequest = postModelRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostModelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PostModelRequest != nil {
		if err := r.SetBodyParam(o.PostModelRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
