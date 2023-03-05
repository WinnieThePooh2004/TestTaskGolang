package Domain

type IHttpService interface {
	Price(url string) float64
}
