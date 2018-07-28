package read

import (
	"testing"
	"plugin"
)

func TestMain(m *testing.M) {
	m.Run()
}

func Test_read_and_apply(t *testing.T) {
	p, err := plugin.Open("plugin-read-test.so")
	if err != nil {
		t.Fatal(err)
	}
	m0 := make(map[string]interface{})
	m0["one"] = "uno"
	m0["two"] = "dos"
	m0["three"] = []string{"tres", "trois"}
	save(p, m0)
	m := make(map[string]interface{})
	apply(p, m)

	if m["one"] != "uno" {
		t.Errorf("expected: %v, actual: %v", "uno", m["one"])
	}
	if m["two"] != "dos" {
		t.Errorf("expected: %v, actual: %v", "dos", m["two"])
	}
	if m["three"].([]string)[0] != "tres" || m["three"].([]string)[1] != "trois" {
		t.Errorf("expected: %v, actual: %v", "three", m["three"])
	}
}
