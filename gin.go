package gin

import "net/http"

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method, path string, handler HandlerFunc) {
	key := method + "-" + path
	engine.router[key] = handler
}

func (engine *Engine) GET(path string, handler HandlerFunc) {
	engine.addRoute("GET", path, handler)
}

func (engine *Engine) POST(path string, handler HandlerFunc) {
	engine.addRoute("POST", path, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		http.Error(w, "404 Not Found", http.StatusNotFound)
	}
}
