// 
// 
// 

package user_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aamsur/playground/simple-auth/datastore/model"
	"github.com/aamsur/playground/simple-auth/src/auth"
	"github.com/aamsur/playground/simple-auth/test"

	"git.qasico.com/cuxs/common/tester"
	"github.com/stretchr/testify/assert"
	"git.qasico.com/cuxs/common"
	"net/http"
)

func TestMain(m *testing.M) {
	test.Setup()

	// run tests
	res := m.Run()

	// cleanup
	test.DataCleanUp()

	os.Exit(res)
}

func TestARouting(t *testing.T) {
	user := &model.User{Username: "testing_login"}
	user.Save()
	sd, _ := auth.Login(user)
	token := "Bearer " + sd.Token

	dummyuser := &model.User{Username: "dummy_user"}
	dummyuser.Save()

	// sukses
	var routers = []struct {
		endpoint string
		method   string
		expected int
	}{
		{"/v1/user", "GET", 200},
	}

	ng := tester.New()
	ng.SetHeader(tester.H{"Authorization": token})
	for _, ep := range routers {
		ng.Method = ep.method
		ng.Path = ep.endpoint
		ng.Run(test.Router(), func(res tester.HTTPResponse, req tester.HTTPRequest) {
			assert.Equal(t, ep.expected, res.Code, fmt.Sprintf("Should has 'endpoint %s' with method '%s'", ep.endpoint, ep.method))
		})
	}
}

func TestHandler_URLMapping_DELETE_deleteuser(t *testing.T) {
	user := &model.User{Username: "testing_login"}
	user.Save()
	sd, _ := auth.Login(user)
	token := "Bearer " + sd.Token

	dummyuser := &model.User{Username: "testing_dummy"}
	dummyuser.Save()
	euserID := common.Encrypt(dummyuser.ID)

	// dummy user masih ada
	m, e := auth.GetUsername(dummyuser.Username)
	assert.NoError(t, e, "seharusnya tidak error")
	assert.Equal(t, dummyuser.Username, m.Username)

	ng := tester.New()
	ng.SetHeader(tester.H{"Authorization": token})
	request := tester.D{}
	ng.DELETE("/v1/user/" + euserID).SetJSON(request).Run(test.Router(), func(res tester.HTTPResponse, req tester.HTTPRequest) {
		assert.Equal(t, http.StatusOK, res.Code, fmt.Sprintf("\nreason: Validation Not Matched,\ndata: %v , \nresponse: %v", request, res.Body.String()))
	})

	// setelah proses delete, dummy user tidak ada
	m, e = auth.GetUsername(dummyuser.Username)
	assert.Error(t, e, "seharusnya error")
	assert.Nil(t, m)
}