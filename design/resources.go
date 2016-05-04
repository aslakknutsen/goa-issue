package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("project", func() {
	DefaultMedia(Project)
	BasePath("/p")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Get projects for current authorized user")
		Response(OK, func() {
			Media(CollectionOf(Project, func() {
				View("list")
			}))
		})
		Response(NoContent)
	})
	Action("get", func() {
		Routing(
			GET("/:userID"),
		)
		Description("Get a single project for authorized user")
		Params(func() {
			Param("userID", Integer)
		})
		Response(OK)
		Response(NoContent)
	})
})
