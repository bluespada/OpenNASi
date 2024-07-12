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


// create server structure
type OpenNASiServer struct {

}


// initialize OpenNASiServer
func New() *OpenNASiServer {
    fmt.Println("Start OpenNASi Server")
    return &OpenNASiServer{
    }
}

func (server *OpenNASiServer) RunServer(){
    // initialize router
    r := router.New()

    server.webRouter(r)
    server.apiRouter(r)

    // initialize router handler
    fasthttp.ListenAndServe(":8000", r.Handler)
}

func (server *OpenNASiServer) baseHandler(ctx *fasthttp.RequestCtx) {
    fmt.Fprintf(ctx, "Hello World")
}


// initialize required dependency or other
func init(){

}
