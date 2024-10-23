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

// MRecentOMs m recent o ms
//
// swagger:model m_RecentOMs
type MRecentOMs struct {

	// version number
	VersionNumber VersionNumber `json:"VersionNumber,omitempty"`

	// m values
	// Required: true
	MValues []float32 `json:"m_Values"`
}

// Validate validates this m recent o ms
func (m *MRecentOMs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersionNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MRecentOMs) validateVersionNumber(formats strfmt.Registry) error {
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

func (m *MRecentOMs) validateMValues(formats strfmt.Registry) error {

	if err := validate.Required("m_Values", "body", m.MValues); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this m recent o ms based on the context it is used
func (m *MRecentOMs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVersionNumber(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MRecentOMs) contextValidateVersionNumber(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MRecentOMs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MRecentOMs) UnmarshalBinary(b []byte) error {
	var res MRecentOMs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
