package utils

import "time"

func NextMonth(t time.Time) *time.Time {
	date := t

	day := t.Day()
	if t.Day() > 28 {
		day = 28
	}

	switch t.Month() {
	case time.January:
		date = time.Date(t.Year(), time.February, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.February:
		date = time.Date(t.Year(), time.March, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.March:
		date = time.Date(t.Year(), time.April, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.April:
		date = time.Date(t.Year(), time.May, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.May:
		date = time.Date(t.Year(), time.June, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.June:
		date = time.Date(t.Year(), time.July, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.July:
		date = time.Date(t.Year(), time.August, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.August:
		date = time.Date(t.Year(), time.September, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.September:
		date = time.Date(t.Year(), time.October, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.October:
		date = time.Date(t.Year(), time.November, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.November:
		date = time.Date(t.Year(), time.December, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.December:
		date = time.Date(t.Year()+1, time.January, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	}

	return &date
}

func LastMonth(t time.Time) *time.Time {
	date := t

	day := t.Day()
	if t.Day() > 28 {
		day = 28
	}

	switch t.Month() {
	case time.January:
		date = time.Date(t.Year()-1, time.December, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.February:
		date = time.Date(t.Year(), time.January, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.March:
		date = time.Date(t.Year(), time.February, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.April:
		date = time.Date(t.Year(), time.March, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.May:
		date = time.Date(t.Year(), time.April, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.June:
		date = time.Date(t.Year(), time.May, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.July:
		date = time.Date(t.Year(), time.June, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.August:
		date = time.Date(t.Year(), time.July, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.September:
		date = time.Date(t.Year(), time.August, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.October:
		date = time.Date(t.Year(), time.September, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.November:
		date = time.Date(t.Year(), time.October, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	case time.December:
		date = time.Date(t.Year(), time.November, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	}

	return &date
}
