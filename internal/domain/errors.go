package domain

import "errors"

var (
	ErrWrongID     = errors.New("некорректный ID задачи")
	ErrWrongStatus = errors.New("некорректный статус задачи")
	ErrNotExistKey = errors.New("несуществующий индекс")
)
