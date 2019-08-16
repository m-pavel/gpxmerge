package gpxmerge

import (
	"github.com/m-pavel/gpxgo/gpx"
	"log"
	"fmt"
	"crypto/md5"
	"io"
)

type GpxMerge struct {

}

func (gm GpxMerge) Merge(files []*gpx.GPX) (*gpx.GPX, error){
	res := gpx.GPX{}
	cfct := 0
	total := 0
	pmap := make(map[string]gpx.GPXPoint)
	for _,gpx := range files{
		gm.mergeHeader(gpx, &res)
		for _, wp := range gpx.Waypoints {
			pkey := key(&wp)

			if _, ok := pmap[pkey]; ok {
				cfct ++
			}
			pmap[pkey] = wp
			total ++
		}
	}
	for _, wp := range pmap {
		res.AppendWaypoint(&wp)
	}
	log.Printf("Total %d, Dupes %d", total, cfct)
	return &res, nil
}

func key(gp *gpx.GPXPoint) string {
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%s:%s", gp.Name))
	key :=  fmt.Sprintf("%x", h.Sum(nil))
	return key
}

func (gm GpxMerge) mergeHeader(gpx, res *gpx.GPX) {

	if res.XMLNs != "" && res.XMLNs != gpx.XMLNs {
		log.Printf("XMLNs conflict %s %s\n", res.XMLNs, gpx.XMLNs)
	} else {
		res.XMLNs = gpx.XMLNs
	}
	if res.XmlNsXsi != "" && res.XmlNsXsi != gpx.XmlNsXsi {
		log.Printf("XmlNsXsi conflict %s %s\n", res.XmlNsXsi, gpx.XmlNsXsi)
	} else {
		res.XmlNsXsi = gpx.XmlNsXsi
	}
	if res.XmlSchemaLoc != "" && res.XmlSchemaLoc != gpx.XmlSchemaLoc {
		log.Printf("XmlSchemaLoc conflict %s %s\n", res.XmlSchemaLoc, gpx.XmlSchemaLoc)
	} else {
		res.XmlSchemaLoc = gpx.XmlSchemaLoc
	}
	if res.Version != "" && res.Version != gpx.Version {
		log.Printf("Version conflict %s %s\n", res.Version, gpx.Version)
	} else {
		res.Version = gpx.Version
	}
	if res.Creator != "" && res.Creator != gpx.Creator {
		log.Printf("Creator conflict %s %s\n", res.Creator, gpx.Creator)
	} else {
		res.Creator = gpx.Creator
	}
	if res.Name != "" && res.Name != gpx.Name {
		log.Printf("Name conflict %s %s\n", res.Name, gpx.Name)
	} else {
		res.Name = gpx.Name
	}
	if res.Description != "" && res.Description != gpx.Description {
		log.Printf("Description conflict %s %s\n", res.Description, gpx.Description)
	} else {
		res.Description = gpx.Description
	}
	if res.AuthorName != "" && res.AuthorName != gpx.AuthorName {
		log.Printf("AuthorName conflict %s %s\n", res.AuthorName, gpx.AuthorName)
	} else {
		res.AuthorName = gpx.AuthorName
	}
	if res.AuthorEmail != "" && res.AuthorEmail != gpx.AuthorEmail {
		log.Printf("AuthorEmail conflict %s %s\n", res.AuthorEmail, gpx.AuthorEmail)
	} else {
		res.AuthorEmail = gpx.AuthorEmail
	}
	if res.AuthorLink != "" && res.AuthorLink != gpx.AuthorLink {
		log.Printf("AuthorLink conflict %s %s\n", res.AuthorLink, gpx.AuthorLink)
	} else {
		res.AuthorLink = gpx.AuthorLink
	}
	if res.AuthorLinkText != "" && res.AuthorLinkText != gpx.AuthorLinkText {
		log.Printf("AuthorLinkText conflict %s %s\n", res.AuthorLinkText, gpx.AuthorLinkText)
	} else {
		res.AuthorLinkText = gpx.AuthorLinkText
	}
	if res.AuthorLinkType != "" && res.AuthorLinkType != gpx.AuthorLinkType {
		log.Printf("AuthorLinkType conflict %s %s\n", res.AuthorLinkType, gpx.AuthorLinkType)
	} else {
		res.AuthorLinkType = gpx.AuthorLinkType
	}
	if res.Copyright != "" && res.Copyright != gpx.Copyright {
		log.Printf("Copyright conflict %s %s\n", res.Copyright, gpx.Copyright)
	} else {
		res.Copyright = gpx.Copyright
	}
	if res.CopyrightYear != "" && res.CopyrightYear != gpx.CopyrightYear {
		log.Printf("CopyrightYear conflict %s %s\n", res.CopyrightYear, gpx.CopyrightYear)
	} else {
		res.CopyrightYear = gpx.CopyrightYear
	}
	if res.CopyrightLicense != "" && res.CopyrightLicense != gpx.CopyrightLicense {
		log.Printf("CopyrightLicense conflict %s %s\n", res.CopyrightLicense, gpx.CopyrightLicense)
	} else {
		res.CopyrightLicense = gpx.CopyrightLicense
	}
	if res.Link != "" && res.Link != gpx.Link {
		log.Printf("Link conflict %s %s\n", res.Link, gpx.Link)
	} else {
		res.Link = gpx.Link
	}
	if res.LinkText != "" && res.LinkText != gpx.LinkText {
		log.Printf("LinkText conflict %s %s\n", res.LinkText, gpx.LinkText)
	} else {
		res.LinkText = gpx.LinkText
	}
	if res.LinkType != "" && res.LinkType != gpx.LinkType {
		log.Printf("LinkType conflict %s %s\n", res.LinkType, gpx.LinkType)
	} else {
		res.LinkType = gpx.LinkType
	}
	if res.Time != nil && res.Time != gpx.Time {
		log.Printf("Time conflict %s %s\n", res.Time, gpx.Time)
	} else {
		res.Time = gpx.Time
	}
	if res.Keywords != "" && res.Keywords != gpx.Keywords {
		log.Printf("Keywords conflict %s %s\n", res.Keywords, gpx.Keywords)
	} else {
		res.Keywords = gpx.Keywords
	}
}