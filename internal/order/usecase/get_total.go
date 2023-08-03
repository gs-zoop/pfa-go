package usecase

import (
	"fmt"

	"github.com/jgsouzadev/pfa-go/internal/order/entity"
)

type GetTotalOutputDto struct {
	Total int
}

type GetTotalUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetTotalUseCase(or entity.OrderRepositoryInterface) *GetTotalUseCase {
	return &GetTotalUseCase{OrderRepository: or}
}

func (c *GetTotalUseCase) Execute() (*GetTotalOutputDto, error) {
	total, err := c.OrderRepository.GetTotal()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &GetTotalOutputDto{Total: total}, nil
}
