package path

import (
	"testing"
)

func TestGetRoot(t *testing.T) {
	g, err := GetRoot()
	t.Log(g)
	if err != nil {
		t.Error(err)
		return
	}

}
