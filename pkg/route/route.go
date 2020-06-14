package route

// Message is message
type Message struct {
	Identification string
	Method         string
	Time           string
	Content        string
	Size           int
}

// Handler is handler
type Handler func(res, req *Message)

// Middleware is public middleware
type middleware func(Handler) Handler

// Router is router
type Router struct {
	middlewareChain []middleware
	mux             map[string]Handler
}

// NewRouter new a router
func NewRouter() *Router {
	return &Router{
		mux: make(map[string]Handler),
	}
}

// add a route
func (r *Router) add(route string, h Handler) {
	var mergedHandler = h
	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergedHandler = r.middlewareChain[i](mergedHandler)
	}
	r.mux[route] = mergedHandler
}

// Get add a get method pattern
func (r *Router) Get(route string, h Handler) {
	r.add("get:"+route, h)
}

// Put add a put method pattern
func (r *Router) Put(route string, h Handler) {
	r.add("put:"+route, h)
}

// Post add a post method pattern
func (r *Router) Post(route string, h Handler) {
	r.add("post:"+route, h)
}

// Delete add a delete method pattern
func (r *Router) Delete(route string, h Handler) {
	r.add("delete:"+route, h)
}

// Use add middleware
func (r *Router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}

// Run router
func (r *Router) Run(res, req *Message) error {
	route := req.Method + ":" + req.Identification
	handler, exists := r.mux[route]
	if exists {
		handler(res, req)
	}
	return nil
}
