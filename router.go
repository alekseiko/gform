package httpway

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request)

func New() *Router {
	return &Router{
		Router: httprouter.New(),
	}
}

type Router struct {
	*httprouter.Router
	SessionManager SessionManager
	Logger         Logger

	prev   *Router
	handle Handler
}

// register a GET handler with path
func (r *Router) GET(path string, handle Handler) {
	r.Handle("GET", path, handle)
}

// register a HEAD handler with path
func (r *Router) HEAD(path string, handle Handler) {
	r.Handle("HEAD", path, handle)
}

// register a OPTIONS handler with path
func (r *Router) OPTIONS(path string, handle Handler) {
	r.Handle("OPTIONS", path, handle)
}

// register a POST handler with path
func (r *Router) POST(path string, handle Handler) {
	r.Handle("POST", path, handle)
}

// register a PUT handler with path
func (r *Router) PUT(path string, handle Handler) {
	r.Handle("PUT", path, handle)
}

// register a PATCH handler with path
func (r *Router) PATCH(path string, handle Handler) {
	r.Handle("PATCH", path, handle)
}

// register a DELETE handler with path
func (r *Router) DELETE(path string, handle Handler) {
	r.Handle("DELETE", path, handle)
}

// Handle registers a new request handle with the given path and method.
//
// For GET, POST, PUT, PATCH and DELETE requests the respective shortcut
// functions can be used.
//
// This function is intended for bulk loading and to allow the usage of less
// frequently used, non-standardized or custom methods (e.g. for internal
// communication with a proxy).
func (r *Router) Handle(method, path string, handle Handler) {
	newHandle := r.GenerateChainHandler(handle)

	r.Router.Handle(method, path, newHandle)
}

//Add a middleware before (and after) the handler run
//   router := httpway.New()
//   public := router.Middleware(AccessLogger)
//   private := public.Middleware(AuthCheck)
//
//   public.GET("/public", somePublicHandler)
//   private.GET("/private", somePrivateHandler)
//
//  func AccessLogger(w http.ResponseWriter, r *http.Request) {
//  	startTime:=time.Now()
//
//	httpway.GetContext(r).Next(w, r)
//
//  	fmt.Printf("Request: %s duration: %s\n", r.URL.EscapedPath(), time.Since(startTime))
//  }
//
//  func AuthCheck(w http.ResponseWriter, r *http.Request) {
//	ctx := httpway.GetContext(r)
//
//  	if !ctx.Session().IsAuth() {
//		http.Error(w, "Auth required", 401)
//		return
//  	}
//	ctx.Next(w, r)
//  }
//
func (r *Router) Middleware(handle Handler) *Router {
	rt := &Router{
		prev:           r,
		handle:         handle,
		Router:         r.Router,
		Logger:         r.Logger,
		SessionManager: r.SessionManager,
	}

	return rt
}

//get httprouter handler with all the middlewares chained
func (router *Router) GenerateChainHandler(handle Handler) httprouter.Handle {
	if router.prev == nil {
		return func(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
			w = CreateContext(router, w, r, nil, nil, &pr)
			handle(w, r)
		}
	}

	var (
		lastMiddleware Handler
		middlewareList = make([]Handler, 0)
	)

	mid := router
	middlewareList = append(middlewareList, handle)

	for mid.prev != nil {
		if mid.prev.handle == nil {
			lastMiddleware = mid.handle
			break
		}
		middlewareList = append(middlewareList, mid.handle)
		mid = mid.prev
	}
	middlewareListLen := len(middlewareList)

	httprouterHandler := func(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
		w = CreateContext(router, w, r, &middlewareList, &middlewareListLen, &pr)

		lastMiddleware(w, r)
	}

	return httprouterHandler
}
