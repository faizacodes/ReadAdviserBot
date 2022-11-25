package scraping

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"os"
)

type Items struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func ScrapingFunc() {
	c := colly.NewCollector(
		colly.AllowedDomains("j2store.net"),
	)

	var items []Items
	c.OnHTML("div.col-sm-9 div[itemprop=itemListElement]", func(h *colly.HTMLElement) {
		item := Items{
			Name:   h.ChildText("h2.product-title"),
			Price:  h.ChildText("div.sale-price"),
			ImgUrl: h.ChildAttr("img", "src"),
		}

		items = append(items, item)
	})

	c.OnHTML("[title=Next]", func(h *colly.HTMLElement) {
		next_page := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(next_page)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())

	})

	c.Visit("https://j2store.net/v3/index.php/shop")

	content, err := json.Marshal(items)
	if err != nil {
		fmt.Println(err.Error())
	}
	ioutil.WriteFile("products.json", content, 0664)
}

func GetJsonData() []Items {
	file, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	var items []Items
	err = json.NewDecoder(file).Decode(&items)
	if err != nil {
		fmt.Println(err.Error())
	}
	return items
}
