package repository

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrNotFound   = errors.New("data not found")
	ErrDuplicate  = errors.New("duplicate data")
	ErrInternal   = errors.New("internal error")
	ErrInvalidArg = errors.New("invalid argument")
)

// Deteksi duplicate key (Postgres contoh)
func isDuplicateKey(err error) bool {
	return errors.Is(err, gorm.ErrDuplicatedKey) // gorm v2 support ini
}

func isRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

// func isInsufficient(err error) bool{
// 	return errors.Is(err, gorm.Err)
// }
