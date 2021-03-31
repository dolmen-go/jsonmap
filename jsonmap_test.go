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

package jsonmap_test

import (
	"encoding/json"
	"testing"

	"github.com/dolmen-go/jsonmap"
)

func TestOrdered(t *testing.T) {
	orig := []byte(`{
		"name": "John",
		"age": 20
	  }`)

	var m jsonmap.Ordered
	if err := json.Unmarshal(orig, &m); err != nil {
		t.Fatal(err)
	}
	out, err := json.Marshal(&m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", out)

	m.Order = nil
	out, err = json.Marshal(&m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", out)
}
