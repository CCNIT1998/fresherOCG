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

type product struct {
	Name  string
	Price string
	Url   string
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

func getProducts(url string) (Products []product) {
	c := colly.NewCollector()
	c.OnHTML("body", func(e *colly.HTMLElement) {
		rootlink := "https://template-homedecor.onshopbase.com/collections/new-arrivals/products/"
		reName := regexp.MustCompile(`"title":"(.*?)"`)
		rePrice := regexp.MustCompile(`"price":(.*?)\d+`)
		reUrl := regexp.MustCompile(`"handle":"(.*?)"`)

		txtdata := regexp.MustCompile(`{"id":10000000(.*?)"price":\d+},*`)
		data := txtdata.FindAllString(e.Text, -1)
		for _, i := range data {
			txtName := strings.Replace(reName.FindString(i), "\"title\":\"", "", -1)
			Price := strings.Replace(rePrice.FindString(i), "\"price\":", "", -1)
			txtUrl := strings.Replace(reUrl.FindString(i), "\"handle\":\"", "", -1)
			Name := strings.Replace(txtName, "\"", "", -1)
			Url := strings.Replace(txtUrl, "\"", "", -1)
			Product := product{
				Name:  Name,
				Price: Price,
				Url:   rootlink + Url,
			}
			Products = append(Products, Product)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.Visit(url)
	return Products
}

func insertdb(user, password, dbname string, Movies []movie, Products []product) {
	db, err := sql.Open("mysql", user+":"+password+"@/"+dbname)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Movies(id int NOT NULL AUTO_INCREMENT PRIMARY KEY, Title varchar(100), Rating varchar(4),Url varchar(255));")
	if err != nil {
		fmt.Println("Cannot create table: ", err)
		return
	}

	for _, v := range Movies {
		_, err = db.Exec("INSERT INTO Movies (Title, Rating, Url) VALUES (?, ?, ?)", v.Title, v.Rating, v.Url)
		if err != nil {
			panic(err)
		}
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Prodcuts(id int NOT NULL AUTO_INCREMENT PRIMARY KEY, Name varchar(100), Price varchar(4),Url varchar(255));")
	if err != nil {
		fmt.Println("Cannot create table: ", err)
		return
	}

	for _, product := range Products {
		_, err = db.Exec("INSERT INTO Prodcuts (Name, Price, Url) VALUES (?, ?, ?)", product.Name, product.Price, product.Url)
		if err != nil {
			panic(err)
		}
	}
}



func main() {
	link := "https://www.imdb.com/chart/top/?ref_=nv_mv_250"
	Movies := getMovies(link)
	fmt.Println(Movies)

	url := "https://template-homedecor.onshopbase.com/collections/new-arrivals?sortby=name%3Aasc"
	Products := getProducts(url)
	fmt.Println(Products)
	insertdb(user, password, dbname, Movies, Products)
}
