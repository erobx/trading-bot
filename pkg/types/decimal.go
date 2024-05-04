package types

import (
	"database/sql/driver"
	"errors"

	"github.com/shopspring/decimal"
)

type DbDecimal decimal.Decimal

func (v *DbDecimal) Scan(value interface{}) error {
	if value == nil {
		*v = DbDecimal(decimal.Zero)
		return nil
	}
	if sv, err := driver.String.ConvertValue(value); err == nil {
		if vv, ok := sv.(string); ok {
			if vvv, err := decimal.NewFromString(vv); err == nil {
				*v = DbDecimal(vvv)
				return nil
			}
		}
	}
	return errors.New("cannot convert to decimal")
}

func (v DbDecimal) Value() (driver.Value, error) {
	dec := decimal.Decimal(v)
	price, _ := dec.Float64()
	return price, nil
}

func (v DbDecimal) String() string {
	return decimal.Decimal(v).String()
}
