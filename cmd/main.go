// (c) Copyright Bluespada <pentingmain@gmail.com>
// (c) Copyright OpenNASi Developer
//
// This software under GPL-3.0 license please read accompanying
// file CoPY or read online at : https://gnu.org/licenses/gpl-3.html

package main

import (
    "github.com/bluespada/OpenNASi/pkg/server"
)


// main function to run OpenNASi
func main(){
    var NasiServer = server.New()

    NasiServer.RunServer()

}
