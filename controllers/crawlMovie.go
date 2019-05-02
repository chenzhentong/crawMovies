package controllers

import (
	"crawl_movie/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"strings"
	"time"
)

type CrawlMovieController struct {
	beego.Controller
}
func (c *CrawlMovieController) Get(){
	//连接redis
	models.ConnectRedis("127.0.0.1:6379")

	sUrl:="https://movie.douban.com/subject/33393084/?tag=热门&amp;from=gaia_video"
	//存储电影信息
	var movieInfo models.MovieInfo
	//把初始url加入队列
	models.PutinQueue(sUrl)

	for{
		lenth:=models.GetQueueLength()
		//如果队列为空，停止
		if lenth==0{
			break
		}
		//成功获取到url
		sUrl=models.PopfromQueue()
		if models.IsVisit(sUrl){
			continue
		}
		movieHtml:=httplib.Get(sUrl)
		//将网页html转换为string
		sMovieHtml,err:=movieHtml.String()
		if err!=nil{
			c.Ctx.WriteString(fmt.Sprintf("err:%v\n",err))
		}


		movieInfo.Movie_name=models.GetMovieName(sMovieHtml)
		//如果当前链接是电影，提取出电影信息并入库
		if movieInfo.Movie_name!=""{
			c.Ctx.WriteString("正在从"+sUrl+"爬取电影"+movieInfo.Movie_name+"\n")
			movieInfo.Movie_release_time=models.GetMovieReleaseTime(sMovieHtml)
			movieInfo.Movie_director=models.GetMovieDirector(sMovieHtml)
			movieInfo.Movie_main_character=models.GetMainCharacter(sMovieHtml)
			movieInfo.Movie_on_time=models.GetOnTime(sMovieHtml)
			movieInfo.Movie_country=models.GetMovieCountry(sMovieHtml)
			movieInfo.Movie_type=models.GetMovieType(sMovieHtml)
			movieInfo.Movie_writer=models.GetMovieWriter(sMovieHtml)
			movieInfo.Movie_grade=models.GetMovieMark(sMovieHtml)
			movieInfo.Movie_length=models.GetMovieLength(sMovieHtml)
			movieInfo.Movie_introduction=models.GetMovieIntroduction(sMovieHtml)
			movieInfo.Movie_language=models.GetMovieLanguage(sMovieHtml)
			movieInfo.Movie_other_name=models.GetMovieOtherName(sMovieHtml)
			/*c.Ctx.WriteString(movieInfo.Movie_release_time+" ")
			c.Ctx.WriteString(movieInfo.Movie_director+" ")
			c.Ctx.WriteString(movieInfo.Movie_writer+" ")
			c.Ctx.WriteString(movieInfo.Movie_main_character+" ")
			c.Ctx.WriteString(movieInfo.Movie_type+" ")
			c.Ctx.WriteString(movieInfo.Movie_country+" ")
			c.Ctx.WriteString(movieInfo.Movie_on_time+" ")
			c.Ctx.WriteString(movieInfo.Movie_length+" ")
			c.Ctx.WriteString(movieInfo.Movie_grade+" ")
			c.Ctx.WriteString(movieInfo.Movie_introduction+" ")
			c.Ctx.WriteString(movieInfo.Movie_language+" ")
			c.Ctx.WriteString(movieInfo.Movie_other_name+" ")*/
			_,err:=models.AddMovie(&movieInfo)
			if err!=nil{
				panic(err)
			}else{
				c.Ctx.WriteString(sUrl+"爬取完成\n")
			}

		}
		//获取链接中的url并加入到队列中
		movies:=models.GetMovieUrls(sMovieHtml)
		//c.Ctx.WriteString(fmt.Sprintf("%v",movies))
		for _,url:=range movies{
			url=strings.Trim(url,`"`)

			models.PutinQueue(url)
			//fmt.Println(i,":",url)
		}
		//将访问过的sUrl记录
		models.AddToSet(sUrl)
		time.Sleep(1)
		//c.Ctx.WriteString(sMovieHtml)
	}
	c.Ctx.WriteString("end of crawling")


}


