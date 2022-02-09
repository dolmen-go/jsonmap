//    Copyright 2021 Olivier Mengué
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package jsonmap provides tools to serialize a map as JSON using a given order for keys.
package jsonmap

import (
	"bytes"
	"encoding/json"
	"errors"
	"sort"
)

// Ordered wraps a map[string]interface{}.
type Ordered struct {
	Order []string // Order of keys
	Data  map[string]interface{}
}

// MarshalJSON implements interface encoding/json.Marshaler.
// MarshalJSON returns o.Data serialized as JSON.
// The keys in o.Order are serialized first.
// If keys not in o.Order remain, they are sorted.
func (o Ordered) MarshalJSON() ([]byte, error) {
	if o.Data == nil {
		return []byte("null"), nil
	}
	if len(o.Data) == 0 {
		return []byte("{}"), nil
	}
	if len(o.Order) == 0 {
		return json.Marshal(o.Data)
	}

	var buf bytes.Buffer
	seen := make([]int, 0, len(o.Order))

	buf.WriteByte('{')
	for i, key := range o.Order {
		val, exists := o.Data[key]
		if !exists {
			continue
		}

		if len(seen) > 0 {
			buf.WriteByte(',')
		}
		seen = append(seen, i)

		k, err := json.Marshal(key)
		if err != nil {
			return nil, err
		}
		buf.Write(k)
		buf.WriteByte(':')
		v, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}
		buf.Write(v)
	}
	if len(seen) < len(o.Data) {
		if len(seen) == 0 {
			return json.Marshal(o.Data)
		}

		others := make([]string, 0, len(o.Data)-len(seen))
	FindMissingKeys:
		for key := range o.Data {
			for i := range seen {
				if o.Order[seen[i]] == key {
					continue FindMissingKeys
				}
			}
			// Not seen

			others = append(others, key)
		}

		sort.Strings(others)

		start := len(seen) == 0
		for _, key := range others {
			k, err := json.Marshal(key)
			if err != nil {
				return nil, err
			}
			v, err := json.Marshal(o.Data[key])
			if err != nil {
				return nil, err
			}
			if !start {
				buf.WriteByte(',')
				start = false
			}
			buf.Write(k)
			buf.WriteByte(':')
			buf.Write(v)
		}
	}

	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func (o *Ordered) UnmarshalJSON(b []byte) error {
	if len(b) < 6 { // {"":0}
		switch b[0] {
		case 'n':
			if len(b) == 4 && b[1] == 'u' && b[2] == 'l' && b[3] == 'l' {
				o.Data = nil
				return nil
			}
		case '{':
			if len(b) == 2 && b[1] == '}' {
				o.Data = map[string]interface{}{}
				return nil
			}
		}
	}
	dec := json.NewDecoder(bytes.NewBuffer(b))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return err
	}
	if tok == nil {
		o.Data = nil
		return nil
	}
	if tok != json.Delim('{') {
		return errors.New("object expected") // FIXME
	}
	if o.Order != nil {
		o.Order = o.Order[:0]
	}
	o.Data = make(map[string]interface{})
	for dec.More() {
		tok, err = dec.Token()
		if err != nil {
			return err
		}
		key := tok.(string)
		var val interface{}
		if err = dec.Decode(&val); err != nil {
			return err
		}
		o.Order = append(o.Order, key)
		o.Data[key] = val
	}
	return nil
}
