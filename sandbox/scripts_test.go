package sandbox

import "testing"

func TestCreateScript(t *testing.T) {
	s := createScript("test", "Hello {{ .Name }} !", map[string]interface{}{"Name": "World"})
	if s != "Hello World !" {
		t.Error("failed")
	}
}
