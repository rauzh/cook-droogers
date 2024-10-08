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

// RequestDTO request d t o
//
// swagger:model RequestDTO
type RequestDTO struct {

	// applier id
	ApplierID uint64 `json:"applier_id,omitempty"`

	// date
	// Format: date
	Date strfmt.Date `json:"date,omitempty"`

	// manager id
	ManagerID uint64 `json:"manager_id,omitempty"`

	// request id
	RequestID uint64 `json:"request_id,omitempty"`

	// status
	// Enum: ["New","Processing","On approval","Closed"]
	Status string `json:"status,omitempty"`

	// type
	// Enum: ["Publish","Sign"]
	Type string `json:"type,omitempty"`
}

// Validate validates this request d t o
func (m *RequestDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestDTO) validateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("date", "body", "date", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

var requestDTOTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["New","Processing","On approval","Closed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		requestDTOTypeStatusPropEnum = append(requestDTOTypeStatusPropEnum, v)
	}
}

const (

	// RequestDTOStatusNew captures enum value "New"
	RequestDTOStatusNew string = "New"

	// RequestDTOStatusProcessing captures enum value "Processing"
	RequestDTOStatusProcessing string = "Processing"

	// RequestDTOStatusOnApproval captures enum value "On approval"
	RequestDTOStatusOnApproval string = "On approval"

	// RequestDTOStatusClosed captures enum value "Closed"
	RequestDTOStatusClosed string = "Closed"
)

// prop value enum
func (m *RequestDTO) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, requestDTOTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RequestDTO) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var requestDTOTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Publish","Sign"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		requestDTOTypeTypePropEnum = append(requestDTOTypeTypePropEnum, v)
	}
}

const (

	// RequestDTOTypePublish captures enum value "Publish"
	RequestDTOTypePublish string = "Publish"

	// RequestDTOTypeSign captures enum value "Sign"
	RequestDTOTypeSign string = "Sign"
)

// prop value enum
func (m *RequestDTO) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, requestDTOTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RequestDTO) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this request d t o based on context it is used
func (m *RequestDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestDTO) UnmarshalBinary(b []byte) error {
	var res RequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
