package goparsetime

import (
	"fmt"
	"testing"
	"time"
)

func testParsetimeInvalid(t *testing.T) {
	_, err := Parsetime("invalid string")
	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestParsetimeDateFormat(t *testing.T) {
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
}

func TestParsetimeDayAddition(t *testing.T) {
	// Date addition
	f := "15:04:05 January 02 2006"
	testDate := time.Now().Truncate(time.Second)
	dateAdd := fmt.Sprintf("%s + 5 days", testDate.Format(f))
	parseDate, err := Parsetime(dateAdd)
	if err != nil {
		t.Fatal(err)
	}
	expectedDate := testDate.Add(time.Duration(120 * time.Hour))
	if expectedDate != parseDate {
		t.Errorf("Actual: %#v, expected: %#v\n", parseDate, expectedDate)
	}
}

func TestParsetimeMultipleAddition(t *testing.T) {
	f := "15:04:05 January 02 2006"
	testDate := time.Now().Truncate(time.Second)
	dateAdd := fmt.Sprintf("%s + 5minutes - 10 seconds", testDate.Format(f))
	parseDate, err := Parsetime(dateAdd)
	if err != nil {
		t.Fatal(err)
	}
	expectedDate := testDate.Add(time.Duration(5 * time.Minute)).Add(time.Duration(-10 * time.Second))
	if expectedDate != parseDate {
		t.Errorf("Actual: %#v, expected: %#v\n", parseDate, expectedDate)
	}
}

func TestParsetimeNow(t *testing.T) {
	before := time.Now().Add(5 * time.Minute).Truncate(time.Second)
	actual, err := Parsetime("now+5minute")
	after := time.Now().Add(5 * time.Minute).Truncate(time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if before.After(actual) || after.Before(actual) {
		t.Errorf("Actual: %#v, expected between %#v and %#v", actual, before, after)
	}
}
