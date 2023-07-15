package main

import (
	"github.com/kamaal111/pocket-slate-api/src/routers"
)

//	@title			Pocket slate API
//	@version		1.0
//	@description	API for pocket slate

//	@license.name	MIT
//	@license.url	https://github.com/kamaal111/pocket-slate-api/blob/main/LICENSE

//	@BasePath	/api/v1
func main() {
	routers.Start()
}
