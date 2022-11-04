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

// PostStreamResponse post stream response
//
// swagger:model postStreamResponse
type PostStreamResponse struct {
	StreamStatus

	// a d
	// Required: true
	AD Uint16Array `json:"AD"`

	// a h
	// Required: true
	AH Uint16Array `json:"AH"`

	// a m
	// Required: true
	AM Float32Array `json:"AM"`

	// a w
	// Required: true
	AW Uint16Array `json:"AW"`

	// ID
	// Required: true
	ID Int32Array `json:"ID"`

	// n i
	// Required: true
	NI Uint16Array `json:"NI"`

	// n s
	// Required: true
	NS Uint16Array `json:"NS"`

	// n w
	// Required: true
	NW Float32Array `json:"NW"`

	// r i
	// Required: true
	RI Uint16Array `json:"RI"`

	// s i
	// Required: true
	SI Uint16Array `json:"SI"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PostStreamResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 StreamStatus
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.StreamStatus = aO0

	// AO1
	var dataAO1 struct {
		AD Uint16Array `json:"AD"`

		AH Uint16Array `json:"AH"`

		AM Float32Array `json:"AM"`

		AW Uint16Array `json:"AW"`

		ID Int32Array `json:"ID"`

		NI Uint16Array `json:"NI"`

		NS Uint16Array `json:"NS"`

		NW Float32Array `json:"NW"`

		RI Uint16Array `json:"RI"`

		SI Uint16Array `json:"SI"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AD = dataAO1.AD

	m.AH = dataAO1.AH

	m.AM = dataAO1.AM

	m.AW = dataAO1.AW

	m.ID = dataAO1.ID

	m.NI = dataAO1.NI

	m.NS = dataAO1.NS

	m.NW = dataAO1.NW

	m.RI = dataAO1.RI

	m.SI = dataAO1.SI

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PostStreamResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.StreamStatus)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		AD Uint16Array `json:"AD"`

		AH Uint16Array `json:"AH"`

		AM Float32Array `json:"AM"`

		AW Uint16Array `json:"AW"`

		ID Int32Array `json:"ID"`

		NI Uint16Array `json:"NI"`

		NS Uint16Array `json:"NS"`

		NW Float32Array `json:"NW"`

		RI Uint16Array `json:"RI"`

		SI Uint16Array `json:"SI"`
	}

	dataAO1.AD = m.AD

	dataAO1.AH = m.AH

	dataAO1.AM = m.AM

	dataAO1.AW = m.AW

	dataAO1.ID = m.ID

	dataAO1.NI = m.NI

	dataAO1.NS = m.NS

	dataAO1.NW = m.NW

	dataAO1.RI = m.RI

	dataAO1.SI = m.SI

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post stream response
func (m *PostStreamResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with StreamStatus
	if err := m.StreamStatus.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAD(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAH(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAM(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAW(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNW(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostStreamResponse) validateAD(formats strfmt.Registry) error {

	if err := validate.Required("AD", "body", m.AD); err != nil {
		return err
	}

	if err := m.AD.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AD")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AD")
		}
		return err
	}

	return nil
}

func (m *PostStreamResponse) validateAH(formats strfmt.Registry) error {

	if err := validate.Required("AH", "body", m.AH); err != nil {
		return err
	}

	if err := m.AH.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AH")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AH")
		}
		return err
	}

	return nil
}

func (m *PostStreamResponse) validateAM(formats strfmt.Registry) error {

	if err := validate.Required("AM", "body", m.AM); err != nil {
		return err
	}

	if err := m.AM.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AM")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AM")
		}
		return err
	}

	return nil
}

func (m *PostStreamResponse) validateAW(formats strfmt.Registry) error {

	if err := validate.Required("AW", "body", m.AW); err != nil {
		return err
	}

	if err := m.AW.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AW")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AW")
		}
		return err
	}

	return nil
}

func (m *PostStreamResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("ID", "body", m.ID); err != nil {
		return err
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ID")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ID")
		}
		return err
	}

	return nil
}

func (m *PostStreamResponse) validateNI(formats strfmt.Registry) error {

	if err := validate.Required("NI", "body", m.NI); err != nil {
		return err
	}

	if err := m.NI.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("NI")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("NI")
		}
		return err
	}

	return nil
}

func (m *PostStreamResponse) validateNS(formats strfmt.Registry) error {

	if err := validate.Required("NS", "body", m.NS); err != nil {
		return err
	}

	if err := m.NS.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("NS")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("NS")
		}
		return err
	}

	return nil
}

func (m *PostStreamResponse) validateNW(formats strfmt.Registry) error {

	if err := validate.Required("NW", "body", m.NW); err != nil {
		return err
	}

	if err := m.NW.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("NW")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("NW")
		}
		return err
	}

	return nil
}

func (m *PostStreamResponse) validateRI(formats strfmt.Registry) error {

	if err := validate.Required("RI", "body", m.RI); err != nil {
		return err
	}

	if err := m.RI.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("RI")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("RI")
		}
		return err
	}

	return nil
}

func (m *PostStreamResponse) validateSI(formats strfmt.Registry) error {

	if err := validate.Required("SI", "body", m.SI); err != nil {
		return err
	}

	if err := m.SI.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SI")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("SI")
		}
		return err
	}

	return nil
}

// ContextValidate validate this post stream response based on the context it is used
func (m *PostStreamResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with StreamStatus
	if err := m.StreamStatus.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAD(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAH(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAM(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAW(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNW(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostStreamResponse) contextValidateAD(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AD.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AD")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AD")
		}
		return err
	}

	return nil
}

func (m *PostStreamResponse) contextValidateAH(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AH.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AH")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AH")
		}
		return err
	}

	return nil
}

func (m *PostStreamResponse) contextValidateAM(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AM.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AM")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AM")
		}
		return err
	}

	return nil
}

func (m *PostStreamResponse) contextValidateAW(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AW.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AW")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AW")
		}
		return err
	}

	return nil
}

func (m *PostStreamResponse) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ID")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ID")
		}
		return err
	}

	return nil
}

func (m *PostStreamResponse) contextValidateNI(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NI.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("NI")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("NI")
		}
		return err
	}

	return nil
}

func (m *PostStreamResponse) contextValidateNS(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NS.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("NS")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("NS")
		}
		return err
	}

	return nil
}

func (m *PostStreamResponse) contextValidateNW(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NW.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("NW")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("NW")
		}
		return err
	}

	return nil
}

func (m *PostStreamResponse) contextValidateRI(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RI.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("RI")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("RI")
		}
		return err
	}

	return nil
}

func (m *PostStreamResponse) contextValidateSI(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SI.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SI")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("SI")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostStreamResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostStreamResponse) UnmarshalBinary(b []byte) error {
	var res PostStreamResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
