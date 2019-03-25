package main

import (
	"log"
	"net/http"
	"os"

	"github.com/andrysds/panera/connection"
	"github.com/andrysds/panera/entity"
	"github.com/andrysds/panera/route"
	"github.com/andrysds/panera/template"
	"github.com/subosito/gotenv"
)

func main() {
	log.Println("Panera starting...")

	gotenv.Load()
	connection.Init()
	entity.InitCollection()
	template.Init()

	log.Println("* Listening")
	log.Println("Use Ctrl-C to stop")
	http.ListenAndServe(":"+os.Getenv("PORT"), route.Router())
}
