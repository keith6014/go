package design

import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("dbexample", func() { // API defines the microservice endpoint and
	Title("db example")                 // other global properties. There should be one
	Description("A simple goa service") // and exactly one API definition appearing in
	Scheme("http")                      // the design.
	Host("localhost:8080")
})

var _ = Resource("account", func() {
	DefaultMedia(Account)
	BasePath("/accounts")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all accounts.")
		Response(OK, CollectionOf(Account))
	})

	Action("add", func() {
		Routing(GET("/add/:left/:right"))
		Description("Add 2 entries")
		Params(func() {
			Param("left", Integer, "Left")
			Param("right", Integer, "Right")

		})
		Response(OK, "text/plain")
	})
})

var Account = MediaType("application/vnd.account+json", func() {
	Description("A tenant account")
	Attributes(func() {
		Attribute("id", Integer, "ID of account", func() {
			Example(1)
		})
		Attribute("cid", Integer, "cid of account", func() {
			Example(2)
		})
		Required("id", "cid")
	})
	View("default", func() {
		Attribute("id")
		Attribute("cid")
	})

})

var _ = Resource("bottle", func() { // Resources group related API endpoints
	BasePath("/bottles")      // together. They map to REST resources for REST
	DefaultMedia(BottleMedia) // services.

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get bottle by id") // with its path, parameters (both path
		Routing(GET("/:bottleID"))      // parameters and querystring values) and payload
		Params(func() {                 // (shape of the request body).
			Param("bottleID", Integer, "Bottle ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})
})

// BottleMedia defines the media type used to render bottles.
var BottleMedia = MediaType("application/vnd.goa.example.bottle+json", func() {
	Description("A bottle of wine")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique bottle ID")
		Attribute("href", String, "API href for making requests on the bottle")
		Attribute("name", String, "Name of wine")
		Required("id", "href", "name")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id")   // Media types may have multiple views and must
		Attribute("href") // have a "default" view.
		Attribute("name")
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Files("/swagger.json", "swagger/swagger.json")
})
