// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StreamingParameters streaming parameters
//
// swagger:model streamingParameters
type StreamingParameters struct {

	// the number of samples to use when calculating AH
	AnomalyHistoryWindow *uint32 `json:"anomalyHistoryWindow,omitempty"`

	// learning graduation requirement for stopping learning upon reaching this cluster count
	LearningMaxClusters *uint16 `json:"learningMaxClusters,omitempty"`

	// learning graduation requirement for stopping learning after acquiring this many samples
	LearningMaxSamples *uint64 `json:"learningMaxSamples,omitempty"`

	// enables graduation requirements for learning
	LearningRateDenominator *uint64 `json:"learningRateDenominator,omitempty"`

	// enables graduation requirements for learning
	LearningRateNumerator *uint64 `json:"learningRateNumerator,omitempty"`
}

// Validate validates this streaming parameters
func (m *StreamingParameters) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this streaming parameters based on context it is used
func (m *StreamingParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StreamingParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StreamingParameters) UnmarshalBinary(b []byte) error {
	var res StreamingParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}