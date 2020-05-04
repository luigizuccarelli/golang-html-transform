package schema

// SchemaInterface - acts as an interface wrapper
// All the go microservices will using this schema
type SchemaInterface struct {
	LastUpdate  int64      `json:"lastupdate,omitempty"`
	MetaInfo    string     `json:"metainfo,omitempty"`
	File        string     `json:"file,omitempty"`
	Html        string     `json:"html,omitempty"`
	ContentData HtmlSchema `json:"contentdata,omitempty"`
}

type HtmlSchema struct {
	Valued                string `json:"valued"`
	ValuedDescription     string `json:"valueddescription"`
	Title                 string `json:"title"`
	Description           string `json:"description"`
	ValuedUrl             string `json:"valuedurl"`
	ValuedBtn             string `json:"valuedbtn"`
	NextUrl               string `json:"nexturl"`
	NextBtn               string `json:"nextbtn"`
	Service               string `json:"service"`
	ServiceDescription    string `json:"servicedescription"`
	OptionA               string `json:"optionA"`
	OptionADescription    string `json:"optionADescription"`
	OptionB               string `json:"optionB"`
	OptionBDescription    string `json:"optionBDescription"`
	SubOption             string `json:"subOption"`
	SubOptionDescription  string `json:"subOptionDescription"`
	SubOptionADescription string `json:"subOptionADescription"`
	SubOptionAUrl         string `json:"subOptionAUrl"`
	SubOptionABtn         string `json:"subOptionABtn"`
	SubOptionBDescription string `json:"subOptionBDescription"`
	SubOptionBUrl         string `json:"subOptionBUrl"`
	SubOptionBBtn         string `json:"subOptionBBtn"`
	SubOptionCDescription string `json:"subOptionCDescription"`
	SubOptionCUrl         string `json:"subOptionCUrl"`
	SubOptionCBtn         string `json:"subOptionCBtn"`
	SubOptionDDescription string `json:"subOptionDDescription"`
	SubOptionDUrl         string `json:"subOptionDUrl"`
	SubOptionDBtn         string `json:"subOptionDBtn"`
	DataA                 string `json:"dataA"`
	DataATitle            string `json:"dataATitle"`
	DataB                 string `json:"dataB"`
	DataBTitle            string `json:"dataBTitle"`
	DataC                 string `json:"dataC"`
	DataCTitle            string `json:"dataCTitle"`
	Pricing               string `json:"pricing"`
	PricingDescription    string `json:"pricingDescription"`
	PlanA                 string `json:"planA"`
	PlanADescription      string `json:"planADescription"`
	PlanADetails          string `json:"planADetails"`
	PlanAUrl              string `json:"planAUrl"`
	PlanABtn              string `json:"planABtn"`
	AS                    string `json:"AS"`
	AP                    string `json:"AP"`
	AM                    string `json:"AM"`
	PlanB                 string `json:"planB"`
	PlanBDescription      string `json:"planBDescription"`
	PlanBDetails          string `json:"planBDetails"`
	PlanBUrl              string `json:"planBUrl"`
	PlanBBtn              string `json:"planBBtn"`
	BS                    string `json:"BS"`
	BP                    string `json:"BP"`
	BM                    string `json:"BM"`
	Address               string `json:"address"`
}

// Response schema
type Response struct {
	Name       string `json:"name"`
	StatusCode string `json:"statuscode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Payload    string `json:"payload"`
}
