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
