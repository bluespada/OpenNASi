// (c) Copyright Bluespada <pentingmain@gmail.com>
// (c) Copyright OpenNASi Developer
//
// This software under GPL-3.0 license please read accompanying
// file CoPY or read online at : https://gnu.org/licenses/gpl-3.html

package module

// define module assets
type AssetType struct {
    // asset name
    AssetName string
    // asset embeded data
    AssetData string
    // asset extension
    AssetExtension int
}

// define module interface
type Module struct {
    // module id
    MID string `json:"mid"`
    // module name
    Name string `json:"name"`
    // module description
    Description string `json:"description"`
    // module version
    Version string `json:"version"`
    // module license
    License string
    // module Authors
    Auhtor string `json:"author"`
    // module maintainers
    Maintainers []string `json:"maintainers"`
    // module assets
    Assets []AssetType `json:"AssetType"`
    // is application module or extension module
    Application bool
    // module category
    Category []string
    AutoInstall bool
}
