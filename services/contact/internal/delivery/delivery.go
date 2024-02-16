package delivery

import (
	"architecture_go/services/contact/internal/usecase"
)

type ContactDeliveryImpl struct {
	usecase usecase.ContactUseCase
}

func NewContactDelivery(usecase usecase.ContactUseCase) ContactDelivery {
	return &ContactDeliveryImpl{
		usecase: usecase,
	}
}

type ContactDelivery interface{}
