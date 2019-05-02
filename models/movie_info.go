package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
)
var db orm.Ormer


type MovieInfo struct {
	Id         int `orm:"column(id);pk"`
	Movie_id   int
	Movie_name string
	Movie_director string
	Movie_writer string
	Movie_main_character string
	Movie_length string
	Movie_type string
	Movie_country string
	Movie_language string
	Movie_on_time string
	Movie_span string
    Movie_grade string
	Movie_other_name string
	Movie_introduction string
	Movie_release_time string
}
func init(){
	orm.Debug=true
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/testgo?charset=utf8",30)
	orm.RegisterModel(new(MovieInfo))
	db=orm.NewOrm()
}
func AddMovie(movie_info *MovieInfo) (int64,error){
	return db.Insert(movie_info)
}
func DeleteMovie(movie_info *MovieInfo) (int64,error){
	return db.Delete(movie_info)
}
func UpdateMovie(movie_info *MovieInfo) (int64,error){
	return db.Update(movie_info)
}
func ReadAllMovie() (int64,error,[]MovieInfo){
	movieInfos:=[]MovieInfo{}
	qb,_:=orm.NewQueryBuilder("mysql")
	qb.Select("*").From("user_info")
	sql:=qb.String()
	num,err:= db.Raw(sql).QueryRows(&movieInfos)
	return num,err,movieInfos
}
func GetMovieDirector(movieHtml string) string{
	if movieHtml==""{
		return ""
	}else{
		reg:=regexp.MustCompile(`<a.*?rel="v:directedBy">(.*?)</a>`)
		result:=reg.FindAllStringSubmatch(movieHtml,-1)
		if len(result)==0{
			return ""
		}
		director:=""
		for i,_:=range  result{
			director+=result[i][1]+" "
		}

		return director
	}
}
func GetMovieName(movieHtml string) string{
	if movieHtml==""{
		return ""
	}else{
		reg:=regexp.MustCompile(`<span.*property="v:itemreviewed">(.*)</span>`)
		result:=reg.FindAllStringSubmatch(movieHtml,-1)
		if len(result)==0{
			return ""
		}
		return string(result[0][1])
	}
}
func GetMainCharacter(movieHtml string) string{
	if movieHtml==""{
		return ""
	}else{
		reg:=regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
		result:=reg.FindAllStringSubmatch(movieHtml,-1)
		if len(result)==0{
			return ""
		}
		///fmt.Println(result)
		mainChacter:=""
		for i,_:=range  result{
			mainChacter+=result[i][1]+" "
		}
		//fmt.Println(reflect.TypeOf(result))
		return mainChacter
	}
}
func GetMovieType(movieHtml string) string{
	if movieHtml==""{
		return ""
	}else{
		reg:=regexp.MustCompile(`<span.*?property="v:genre">(.*?)</span>`)
		result:=reg.FindAllStringSubmatch(movieHtml,-1)
		if len(result)==0{
			return ""
		}
		movieType:=""
		for i,_:=range result{
			movieType+=result[i][1]+" "
		}
		return movieType

	}
}
func GetMovieCountry(movieHtml string) string{
	if movieHtml==""{
		return ""
	}else{
		reg:=regexp.MustCompile(`<span.*class="pl">制片国家/地区:</span>.*?(.*)<br/>`)
		result:=reg.FindAllStringSubmatch(movieHtml,-1)
		if len(result)==0{
			return ""
		}
		movieCountry:=""
		for i,_:=range result{
			movieCountry+=result[i][1]+" "
		}
		//fmt.Println(result)
		return movieCountry
	}
}
func GetOnTime(movieHtml string) string{
	if movieHtml==""{
		return ""
	}else{
		reg:=regexp.MustCompile(`<span.*?property="v:initialReleaseDate".*?>(.*?)</span>`)
		result:=reg.FindAllStringSubmatch(movieHtml,-1)
		onTime:=""
		for i,_:=range result{
			onTime+=result[i][1]+" "
		}
		return onTime
	}
}
func GetMovieLength(movieHtml string) string{
	if movieHtml==""{
		return ""
	}else{
		reg:=regexp.MustCompile(`<span.*?property="v:runtime".*?>(.*?)</span>`)
		result:=reg.FindAllStringSubmatch(movieHtml,-1)
		if len(result)==0{
			return ""
		}
		length:=""
		for i,_:=range result{
			length+=result[i][1]+" "
		}
		return length
	}
}
func GetMovieMark(movieHtml string) string{
	if movieHtml==""{
		return ""
	}else{
		reg:=regexp.MustCompile(`<strong.*?class="ll rating_num" property="v:average">(.*)</strong>`)
		result:=reg.FindAllStringSubmatch(movieHtml,-1)
		if len(result)==0{
			return ""
		}
		return result[0][1]

	}
}
func GetMovieReleaseTime(movieHtml string) string{
	if movieHtml==""{
		return ""
	}else{
		reg:=regexp.MustCompile(`<span.*class="year">(.*?)</span>`)
		result:=reg.FindAllStringSubmatch(movieHtml,-1)
		if len(result)==0{
			return ""
		}
		releaseTime:=""
		for i,_:=range result{
			releaseTime+=result[i][1]+" "
		}
		return releaseTime
	}
}

func GetMovieWriter(movieHtml string) string{
	if movieHtml==""{
		return ""
	}else{
		reg:=regexp.MustCompile(`<a.*?href="/celebrity/.*?/">(.*?)</a>`)
		result:=reg.FindAllStringSubmatch(movieHtml,-1)
		/*fmt.Println(result)
		fmt.Println("编剧")*/
		writer:=""
		for i,_:=range result{
			writer+=result[i][1]+" "
		}
		return writer
	}
}
func GetMovieIntroduction(movieHtml string) string{
	if movieHtml==""{
		return ""
	}else{
		//reg:=regexp.MustCompile(`<span.*property="v:summary">\s*(.*)`)
		reg:=regexp.MustCompile(`(?s)<span.*property="v:summary">(.*?)</span>`)
		result:=reg.FindAllStringSubmatch(movieHtml,-1)
		if len(result)==0{
			return ""
		}
		fmt.Println("简介")
		introduction:=""
		for i,_:=range result{
			introduction+=result[i][1]+" "
		}
		return introduction
	}
}
func GetMovieUrls(movieHtml string) []string{
	var movies []string
	if movieHtml==""{
		return movies
	}else{
		//reg:=regexp.MustCompile(`"https://movie.douban.com/subject/(.*?)from=subject-page"`)
		//匹配所有豆瓣链接的url
		reg:=regexp.MustCompile(`"https://movie.douban.com/(.*?)/"`)
		result:=reg.FindAllString(movieHtml,-1)
		/*fmt.Println(result)
		fmt.Println(reflect.TypeOf(result))*/
		//fmt.Println(len(result))
		if len(result)==0{
			return result
		}

		for i,_:=range result{
			fmt.Println(i,":",result[i])

		}
		return result


	}
}
func GetMovieLanguage(movieHtml string) string{
	if movieHtml==""{
		return ""
	}else{
		reg:=regexp.MustCompile(`<span.*class="pl">语言:</span>(.*?)<br/>`)
		result:=reg.FindAllStringSubmatch(movieHtml,-1)
		if len(result)==0{
			return ""
		}
		/*fmt.Println(result)
		fmt.Println("语言")*/
		language:=""
		for i,_:=range result{
			language+=result[i][1]+" "
		}
		return language
	}
}
func GetMovieOtherName(movieHtml string)string{
	if movieHtml==""{
		return ""
	}else{
		reg:=regexp.MustCompile(`<span.*class="pl">又名:</span>(.*?)<br/>`)
		result:=reg.FindAllStringSubmatch(movieHtml,-1)
		if len(result)==0{
			return ""
		}
		/*fmt.Println(result)
		fmt.Println("又名")*/
		otherName:=""
		for i,_:=range result{
			otherName+=result[i][1]+" "
		}
		return otherName
	}
}