package access1

import (
	"fmt"
	"news-spider/src/models"
	"news-spider/src/utils"
)

func Save(movie *models.Movie) {
	result, err := db.Exec("INSERT INTO movie (id, `name`,director,staring,script_writer,`year`,movie_type,add_date ) values (?,?,?,?,?,?,?,?)", nil, movie.Name, movie.Director, movie.Staring, movie.Scriptwriter, movie.Year, movie.MovieType, movie.AddDate)
	utils.CheckErr(err)
	if nil != result {
		fmt.Println("保存成功")
		return
	}
	fmt.Println("保存失败")

}
