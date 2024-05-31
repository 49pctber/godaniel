package godaniel

import (
	_ "embed"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var DefaultName string
var Port int

type TemplateData struct {
	Name         string
	Greeting     string
	Affirmations []string
	Now          time.Time
	Farewell     string
}

func init() {
	flag.StringVar(&DefaultName, "name", "Daniel", "the name of the person to affirm")
	flag.IntVar(&Port, "port", 8052, "port number to use for the server")
	flag.Parse()
}

//go:embed static/index.html
var templatestr string

func (td *TemplateData) UpdateData(name string) {
	td.Name = name
	loc, err := time.LoadLocation("America/Denver")
	if err != nil {
		panic(err)
	}
	td.Now = time.Now().In(loc)
	td.getGreeting()
	td.getAffirmations()
	td.getFarewell()
}

func (td *TemplateData) getGreeting() {
	if td.Now.Hour() < 5 {
		td.Greeting = fmt.Sprintf("😴 Bro, why are you awake? Go to bed, %s...", td.Name)
	} else if td.Now.Hour() < 12 {
		td.Greeting = fmt.Sprintf("🌞 Good morning, %s!", td.Name)
	} else if td.Now.Hour() < 18 {
		td.Greeting = fmt.Sprintf("👋 Good afternoon, %s!", td.Name)
	} else {
		td.Greeting = fmt.Sprintf("🌛 Good evening, %s!", td.Name)
	}
}

func (td *TemplateData) getFarewell() {
	td.Farewell = fmt.Sprintf("🎉 GO, %s! 🎉", strings.ToUpper(td.Name))
}

func GetTemplateData(name string) TemplateData {
	// get current time
	td := TemplateData{}
	td.UpdateData(name)
	return td
}

func PrintAffirmations(td TemplateData) {
	fmt.Printf("\n%s\n\n", td.Greeting)
	time.Sleep(1 * time.Second)
	for i := 0; i < len(td.Affirmations); i++ {
		fmt.Printf("  - %s\n", td.Affirmations[i])
		time.Sleep(3 * time.Second)
	}
	fmt.Printf("\n🎉 GO, %s! 🎉\n\n", strings.ToUpper(td.Name))
	time.Sleep(2 * time.Second)
}

func Handler(w http.ResponseWriter, req *http.Request) {

	var td TemplateData
	rname := req.URL.Query().Get("name")
	caser := cases.Title(language.English)
	rname = caser.String(rname)

	if len(rname) != 0 {
		td = GetTemplateData(rname)
	} else {
		td = GetTemplateData(DefaultName)
	}

	// render template data
	tmpl, err := template.New("template").Parse(templatestr)
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, td)
}
