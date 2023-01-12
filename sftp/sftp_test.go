package sftp

import "testing"

func TestRandomKey(t *testing.T) {
	if err:= RandomKey(); err != nil {
		t.Fail()
	}
}