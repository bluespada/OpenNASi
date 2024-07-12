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
    "github.com/bluespada/OpenNASi/pkg/interfaces"
    "github.com/bytedance/sonic"
)

func (server *OpenNASiServer) apiRouter(r *router.Router){
    api := r.Group("/api")

    api.GET("/", func(ctx *fasthttp.RequestCtx){
        // create reponse
        var response = &interfaces.ApiResponse{
            Error: false,
            Messages: "",
            Data: []string{},
        };

        // set header as application json
        ctx.SetStatusCode(fasthttp.StatusOK)

        // create json response
        data, err := sonic.Marshal(response)

        // set internal server error because parsing error
        if err != nil {
            ctx.SetStatusCode(fasthttp.StatusInternalServerError)
            fmt.Fprintf(ctx, fmt.Sprintf("Internal Error : %s", err.Error()))
            return
        }
        fmt.Fprintf(ctx, string(data))
    })
}
