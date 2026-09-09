package sqlite

import (
	"context"

	"gorm.io/gorm"

	v1 "{{ .App.ModuleName }}/pkg/api/apiserver/v1"
	metav1 "{{ .App.ModuleName }}/pkg/api/meta/v1"
	"{{ .App.ModuleName }}/pkg/code"
	"{{ .App.ModuleName }}/pkg/util/gormutil"
	"{{ .App.ModuleName }}/pkg/xerr"
)

type users struct {
	db *gorm.DB
}

func newUsers(ds *datastore) *users {
	return &users{ds.db}
}

func (u *users) Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error {
	return u.db.Create(user).Error
}

func (u *users) Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error {
	return u.db.Session(&gorm.Session{NewDB: true}).Transaction(func(tx *gorm.DB) error {
		return tx.Save(user).Error
	})
}

func (u *users) Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		u.db = u.db.Unscoped()
	}

	user := &v1.User{}
	if err := u.db.Where("name = ?", username).First(user).Error; err != nil {
		return xerr.WithCode(err, code.ErrDatabase)
	}

	return u.db.Session(&gorm.Session{NewDB: true}).Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", user.ID).Delete(&v1.User{}).Error
	})
}

func (u *users) DeleteList(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error {
	return nil
}

func (u *users) Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error) {
	user := &v1.User{}
	if err := u.db.Where("name = ?", username).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *users) List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error) {
	ret := &v1.UserList{}
	ol := gormutil.DePointer(opts.Offset, opts.Limit)
	d := u.db.Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.Total)
	return ret, d.Error
}
