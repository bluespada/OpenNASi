// (c) Copyright Bluespada <pentingmain@gmail.com>
// (c) Copyright OpenNASi Developer
//
// This software under GPL-3.0 license please read accompanying
// file CoPY or read online at : https://gnu.org/licenses/gpl-3.html

package server

import (
    "fmt"

    "github.com/valyala/fasthttp"
    "github.com/fasthttp/router"
)

// server structure
type Server struct {
}

// initialize server
func New() *Server {
    return &Server{
    }
}

// start and run http server
func(s *Server) ListenAndRun(port int){
    
    // initialize fasthttp and the route
    fasthttp.ListenAndServe(fmt.Sprintf(":%d", port), s.initRoute().Handler)
}


// this function is to handle OpenNASi Web Server Route
// including to serve global middleware and the assets.
func (s *Server) initRoute() *router.Router {
    r := router.New()
    // initialize route

    // handle global middleware

    // handle assets route

    // handle website route

    // handle api route

    // return router
    return r
}

