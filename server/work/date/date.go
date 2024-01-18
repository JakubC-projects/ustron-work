package date

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"cloud.google.com/go/civil"
)

type Date struct {
	civil.Date
}

var _ sql.Scanner = (*Date)(nil)
var _ driver.Valuer = (*Date)(nil)

func (d *Date) Scan(src any) error {
	v, ok := src.(time.Time)
	if !ok {
		return fmt.Errorf("invalid data type: %T", src)
	}
	d.Date = civil.DateOf(v)

	return nil
}

func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}

func DateOf(t time.Time) Date {
	return Date{civil.DateOf(t)}
}
