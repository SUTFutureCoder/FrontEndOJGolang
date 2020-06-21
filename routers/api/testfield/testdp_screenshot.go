package testfield

import (
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"context"
	"flock/util/logs"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
)

func ScreenShot(c *gin.Context) {

	appG := app.Gin{C: c}
	
	url, _ := c.GetPostForm("url")
	quantility, _ := c.GetPostForm("quanlity")

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()


	var buf []byte
	if err := chromedp.Run(ctx, elementScreenShot(url, "#app", &buf)); err != nil {
		logs.Fatalf("chromedp run error:%v", err)
	}
	if err := ioutil.WriteFile("elementScreenShot.png", buf, 0777); err != nil {
		logs.Fatalf("chrome dp save file error:%v", err)
	}

	int64, _ := strconv.ParseInt(quantility, 10, 64)
	if err := chromedp.Run(ctx, elementFullScreenShot(url, int64, &buf)); err != nil {
		logs.Fatalf("chromedp run error:%v", err)
	}
	if err := ioutil.WriteFile("elementScreenShotFullScreen.png", buf, 0777); err != nil {
		logs.Fatalf("chrome dp save file error:%v", err)
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

func elementFullScreenShot(url string, quantity int64, buf *[]byte) chromedp.Action {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady("#app", chromedp.ByID),
		//chromedp.Sleep(60 * time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type: emulation.OrientationTypePortraitPrimary,
					Angle: 0,
			}).Do(ctx)
			if err != nil {
				return err
			}

			*buf, err = page.CaptureScreenshot().
				WithQuality(quantity).
				WithClip(&page.Viewport{
					X:	contentSize.X,
					Y:  contentSize.Y,
					Width: contentSize.Width,
					Height: contentSize.Height,
					Scale: 1,
			}).Do(ctx)

			if err != nil {
				return err
			}
			return nil
		}),
	}
}

func elementScreenShot(url, sel string, buf *[]byte) chromedp.Action {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(sel, chromedp.ByID),
		chromedp.Screenshot(sel, buf, chromedp.NodeVisible, chromedp.ByID),
	}
}
