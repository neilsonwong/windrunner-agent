package api

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)

type proxyServer struct {
	prefix string
	route  string
	url    *url.URL
	server *httputil.ReverseProxy
}

// ProxyRouter handles proxying to the local server
func ProxyRouter() http.Handler {
	proxy := newProxy()
	r := chi.NewRouter()
	r.Use(fixProxyHeadersMiddleware(&proxy))

	// proxy all routes
	r.Handle("/*", handleWithProxy(&proxy))

	return r
}

func newProxy() proxyServer {
	proxyAddr := viper.GetString("listing_server")
	proxyPrefix := viper.GetString("proxy_prefix")

	proxyURL, _ := url.Parse(proxyAddr)
	reverseProxy := httputil.NewSingleHostReverseProxy(proxyURL)

	proxy := proxyServer{}
	proxy.prefix = strings.TrimSuffix(proxyPrefix, "/")
	proxy.route = proxy.prefix + "/"
	proxy.url = proxyURL
	proxy.server = reverseProxy
	return proxy
}

// HandleProxy will proxy a request to the appropriate url
func handleWithProxy(ps *proxyServer) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ps.server.ServeHTTP(res, req)
	})
}

// not beautiful double wrapped function, but o well
func fixProxyHeadersMiddleware(ps *proxyServer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			// rewrite the req path
			req.URL.Host = ps.url.Host
			req.URL.Scheme = ps.url.Scheme
			// use the requestURI as that preserves the "non sanitation" needed to preserve urlEncodes
			req.URL.Path = strings.Replace(req.URL.RequestURI(), ps.prefix, "", 1)
			req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
			req.Host = ps.url.Host
			next.ServeHTTP(res, req)
		})
	}
}
