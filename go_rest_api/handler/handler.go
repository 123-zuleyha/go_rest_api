package handler

import (
	"github.com/123-zuleyha/go_rest_api/database"
	"github.com/123-zuleyha/go_rest_api/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CREATE USER
func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "something's wrong with your input",
			"data":    err,
		})
	}
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": user})
}

//GET ALL USERS

func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []model.User

	db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Users Found", "data": users})

}

// GET SİNGLE USER fromdb
func GetSingleUser(c *fiber.Ctx) error {
	db := database.DB.Db

	id := c.Params("id")
	var user model.User

	db.Find(&user, "id= ?", id)

	if user.ID == uuid.Nil {
		return c.Status(200).JSON(fiber.Map{"status": "success", "message": "user not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "user found", "data": user})
}

// UPDATE USER
func UpdateUser(c *fiber.Ctx) error {

	type UpdateUser struct {
		Username string `json:"username"`
	}
	db := database.DB.Db
	var user model.User

	id := c.Params("id")

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "success", "message": "user not found", "data": nil})

	}
	var updateUserData UpdateUser
	err := c.BodyParser(&updateUserData)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "something wrong with your input", "data": err})

	}
	user.Username = updateUserData.Username
	db.Save(&user)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users Found", "data": user})

}

// DELETE A USER-- delete user in db by id
func DeleteUserByID(c *fiber.Ctx) error {
	db := database.DB.Db

	var user model.User

	id := c.Params("id")

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})

	}
	err := db.Delete(&user, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}
