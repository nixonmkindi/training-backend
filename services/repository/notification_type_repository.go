package repository

import (
	"context"
	"fmt"
	"os"
	"time"
	"training-backend/services/database"
	"training-backend/services/entity"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

// NotificationTypeConn Initializes connection to DB
type NotificationTypeConn struct {
	conn *pgxpool.Pool
}

// NewNotificationType Connects to DB
func NewNotificationType() *NotificationTypeConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &NotificationTypeConn{
		conn: conn,
	}
}

// Create Inserts new record to DB
func (con *NotificationTypeConn) Create(e *entity.NotificationType) (int32, error) {
	var notificationTypeID int32
	query := "INSERT INTO notification_type (name, description, created_by,created_at) VALUES($1, $2, $3, $4) RETURNING id"
	err := con.conn.QueryRow(context.Background(), query, e.Name, e.Description, e.CreatedBy, time.Now().Local()).Scan(&notificationTypeID)
	return notificationTypeID, err
}

// CheckNotificationType Checks if record exists in DB
func (con *NotificationTypeConn) CheckNotificationType(name string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM notification_type WHERE name  = $1)"
	err := con.conn.QueryRow(context.Background(), query, name).Scan(&exists)
	return exists, err
}

// List Lists all records
func (con *NotificationTypeConn) List() ([]*entity.NotificationType, error) {
	var id int32
	var name, description pgtype.Text
	var createdBy, updatedBy, deletedBy pgtype.Int4
	var createdAt, updatedAt, deletedAt pgtype.Timestamp
	var notificationType []*entity.NotificationType
	query := `SELECT id, name, description, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at FROM notification_type WHERE deleted_at IS NULL`
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.NotificationType{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&id, &name, &description, &createdBy, &createdAt, &updatedBy, &updatedAt, &deletedBy, &deletedAt); err != nil {
			return []*entity.NotificationType{}, err
		}
		fileType := &entity.NotificationType{
			ID:          id,
			Name:        name.String,
			Description: description.String,
			CreatedBy:   createdBy.Int,
			CreatedAt:   createdAt.Time.Local(),
		}
		notificationType = append(notificationType, fileType)
	}
	return notificationType, err
}

// Get Gets single record by ID field
func (con *NotificationTypeConn) Get(id int32) (*entity.NotificationType, error) {
	var name, description pgtype.Text
	var createdBy, updatedBy, deletedBy pgtype.Int4
	var createdAt, updatedAt, deletedAt pgtype.Timestamp
	var notificationType *entity.NotificationType

	query := `SELECT id, name, description, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at FROM notification_type WHERE id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(&id, &name, &description, &createdBy, &createdAt,
		&updatedBy, &updatedAt, &deletedBy, &deletedAt)

	if err != nil {
		return notificationType, err
	}

	fileType := &entity.NotificationType{
		ID:          id,
		Name:        name.String,
		Description: description.String,
		CreatedBy:   createdBy.Int,
		CreatedAt:   createdAt.Time.Local(),
		UpdatedBy:   updatedBy.Int,
		UpdatedAt:   updatedAt.Time.Local(),
		DeletedBy:   deletedBy.Int,
		DeletedAt:   deletedAt.Time.Local(),
	}
	return fileType, err

}

// Update Updates single record by ID field
func (con *NotificationTypeConn) Update(e *entity.NotificationType) (int32, error) {
	query := "UPDATE notification_type SET name=$1, description=$2, updated_by = $3, updated_at = $4 WHERE id = $5"
	_, err := con.conn.Exec(context.Background(), query, e.Name, e.Description, e.UpdatedBy, time.Now().Local(), e.ID)
	if err != nil {
		return e.ID, err
	}
	return e.ID, err
}

// SoftDelete Softly delete single record by ID
func (con *NotificationTypeConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE notification_type SET deleted_by = $1, deleted_at = $2 WHERE id = $3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	if err != nil {
		return err
	}
	return err
}

// HardDelete Permanently delete single record by ID
func (con *NotificationTypeConn) HardDelete(id int32) error {
	query := "DELETE FROM notification_type WHERE id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return err
}
