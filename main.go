package main

import (
	"fmt"
	"news-spider/src/service"
	"strconv"
)

func search() {
	url := "https://movie.douban.com/top250?start="
	var urls []string
	var newUrl string
	fmt.Println("%v", urls)
	for i := 0; i < 10; i++ {
		start := i * 25
		newUrl = url + strconv.Itoa(start)
		urls = service.GetToplist(newUrl)

		if len(urls) <= 0 {
			return
		}
		//movies := models.Movies{}
		for _, url := range urls {
			service.SearchMovie(url)
			//movies.Items = append(movies.Items, &movie)
		}

		//for _, item := range movies.Items {
		//	service.SaveMovie(item)
		//}

	}
}
func main() {
	//for i := 0; i <= 10; i++ {
	search()
	//}
}
