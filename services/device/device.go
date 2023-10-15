package device

import (
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxentities"
)

type Service interface {
	Insert(entities.Device) error
	FindAllConnected() ([]entities.Device, error)
	FindByMacAddress(address string) (*entities.Device, error)
}
