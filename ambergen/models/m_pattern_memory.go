// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MPatternMemory m pattern memory
//
// swagger:model m_PatternMemory
type MPatternMemory struct {

	// version number
	VersionNumber VersionNumber `json:"VersionNumber,omitempty"`
}

// Validate validates this m pattern memory
func (m *MPatternMemory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersionNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MPatternMemory) validateVersionNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.VersionNumber) { // not required
		return nil
	}

	if err := m.VersionNumber.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("VersionNumber")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("VersionNumber")
		}
		return err
	}

	return nil
}

// ContextValidate validate this m pattern memory based on the context it is used
func (m *MPatternMemory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVersionNumber(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MPatternMemory) contextValidateVersionNumber(ctx context.Context, formats strfmt.Registry) error {

	if err := m.VersionNumber.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("VersionNumber")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("VersionNumber")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MPatternMemory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MPatternMemory) UnmarshalBinary(b []byte) error {
	var res MPatternMemory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}