package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Artiklar struct {
	XMLName   xml.Name `xml:"artiklar"`
	Text      string   `xml:",chardata"`
	Xsd       string   `xml:"xsd,attr"`
	Xsi       string   `xml:"xsi,attr"`
	SkapadTid string   `xml:"skapad-tid"`
	Info      struct {
		Text       string `xml:",chardata"`
		Meddelande string `xml:"meddelande"`
	} `xml:"info"`
	Artikel []Artikel `xml:"artikel"`
}

type Artikel struct {
	Text               string `xml:",chardata" json:"-"`
	Nr                 int    `xml:"nr" json:"N"`
	Artikelid          int    `xml:"Artikelid" json:"A"`
	Varnummer          int    `xml:"Varnummer" json:"V"`
	Namn               string `xml:"Namn" json:"-"`
	Namn2              string `xml:"Namn2" json:"-"`
	Prisinklmoms       string `xml:"Prisinklmoms" json:"-"`
	Volymiml           string `xml:"Volymiml" json:"-"`
	PrisPerLiter       string `xml:"PrisPerLiter" json:"-"`
	Saljstart          string `xml:"Saljstart" json:"-"`
	Utgått             string `xml:"Utgått" json:"-"`
	Varugrupp          string `xml:"Varugrupp" json:"-"`
	Typ                string `xml:"Typ" json:"-"`
	Stil               string `xml:"Stil" json:"-"`
	Forpackning        string `xml:"Forpackning" json:"-"`
	Forslutning        string `xml:"Forslutning" json:"-"`
	Ursprung           string `xml:"Ursprung" json:"-"`
	Ursprunglandnamn   string `xml:"Ursprunglandnamn" json:"-"`
	Producent          string `xml:"Producent" json:"-"`
	Leverantor         string `xml:"Leverantor" json:"-"`
	Argang             string `xml:"Argang" json:"-"`
	Provadargang       string `xml:"Provadargang" json:"-"`
	Alkoholhalt        string `xml:"Alkoholhalt" json:"-"`
	Sortiment          string `xml:"Sortiment" json:"-"`
	SortimentText      string `xml:"SortimentText" json:"-"`
	Ekologisk          string `xml:"Ekologisk" json:"-"`
	Etiskt             string `xml:"Etiskt" json:"-"`
	Koscher            string `xml:"Koscher" json:"-"`
	RavarorBeskrivning string `xml:"RavarorBeskrivning" json:"-"`
}

func main() {

	resp, err := http.Get("https://www.systembolaget.se/api/assortment/products/xml")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	artMap := make(map[int]Artikel)
	am2 := make(map[int]Artikel)

	atk := Artiklar{}

	fmt.Println("Downloading XML file from Systembolaget")

	dec := xml.NewDecoder(resp.Body)

	dec.Decode(&atk)

	for _, v := range atk.Artikel {
		artMap[v.Varnummer] = v
	}

	fmt.Println(artMap[11392])

	file, _ := json.Marshal(artMap)

	_ = ioutil.WriteFile("bolaget-index.json", file, 0644)

	f, _ := os.Open("bolaget-index.json")

	me := json.NewDecoder(f)
	me.Decode(&am2)

	fmt.Println(artMap[11392])

}
