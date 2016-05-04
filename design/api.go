package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("alm", func() {
	Title("xxx")
	Description("xxxx")
	Version("1.0")
	Host("test.io")
	Scheme("http")
	BasePath("/api")
	Consumes("application/json")
	Produces("application/json")

})
