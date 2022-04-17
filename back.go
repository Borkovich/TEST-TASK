package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type user struct {
	Name   string
	weight int
	color  string
}

func saveCsv(name string, weight int, color string, id string) {

	_, er := os.Stat("valid.csv")
	if er != nil {
		if os.IsNotExist(er) {
			csvfile, _ := os.OpenFile("valid.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
			csvfile.WriteString("id,name,weight,color\n")
			defer csvfile.Close()
		}
	}
	var data = []string{id, name, strconv.Itoa(weight), color}
	csvfile, err := os.OpenFile("valid.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error", err)
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)
	defer writer.Flush()
	writer.Write(data)

}

func validation(name string, weight int) (string, error) {

	if weight < 1 || weight > 500 {
		return "Валидация не пройдена", errors.New("weight is too small or high")
	} else if len(name) < 3 || len(name) > 120 {
		return "Валидация не пройдена", errors.New("name too long or short")
	} else {
		return "Валидация пройдена", nil
	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("index.html"))
	var mess = map[string]string{"mess": ""}
	r.ParseForm()
	fmt.Println(r.Form)
	weight, _ := strconv.Atoi(r.FormValue("height"))
	fmt.Println(r.FormValue("color"))
	if r.Method != http.MethodPost {
		tpl.ExecuteTemplate(w, "messages", mess)
		return
	}
	details := user{
		Name:   r.FormValue("title"),
		weight: weight,
		color:  r.FormValue("color"),
	}
	uuidWithHyphen := uuid.New()
	id := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	_ = details
	_, err := validation(details.Name, weight)
	mess = map[string]string{"mess": "Добавлена запись ID=" + id}
	if err != nil {
		mess = map[string]string{"mess": err.Error()}
	} else {
		saveCsv(details.Name, details.weight, details.color, id)
	}
	tpl.ExecuteTemplate(w, "messages", mess)

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe("0.0.0.0:80", nil)

}
