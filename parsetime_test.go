package goparsetime

import (
	"testing"
	"time"
)

func TestParsetime(t *testing.T) {

	f := "15:04:05 UTC January 02 2006"

	n := time.Now().Truncate(time.Second)
	d, err := Parsetime(n.Format(f))
	if err != nil {
		t.Fatal(err)
	}
	if d != n {
		t.Errorf("Actual: %#v, expected: %#v\n", n, d)
	}

}
