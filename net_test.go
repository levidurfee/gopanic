package gopanic

import "testing"

func TestMapErr(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		} else {
			t.Errorf("The code panicked")
		}
	}()

	err := func() error {
		return nil
	}()

	// See https://github.com/golang/go/blob/go1.20.3/src/net/lookup.go#L351
	_ = mapErr(err).Error()
}
