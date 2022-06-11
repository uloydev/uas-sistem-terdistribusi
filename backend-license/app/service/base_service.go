package service

type BaseService[Req any, Resp any] interface {
	Create(request Req) (response Resp)
	List() (responses []Resp)
}
