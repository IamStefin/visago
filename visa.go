package visago

import (
	"fmt"
	"github.com/anaskhan96/soup"
)

func VisaFetcher(origin string,destination string)(Visa){
	var Newurl string
	resp, err := soup.Get("https://en.wikipedia.org/w/index.php?sort=relevance&search=Visa+requirements+for+"+origin+"+citizens&ns0=1")
	if err != nil {
		fmt.Println("Not Connected!")
	}else{
		doc := soup.HTMLParse(resp)
		links := doc.FindAll("div","class","mw-search-result-heading")
		for i, link := range links {
			if i == 0{
				Newurl = link.Find("a").Attrs()["href"]
			}
		}
		resp, err = soup.Get("https://en.wikipedia.org/"+Newurl)
		doc = soup.HTMLParse(resp)
		links = doc.Find("tbody").FindAll("tr")
		for i, n := range links {
			if i > 0 {
				if (destination == n.Find("td").Find("a").Text()){
					return Visa{destination,origin,n.FindAll("td")[1].Text()}
				}
			}
		}
	}
	return Visa{destination,origin,"s"}
}
