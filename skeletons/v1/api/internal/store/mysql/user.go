package mysql

import (
	"context"

	"github.com/shipengqi/errors"
	"gorm.io/gorm"

	v1 "github.com/jaguar/apiskeleton/pkg/api/apiserver/v1"
	metav1 "github.com/jaguar/apiskeleton/pkg/api/meta/v1"
	"github.com/jaguar/apiskeleton/pkg/code"
	"github.com/jaguar/apiskeleton/pkg/util/gormutil"
)

type users struct {
	db *gorm.DB
}

func newUsers(ds *datastore) *users {
	return &users{ds.db}
}

// Create creates a new user account.
func (u *users) Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error {
	// return u.db.Transaction(func(tx *gorm.DB) error {
	// 	err := tx.Create(user).Error
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// do more operations in transaction
	// 	return nil
	// })

	return u.db.Create(user).Error
}

// Update updates a user account information.
func (u *users) Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) (err error) {
	return u.db.Session(&gorm.Session{NewDB: true}).Transaction(func(tx *gorm.DB) error {
		return tx.Save(user).Error
	})
}

// Delete deletes the user by the user identifier.
func (u *users) Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		u.db = u.db.Unscoped()
	}

	user := &v1.User{}
	err := u.db.Where("name = ?", username).First(&user).Error
	if err != nil {
		return errors.WithCode(err, code.ErrDatabase)
	}

	return u.db.Session(&gorm.Session{NewDB: true}).Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", user.ID).Delete(&v1.User{}).Error
	})
}

// DeleteList batch deletes the users.
func (u *users) DeleteList(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error {
	return nil
}

// Get returns a user by the user identifier.
func (u *users) Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error) {
	user := &v1.User{}
	err := u.db.Where("name = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// List return all users.
func (u *users) List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error) {
	ret := &v1.UserList{}

	// Todo order, selector, add status option
	ol := gormutil.DePointer(opts.Offset, opts.Limit)
	d := u.db.Where("org_id = ? and status = 1", 1).
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.Total)

	return ret, d.Error
}
