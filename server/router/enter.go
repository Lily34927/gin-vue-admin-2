package router

import (
	"server/router/example"
	"server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupAPP = new(RouterGroup)
