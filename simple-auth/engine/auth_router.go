// 
// 
// 

package engine

import "github.com/aamsur/playground/simple-auth/src/auth"

func init() {
	handlers["auth"] = &auth.Handler{}
}
