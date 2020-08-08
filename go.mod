module FrontEndOJGolang

go 1.13

require (
	FrontEndOJudger v0.0.0-incompatible
	github.com/chromedp/cdproto v0.0.0-20200116234248-4da64dd111ac
	github.com/chromedp/chromedp v0.5.3
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.57.0
	github.com/go-sql-driver/mysql v1.5.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
)

replace FrontEndOJudger => ../FrontEndOJudger
