package route

// Message is message
type Message struct {
	Time    string
	Content string
	Size    int
}

// Handler is query handler
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

func (r *Router) Add(route string, h Handler) {
	var mergedHandler = h
	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergedHandler = r.middlewareChain[i](mergedHandler)
	}
	r.mux[route] = mergedHandler
}

// Get add a get method pattern
func (r *Router) Get(dataType string, h Handler) {
	r.Add(dataType, h)
}

// Put add a put method pattern
func (r *Router) Put(dataType string, h Handler) {
	r.Add(dataType, h)
}

// Post add a post method pattern
func (r *Router) Post(dataType string, h Handler) {
	r.Add(dataType, h)
}

// Delete add a delete method pattern
func (r *Router) Delete(dataType string, h Handler) {
	r.Add(dataType, h)
}

// Use add middleware
func (r *Router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}

// Run router
func (r *Router) Run(res, req *Message) error {
	// route := string(message)
	handler, exists := r.mux["hello"]
	if exists {
		handler(res, req)
	}
	return nil
}
