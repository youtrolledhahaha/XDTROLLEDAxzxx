package device

import (
	"github.com/youtrolledhahaha/youdmmmbaa/entities"
	"time"
)

type Repository interface {
	Insert(device entities.Device) error
	Update(device entities.Device) error
	FindByMacAddress(address string) (*entities.Device, error)
	FindAll(fetchedAt time.Time) ([]entities.Device, error)
}
