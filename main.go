package main

import (
	_ "embed"
	"newsletter-builder/listmonk"
	"newsletter-builder/loadenv"
	"newsletter-builder/rss"
)

func main() {

	loadenv.LoadEnv()
	rss.Rss()
	listmonk.Listmonk()

}
