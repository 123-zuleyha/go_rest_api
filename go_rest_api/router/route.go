package router
import (	
  "github.com/123-zuleyha/go_rest_api/handler"
"github.com/gofiber/fiber/v2"
 )

  //SETUP ROUTES FUNC
  func SetupRoutes(app *fiber.App){
	api :=app.Group("/api")
	v1 :=api.Group("/user")
 
  //routes
  v1.Get("/",handler.GetAllUsers)
  v1.Get("/:id",handler.GetSingleUser)
  v1.Post("/",handler.CreateUser)
  v1.Put("/:id", handler.UpdateUser)
  v1.Delete("/:id", handler.DeleteUserByID)
 }