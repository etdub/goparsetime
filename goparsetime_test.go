package goparsetime

import (
	"fmt"
	"testing"
	"time"
)

func TestParsetime(t *testing.T) {

	f := "15:04:05 January 02 2006"
	testDate := time.Now().Truncate(time.Second)
	parseDate, err := Parsetime(testDate.Format(f))
	if err != nil {
		t.Fatal(err)
	}

	if testDate != parseDate {
		t.Errorf("Actual: %#v, expected: %#v\n", parseDate, testDate)
	}

	parseDate, err = Parsetime("18:00 UTC May 5 2003")
	if err != nil {
		t.Fatal(err)
	}
	if parseDate.Unix() != 1052157600 {
		t.Errorf("Actual %d, expected %d", parseDate.Unix(), 1052157600)
	}

	local, _ := time.LoadLocation("Local")
	testDate = time.Date(2015, 12, 25, 10, 0, 0, 0, local)
	parseDate, err = Parsetime(testDate.Format(f))
	if err != nil {
		t.Fatal(err)
	}
	if testDate != parseDate {
		t.Errorf("Actual: %#v, expected: %#v\n", parseDate, testDate)
	}

	// Date addition
	dateAdd := fmt.Sprintf("%s + 5 days", testDate.Format(f))
	parseDate, err = Parsetime(dateAdd)
	if err != nil {
		t.Fatal(err)
	}
	expectedDate := testDate.Add(time.Duration(120 * time.Hour))
	if expectedDate != parseDate {
		t.Errorf("Actual: %#v, expected: %#v\n", parseDate, expectedDate)
	}
}
