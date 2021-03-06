package main

import (
	"fmt"

	"github.com/kataras/go-template/django"
	"github.com/kataras/iris"
)

type Visitor struct {
	Username string
	Mail     string
	Data     []string `form:"mydata"`
}

func main() {
	// Yes you see right, only one line change and ou'ready to go, defaults to ./templates as directory and .html as file extensions
	iris.UseTemplate(django.New())

	iris.Get("/", func(ctx *iris.Context) {
		if err := ctx.Render("form.html", nil); err != nil {
			iris.Logger.Panic(err.Error())
		}
	})

	iris.Post("/form_action", func(ctx *iris.Context) {
		visitor := Visitor{}
		err := ctx.ReadForm(&visitor)
		if err != nil {
			fmt.Println("Error when reading form: " + err.Error())
		}
		fmt.Printf("\n Visitor: %#v", visitor)
		ctx.Writef("%#v", visitor)
	})

	iris.Listen(":8080")

}
