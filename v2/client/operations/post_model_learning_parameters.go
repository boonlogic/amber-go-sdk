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

// NewPostModelLearningParams creates a new PostModelLearningParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostModelLearningParams() *PostModelLearningParams {
	return &PostModelLearningParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostModelLearningParamsWithTimeout creates a new PostModelLearningParams object
// with the ability to set a timeout on a request.
func NewPostModelLearningParamsWithTimeout(timeout time.Duration) *PostModelLearningParams {
	return &PostModelLearningParams{
		timeout: timeout,
	}
}

// NewPostModelLearningParamsWithContext creates a new PostModelLearningParams object
// with the ability to set a context for a request.
func NewPostModelLearningParamsWithContext(ctx context.Context) *PostModelLearningParams {
	return &PostModelLearningParams{
		Context: ctx,
	}
}

// NewPostModelLearningParamsWithHTTPClient creates a new PostModelLearningParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostModelLearningParamsWithHTTPClient(client *http.Client) *PostModelLearningParams {
	return &PostModelLearningParams{
		HTTPClient: client,
	}
}

/*
PostModelLearningParams contains all the parameters to send to the API endpoint

	for the post model learning operation.

	Typically these are written to a http.Request.
*/
type PostModelLearningParams struct {

	// ModelID.
	ModelID string

	/* PostLearningRequest.

	   updates to apply
	*/
	PostLearningRequest *models.PostLearningRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post model learning params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostModelLearningParams) WithDefaults() *PostModelLearningParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post model learning params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostModelLearningParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post model learning params
func (o *PostModelLearningParams) WithTimeout(timeout time.Duration) *PostModelLearningParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post model learning params
func (o *PostModelLearningParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post model learning params
func (o *PostModelLearningParams) WithContext(ctx context.Context) *PostModelLearningParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post model learning params
func (o *PostModelLearningParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post model learning params
func (o *PostModelLearningParams) WithHTTPClient(client *http.Client) *PostModelLearningParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post model learning params
func (o *PostModelLearningParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModelID adds the modelID to the post model learning params
func (o *PostModelLearningParams) WithModelID(modelID string) *PostModelLearningParams {
	o.SetModelID(modelID)
	return o
}

// SetModelID adds the modelId to the post model learning params
func (o *PostModelLearningParams) SetModelID(modelID string) {
	o.ModelID = modelID
}

// WithPostLearningRequest adds the postLearningRequest to the post model learning params
func (o *PostModelLearningParams) WithPostLearningRequest(postLearningRequest *models.PostLearningRequest) *PostModelLearningParams {
	o.SetPostLearningRequest(postLearningRequest)
	return o
}

// SetPostLearningRequest adds the postLearningRequest to the post model learning params
func (o *PostModelLearningParams) SetPostLearningRequest(postLearningRequest *models.PostLearningRequest) {
	o.PostLearningRequest = postLearningRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostModelLearningParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param modelId
	if err := r.SetPathParam("modelId", o.ModelID); err != nil {
		return err
	}
	if o.PostLearningRequest != nil {
		if err := r.SetBodyParam(o.PostLearningRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
