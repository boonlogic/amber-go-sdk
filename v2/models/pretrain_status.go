// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PretrainStatus pretrain status
//
// swagger:model pretrainStatus
type PretrainStatus struct {

	// Pretraining status description.
	Message string `json:"message,omitempty"`

	// Pretraining status of the model. One of: `None`, `Chunking`, `Pretraining`, `Pretrained`.
	// Enum: [None Chunking Pretraining Pretrained]
	Status string `json:"status,omitempty"`
}

// Validate validates this pretrain status
func (m *PretrainStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var pretrainStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Chunking","Pretraining","Pretrained"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pretrainStatusTypeStatusPropEnum = append(pretrainStatusTypeStatusPropEnum, v)
	}
}

const (

	// PretrainStatusStatusNone captures enum value "None"
	PretrainStatusStatusNone string = "None"

	// PretrainStatusStatusChunking captures enum value "Chunking"
	PretrainStatusStatusChunking string = "Chunking"

	// PretrainStatusStatusPretraining captures enum value "Pretraining"
	PretrainStatusStatusPretraining string = "Pretraining"

	// PretrainStatusStatusPretrained captures enum value "Pretrained"
	PretrainStatusStatusPretrained string = "Pretrained"
)

// prop value enum
func (m *PretrainStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pretrainStatusTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PretrainStatus) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this pretrain status based on context it is used
func (m *PretrainStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PretrainStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PretrainStatus) UnmarshalBinary(b []byte) error {
	var res PretrainStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
