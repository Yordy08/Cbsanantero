package controllers

import (
	"context"
	"log"
	"sync"

	"github.com/cbsanantero/config"
	. "github.com/cbsanantero/db"
	"github.com/cbsanantero/db/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCustumer(c *fiber.Ctx) error {
	customer := Instance.Database.Collection("Customer")

	busqueda, err := customer.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	defer busqueda.Close(context.Background())

	var customers []models.Customer

	if err = busqueda.All(context.Background(), &customers); err != nil {
		log.Println(err)
	}

	return c.Status(fiber.StatusAccepted).JSON(customers)
}

func GetCustumerById(c *fiber.Ctx) error {
	customers := Instance.Database.Collection("Customer")

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}

	var customer bson.M

	err = customers.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&customer)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el cliente"})
	}

	return c.Status(fiber.StatusAccepted).JSON(customer)
}

func CreateCustomer(c *fiber.Ctx) error {
	customers := Instance.Database.Collection("Customer")

	customer := new(models.Customer)

	if err := c.BodyParser(customer); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(Message{Msg: "Error al analizar los datos"})
	}

	var customerExists models.Customer
	err := customers.FindOne(context.Background(), bson.M{"email": customer.Email}).Decode(&customerExists)
	if err == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "El correo ya está registrado"})
	}

	// Manejo del formulario multipart para la imagen
	form, err := c.MultipartForm()
	if err != nil {
		log.Println(err)
	}

	files := form.File["image"]
	if len(files) > 0 {
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

		customer.Image = UrlCloudinary
	} else {
		// Establece una imagen predeterminada si no se proporciona ninguna
		customer.Image = "https://i.pinimg.com/736x/76/f3/f3/76f3f3007969fd3b6db21c744e1ef289.jpg" // Reemplaza con la URL de la imagen predeterminada
	}

	// Asignar tipo de negocio si existe
	if tipoNegocios, ok := form.Value["tipo_negocio"]; ok && len(tipoNegocios) > 0 {
		customer.TipoNegocio = tipoNegocios[0]
	}

	customer.Status = "Inactivo"

	result, err := customers.InsertOne(context.Background(), customer)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo insertar el cliente"})
	}
	if result.InsertedID == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo insertar el cliente"})
	}

	return c.Status(fiber.StatusCreated).JSON(Message{Msg: "El usuario se ha registrado correctamente"})
}

func UpdateCustomer(c *fiber.Ctx) error {
	customers := Instance.Database.Collection("Customer")

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}

	customer := new(models.Customer)

	if err := c.BodyParser(customer); err != nil {
		log.Println(err)
	}
	if customer.Image == "" {
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
		customer.Image = UrlCloudinary
	}

	update := bson.M{
		"$set": customer,
	}

	result, err := customers.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		log.Println(err)
	}

	if result.ModifiedCount == 0 {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo actualizar el cliente"})
	} else {
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Cliente actualizado"})
	}
}

func DeleteCustomer(c *fiber.Ctx) error {
	customers := Instance.Database.Collection("Customer")
	artesanias := Instance.Database.Collection("Artesanias")
	bares := Instance.Database.Collection("Bares")
	hospedaje := Instance.Database.Collection("Hospedaje")
	hoteles := Instance.Database.Collection("Hoteles")
	recreacion := Instance.Database.Collection("Recreacion")
	restaurantes := Instance.Database.Collection("Restaurantes")
	tour := Instance.Database.Collection("Tour")
	transportes := Instance.Database.Collection("Traporte")
	galeria := Instance.Database.Collection("Galeria")
	servicio := Instance.Database.Collection("Servicio")

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}

	var customer bson.M

	filter := bson.M{"customer_id": id}

	err = customers.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&customer)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo encontrar el cliente"})
	}

	result, err := customers.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		log.Println(err)
	}

	_, err = artesanias.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}

	_, err = bares.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	_, err = hospedaje.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	_, err = hoteles.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	_, err = recreacion.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	_, err = restaurantes.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	_, err = tour.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	_, err = transportes.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	_, err = galeria.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	_, err = servicio.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotAcceptable).JSON(Message{Msg: "No se pudo eliminar el cliente"})
	} else {
		config.DeleteImage(customer["image"].(string))
		return c.Status(fiber.StatusAccepted).JSON(Message{Msg: "Cliente eliminado"})
	}
}

func UpdateCustomerStatus(c *fiber.Ctx) error {
	type Status struct {
		Status bool `json:"status,omitempty" bson:"status,omitempty"`
	}

	customers := Instance.Database.Collection("Customer")
	artesanias := Instance.Database.Collection("Artesanias")
	bares := Instance.Database.Collection("Bares")
	hospedaje := Instance.Database.Collection("Hospedaje")
	hoteles := Instance.Database.Collection("Hoteles")
	recreacion := Instance.Database.Collection("Recreacion")
	restaurantes := Instance.Database.Collection("Restaurantes")
	tour := Instance.Database.Collection("Tour")
	transportes := Instance.Database.Collection("Traporte")
	galeria := Instance.Database.Collection("Galeria")
	servicio := Instance.Database.Collection("Servicio")

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var customerStatus Status

	if err := c.BodyParser(&customerStatus); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var Estado string

	if customerStatus.Status {
		Estado = "Activo"
	} else {
		Estado = "Inactivo"
	}

	result, err := customers.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": bson.M{"status": Estado}})
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update customer"})
	}

	filter := bson.M{"customer_id": id}
	update := bson.M{"$set": bson.M{"status": Estado}}

	_, err = artesanias.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update artesanias"})
	}
	_, err = bares.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update bares"})
	}
	_, err = hospedaje.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update hospedaje"})
	}
	_, err = hoteles.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update hoteles"})
	}
	_, err = recreacion.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update recreacion"})
	}
	_, err = restaurantes.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update restaurantes"})
	}
	_, err = tour.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update tour"})
	}
	_, err = transportes.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update transportes"})
	}
	_, err = galeria.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update galeria"})
	}
	_, err = servicio.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update servicio"})
	}

	if result.ModifiedCount == 0 {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"message": "No se actualizo el cliente"})
	} else {
		if Estado == "Activo" {
			return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Cliente status actualizado a activo"})
		} else {
			return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Cliente status actualizado a inactivo"})
		}
	}
}

func UpdateCustomerRol(c *fiber.Ctx) error {
	type Rol struct {
		Rol string `json:"rol,omitempty" bson:"rol,omitempty"`
	}

	customers := Instance.Database.Collection("Customer")

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var customerRol Rol

	if err := c.BodyParser(&customerRol); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	result, err := customers.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": Rol{Rol: customerRol.Rol}})
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update rol"})
	}

	if result.ModifiedCount == 0 {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"message": "No se pudo actualizar el cliente"})
	} else {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Cliente actualizado"})
	}
}
