// 
// 
// 

package engine

import "github.com/aamsur/playground/simple-auth/src/user"

func init() {
	handlers["user"] = &user.Handler{}
}
