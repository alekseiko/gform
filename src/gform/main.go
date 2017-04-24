package main

import (
	"fmt"
	"time"
    "os"

	"gopkg.in/corneldamian/httpwaymid.v0"
	"gopkg.in/corneldamian/httpway.v0"
    "gopkg.in/mgo.v2"
    "gform/controllers"
)


func main() {
    var server *httpway.Server
    session, err := mgo.Dial(env(DB_CONNECTION,"localhost")) // mongodb://<dbuser>:<dbpassword>@ds115671.mlab.com:15671
    if err != nil {
        panic(err)
    }
    defer session.Close()

    session.SetMode(mgo.Monotonic, true)

	router := httpway.New()

    landing := router.Middleware(httpwaymid.TemplateRenderer(env("TEMPLATE_DIR", "templates"), "tmpl", "vars", "status"))

	landing.GET("/", controllers.Index)
    landing.POST("/", controllers.StoreForm("1"))

    router.GET("/f/:id", controllers.FormHandler("1"))
    router.POST("/f/:id", controllers.FormHandler("1"))
    router.PUT("/f/:id", controllers.FormHandler("1"))

	server = httpway.NewServer(nil)
	server.Addr = fmt.Sprintf(":%s", env("SERVER_PORT", "8080"))
	server.Handler = router

	if err := server.Start(); err != nil {
		fmt.Println("Error", err)
		return
	}

	if err := server.WaitStop(10 * time.Second); err != nil {
		fmt.Println("Error", err)
	}
}

func env(key, defaultValue string) string {
    val := defaultValue
    if envVal := os.Getenv(key); envVal != "" {
       val=envVal
    }
    return val
}