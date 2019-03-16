// 
// 
// 

package user

import (
	"time"

	"github.com/aamsur/playground/simple-auth/datastore/model"

	"regexp"

	"git.qasico.com/cuxs/common"
	"git.qasico.com/cuxs/validation"
)

// createRequest data struct that stored request data when requesting an create user process.
// All data must be provided and must be match with specification validation below.
// handler function should be bind this with context to matches incoming request
// data keys to the defined json tag.
type createRequest struct {
	Username        string `json:"username" valid:"required"`
	FullName        string `json:"full_name" valid:"required"`
	Email           string `json:"email" valid:"required"`
	Address         string `json:"address" valid:"required"`
	Password        string `json:"password" valid:"required|gte:5"`
	ConfirmPassword string `json:"confirm_password" valid:"required"`
}

// Validate implement validation.Requests interfaces.
func (r *createRequest) Validate() *validation.Output {
	o := &validation.Output{Valid: true}

	// validasi username
	user := &model.User{Username: r.Username}
	if e := user.Read("Username"); e == nil {
		o.Failure("username", "username has been registered")
	}

	res, err := regexp.MatchString("^[a-zA-Z0-9_]*$", r.Username)
	if res == false && err == nil {
		o.Failure("username", "Can't use space or special character in this input field")
	}

	// validasi FullName
	res, err = regexp.MatchString("^[a-zA-Z0-9 ]*$", r.FullName)
	if res == false && err == nil {
		o.Failure("fullname", "Can't use space or special character in this input field")
	}

	// validasi confirm password
	if r.ConfirmPassword != r.Password {
		o.Failure("confirm_password", "confirm_password not same with password")
	}
	return o
}

// Messages implement validation.Requests interfaces
// return custom messages when validation fails.
func (r *createRequest) Messages() map[string]string {
	return map[string]string{}
}

// Transform transforming request into model.
func (r *createRequest) Transform() *model.User {
	pwd, _ := common.PasswordHasher(r.Password)
	user := &model.User{
		Username:  r.Username,
		FullName:  r.FullName,
		Email:     r.Email,
		Address:   r.Address,
		Password:  pwd,
		CreatedAt: time.Now(),
	}
	return user
}

// updateRequest data struct that stored request data when requesting an update partnership process.
// All data must be provided and must be match with specification validation below.
// handler function should be bind this with context to matches incoming request
// data keys to the defined json tag.
type updateRequest struct {
	Username string `json:"username" valid:"required"`
	FullName string `json:"full_name" valid:"required"`
	Email    string `json:"email" valid:"required"`
	Address  string `json:"address" valid:"required"`

	OldUser *model.User `json:"-"`
}

// Validate implement validation.Requests interfaces.
func (r *updateRequest) Validate() *validation.Output {
	o := &validation.Output{Valid: true}

	return o
}

// Messages implement validation.Requests interfaces
// return custom messages when validation fails.
func (r *updateRequest) Messages() map[string]string {
	return map[string]string{}
}

// Transform untuk mengubah isi model.
func (r *updateRequest) Transform() *model.User {
	user := &model.User{
		ID:        r.OldUser.ID,
		Username:  r.Username,
		FullName:  r.FullName,
		Email:     r.Email,
		Address:   r.Address,
		UpdatedAt: time.Now(),
	}

	return user
}
