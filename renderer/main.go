package main

import (
	"log"
	"net/http"

	"github.com/devetek/go-the-next-renderer/installer"
	"github.com/devetek/go-the-next-renderer/router"
	"github.com/gorilla/mux"
)

func main() {
	// Template Generator
	installer.New()
	// Template Generator
	r := mux.NewRouter()

	r.HandleFunc("/favicon.ico", router.AssetsHandler)
	r.HandleFunc("/thirteen.svg", router.AssetsHandler)
	r.HandleFunc("/next.svg", router.AssetsHandler)
	r.HandleFunc("/vercel.svg", router.AssetsHandler)
	r.HandleFunc("/_next/{page1}", router.AssetsHandler)
	r.HandleFunc("/_next/{page1}/{page2}", router.AssetsHandler)
	r.HandleFunc("/_next/{page1}/{page2}/{page3}", router.AssetsHandler)
	r.HandleFunc("/_next/{page1}/{page2}/{page3}/{page4}", router.AssetsHandler)
	r.HandleFunc("/_next/{page1}/{page2}/{page3}/{page4}/{page5}", router.AssetsHandler)
	r.HandleFunc("/example", router.ExampleHandler)
	r.HandleFunc("/loop", router.LoopHandler)
	r.HandleFunc("/", router.HomeHandler)

	http.Handle("/", r)

	log.Println("Service run in port 9000")
	http.ListenAndServe(":9000", r)
}
