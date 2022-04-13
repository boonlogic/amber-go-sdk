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

// PutStreamResponse put stream response
//
// swagger:model putStreamResponse
type PutStreamResponse struct {

	// streaming results when available
	Results *PostStreamResponse `json:"results,omitempty"`

	// updated sensor fusion vector
	// Required: true
	Vector MayContainNullsArray `json:"vector"`

	// updated sensor fusion vector as a string of comma-separated values
	// Required: true
	VectorCSV *string `json:"vectorCSV"`
}

// Validate validates this put stream response
func (m *PutStreamResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVectorCSV(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutStreamResponse) validateResults(formats strfmt.Registry) error {
	if swag.IsZero(m.Results) { // not required
		return nil
	}

	if m.Results != nil {
		if err := m.Results.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("results")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("results")
			}
			return err
		}
	}

	return nil
}

func (m *PutStreamResponse) validateVector(formats strfmt.Registry) error {

	if err := validate.Required("vector", "body", m.Vector); err != nil {
		return err
	}

	if err := m.Vector.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vector")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("vector")
		}
		return err
	}

	return nil
}

func (m *PutStreamResponse) validateVectorCSV(formats strfmt.Registry) error {

	if err := validate.Required("vectorCSV", "body", m.VectorCSV); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this put stream response based on the context it is used
func (m *PutStreamResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutStreamResponse) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	if m.Results != nil {
		if err := m.Results.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("results")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("results")
			}
			return err
		}
	}

	return nil
}

func (m *PutStreamResponse) contextValidateVector(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Vector.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vector")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("vector")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PutStreamResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutStreamResponse) UnmarshalBinary(b []byte) error {
	var res PutStreamResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
