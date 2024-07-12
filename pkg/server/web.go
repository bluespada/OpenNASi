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

func (server *OpenNASiServer) webRouter(r *router.Router){
    
    r.GET("/", server.baseHandler)

    r.NotFound = func(ctx *fasthttp.RequestCtx){
        fmt.Fprintf(ctx, "Page Not Found")
    }
}
