package godaniel

import (
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var DefaultName string = "Daniel"

type TemplateData struct {
	Name         string
	Greeting     string
	Affirmations []string
	Now          time.Time
	Farewell     string
}

//go:embed static/*.html
var templates embed.FS

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
	caser := cases.Title(language.English)
	name = caser.String(name)
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

var re *regexp.Regexp = regexp.MustCompile(`[^a-zA-Z]+`)

func removeNonLetters(input string) string {
	return re.ReplaceAllString(input, "")
}

func GoDanielHandler(w http.ResponseWriter, req *http.Request) {

	var td TemplateData
	rname := removeNonLetters(req.URL.Query().Get("name"))

	if len(rname) != 0 {
		// render template for name
		td = GetTemplateData(rname)
		tmpl, err := template.ParseFS(templates, "static/base.html", "static/godaniel.html")
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, td)
	} else {
		// get name
		tmpl, err := template.ParseFS(templates, "static/base.html", "static/getname.html")
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, nil)
	}

}
