package main

import (
	"os"
	"log"
	"github.com/m-pavel/gpxgo/gpx"
	"github.com/m-pavel/gpxmerge/pkg"
	"io/ioutil"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("gpxmerge IN... OUT")
		return
	}
	var err error
	gfiles := make([]*gpx.GPX, len(os.Args) -2)
	for i:=1 ; i<len(os.Args) - 1; i++  {
		gfiles[i-1], err = gpx.ParseFile(os.Args[i])
		if err != nil {
			log.Panic(err)
		}
	}

	merge := gpxmerge.GpxMerge{}
	res, err := merge.Merge(gfiles)
	if err != nil {
		log.Panic(err)
	}
	xmlb, err := gpx.ToXml(res, gpx.ToXmlParams{Version:"1.0", Indent:true})
	if err != nil {
		log.Panic(err)
	}
	err = ioutil.WriteFile(os.Args[len(os.Args) - 1], xmlb, 0644)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("%s is written. %d waypoints merged\n", os.Args[len(os.Args) - 1], len(res.Waypoints))
}
