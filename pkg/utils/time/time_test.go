package timeutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	// change implementation of clock in the beginning of the test
	currentDate := time.Date(2000, 12, 15, 17, 8, 00, 0, time.UTC)
	nowFunc = func() time.Time {
		return currentDate
	}

	// after finish with the test, reset the time implementation
	defer resetClockImplementation()

	assert.Equal(t, currentDate, Now(), "Now: should be equals")
}

func TestYesterday(t *testing.T) {
	// change implementation of clock in the beginning of the test
	currentDate := time.Date(2000, 12, 15, 17, 8, 00, 0, time.UTC)
	nowFunc = func() time.Time {
		return currentDate
	}
	expectedDate := time.Date(2000, 12, 14, 17, 8, 00, 0, time.UTC)

	// after finish with the test, reset the time implementation
	defer resetClockImplementation()

	assert.Equal(t, expectedDate, Yesterday(), "Yesterday: should be equals")
}

func TestTomorrow(t *testing.T) {
	// change implementation of clock in the beginning of the test
	currentDate := time.Date(2000, 12, 15, 17, 8, 00, 0, time.UTC)
	nowFunc = func() time.Time {
		return currentDate
	}
	expectedDate := time.Date(2000, 12, 16, 17, 8, 00, 0, time.UTC)

	// after finish with the test, reset the time implementation
	defer resetClockImplementation()

	assert.Equal(t, expectedDate, Tomorrow(), "Tomorrow: should be equals")
}

func TestDateAdd(t *testing.T) {
	currentDate := time.Date(2000, 12, 15, 17, 8, 00, 0, time.UTC)
	expectedDate := time.Date(2000, 12, 18, 17, 8, 00, 0, time.UTC)

	assert.Equal(t, expectedDate, DateAdd(currentDate, 3), "DateAdd: should be equals")
}

func TestHoursAdd(t *testing.T) {
	currentDate := time.Date(2000, 12, 15, 17, 8, 00, 0, time.UTC)
	expectedDate := time.Date(2000, 12, 15, 20, 8, 00, 0, time.UTC)

	assert.Equal(t, expectedDate, HoursAdd(currentDate, 3), "HoursAdd: should be equals")
}

func TestMinutesAdd(t *testing.T) {
	currentDate := time.Date(2000, 12, 15, 17, 8, 00, 0, time.UTC)
	expectedDate := time.Date(2000, 12, 15, 17, 11, 00, 0, time.UTC)

	assert.Equal(t, expectedDate, MinutesAdd(currentDate, 3), "MinutesAdd: should be equals")
}

func TestNowStr(t *testing.T) {
	// change implementation of clock in the beginning of the test
	currentDate := time.Date(2000, 12, 15, 17, 8, 00, 0, time.UTC)
	nowFunc = func() time.Time {
		return currentDate
	}

	// after finish with the test, reset the time implementation
	defer resetClockImplementation()

	assert.Equal(t, "2000-12-15 17:08:00", NowStr(ISO8601TimeWithoutZone), "NowStr: should be equals")
	assert.Equal(t, "2000-12-15", NowStr(ISO8601TimeDate), "NowStr: should be equals")
	assert.Equal(t, "2000-12-15T17:08:00+07:00", NowStr(ISO8601TimeDateWithTimeZone), "NowStr: should be equals")
}

func TestStrFormat(t *testing.T) {
	currentDate := time.Date(2000, 12, 15, 17, 8, 00, 0, time.UTC)

	assert.Equal(t, "2000-12-15 17:08:00", StrFormat(currentDate, ISO8601TimeWithoutZone), "StrFormat: should be equals")
	assert.Equal(t, "2000-12-15", StrFormat(currentDate, ISO8601TimeDate), "StrFormat: should be equals")
	assert.Equal(t, "2000-12-15T17:08:00+07:00", StrFormat(currentDate, ISO8601TimeDateWithTimeZone), "StrFormat: should be equals")
}

func TestDateDifferenceCounter(t *testing.T) {
	date1 := time.Date(2000, 12, 15, 17, 8, 00, 0, time.UTC)
	date2 := time.Date(2000, 12, 13, 17, 8, 00, 0, time.UTC)

	assert.Equal(t, 2, DateDifferenceCounter(date1, date2), "DateDifferenceCounter: should be equals")
}

func TestMaturityDate(t *testing.T) {
	currentDate := time.Date(2000, 12, 15, 17, 8, 00, 0, time.UTC)
	expectedDate := time.Date(2000, 12, 16, 17, 8, 00, 0, time.UTC)

	assert.Equal(t, expectedDate, MaturityDate(currentDate, 2), "MaturityDate: should be equals")
	assert.Equal(t, 16, MaturityDate(currentDate, 2).Day(), "MaturityDate: should be equals")
}

func TestParse(t *testing.T) {
	//change implementation of clock in the beginning of the test
	currentDate := time.Date(2000, 12, 15, 17, 8, 00, 0, time.UTC)
	nowFunc = func() time.Time {
		return currentDate
	}

	// after finish with the test, reset the time implementation
	defer resetClockImplementation()

	parseResult, err := time.Parse(ISO8601TimeDate, now().Format(ISO8601TimeDate))
	assert.NoError(t, err)
	assert.Equal(t, "2000-12-15 00:00:00 +0000 UTC", fmt.Sprint(parseResult), "Parse: should be equals")
}
