package router

import (
	"go-mongodb/common"
	"go-mongodb/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddPlayerGroup(app *fiber.App) {
	playerGroup := app.Group("/players")

	playerGroup.Get("/", getPlayers)
	playerGroup.Get("/:id", getPlayer)
	playerGroup.Post("/", createPlayer)
	playerGroup.Put("/:id", updatePlayer)
	playerGroup.Delete("/:id", deletePlayer)
}

// Get players func
func getPlayers(c *fiber.Ctx) error {
	coll := common.GetDBCollection("books")

	// find all players
	players := make([]model.Player, 0)
	cursor, err := coll.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// iterate over the cursor
	for cursor.Next(c.Context()) {
		player := model.Player{}
		err := cursor.Decode(&player)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		players = append(players, player)
	}

	return c.Status(200).JSON(fiber.Map{"data": players})
}

// Get player func
func getPlayer(c *fiber.Ctx) error {
	coll := common.GetDBCollection("players")

	// find the player
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	player := model.Player{}

	err = coll.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&player)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{"data": player})
}

type createDTO struct {
	Nickname string `json:"nickname" bson:"nickname"`
	Fraction string `json:"fraction" bson:"fraction"`
	Level    string `json:"level" bson:"level:"`
}

// Create player func
func createPlayer(c *fiber.Ctx) error {
	// validate the body
	b := new(createDTO)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	// create the player
	coll := common.GetDBCollection("players")
	result, err := coll.InsertOne(c.Context(), b)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create player",
			"message": err.Error(),
		})
	}

	// return the player
	return c.Status(201).JSON(fiber.Map{
		"result": result,
	})
}

type updateDTO struct {
	Nickname string `json:"nickname" bson:"nickname"`
	Fraction string `json:"fraction" bson:"fraction"`
	Level    string `json:"level" bson:"level:"`
}

// Update player func
func updatePlayer(c *fiber.Ctx) error {
	// validate the body
	b := new(updateDTO)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	// get the id
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id",
		})
	}

	// update the player
	coll := common.GetDBCollection("books")
	result, err := coll.UpdateOne(c.Context(), bson.M{"_id": objectId}, bson.M{"$set": b})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update player",
			"message": err.Error(),
		})
	}

	// return the player
	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})
}

func deletePlayer(c *fiber.Ctx) error {
	// get the id
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id",
		})
	}

	// delete the player
	coll := common.GetDBCollection("books")
	result, err := coll.DeleteOne(c.Context(), bson.M{"_id": objectId})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete player",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})
}
