package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (manager *Manager) Use(middlewares ...Middleware) {
	manager.globalMiddlewares = append(manager.globalMiddlewares, middlewares...)
}

func (manager *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next

	// for i := len(middlewares) - 1; i >= 0; i-- {
	// 	n = middlewares[i](n)
	// }

	for _, middleware := range middlewares {
		n = middleware(n)
	}

	for _, middleware := range manager.globalMiddlewares {
		n = middleware(n)
	}

	return n
}
