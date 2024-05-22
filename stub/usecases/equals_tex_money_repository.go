package usecases

import "stub/domain"

type EqualsTexMoney interface {
	IsEqualStatusCash(a, b domain.StatusCash) bool
}
