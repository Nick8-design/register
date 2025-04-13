package handles

import (
	"crypto/sha256"
	"encoding/hex"

	"kt_go/db"
	"kt_go/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx)error{
var regUser models.RegModel

if err:=c.BodyParser(&regUser);err!=nil{
	return c.Status(400).JSON(fiber.Map{"msg":"Bad Request"});
}

    psd:=sha256.Sum256([]byte(regUser.Password))
	regUser.Password = hex.EncodeToString(psd[:])

	if err:=db.Db.Create(&regUser).Error;err!=nil{
		return c.Status(500).JSON(fiber.Map{"msg":"Failed to save to the database"})

	}

	return c.Status(200).JSON(fiber.Map{"msg":"Successfully registerd"})
}

func Login(c *fiber.Ctx)error{
	var logUser models.Login

	if err:=c.BodyParser(&logUser);err!=nil{
		return c.Status(401).JSON(fiber.Map{"msg":"Invalid Credentials"});
	}
	 
	var regUser models.RegModel

	if err:=db.Db.Where("Email = ?",logUser.Email).First(&regUser).Error;err!=nil{
		return c.Status(404).JSON(fiber.Map{"msg":"Invalid Credentials"});
	}



	 psd:=sha256.Sum256([]byte(logUser.Password))
	if regUser.Password == hex.EncodeToString(psd[:]) {
		return c.Status(200).JSON(fiber.Map{"msg":"ok"});
	}

	return c.Status(401).JSON(fiber.Map{"msg":"Invalid Password"});
}

func Ping(c *fiber.Ctx)error{
	return c.Status(200).JSON(fiber.Map{"msg":"pong"});
}