module FrontEndOJGolang

go 1.13

require (
	FrontEndOJudger v0.0.0-incompatible
	github.com/chromedp/cdproto v0.0.0-20200116234248-4da64dd111ac
	github.com/chromedp/chromedp v0.5.3
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.57.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorilla/sessions v1.2.0
	github.com/kr/pretty v0.1.0 // indirect
	golang.org/x/crypto v0.0.0-20201217014255-9d1352758620
	golang.org/x/sys v0.0.0-20201218084310-7d0127a74742 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

replace FrontEndOJudger => ../FrontEndOJudger
