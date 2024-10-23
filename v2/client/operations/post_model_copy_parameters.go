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

// NewPostModelCopyParams creates a new PostModelCopyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostModelCopyParams() *PostModelCopyParams {
	return &PostModelCopyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostModelCopyParamsWithTimeout creates a new PostModelCopyParams object
// with the ability to set a timeout on a request.
func NewPostModelCopyParamsWithTimeout(timeout time.Duration) *PostModelCopyParams {
	return &PostModelCopyParams{
		timeout: timeout,
	}
}

// NewPostModelCopyParamsWithContext creates a new PostModelCopyParams object
// with the ability to set a context for a request.
func NewPostModelCopyParamsWithContext(ctx context.Context) *PostModelCopyParams {
	return &PostModelCopyParams{
		Context: ctx,
	}
}

// NewPostModelCopyParamsWithHTTPClient creates a new PostModelCopyParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostModelCopyParamsWithHTTPClient(client *http.Client) *PostModelCopyParams {
	return &PostModelCopyParams{
		HTTPClient: client,
	}
}

/*
PostModelCopyParams contains all the parameters to send to the API endpoint

	for the post model copy operation.

	Typically these are written to a http.Request.
*/
type PostModelCopyParams struct {

	// ModelID.
	ModelID string

	/* PostModelCopyRequest.

	   copy request params
	*/
	PostModelCopyRequest *models.PostModelCopyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post model copy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostModelCopyParams) WithDefaults() *PostModelCopyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post model copy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostModelCopyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post model copy params
func (o *PostModelCopyParams) WithTimeout(timeout time.Duration) *PostModelCopyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post model copy params
func (o *PostModelCopyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post model copy params
func (o *PostModelCopyParams) WithContext(ctx context.Context) *PostModelCopyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post model copy params
func (o *PostModelCopyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post model copy params
func (o *PostModelCopyParams) WithHTTPClient(client *http.Client) *PostModelCopyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post model copy params
func (o *PostModelCopyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModelID adds the modelID to the post model copy params
func (o *PostModelCopyParams) WithModelID(modelID string) *PostModelCopyParams {
	o.SetModelID(modelID)
	return o
}

// SetModelID adds the modelId to the post model copy params
func (o *PostModelCopyParams) SetModelID(modelID string) {
	o.ModelID = modelID
}

// WithPostModelCopyRequest adds the postModelCopyRequest to the post model copy params
func (o *PostModelCopyParams) WithPostModelCopyRequest(postModelCopyRequest *models.PostModelCopyRequest) *PostModelCopyParams {
	o.SetPostModelCopyRequest(postModelCopyRequest)
	return o
}

// SetPostModelCopyRequest adds the postModelCopyRequest to the post model copy params
func (o *PostModelCopyParams) SetPostModelCopyRequest(postModelCopyRequest *models.PostModelCopyRequest) {
	o.PostModelCopyRequest = postModelCopyRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostModelCopyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param modelId
	if err := r.SetPathParam("modelId", o.ModelID); err != nil {
		return err
	}
	if o.PostModelCopyRequest != nil {
		if err := r.SetBodyParam(o.PostModelCopyRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
