package route

// Message is message
type Message struct {
	Identification string
	Time           string
	Content        string
	Size           int
}

// Handler is handler
type Handler func(res, req *Message)

// Router is router
type Router struct {
	mux map[string]Handler
}

// NewRouter new a router
func NewRouter() *Router {
	return &Router{
		mux: make(map[string]Handler),
	}
}

// Add a route
func (r *Router) Add(route string, h Handler) {
	var mergedHandler = h
	r.mux[route] = mergedHandler
}

// Run router
func (r *Router) Run(res, req *Message) error {
	route := req.Identification
	handler, exists := r.mux[route]
	if exists {
		handler(res, req)
	}
	return nil
}
