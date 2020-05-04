package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	"gitea-cicd.apps.aws2-dev.ocp.14west.io/cicd/golang-html-transform/pkg/connectors"
	"gitea-cicd.apps.aws2-dev.ocp.14west.io/cicd/golang-html-transform/pkg/schema"
)

const (
	CONTENTTYPE     string = "Content-Type"
	APPLICATIONJSON string = "application/json"
)

func TransformHandler(w http.ResponseWriter, r *http.Request, con connectors.Clients) {
	var response *schema.Response
	var data *schema.SchemaInterface
	//var htmldata *schema.HtmlSchema

	addHeaders(w, r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response = &schema.Response{Name: os.Getenv("NAME"), StatusCode: "500", Status: "KO", Message: fmt.Sprintf("TransformHandler could not read body data %v\n", err), Payload: ""}
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		e := json.Unmarshal(body, &data)
		if e != nil {
			response = &schema.Response{Name: os.Getenv("NAME"), StatusCode: "500", Status: "KO", Message: fmt.Sprintf("TransformHandler could not unmarshal data %v\n", e), Payload: ""}
			w.WriteHeader(http.StatusInternalServerError)
			con.Error("TransformHandler %v\n", e)
		} else {
			con.Debug("TransformHandler data %v\n", data.ContentData)
			// perform io operations
			var bb bytes.Buffer
			tmpl, err := template.New("transform").Parse(string(data.Html))
			tmpl.Execute(&bb, data.ContentData)
			err = ioutil.WriteFile(data.File, bb.Bytes(), 0755)
			if err != nil {
				response = &schema.Response{Name: os.Getenv("NAME"), StatusCode: "500", Status: "KO", Message: fmt.Sprintf("TransformHandler could not writefile %v\n", err), Payload: ""}
				w.WriteHeader(http.StatusInternalServerError)
				con.Error("TransformHandler %v\n", err)
			} else {
				response = &schema.Response{Name: os.Getenv("NAME"), StatusCode: "200", Status: "OK", Message: "Data processed succesfully", Payload: ""}
				w.WriteHeader(http.StatusOK)
			}
		}

		b, _ := json.MarshalIndent(response, "", "	")
		con.Debug(fmt.Sprintf("TransformHandler response : %s", string(b)))
		fmt.Fprintf(w, string(b))
	}
}

func IsAlive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{ \"version\" : \"1.0.2\" , \"name\": \""+os.Getenv("NAME")+"\" }")
}

// headers (with cors) utility
func addHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(CONTENTTYPE, APPLICATIONJSON)
	// use this for cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
