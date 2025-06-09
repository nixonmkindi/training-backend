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

// PositionConn Initializes connection to DB
type PositionConn struct {
	conn *pgxpool.Pool
}

// NewPosition Connects to DB
func NewPosition() *PositionConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &PositionConn{
		conn: conn,
	}
}

// Create Inserts new record to DB
func (con *PositionConn) Create(e *entity.Position) (int32, error) {
	var id int32
	query := `INSERT INTO position (name, description, created_by, created_at) VALUES($1, $2, $3, $4) RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.Name,
		e.Description,
		e.CreatedBy,
		time.Now().Local()).Scan(&id)
	return id, err
}

// CheckPosition Checks if record exists in DB
func (con *PositionConn) CheckPosition(name string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM position WHERE name = $1)"
	err := con.conn.QueryRow(context.Background(), query, name).Scan(&exists)
	return exists, err
}

func positionSelectQuery() string {
	return `SELECT
				id,
				name,
				description,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM position WHERE deleted_at IS NULL`
}

// List Lists all records
func (con *PositionConn) List() ([]*entity.Position, error) {
	var id pgtype.Int4
	var name pgtype.Text
	var description pgtype.Text
	var createdBy pgtype.Int4
	var updatedBy pgtype.Int4
	var deletedBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var items []*entity.Position

	query := positionSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.Position{}, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&name,
			&description,
			&createdBy,
			&updatedBy,
			&deletedBy,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return []*entity.Position{}, err
		}
		item := &entity.Position{
			ID:          id.Int,
			Name:        name.String,
			Description: description.String,
			CreatedBy:   createdBy.Int,
			UpdatedBy:   updatedBy.Int,
			DeletedBy:   deletedBy.Int,
			CreatedAt:   createdAt.Time.Local(),
			UpdatedAt:   updatedAt.Time.Local(),
			DeletedAt:   deletedAt.Time.Local(),
		}
		items = append(items, item)
	}

	return items, err
}

// Get Gets single record by ID field
func (con *PositionConn) Get(id int32) (*entity.Position, error) {
	var name pgtype.Text
	var description pgtype.Text
	var createdBy pgtype.Int4
	var updatedBy pgtype.Int4
	var deletedBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var item *entity.Position

	query := positionSelectQuery() + ` AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id,
		&name,
		&description,
		&createdBy,
		&updatedBy,
		&deletedBy,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)

	if err != nil {
		return item, err
	}

	item = &entity.Position{
		ID:          id,
		Name:        name.String,
		Description: description.String,
		CreatedBy:   createdBy.Int,
		UpdatedBy:   updatedBy.Int,
		DeletedBy:   deletedBy.Int,
		CreatedAt:   createdAt.Time.Local(),
		UpdatedAt:   updatedAt.Time.Local(),
		DeletedAt:   deletedAt.Time.Local(),
	}
	return item, err
}

// Update Updates single record by ID field
func (con *PositionConn) Update(e *entity.Position) (int32, error) {
	query := `UPDATE position SET 
				name = $1, 
				description = $2, 
				updated_by = $3, 
				updated_at = $4 WHERE id = $5`
	_, err := con.conn.Exec(context.Background(), query,
		e.Name,
		e.Description,
		e.UpdatedBy,
		time.Now().Local(),
		e.ID)
	if err != nil {
		return e.ID, err
	}
	return e.ID, err
}

// SoftDelete Softly delete single record by ID
func (con *PositionConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE position SET deleted_by = $1, deleted_at = $2 WHERE id = $3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	if err != nil {
		return err
	}
	return err
}

// HardDelete Permanently delete single record by ID
func (con *PositionConn) HardDelete(id int32) error {
	query := "DELETE FROM position WHERE id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return err
}
