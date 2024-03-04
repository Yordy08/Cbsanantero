package controllers

import (
	"context"
	"time"

	"github.com/cbsanantero/config"
	"github.com/cbsanantero/db"
	"github.com/cbsanantero/db/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetRestaurante(c *fiber.Ctx) error {
	restaurante := db.Restaurantes

	busqueda, err := restaurante.Find(context.TODO(), bson.M{"status": "Activo"})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el restaurante"})
	}
	defer busqueda.Close(context.Background())

	var restaurantes []bson.M
	if err = busqueda.All(context.Background(), &restaurantes); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el restaurante"})
	}
	return c.Status(fiber.StatusAccepted).JSON(restaurantes)
}

func GetRestauranteById(c *fiber.Ctx) error {
	restaurantes := db.Restaurantes

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el restaurante"})
	}

	var restaurante bson.M

	err = restaurantes.FindOne(context.TODO(), bson.M{"_id": objID, "status": "Activo"}).Decode(&restaurante)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el restaurante"})
	}

	return c.Status(fiber.StatusAccepted).JSON(restaurante)
}

func CreateRestaurante(c *fiber.Ctx) error {
	restaurante := db.Restaurantes
	customer := db.Customer

	data := new(models.Restaurante)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el restaurante"})
	}
	idc := data.CustomerID

	objID, err := primitive.ObjectIDFromHex(idc)

	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "El _id no es valido"})
	}

	busqueda := customer.FindOne(context.Background(), bson.M{"_id": objID})

	var customerData models.Customer

	if err = busqueda.Decode(&customerData); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el cliente"})

	}
	data.Status = "Activo"
	go config.UploadImageLocal(data.Image)
	time.Sleep(1 * time.Second)
	data.Image = config.UploadImage()

	insertion, err := restaurante.InsertOne(context.Background(), data)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el restaurante"})
	}

	if insertion.InsertedID == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el restaurante"})
	} else {
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Restaurante creado"})
	}
}

func UpdateRestaurante(c *fiber.Ctx) error {
	restaurantes := db.Restaurantes

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el restaurante"})
	}

	data := new(models.Restaurante)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el restaurante"})
	}
	if data.Image != "" {
		go config.UploadImageLocal(data.Image)
		time.Sleep(1 * time.Second)
		data.Image = config.UploadImage()
	}
	_, err = restaurantes.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": data})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el restaurante"})
	}

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Restaurante actualizado"})
}

func DeleteRestaurante(c *fiber.Ctx) error {
	restaurantes := db.Restaurantes

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar el restaurante"})
	}

	_, err = restaurantes.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar el restaurante"})
	}

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Restaurante eliminado"})
}
