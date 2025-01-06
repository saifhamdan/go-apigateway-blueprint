// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package v1

import "time"

type CommonModel struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	CreatedBy int32     `gorm:"column:created_by" json:"created_by,omitempty"`
	UpdatedBy int32     `gorm:"column:updated_by" json:"updated_by,omitempty"`
}
