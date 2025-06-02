/**
 * @author Yohana Kangwe
 * @email yonakangwe@gmail.com
 * @create date 2024-03-22 05:04:43
 * @modify date 2024-03-22 05:04:43
 * @desc [description]
 */

package models

type ID struct {
	ID int32 `json:"id" form:"id" validate:"required"`
}

type DeleteIDs struct {
	ID        int32 `json:"id" validate:"required"`
	DeletedBy int32 `json:"deleted_by" validate:"required"`
}

type UpdateIDs struct {
	ID        int32 `json:"id,omitempty" form:"id" validate:"numeric,required"`
	UpdatedBy int32 `json:"updated_by,omitempty" form:"updated_by" validate:"numeric,required"`
}

type Name struct {
	Name string `json:"name"`
}

type Exists struct {
	Exists bool `json:"exists"`
}