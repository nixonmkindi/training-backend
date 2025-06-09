package position

import "training-backend/services/entity"

// Reader interface
type Reader interface {
	Get(id int32) (*entity.Position, error)
	List() ([]*entity.Position, error)
	CheckPosition(name string) (bool, error)
}

// Writer interface
type Writer interface {
	Create(e *entity.Position) (int32, error)
	Update(e *entity.Position) (int32, error)
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	CreatePosition(name string, description string, createdBy int32) (int32, error)
	ListPosition() ([]*entity.Position, error)
	GetPosition(id int32) (*entity.Position, error)
	UpdatePosition(e *entity.Position) (int32, error)
	SoftDeletePosition(id, deletedBy int32) error
	HardDeletePosition(id int32) error
}
