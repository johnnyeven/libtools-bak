package sqlx

import (
	"errors"
)

var (
	ErrSqlInvalid        = errors.New("sql invalid")
	ErrInvalidScanTarget = errors.New("can not scan to a none pointer value")
	ErrNotFound          = errors.New("record is not found")
	ErrSelectShouldOne   = errors.New("more than one records found, but only one")
	ErrConflict          = errors.New("record conflict")
)

var DuplicateEntryErrNumber uint16 = 1062

func DBErr(err error) *dbErr {
	return &dbErr{
		err: err,
	}
}

type dbErr struct {
	err error

	errDefault  error
	errNotFound error
	errConflict error
}

func (r dbErr) WithNotFound(err error) *dbErr {
	r.errNotFound = err
	return &r
}

func (r dbErr) WithDefault(err error) *dbErr {
	r.errDefault = err
	return &r
}

func (r dbErr) WithConflict(err error) *dbErr {
	r.errConflict = err
	return &r
}

func (r *dbErr) IsNotFound() bool {
	return r.err == ErrNotFound
}

func (r *dbErr) IsConflict() bool {
	return r.err == ErrConflict
}

func (r *dbErr) Err() error {
	if r.err == nil {
		return nil
	}
	switch r.err {
	case ErrNotFound:
		if r.errNotFound != nil {
			return r.errNotFound
		}
	case ErrConflict:
		if r.errConflict != nil {
			return r.errConflict
		}
	}
	if r.errDefault != nil {
		return r.errDefault
	}
	return r.err
}
