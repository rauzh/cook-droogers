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

// StatsDTO stats d t o
//
// swagger:model StatsDTO
type StatsDTO struct {

	// date
	// Format: date
	Date strfmt.Date `json:"date,omitempty"`

	// likes
	Likes uint64 `json:"likes,omitempty"`

	// stat id
	StatID uint64 `json:"stat_id,omitempty"`

	// streams
	Streams uint64 `json:"streams,omitempty"`

	// track id
	TrackID uint64 `json:"track_id,omitempty"`
}

// Validate validates this stats d t o
func (m *StatsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatsDTO) validateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("date", "body", "date", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this stats d t o based on context it is used
func (m *StatsDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StatsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatsDTO) UnmarshalBinary(b []byte) error {
	var res StatsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}