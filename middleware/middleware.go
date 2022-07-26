package main

func NewMiddleware(service string) *Middleware {
	return &Middleware{
		Service: service,
	}
}
