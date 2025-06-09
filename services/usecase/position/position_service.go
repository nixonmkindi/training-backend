package position

import (
	"training-backend/package/log"
	"training-backend/services/entity"
	"training-backend/services/error_message"
	"training-backend/services/repository"
)

// Service Initialize repository
type Service struct {
	repo Repository
}

// NewService Instantiate new service
func NewService() UseCase {
	repo := repository.NewPosition()
	return &Service{
		repo: repo,
	}
}

// CreatePosition Calls create new record repository
func (s *Service) CreatePosition(name string, description string, createdBy int32) (int32, error) {
	var id int32
	position, err := entity.NewPosition(name, description, createdBy)
	if err != nil {
		log.Error(err)
		return id, err
	}

	exists, _ := s.repo.CheckPosition(position.Name)
	if !exists {
		id, err = s.repo.Create(position)
		if err != nil {
			log.Errorf("error creating position: %v", err)
			return id, error_message.ErrCannotBeCreated
		}
	}
	return id, err
}

// ListPosition Calls list records repository
func (s *Service) ListPosition() ([]*entity.Position, error) {
	positions, err := s.repo.List()
	if err != nil && err.Error() != error_message.ErrNoResultSet.Error() {
		log.Error(err)
		return positions, err
	}
	return positions, err
}

// GetPosition Calls get single record repository
func (s *Service) GetPosition(id int32) (*entity.Position, error) {
	position, err := s.repo.Get(id)
	if err != nil && err.Error() != error_message.ErrNoResultSet.Error() {
		log.Error(err)
		return position, err
	}
	return position, err
}

// UpdatePosition Calls updates single record by ID field repository
func (s *Service) UpdatePosition(e *entity.Position) (int32, error) {
	err := e.ValidateUpdatePosition()
	if err != nil {
		log.Error(err)
		return e.ID, error_message.ErrCannotBeUpdated
	}
	_, err = s.repo.Update(e)
	if err != nil {
		log.Error(err)
		return e.ID, error_message.ErrNotFound
	}
	return e.ID, err
}

// SoftDeletePosition Calls soft delete function for single record by ID repository
func (s *Service) SoftDeletePosition(id, deletedBy int32) error {
	_, err := s.GetPosition(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	errDelete := s.repo.SoftDelete(id, deletedBy)
	if errDelete != nil {
		log.Error(errDelete)
		return error_message.ErrCannotBeDeleted
	}
	return errDelete
}

// HardDeletePosition Calls hard delete function for single record by ID repository
func (s *Service) HardDeletePosition(id int32) error {
	_, err := s.GetPosition(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	errDelete := s.repo.HardDelete(id)
	if errDelete != nil {
		log.Error(errDelete)
		return error_message.ErrCannotBeDeleted
	}
	return errDelete
}
