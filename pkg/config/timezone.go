package config

import "time"

func SetTimeZoneUTC() {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		panic("Failed to set time zone")
	}
	time.Local = loc // -> this is setting the global timezone
}
