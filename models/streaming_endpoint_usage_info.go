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

// StreamingEndpointUsageInfo streaming endpoint usage info
//
// swagger:model streamingEndpointUsageInfo
type StreamingEndpointUsageInfo struct {

	// number of calls to this endpoint during the current billing period
	// Required: true
	CallsThisPeriod *uint64 `json:"callsThisPeriod"`

	// total number of calls to this endpoint
	// Required: true
	CallsTotal *uint64 `json:"callsTotal"`

	// ISO formatted time of last call to this endpoint
	// Required: true
	LastCalled *string `json:"lastCalled"`

	// number of samples processed during the current billing period
	// Required: true
	SamplesThisPeriod *uint64 `json:"samplesThisPeriod"`

	// total number of samples processed
	// Required: true
	SamplesTotal *uint64 `json:"samplesTotal"`
}

// Validate validates this streaming endpoint usage info
func (m *StreamingEndpointUsageInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallsThisPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallsTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastCalled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSamplesThisPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSamplesTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StreamingEndpointUsageInfo) validateCallsThisPeriod(formats strfmt.Registry) error {

	if err := validate.Required("callsThisPeriod", "body", m.CallsThisPeriod); err != nil {
		return err
	}

	return nil
}

func (m *StreamingEndpointUsageInfo) validateCallsTotal(formats strfmt.Registry) error {

	if err := validate.Required("callsTotal", "body", m.CallsTotal); err != nil {
		return err
	}

	return nil
}

func (m *StreamingEndpointUsageInfo) validateLastCalled(formats strfmt.Registry) error {

	if err := validate.Required("lastCalled", "body", m.LastCalled); err != nil {
		return err
	}

	return nil
}

func (m *StreamingEndpointUsageInfo) validateSamplesThisPeriod(formats strfmt.Registry) error {

	if err := validate.Required("samplesThisPeriod", "body", m.SamplesThisPeriod); err != nil {
		return err
	}

	return nil
}

func (m *StreamingEndpointUsageInfo) validateSamplesTotal(formats strfmt.Registry) error {

	if err := validate.Required("samplesTotal", "body", m.SamplesTotal); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this streaming endpoint usage info based on context it is used
func (m *StreamingEndpointUsageInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StreamingEndpointUsageInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StreamingEndpointUsageInfo) UnmarshalBinary(b []byte) error {
	var res StreamingEndpointUsageInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
