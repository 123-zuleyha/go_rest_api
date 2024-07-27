package main 
import (
	
"github.com/123-zuleyha/go_rest_api/database" 
 "github.com/123-zuleyha/go_rest_api/router" 
 "github.com/gofiber/fiber/v2" 
 "github.com/gofiber/fiber/v2/middleware/cors" 
 "github.com/gofiber/fiber/v2/middleware/logger"
 "github.com/lib/pq" 

)

func main (){
	database.Connect()
	app:=fiber.New() // yeni bir fiber uygulaması oluşturuyoruz 
	app.Use(logger.New())// Logger middleware'ini uygulamaya ekler, istekleri loglar
     app.Use(cors.New())// cors middleware ini uygulamaya ekler
	 router.SetupRoutes(app)

	 app.Use(func(c*fiber.Ctx) error{
		return c.SendStatus(404)
	 })
	 app.Listen(":8080")
}
