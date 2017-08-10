/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
   

package models

 
 

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// KeyGetResponse key get response
// swagger:model KeyGetResponse
type KeyGetResponse struct {
	KeyCreate

	// Id of the key.
	ID strfmt.UUID `json:"id,omitempty"`

	// Identifies what kind of resource this is. Value: the fixed string "weaviate#keyGetResponse".
	Kind *string `json:"kind,omitempty"`

	// Parent key. A parent allways has access to a child. Root key has parent value 0. Only a user with a root of 0 can set a root key.
	Parent string `json:"parent,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *KeyGetResponse) UnmarshalJSON(raw []byte) error {

	var aO0 KeyCreate
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.KeyCreate = aO0

	var data struct {
		ID strfmt.UUID `json:"id,omitempty"`

		Kind *string `json:"kind,omitempty"`

		Parent string `json:"parent,omitempty"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.ID = data.ID

	m.Kind = data.Kind

	m.Parent = data.Parent

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m KeyGetResponse) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.KeyCreate)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var data struct {
		ID strfmt.UUID `json:"id,omitempty"`

		Kind *string `json:"kind,omitempty"`

		Parent string `json:"parent,omitempty"`
	}

	data.ID = m.ID

	data.Kind = m.Kind

	data.Parent = m.Parent

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this key get response
func (m *KeyGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.KeyCreate.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *KeyGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyGetResponse) UnmarshalBinary(b []byte) error {
	var res KeyGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
