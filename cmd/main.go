package main

import (
	"log"

	"fbnoi.com/render"
)

func main() {
	if t, err := render.Parse("./var/template/base.html.tpl"); err != nil {
		log.Fatalln(err)
		log.Fatalln("Error")
	} else {
		log.Fatalln(t.ExecuteTemplate(log.Writer(), "base.html.tpl", nil))
		log.Println("123123123")
	}
	// if t, err := template.New("test").ParseFiles("./var/template/base.html.tpl"); err != nil {
	// 	log.Fatalln(err)
	// } else {
	// 	log.Println(t.ExecuteTemplate(log.Writer(), "base.html.tpl", nil))
	// }
}
