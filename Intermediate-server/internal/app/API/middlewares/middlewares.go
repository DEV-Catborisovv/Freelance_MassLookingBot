package middlewares

// Here i'm using fabric-pattern

type IMiddleware interface{}

type Middleware struct {
	IMiddleware
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}
