package domain

import (
	"errors"
	"log"
	"time"

	"database/sql/driver"
)

// Time is custom time, so that I could add scanner on it
type Time time.Time

// Scan is to conform to StructScan
func (t *Time) Scan(v interface{}) error {
	stringTime, ok := v.(string)
	if ok {
		vt, err := time.Parse(DateFormat, stringTime)

		if err != nil {
			return err
		}

		*t = Time(vt)
		return nil
	}
	log.Printf("Error parsing time %v", v)
	return errors.New("Could not parse")
}

func (t Time) String() string {
	return time.Time(t).Format(DateFormat)
}

// MarshalJSON is json encoding for Time
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte("\"" + t.String() + "\""), nil

}

// Value driver
func (t Time) Value() (driver.Value, error) {
	return time.Time(t), nil
}

// Duration is just an alias for time.Duration
type Duration time.Duration

// Scan is for use in StructScan in repository layers.
func (d *Duration) Scan(v interface{}) error {
	stringDuration, ok := v.(string)
	if ok {
		vd, err := time.ParseDuration(stringDuration)
		if err != nil {
			log.Printf("Error parsing duration %v", v)
			return err
		}

		*d = Duration(vd)
		return nil
	}
	log.Printf("Error parsing duration %v", v)
	return errors.New("Could not parse")
}

func (d Duration) String() string {
	return time.Duration(d).String()
}

// MarshalJSON is json encoding for Duration.
// Example: A duration of "1h10m" is returned as it is, i.e., "1h10m"
func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.String() + "\""), nil

}

// Value driver
func (d Duration) Value() (driver.Value, error) {
	return d.String(), nil
}

