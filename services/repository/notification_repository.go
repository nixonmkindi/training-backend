/**
 * @author Grace Remce
 * @email graceremce99@gmail.com
 * @create date 2024-03-24 08:43:04
 * @modify date 2024-03-24 08:43:04
 * @desc [description]
 */

package repository

import (
	"context"
	"fmt"
	"training-backend/package/log"
	"training-backend/services/database"
	"training-backend/services/entity"

	"os"
	"time"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

type NotificationConn struct {
	conn *pgxpool.Pool
}

func NewNotification() *NotificationConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &NotificationConn{
		conn: conn,
	}
}

func getNotificationQuery() string {
	return `SELECT 	notification.id, 
					 notification.notification_type_id, 
					 notification.subsystem_id, 
					 notification.from_id, 
					 notification.to_id,
					 notification.content,
					 notification.is_opened,
					 notification_type.name,
					 notification.created_by, 
					 notification.created_at, 
					 notification.updated_by, 
					 notification.updated_at, 
					 notification.deleted_by, 
					 notification.deleted_at 
				 FROM notification
				 INNER JOIN notification_type ON notification.notification_type_id = notification_type.id
				 WHERE notification.deleted_at IS NULL`
}

func (con *NotificationConn) Create(e *entity.Notification) (int32, error) {
	var notificationID int32
	query := `INSERT INTO notification (
						 notification_type_id, 
						 subsystem_id, 
						 from_id, 
						 to_id,
						 content,
						 is_opened,
						 created_by,
						 created_at) 
						 VALUES($1,$2,$3,$4,$5,$6,$7,$8) 
						 RETURNING id`
	err := con.conn.QueryRow(context.Background(), query,
		e.NotificationTypeID,
		e.SubsystemID,
		e.FromID,
		e.ToID,
		e.Content,
		false,
		e.CreatedBy,
		time.Now()).Scan(&notificationID)
	if err != nil {
		log.Errorf("error creating notification: %v", err)
	}
	return notificationID, err
}

func (con *NotificationConn) ListByFromID(fromID int32) ([]*entity.Notification, error) {
	var id int32
	var notificationTypeID, subsystemID, toID pgtype.Int4
	var content, notificationTypeName pgtype.GenericText
	var isOpened pgtype.Bool
	var createdBy, updatedBy, deletedBy pgtype.Int4
	var createdAt, updatedAt, deletedAt pgtype.Timestamp

	var notifications []*entity.Notification

	var query = getNotificationQuery() + ` AND notification.from_id = $1 ORDER BY notification.created_at DESC`
	rows, err := con.conn.Query(context.Background(), query, fromID)
	if err != nil {
		log.Errorf("error querying notification by from id %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(
			&id,
			&notificationTypeID,
			&subsystemID,
			&fromID,
			&toID,
			&content,
			&isOpened,
			&notificationTypeName,
			&createdBy,
			&createdAt,
			&updatedBy,
			&updatedAt,
			&deletedBy,
			&deletedAt); err != nil {
			log.Errorf("error scaning notification %v", err)
			return nil, err
		}

		notificationType := &entity.NotificationType{
			Name: notificationTypeName.String,
		}

		notification := &entity.Notification{
			ID:                 id,
			NotificationTypeID: notificationTypeID.Int,
			SubsystemID:        subsystemID.Int,
			FromID:             fromID,
			ToID:               toID.Int,
			Content:            content.String,
			IsOpened:           isOpened.Bool,
			NotificationType:   notificationType,
			CreatedBy:          createdBy.Int,
			CreatedAt:          createdAt.Time,
			UpdatedBy:          updatedBy.Int,
			UpdatedAt:          updatedAt.Time,
			DeletedBy:          deletedBy.Int,
			DeletedAt:          deletedAt.Time,
		}
		notifications = append(notifications, notification)
	}

	return notifications, err
}

func (con *NotificationConn) ListByToID(toID int32) ([]*entity.Notification, error) {
	var id int32
	var notificationTypeID, subsystemID, fromID pgtype.Int4
	var content, notificationTypeName pgtype.GenericText
	var isOpened pgtype.Bool
	var createdBy, updatedBy, deletedBy pgtype.Int4
	var createdAt, updatedAt, deletedAt pgtype.Timestamp

	var notifications []*entity.Notification

	var query = getNotificationQuery() + ` AND notification.to_id = $1 ORDER BY notification.created_at DESC`
	rows, err := con.conn.Query(context.Background(), query, toID)
	if err != nil {
		log.Errorf("error querying notification by to id %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(
			&id,
			&notificationTypeID,
			&subsystemID,
			&fromID,
			&toID,
			&content,
			&isOpened,
			&notificationTypeName,
			&createdBy,
			&createdAt,
			&updatedBy,
			&updatedAt,
			&deletedBy,
			&deletedAt); err != nil {
			log.Errorf("error scaning notification %v", err)
			return nil, err
		}

		notificationType := &entity.NotificationType{
			Name: notificationTypeName.String,
		}

		notification := &entity.Notification{
			ID:                 id,
			NotificationTypeID: notificationTypeID.Int,
			SubsystemID:        subsystemID.Int,
			FromID:             fromID.Int,
			ToID:               toID,
			Content:            content.String,
			IsOpened:           isOpened.Bool,
			NotificationType:   notificationType,
			CreatedBy:          createdBy.Int,
			CreatedAt:          createdAt.Time,
			UpdatedBy:          updatedBy.Int,
			UpdatedAt:          updatedAt.Time,
			DeletedBy:          deletedBy.Int,
			DeletedAt:          deletedAt.Time,
		}
		notifications = append(notifications, notification)
	}

	return notifications, err
}

func (con *NotificationConn) ListByNotificationType(notificationTypeID int32) ([]*entity.Notification, error) {
	var id int32
	var subsystemID, fromID, toID pgtype.Int4
	var content, notificationTypeName pgtype.GenericText
	var isOpened pgtype.Bool
	var createdBy, updatedBy, deletedBy pgtype.Int4
	var createdAt, updatedAt, deletedAt pgtype.Timestamp

	var notifications []*entity.Notification

	var query = getNotificationQuery() + ` AND notification_type.id = $1 ORDER BY notification.created_at DESC`
	rows, err := con.conn.Query(context.Background(), query, notificationTypeID)
	if err != nil {
		log.Errorf("error querying notification by notification type %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(
			&id,
			&notificationTypeID,
			&subsystemID,
			&fromID,
			&toID,
			&content,
			&isOpened,
			&notificationTypeName,
			&createdBy,
			&createdAt,
			&updatedBy,
			&updatedAt,
			&deletedBy,
			&deletedAt); err != nil {
			log.Errorf("error scaning notification %v", err)
			return nil, err
		}

		notificationType := &entity.NotificationType{
			Name: notificationTypeName.String,
		}

		notification := &entity.Notification{
			ID:                 id,
			NotificationTypeID: notificationTypeID,
			SubsystemID:        subsystemID.Int,
			FromID:             fromID.Int,
			ToID:               toID.Int,
			Content:            content.String,
			NotificationType:   notificationType,
			IsOpened:           isOpened.Bool,
			CreatedBy:          createdBy.Int,
			CreatedAt:          createdAt.Time,
			UpdatedBy:          updatedBy.Int,
			UpdatedAt:          updatedAt.Time,
			DeletedBy:          deletedBy.Int,
			DeletedAt:          deletedAt.Time,
		}
		notifications = append(notifications, notification)
	}

	return notifications, err
}

func (con *NotificationConn) Get(id int32) (*entity.Notification, error) {
	var notificationTypeID, subsystemID, fromID, toID pgtype.Int4
	var content, notificationTypeName pgtype.GenericText
	var isOpened pgtype.Bool
	var createdBy, updatedBy, deletedBy pgtype.Int4
	var createdAt, updatedAt, deletedAt pgtype.Timestamp

	query := getNotificationQuery() + ` AND notification.id = $1 ORDER BY notification.created_at DESC`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id,
		&notificationTypeID,
		&subsystemID,
		&fromID,
		&toID,
		&content,
		&isOpened,
		&notificationTypeName,
		&createdBy,
		&createdAt,
		&updatedBy,
		&updatedAt,
		&deletedBy, &deletedAt)

	if err != nil {
		log.Errorf("error getting notification %v", err)
		return nil, err
	}

	notificationType := &entity.NotificationType{
		Name: notificationTypeName.String,
	}

	notification := &entity.Notification{
		ID:                 id,
		NotificationTypeID: notificationTypeID.Int,
		SubsystemID:        subsystemID.Int,
		FromID:             fromID.Int,
		ToID:               toID.Int,
		Content:            content.String,
		NotificationType:   notificationType,
		CreatedBy:          createdBy.Int,
		CreatedAt:          createdAt.Time,
		UpdatedBy:          updatedBy.Int,
		UpdatedAt:          updatedAt.Time,
		DeletedBy:          deletedBy.Int,
		DeletedAt:          deletedAt.Time,
	}

	return notification, err
}

func (con *NotificationConn) UpdateOpenedStatus(id, userID int32) error {
	query := `UPDATE notification SET 
	                             is_opened = $1,
								 updated_by = $2, 
								 updated_at = $3 
								 WHERE id = $4`
	_, err := con.conn.Exec(context.Background(), query, true, userID, time.Now(), id)
	if err != nil {
		log.Errorf("error updating notification opened status: %v", err)
	}
	return err
}

func (con *NotificationConn) SoftDelete(id, deletedBy int32) error {
	query := `UPDATE notification SET 
	                             deleted_by = $1, 
								 deleted_at = $2
								 WHERE id = $3`
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now(), id)
	if err != nil {
		log.Errorf("error soft delete notification by id: %v", err)
	}
	return err
}

func (con *NotificationConn) HardDelete(id int32) error {
	query := `DELETE FROM notification WHERE id = $1`
	_, err := con.conn.Exec(context.Background(), query, id)
	if err != nil {
		log.Errorf("error permanent delete notification by id: %v", err)
	}
	return err
}

func (con *NotificationConn) GetUnreadUserNotification(toID int32) (int32, error) {
	var total pgtype.Int8

	query := `SELECT COUNT(*) FROM notification WHERE to_id = $1 AND is_opened = false`
	err := con.conn.QueryRow(context.Background(), query, toID).Scan(
		&total)

	if err != nil {
		log.Errorf("error counting unread user notification %v", err)
		return 0, err
	}

	return int32(total.Int), err
}
