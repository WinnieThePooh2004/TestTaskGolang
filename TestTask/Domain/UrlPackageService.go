package Domain

import (
	"TestTask/DataAccess/Repositories"
	"TestTask/Requests"
	"TestTask/Responses"
	"net"
)

type UrlPackageService struct {
	repository  Repositories.IUrlPackageRepository
	httpService IHttpService
}

func NewUrlPackageService(repository Repositories.IUrlPackageRepository, httpService IHttpService) IUrlPackageService {
	return UrlPackageService{repository: repository, httpService: httpService}
}

func (service UrlPackageService) GetMaxPrice(request Requests.Request) (*Responses.MaxPriceResponse,
	*Responses.BadIpResponse, *Responses.NotFoundUrlIdResponse, *Responses.RequiredValueIsNil, error) {

	if request.Ip == "" || request.UrlPackage == nil || len(request.UrlPackage) == 0 {
		return nil, nil, nil, &Responses.RequiredValueIsNil{}, nil
	}

	if validIP := net.ParseIP(request.Ip); validIP == nil || validIP.To4() == nil {
		return nil, &Responses.BadIpResponse{}, nil, nil, nil
	}

	urlPackage := removeDuplicates(request.UrlPackage)
	urls := make([]string, len(urlPackage))

	// two loops is needed to avoid calling http when some ids are invalid
	for i := 0; i < len(urlPackage); i++ {
		url := service.repository.GetById(urlPackage[i])
		if url == nil {
			return nil, nil, &Responses.NotFoundUrlIdResponse{}, nil, nil
		}
		urls[i] = url.Url
	}

	maxPrice := 0.0

	for i := 0; i < len(urls); i++ {
		price := service.httpService.Price(urls[i])
		if price > maxPrice {
			maxPrice = price
		}
	}

	return &Responses.MaxPriceResponse{MaxPrice: maxPrice}, nil, nil, nil, nil
}

func removeDuplicates(arr []int) []int {
	seen := make(map[int]bool)
	var result []int

	for _, item := range arr {
		if _, ok := seen[item]; !ok {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}
