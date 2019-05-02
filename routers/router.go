package routers

import (
	"crawl_movie/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/crawlMovie",&controllers.CrawlMovieController{},"get:Get;post:Post")
}
