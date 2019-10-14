package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	access1 "news-spider/src/access"
	"news-spider/src/exception"
	"news-spider/src/models"
	"time"
)

func SearchMovie(url string) models.Movie {
	movie := getMovie(url)
	SaveMovie(&movie)
	return movie
}

func SaveMovie(movie *models.Movie) {
	access1.Save(movie)
}

func getMovie(url string) models.Movie {
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//bodyString, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bodyString))

	if resp.StatusCode != 200 {
		var baseException = exception.BaseException{
			Code:        10000,
			Description: "请求响应错误",
		}
		baseException.Error()
		return models.Movie{}
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	//
	movie := models.Movie{}
	doc.Find("#content h1").Each(func(i int, s *goquery.Selection) {
		// name
		name := s.ChildrenFiltered(`[property="v:itemreviewed"]`).Text()
		fmt.Println("name:" + name)
		// year
		year := s.ChildrenFiltered(`.year`).Text()
		fmt.Println("year:" + year)
		movie.Name = name
		movie.Year = year
	})

	// #info > span:nth-child(1) > span.attrs
	doc.Find("#info span:nth-child(1) span.attrs").Each(func(i int, s *goquery.Selection) {
		// 导演
		director := s.Text()
		//fmt.Println(s.Text())
		movie.Director = director
		fmt.Println("导演:" + director)
	})
	//fmt.Println("\n")

	doc.Find("#info span:nth-child(3) span.attrs").Each(func(i int, s *goquery.Selection) {
		pl := s.Text()
		movie.Scriptwriter = pl
		fmt.Println("编剧:" + pl)
	})

	doc.Find("#info span.actor span.attrs").Each(func(i int, s *goquery.Selection) {
		staring := s.Text()
		fmt.Println("主演:" + staring)
		movie.Staring = staring

	})

	doc.Find("#info > span:nth-child(8)").Each(func(i int, s *goquery.Selection) {
		typeStr := s.Text()
		fmt.Println("类型:" + typeStr)
		movie.MovieType = typeStr
	})
	//movie.ID = int32(rand.Float32())
	now := time.Now()
	movie.AddDate = now.Format("2006-01-02 15:04:00")
	return movie

}

func GetToplist(url string) []string {
	var urls []string
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//bodyString, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bodyString))
	if resp.StatusCode != 200 {
		fmt.Println("err")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	doc.Find("#content div div.article ol li div div.info div.hd a").Each(func(i int, s *goquery.Selection) {
		// year
		fmt.Printf("%v", s)
		herf, _ := s.Attr("href")
		urls = append(urls, herf)
	})
	return urls
}
