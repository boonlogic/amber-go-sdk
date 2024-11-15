// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetPretrainResponse get pretrain response
//
// swagger:model getPretrainResponse
type GetPretrainResponse struct {

	// latest pretrain message
	Message string `json:"message,omitempty"`

	// state of pretraining, one of: None, Chunking, Pretraining, Pretrained, Error
	// Required: true
	State *string `json:"state"`
}

// Validate validates this get pretrain response
func (m *GetPretrainResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetPretrainResponse) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get pretrain response based on context it is used
func (m *GetPretrainResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetPretrainResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetPretrainResponse) UnmarshalBinary(b []byte) error {
	var res GetPretrainResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}