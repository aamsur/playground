// Copyright 2017 PT. Qasico Teknologi Indonesia. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package engine

import "github.com/aamsur/playground/simple-auth/src/user"

func init() {
	handlers["user"] = &user.Handler{}
}
