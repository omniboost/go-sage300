package sage300

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

type Time struct {
	time.Time
}

func (t Time) IsEmpty() bool {
	return t.Time.IsZero()
}

type Date struct {
	time.Time
}

func (d Date) MarshalSchema() string {
	return d.Time.Format("2006-01-02")
}

func (d Date) IsEmpty() bool {
	return d.Time.IsZero()
}

type DateTime struct {
	time.Time
}

func (d *Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.Time.Format("2006-01-02T15:04:05"))
}

func (d *Date) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02", value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("02/01/2006", value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02T15:04:05", value)
	return err
}

func (d DateTime) MarshalSchema() string {
	return d.Time.Format(time.RFC3339)
}

func (dt *DateTime) MarshalJSON() ([]byte, error) {
	if dt.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(dt.Time.Format("2006-01-02T15:04:05"))
}

func (dt *DateTime) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	dt.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	dt.Time, err = time.Parse("2006-01-02T15:04:05", value)
	return err
}

type Bool bool

func (b *Bool) UnmarshalJSON(text []byte) (err error) {
	var bl bool
	err = json.Unmarshal(text, &bl)
	if err == nil {
		*b = Bool(bl)
		return nil
	}

	var str string
	err = json.Unmarshal(text, &str)
	if err == nil {
		return nil
	}

	if str == "" {
		return nil
	}

	if str == "F" {
		*b = false
	}

	return errors.New("FML")
}
