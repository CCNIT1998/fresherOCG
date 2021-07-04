package main

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = ""
	dbname   = "ocg"
)

type movie struct {
	Title  string
	Rating string
	Url    string
}

func getMovies(link string) (Movies []movie) {
	c := colly.NewCollector()
	c.OnHTML(".lister-list tr", func(e *colly.HTMLElement) {

		txtTitle := e.ChildText(".titleColumn")
		Title := regexp.MustCompile(`\n\s+`).ReplaceAllString(txtTitle, " ")
		Title = regexp.MustCompile(`\d+.\s+`).ReplaceAllString(Title, "")

		Rating := e.ChildText(".imdbRating")

		txtUrl := e.ChildAttr("a", "href")
		Url := "https://www.imdb.com" + strings.Split(txtUrl, "?")[0]

		Movie := movie{
			Title:  Title,
			Rating: Rating,
			Url:    Url,
		}
		Movies = append(Movies, Movie)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.Visit(link)
	return Movies
}

func insertdb(user, password, dbname string, Movies []movie) {
	db, err := sql.Open("mysql", user+":"+password+"@/"+dbname)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Movies(id int NOT NULL AUTO_INCREMENT PRIMARY KEY, Title varchar(100), Rating varchar(4),Url varchar(50));")
	if err != nil {
		fmt.Println("Cannot create table: ", err)
		return
	}
	for _, k := range Movies {
		fmt.Println(k.Title)
		_, err = db.Exec("INSERT INTO Movies (Title, Rating, Url) VALUES (?, ?, ?)", k.Title, k.Rating, k.Url)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	link := "https://www.imdb.com/chart/top/?ref_=nv_mv_250"
	Movies := getMovies(link)
	insertdb(user, password, dbname, Movies)
}
