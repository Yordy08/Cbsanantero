package controllers

import (
	"context"
	"sync"

	"github.com/cbsanantero/config"
	. "github.com/cbsanantero/db"
	"github.com/cbsanantero/db/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetRestaurante(c *fiber.Ctx) error {
	restaurante := Instance.Database.Collection("Restaurantes")

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
	restaurantes := Instance.Database.Collection("Restaurantes")

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
	restaurante := Instance.Database.Collection("Restaurantes")
	customer := Instance.Database.Collection("Customer")

	data := new(models.Restaurante)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo crear el restaurante"})
	}
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	customerID := form.Value["customer_id"]
	data.CustomerID = customerID[0]
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

	// Obtiene los archivos subidos
	files := form.File["image"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "La imagen es requerida",
		})
	}
	ImageFile := files[0]
	var UrlCloudinary string
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		UrlCloudinary = config.UploadImage(ImageFile)
		wg.Done()
	}()
	wg.Wait()

	if UrlCloudinary == "error al subir la imagen a cloudinary" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error al subir la imagen",
		})
	}

	data.Image = UrlCloudinary

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
	restaurantes := Instance.Database.Collection("Restaurantes")

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el restaurante"})
	}

	data := new(models.Restaurante)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el restaurante"})
	}
	if data.Image == "" {
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}
		files := form.File["imagen"]
		if len(files) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "La imagen es requerida",
			})
		}

		ImageFile := files[0]
		var UrlCloudinary string
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			UrlCloudinary = config.UploadImage(ImageFile)
			wg.Done()
		}()
		wg.Wait()

		if UrlCloudinary == "error al subir la imagen a cloudinary" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al subir la imagen",
			})
		
		}
		data.Image = UrlCloudinary
		_, err = restaurantes.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": data})
		if err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el restaurante"})
		}
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Restaurante actualizado"})
	}
	_, err = restaurantes.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": data})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el restaurante"})
	}

	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Restaurante actualizado"})
}

func DeleteRestaurante(c *fiber.Ctx) error {
	restaurantes := Instance.Database.Collection("Restaurantes")

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar el restaurante"})
	}

	var restaurante bson.M

	err = restaurantes.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&restaurante)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el restaurante"})
	}

	_, err = restaurantes.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar el restaurante"})
	}
	config.DeleteImage(restaurante["image"].(string))
	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Restaurante eliminado"})
}
