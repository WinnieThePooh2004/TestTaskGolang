package Domain

import (
	"TestTask/Responses"
	"encoding/json"
	"net/http"
)

type HttpService struct {
	client *http.Client
}

func NewHttpService() IHttpService {
	client := &http.Client{}
	return HttpService{client: client}
}

func (service HttpService) Price(url string) float64 {
	response, err := service.client.Get(url)

	if err != nil {
		return 0
	}

	var responseContent Responses.PriceResponse
	err = json.NewDecoder(response.Body).Decode(&responseContent)
	if err != nil {
		return 0
	}

	return responseContent.Price
}
