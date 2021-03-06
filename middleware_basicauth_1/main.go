package main

import (
	"github.com/iris-contrib/middleware/basicauth"
	"github.com/kataras/iris"
)

func main() {
	authentication := basicauth.Default(map[string]string{"myusername": "mypassword", "mySecondusername": "mySecondpassword"})

	// to global iris.Use(authentication)
	// to party: iris.Party("/secret", authentication) { ... }

	// to routes
	iris.Get("/secret", authentication, func(ctx *iris.Context) {
		username := ctx.GetString("user") // this can be changed, you will see at the middleware_basic_auth_2 folder
		ctx.Writef("Hello authenticated user: %s ", username)
	})

	iris.Get("/secret/profile", authentication, func(ctx *iris.Context) {
		username := ctx.GetString("user")
		ctx.Writef("Hello authenticated user: %s from localhost:8080/secret/profile ", username)
	})

	iris.Get("/othersecret", authentication, func(ctx *iris.Context) {
		username := ctx.GetString("user")
		ctx.Writef("Hello authenticated user: %s from localhost:8080/othersecret ", username)
	})

	iris.Listen(":8080")
}
