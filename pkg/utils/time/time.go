package timeutil

import (
	"go-skeleton/pkg/utils/errors"
	"go-skeleton/pkg/utils/null"
	"time"
)

type nowFuncT func() time.Time

var nowFunc nowFuncT

func init() {
	resetClockImplementation()
}

func resetClockImplementation() {
	nowFunc = func() time.Time {
		return time.Now()
	}
}

func now() time.Time {
	return nowFunc()
}

func Now() time.Time {
	return now()
}

func Yesterday() time.Time {
	return Now().AddDate(0, 0, -1)
}

func Tomorrow() time.Time {
	return Now().AddDate(0, 0, 1)
}

func DateAdd(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

func HoursAdd(t time.Time, hours int) time.Time {
	return t.Add(time.Hour * time.Duration(hours))
}

func MinutesAdd(t time.Time, minutes int) time.Time {
	return t.Add(time.Minute * time.Duration(minutes))
}

func NowStr(formats ...interface{}) string {
	return StrFormat(Now(), formats...)
}

func StrFormat(t time.Time, formats ...interface{}) string {
	var format string = ISO8601TimeWithoutZone
	if len(formats) > 0 {
		// First parameter is the format.
		var first = formats[0]
		if first != nil {
			var ok bool
			format, ok = formats[0].(string)
			if !ok {
				format = ISO8601TimeWithoutZone
			}
		}
	}

	return t.Format(format)
}

func DateDifferenceCounter(date1 time.Time, date2 time.Time) int {
	return int(date1.Sub(date2).Hours() / 24)
}

func Parse(t string, layout string) (time.Time, error) {
	return time.Parse(layout, t)
}

func ReformatToYYYY_MM_DD(t *string) (*string, error) {
	if null.IsNil(t) {
		return nil, errors.NewGenericError(errors.INTERNAL_SERVER_ERROR)
	}

	var parseError error

	// step 1, Try to parse with layout YYYY-MM-DD
	if _, err := time.Parse(ISO8601TimeDate, *t); err == nil {
		// step 1 success, return without reformatting
		return t, nil
	}

	// Failed.. so it's from DB
	// step 2, try to parse with 24 hour layout
	if dobTime, err := time.Parse(ISO8601TimeDateWithTimeZone, *t); err == nil {
		// Success, reformat to YYYY-MM-DD
		fixedDob := StrFormat(dobTime, ISO8601TimeDate)
		return &fixedDob, nil
	}

	// Failed.. maybe it's in 12 hour format
	// Step 3, try to parse with 12 hour layout
	if dobTime12HourFormat, err := time.Parse(ISO8601TimeDateWithTimeZone12HourFormat, *t); err == nil {
		// Success, reformat to YYYY-MM-DD
		fixedDob := StrFormat(dobTime12HourFormat, ISO8601TimeDate)
		return &fixedDob, nil
	}

	// Hmm.. maybe need to use another layout
	// Step 4, try to parse with layout without zone
	if dobTimeWithoutZone, err := time.Parse(ISO8601TimeWithoutZone, *t); err == nil {
		// Success, reformat to YYYY-MM-DD
		fixedDob := StrFormat(dobTimeWithoutZone, ISO8601TimeDate)
		return &fixedDob, nil
	}

	// Hmm.. weird.. lets try utc layout
	// Step 4, try to parse with utc layout
	if dobTimeUTC, err := time.Parse(ISO8601TimeDateFromUTC, *t); err == nil {
		// Success, reformat to YYYY-MM-DD
		fixedDob := StrFormat(dobTimeUTC, ISO8601TimeDate)
		return &fixedDob, nil
	} else {
		parseError = err
	}

	// Oh my goodness..still error parsed..

	return nil, parseError
}

// today is also counted as loan date - hence the -1
func MaturityDate(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days-1)
}
