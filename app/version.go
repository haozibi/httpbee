package app

import "fmt"

var (
	BuildTime    = ""
	BuildVersion = ""
	BuildAppName = ""
	CommitHash   = ""
)

func ShowLogo() {
	var logo = `
                    .' '.            __
            .       .   .           (__\_
HTTP Bee     .         .         . -{{_(|8)
%s        ' .  . ' ' .  . '     (__/
Github: https://github.com/haozibi/httpbee

`
	fmt.Printf(logo, BuildVersion)
}
