package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Project defines the Project MEdiaType
var Project = MediaType("application/vnd.project+json", func() {
	Description("Describes a project")
	Attributes(func() {
		Attribute("id", String, "The id of the project")
		Attribute("name", String, "The name of the project")
		Required("name", "id")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
	})
	View("create", func() {
		Attribute("name")
	})
	View("list", func() {
		Attribute("id")
		Attribute("name")
	})
})
