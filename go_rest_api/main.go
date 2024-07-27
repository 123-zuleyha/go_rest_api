package main 
import (
	// Veritabanı bağlantı işlemleri için kullanılan paket	
     "github.com/123-zuleyha/go_rest_api/database" 
     "github.com/123-zuleyha/go_rest_api/router" // Rota ayarları için kullanılan paket
 "github.com/gofiber/fiber/v2" // Fiber web framework'ü için kullanılan paket
 "github.com/gofiber/fiber/v2/middleware/cors" // CORS (Cross-Origin Resource Sharing) middleware'i için kullanılan paket
 "github.com/gofiber/fiber/v2/middleware/logger" // Logger middleware'i için kullanılan paket
 _ "github.com/lib/pq" // PostgreSQL sürücüsü, sadece yan etkileri için içe aktarılıyor

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