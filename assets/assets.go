package assets

import "embed"

//go:embed all:css
var Css embed.FS

//go:embed all:js
var Js embed.FS
