package RequestHandlers

import (
	"TestTask/DataFactories"
	"TestTask/Requests"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
)

func HandleGetRequest(factory *DataFactories.DataFactory, context *fiber.Ctx) {
	service, err := factory.GetUrlPackageService()

	if err != nil {
		context.Status(204)
		return
	}

	urlPackage, err := parseUrlPackage(context)

	if err != nil {
		context.Status(204)
		return
	}

	ip := context.Query("ip")
	request := Requests.Request{Ip: ip, UrlPackage: urlPackage}
	parseRequestId(context, request)

	response, badIp, notFoundUrlIdResponse, nilValueResponse, err := service.GetMaxPrice(request)

	if err != nil {
		context.Status(204)
		return
	}

	if badIp != nil || nilValueResponse != nil || notFoundUrlIdResponse != nil {
		context.Status(204)
		return
	}

	jsonBytes, err := json.Marshal(response)

	if err != nil {
		context.Status(204)
		return
	}

	jsonString := string(jsonBytes)

	err = context.SendString(jsonString)

	if err != nil {
		return
	}
}

func parseRequestId(context *fiber.Ctx, request Requests.Request) {
	idAsString := context.Query("request_id")
	if idAsString != "" {
		return
	}

	id, err := strconv.ParseInt(idAsString, 0, 32)
	if err == nil {
		return
	}

	idAsInt := int(id)
	request.RequestId = &idAsInt
}

func parseUrlPackage(context *fiber.Ctx) ([]int, error) {
	var urlPackage []int
	queryString := string(context.Request().URI().QueryString())
	parameters := strings.Split(queryString, "&")

	for i := 0; i < len(parameters); i++ {
		keyValue := strings.Split(parameters[i], "=")
		if !(len(keyValue) == 2 && keyValue[0] == "url_package") {
			continue
		}

		parsed, err := strconv.ParseInt(keyValue[1], 0, 32)
		if err != nil {
			return nil, err
		}

		urlPackage = append(urlPackage, int(parsed))
	}

	return urlPackage, nil
}
