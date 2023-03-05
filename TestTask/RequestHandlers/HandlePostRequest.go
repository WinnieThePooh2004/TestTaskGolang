package RequestHandlers

import (
	"TestTask/DataFactories"
	"TestTask/Requests"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func HandlePostRequest(factory *DataFactories.DataFactory, context *fiber.Ctx) {
	var request Requests.Request
	err := json.Unmarshal(context.Body(), &request)
	if err != nil {
		context.Status(204)
		return
	}

	service, err := factory.GetUrlPackageService()
	if err != nil {
		context.Status(204)
		return
	}

	response, badIdResponse, notFoundUrlResponse, nilValueResponse, err := service.GetMaxPrice(request)

	if badIdResponse != nil || notFoundUrlResponse != nil || nilValueResponse != nil {
		context.Status(204)
		return
	}

	jsonBytes, err := json.Marshal(response)

	if err != nil {
		context.Status(204)
	}

	jsonString := string(jsonBytes)

	err = context.SendString(jsonString)

	if err != nil {
		return
	}
}
