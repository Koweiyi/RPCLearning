package handler

const HelloServiceName = "handler/HelloService"

type HelloService struct {
}

func (s *HelloService) Hello(req string, resp *string) error {
	*resp = "hello, " + req
	return nil
}
