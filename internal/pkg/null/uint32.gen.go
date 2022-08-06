// Code generated by ./internal/cmd/gen/null/main.go. DO NOT EDIT.

// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

package null

import (
	"github.com/goccy/go-json"
)

var _ json.Unmarshaler = (*Uint32)(nil)

// Uint32 is a nullable type.
type Uint32 struct {
	Value uint32
	Set   bool // if json object has this field
}

// NewUint32 creates a new uint32.
func NewUint32(v uint32) Uint32 {
	return Uint32{
		Value: v,
		Set:   true,
	}
}

func (t Uint32) Ptr() *uint32 {
	if t.Set {
		return &t.Value
	}

	return nil
}

// Default return default value its value is Null or not Set.
func (t Uint32) Default(v uint32) uint32 {
	if t.Set {
		return t.Value
	}

	return v
}

func (t Uint32) Interface() any {
	if t.Set {
		return t.Value
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Uint32) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	t.Set = true
	return json.UnmarshalNoEscape(data, &t.Value) //nolint:wrapcheck
}
