// Copyright 2017 PT. Qasico Teknologi Indonesia. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"github.com/aamsur/playground/simple-auth/datastore/model"
	"github.com/aamsur/playground/simple-auth/src/auth"

	"git.qasico.com/cuxs/common"
	"git.qasico.com/cuxs/cuxs"
	"github.com/labstack/echo"
)

// Handler collection handler for user.
type Handler struct{}

// URLMapping declare endpoint with handler function.
func (h *Handler) URLMapping(r *echo.Group) {
	r.POST("", h.create, cuxs.Authorized())
	r.GET("", h.get, cuxs.Authorized())
	r.PUT("/:id", h.update, cuxs.Authorized())
	r.DELETE("/:id", h.delete, cuxs.Authorized())
}

// get endpoint to handle get http method.
func (h *Handler) get(c echo.Context) (e error) {
	ctx := c.(*cuxs.Context)
	// get query string from request
	rq := ctx.RequestQuery()
	var total int64
	var data *[]model.User
	if _, e = auth.UserSession(ctx); e == nil {
		if data, total, e = GetUsers(rq); e == nil {
			ctx.Data(data, total)
		}
	}
	return ctx.Serve(e)
}

// delete to delete user by id.
func (h *Handler) delete(c echo.Context) (e error) {
	ctx := c.(*cuxs.Context)

	var id int64
	var data *model.User
	if _, e = auth.UserSession(ctx); e == nil {
		if id, e = common.Decrypt(ctx.Param("id")); e == nil {
			if data, e = GetUserByID(id); e == nil {
				data.Delete()
			} else {
				e = echo.ErrNotFound
			}
		}
	}

	return ctx.Serve(e)
}

// create endpoint to handle post http method
func (h *Handler) create(c echo.Context) (e error) {
	ctx := c.(*cuxs.Context)
	var r createRequest

	if e = ctx.Bind(&r); e == nil {
		data := r.Transform()
		if e = data.Save(); e == nil {
			ctx.Data(data)
		}
	}

	return ctx.Serve(e)
}

// update endpoint to handle put http method
func (h *Handler) update(c echo.Context) (e error) {
	var id int64
	var r updateRequest
	var u *model.User
	ctx := c.(*cuxs.Context)

	// check session
	if _, e = auth.UserSession(ctx); e == nil {

		// get user id from param
		if id, e = common.Decrypt(ctx.Param("id")); e == nil {

			// get old user data
			if u, e = GetUserByID(id); e == nil {
				r.OldUser = u

				// running the validation
				if e = ctx.Bind(&r); e == nil {

					// transform data
					m := r.Transform()

					// save into database
					if m.Save("Username", "FullName", "Email", "Address", "UpdatedAt"); e == nil {
						m.Read("ID")
						ctx.Data(m)
					}
				}
			} else {
				e = echo.ErrNotFound
			}
		}
	}
	return ctx.Serve(e)
}
