package routes

import "net/http"

func Api(mux *http.ServeMux) {
	routes := map[string]http.HandlerFunc{}

	routes["/"] = func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello World</h1>"))
	}

	for route, handler := range routes {
		mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			if exist := routes[r.URL.Path]; exist != nil {
				handler(w, r)
				return
			}
			func() {
				w.Write([]byte("<h1>404 NotFound</h1>"))
			}()
		})
	}
}
