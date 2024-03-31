package v1

import (
	"time"

	metav1 "github.com/jaguar/apiskeleton/pkg/api/meta/v1"
)

// User represents a user restful resource. It is also used as gorm model.
type User struct {
	// Standard object's metadata.
	metav1.ObjectMeta

	Status int `json:"status" gorm:"column:status;type:int(1);DEFAULT:1;" validate:"omitempty"`

	// Required: true
	Sex int `json:"sex" gorm:"column:sex;type:int(1);" validate:"required"`

	// Required: true
	OrgId uint64 `json:"org_id" gorm:"column:org_id;type:bigint(20);not null;"`

	// Required: true
	Nickname string `json:"nickname" gorm:"column:nickname;type:varchar(64);" validate:"required,min=1,max=30"`

	// Required: true
	Password string `json:"password,omitempty" gorm:"column:password;type:varchar(128)" validate:"required"`

	// Required: true
	Email string `json:"email" gorm:"column:email;type:varchar(128)" validate:"required,email,min=1,max=100"`

	// Required: true
	Phone string `json:"phone" gorm:"column:phone;type:varchar(11);" validate:"required"`

	LoginAt time.Time `json:"login_at,omitempty" gorm:"column:login_at"`
}

// UserList is the whole list of all users which have been stored in storage.
type UserList struct {
	// Standard list metadata.
	// +optional
	metav1.ListMeta `json:",inline"`

	Items []*User `json:"items"`
}

// TableName overwrite table name `users` to `{{example}}_user`.
func (u *User) TableName() string {
	return "{{example}}_user"
}
