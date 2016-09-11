package main

import (
	"log"

	"github.com/Mehedidec/mh-torrent/server"
	"github.com/Mehedidec/opts"
)

var VERSION = "0.0.0-src" //set with ldflags

func main() {
	s := server.Server{
		Title:      "MH Torrent",
		Port:       3000,
		ConfigPath: "mh-torrent.json",
	}

	o := opts.New(&s)
	o.Version(VERSION)
	o.PkgRepo()
	o.LineWidth = 96
	o.Parse()

	if err := s.Run(VERSION); err != nil {
		log.Fatal(err)
	}
}
