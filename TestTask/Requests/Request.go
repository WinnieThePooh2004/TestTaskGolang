package Requests

type Request struct {
	RequestId  *int   `json:"request_id"`
	UrlPackage []int  `json:"url_package"`
	Ip         string `json:"ip"`
}
