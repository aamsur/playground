// 
// 
// 

package user

import (
	"github.com/aamsur/playground/simple-auth/datastore/model"

	"git.qasico.com/cuxs/orm"
)

// GetUserByID untuk get data user berdasarkan id user
// return : data user dan error
func GetUserByID(id int64) (m *model.User, err error) {
	mx := new(model.User)
	o := orm.NewOrm().QueryTable(mx)

	if err = o.Filter("id", id).RelatedSel().Limit(1).One(mx); err != nil {
		return nil, err
	}
	return mx, nil
}

// GetUsers get all data user that matched with query request parameters.
// returning slices of users, total data without limit and error.
func GetUsers(rq *orm.RequestQuery) (m *[]model.User, total int64, err error) {
	// make new orm query
	q, _ := rq.Query(new(model.User))

	// get total data
	if total, err = q.Count(); err != nil || total == 0 {
		return nil, total, err
	}

	// get data requested
	var mx []model.User
	if _, err = q.All(&mx, rq.Fields...); err == nil {
		return &mx, total, nil
	}

	// return error some thing went wrong
	return nil, total, err
}
