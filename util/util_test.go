package util

import (
	"testing"
)

func TestStoreAndLoad(t *testing.T) {
	test := struct {
		Field string
		Value string
	}{
		"Ramon",
		"valor",
	}
	err := Store(test, "/tmp/test.gob")
	if err != nil {
		t.Errorf("The struct cant be store: %v", err)
	}
	test2 := struct {
		Field string
		Value string
	}{}
	err = Load(&test2, "/tmp/test.gob")
	if err != nil {
		t.Errorf("The struct cant be load: %v", err)
	}

	if test2.Field != "Ramon" {
		t.Errorf("The value of the struct is wrong")
	}
}
