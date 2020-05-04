package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"strings"
	"text/template"
)

type HtmlSchema struct {
	Valued            string `json:"valued"`
	ValuedDescription string `json:"valueddescription"`
}

func main() {

	var sch *HtmlSchema
	html, _ := ioutil.ReadFile("../project-flow-master/public/14/template.html")
	err := json.Unmarshal([]byte("{ \"valued\": \"Most Valued Customer\" , \"valueddescription\": \"We really value you. We are offering a <b>20%</b> disount with immediate effect on your current subscriptions click on the link to avail of this offer now\" }"), &sch)
	fmt.Println(fmt.Sprintf("ERROR %v\n", err))

	var data bytes.Buffer
	tmpl, err := template.New("test").Parse(string(html))
	fmt.Println(fmt.Sprintf("DEBUG %v\n", sch))
	tmpl.Execute(&data, sch)

	/*
		"@{model.valued}", "Most Valued Customer",
		"@{model.valuedDescription}", "We really value you. We are offering a <b>20%</b> disount with immediate effect on your current subscriptions click on the link to avail of this offer now",
		"@{model.valuedUrl}", "#",
		"@{model.valuedBtn}", "Avail Now",
		"@{model.title}", "The Great Stuff",
		"@{model.nextUrl}", "www.google.com",
		"@{model.nextBtn}", "Read More",
		"@{model.description}", "We are a group of companies that create wealth portfolios and investment suggestions for the beginner to the advanced investor",
		"@{model.service}", "Wealth Management",
		"@{model.serviceDescription}", "We offer consultancy, learning programs and e-books",
		"@{model.optionA}", "Portoflio's",
		"@{model.optionaADescription}", "We explain the how why and what",
		"@{model.optionB}", "Investments",
		"@{model.optionBDescription}", "We use videos, e-books and tutorials to let you invest with confidence",
		"@{model.subOption}", "Company Profile",
		"@{model.subOptionDescription}", "We have created over $2bn in wealth for all our clients",
		"@{model.subOptionADescription}", "Experts",
		"@{model.subOptionAUrl}", "www.google.com",
		"@{model.subOptionABtn}", "Read More",
		"@{model.subOptionBDescription}", "Front Office",
		"@{model.subOptionBUrl}", "www.google.com",
		"@{model.subOptionBBtn}", "Read More",
		"@{model.subOptionCDescription}", "Back Office",
		"@{model.subOptionCUrl}", "www.google.com",
		"@{model.subOptionCBtn}", "Read More",
		"@{model.subOptionDDescription}", "Advisary",
		"@{model.subOptionDUrl}", "www.google.com",
		"@{model.subOptionDBtn}", "Read More",
		"@{model.dataA}", "5000",
		"@{model.dataATitle}", "Clients",
		"@{model.dataB}", "250",
		"@{model.dataBTitle}", "Portfolio's",
		"@{model.dataC}", "$2BN",
		"@{model.dataCTitle}", "Profit",
		"@{model.pricing}", "Pricing",
		"@{model.pricingDescription}", "Our pricing offerings",
		"@{model.planA}", "Standard",
		"@{model.planADescription}", "The standard package",
		"@{model.planADetails}", "<li>Access to experts</li><li>Monthly e-book</li><li>&nbsp;</li><li>&nbsp;</li>",
		"@{model.planAUrl}", "www.google.com",
		"@{model.planABtn}", "Buy Now",
		"@{model.AS}", "$",
		"@{model.AP}", "35",
		"@{model.AM}", "/month",
		"@{model.planB}", "Premium",
		"@{model.planBDescription}", "The premium package",
		"@{model.planBDetails}", "<li>Access to experts</li><li>Monthly e-book</li><li>Video's</li><li>Online tutorials</li>",
		"@{model.planBUrl}", "www.google.com",
		"@{model.planBBtn}", "Buy Now",
		"@{model.BS}", "$",
		"@{model.BP}", "125",
		"@{model.BM}", "/month",
		"@{model.address}", "Confederation House Block C&D Waterford, Ireland",
	*/

	//formatedHtml := url.Replace(string(html))
	err = ioutil.WriteFile("../project-flow-master/public/14/rendered.html", data.Bytes(), 0755)
	if err != nil {
		fmt.Println("Error writing file")
	}

}
