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

// UserDTO user d t o
//
// swagger:model UserDTO
type UserDTO struct {

	// email
	// Format: email
	Email strfmt.Email `json:"email,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// type
	// Enum: [0,1,2,3]
	Type int64 `json:"type,omitempty"`

	// user id
	UserID uint64 `json:"user_id,omitempty"`
}

// Validate validates this user d t o
func (m *UserDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
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

func (m *UserDTO) validateEmail(formats strfmt.Registry) error {
	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

var userDTOTypeTypePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1,2,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userDTOTypeTypePropEnum = append(userDTOTypeTypePropEnum, v)
	}
}

// prop value enum
func (m *UserDTO) validateTypeEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, userDTOTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserDTO) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user d t o based on context it is used
func (m *UserDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserDTO) UnmarshalBinary(b []byte) error {
	var res UserDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
