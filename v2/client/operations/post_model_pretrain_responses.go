// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/boonlogic/amber-go-sdk/v2/models"
)

// PostModelPretrainReader is a Reader for the PostModelPretrain structure.
type PostModelPretrainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostModelPretrainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostModelPretrainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostModelPretrainAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostModelPretrainBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostModelPretrainUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostModelPretrainNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostModelPretrainInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostModelPretrainOK creates a PostModelPretrainOK with default headers values
func NewPostModelPretrainOK() *PostModelPretrainOK {
	return &PostModelPretrainOK{}
}

/*
	PostModelPretrainOK describes a response with status code 200, with default header values.

OK
*/
type PostModelPretrainOK struct {
	Payload *models.PostPretrainResponse
}

func (o *PostModelPretrainOK) Error() string {
	return fmt.Sprintf("[POST /models/{modelId}/pretrain][%d] postModelPretrainOK  %+v", 200, o.Payload)
}
func (o *PostModelPretrainOK) GetPayload() *models.PostPretrainResponse {
	return o.Payload
}

func (o *PostModelPretrainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostPretrainResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostModelPretrainAccepted creates a PostModelPretrainAccepted with default headers values
func NewPostModelPretrainAccepted() *PostModelPretrainAccepted {
	return &PostModelPretrainAccepted{}
}

/*
	PostModelPretrainAccepted describes a response with status code 202, with default header values.

Operation was accepted
*/
type PostModelPretrainAccepted struct {
	RunAsync bool
	TxnID    string

	Payload *models.PostPretrainResponse
}

func (o *PostModelPretrainAccepted) Error() string {
	return fmt.Sprintf("[POST /models/{modelId}/pretrain][%d] postModelPretrainAccepted  %+v", 202, o.Payload)
}
func (o *PostModelPretrainAccepted) GetPayload() *models.PostPretrainResponse {
	return o.Payload
}

func (o *PostModelPretrainAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Run-Async
	hdrRunAsync := response.GetHeader("Run-Async")

	if hdrRunAsync != "" {
		valrunAsync, err := swag.ConvertBool(hdrRunAsync)
		if err != nil {
			return errors.InvalidType("Run-Async", "header", "bool", hdrRunAsync)
		}
		o.RunAsync = valrunAsync
	}

	// hydrates response header txnId
	hdrTxnID := response.GetHeader("txnId")

	if hdrTxnID != "" {
		o.TxnID = hdrTxnID
	}

	o.Payload = new(models.PostPretrainResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostModelPretrainBadRequest creates a PostModelPretrainBadRequest with default headers values
func NewPostModelPretrainBadRequest() *PostModelPretrainBadRequest {
	return &PostModelPretrainBadRequest{}
}

/*
	PostModelPretrainBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostModelPretrainBadRequest struct {
	Payload *models.Error
}

func (o *PostModelPretrainBadRequest) Error() string {
	return fmt.Sprintf("[POST /models/{modelId}/pretrain][%d] postModelPretrainBadRequest  %+v", 400, o.Payload)
}
func (o *PostModelPretrainBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostModelPretrainBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostModelPretrainUnauthorized creates a PostModelPretrainUnauthorized with default headers values
func NewPostModelPretrainUnauthorized() *PostModelPretrainUnauthorized {
	return &PostModelPretrainUnauthorized{}
}

/*
	PostModelPretrainUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostModelPretrainUnauthorized struct {
	Payload *models.Error
}

func (o *PostModelPretrainUnauthorized) Error() string {
	return fmt.Sprintf("[POST /models/{modelId}/pretrain][%d] postModelPretrainUnauthorized  %+v", 401, o.Payload)
}
func (o *PostModelPretrainUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostModelPretrainUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostModelPretrainNotFound creates a PostModelPretrainNotFound with default headers values
func NewPostModelPretrainNotFound() *PostModelPretrainNotFound {
	return &PostModelPretrainNotFound{}
}

/*
	PostModelPretrainNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type PostModelPretrainNotFound struct {
	Payload *models.Error
}

func (o *PostModelPretrainNotFound) Error() string {
	return fmt.Sprintf("[POST /models/{modelId}/pretrain][%d] postModelPretrainNotFound  %+v", 404, o.Payload)
}
func (o *PostModelPretrainNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostModelPretrainNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostModelPretrainInternalServerError creates a PostModelPretrainInternalServerError with default headers values
func NewPostModelPretrainInternalServerError() *PostModelPretrainInternalServerError {
	return &PostModelPretrainInternalServerError{}
}

/*
	PostModelPretrainInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostModelPretrainInternalServerError struct {
	Payload *models.Error
}

func (o *PostModelPretrainInternalServerError) Error() string {
	return fmt.Sprintf("[POST /models/{modelId}/pretrain][%d] postModelPretrainInternalServerError  %+v", 500, o.Payload)
}
func (o *PostModelPretrainInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostModelPretrainInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
