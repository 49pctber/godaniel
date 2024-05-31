package main

import (
	"github.com/49pctber/godaniel"
)

func main() {
	td := godaniel.GetTemplateData(godaniel.DefaultName)
	godaniel.PrintAffirmations(td)
}
