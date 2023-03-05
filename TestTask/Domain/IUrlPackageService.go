package Domain

import (
	"TestTask/Requests"
	"TestTask/Responses"
)

type IUrlPackageService interface {
	GetMaxPrice(request Requests.Request) (*Responses.MaxPriceResponse, *Responses.BadIpResponse,
		*Responses.NotFoundUrlIdResponse, *Responses.RequiredValueIsNil, error)
}
