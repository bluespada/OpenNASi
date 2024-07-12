// (c) Copyright Bluespada <pentingmain@gmail.com>
// (c) Copyright OpenNASi Developer
//
// This software under GPL-3.0 license please read accompanying
// file CoPY or read online at : https://gnu.org/licenses/gpl-3.html

package interfaces

type ApiResponse struct {
    Error bool `json:"error"`
    Messages string `json:"messages"`
    Data any `json:"any"`
}
