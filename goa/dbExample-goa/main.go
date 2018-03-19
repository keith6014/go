//go:generate goagen bootstrap -d dbExample-goa/design

package main

import (
	"dbExample-goa/app"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {

	// Create service
	service := goa.New("dbexample")

	//	w := log.NewSyncWriter(os.Stderr)
	//	logger := log.NewLogfmtLogger(w)
	//	service.WithLogger(log.New(logger))

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "account" controller
	c := NewAccountController(service)
	app.MountAccountController(service, c)

	// Mount "bottle" controller
	c2 := NewBottleController(service)
	app.MountBottleController(service, c2)
	// Mount "swagger" controller
	c3 := NewSwaggerController(service)
	app.MountSwaggerController(service, c3)

	service.ServeFiles("/*filepath", "./public/")
	service.ServeFiles("/", "./html/")

	service.LogInfo("Starting service")
	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
