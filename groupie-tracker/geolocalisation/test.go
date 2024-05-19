package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"
)

var lien = "https://groupietrackers.herokuapp.com/api/"

var (
	templates  = template.Must(template.ParseFiles("test.html"))
	templates1 = template.Must(template.ParseFiles("geo.html"))
)

const port = ":5555"

type globale struct {
	IdAPI           int
	ImageAPI        string
	NameAPI         string
	MembersAPI      []string
	CreationDateAPI int
	FirstAlbumAPI   string
	LocationAPI     []string
	DateAPI         []string
	RelationAPI     map[string][]string
}

type Artiste struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Location     string
	Date         string
	Relation     string
}

type Location struct {
	Index []location2
}

/*var e l[]ocation2
var t []Artiste //lpas oublier de recuperer les get
//ensuite faire boucle declarer var
var y []globale
//mettre à chaque avancer boucle integrer valeur de a et t dedans y
*/

type location2 struct {
	Id        int
	Locations []string
}
type Date struct {
	Index []date2
}
type date2 struct {
	Id    int
	Dates []string
}

type Relation struct {
	Index []relation2
}

type relation2 struct {
	Id             int
	DatesLocations map[string][]string
}

var (
	artists      []Artiste
	ensemble     []globale
	locationsART Location
	datesART     Date
	relationsART Relation
)

func joindre(ensemble []globale) []globale {
	var globaleAPI globale
	var artist []Artiste
	x, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer x.Body.Close()

	body2, err := ioutil.ReadAll(x.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body2, &artist)
	if err != nil {
		fmt.Println(err)
	}
	// var locations Location
	locationsART = getLocation()
	datesART = getDate()
	relationsART = getRelation()

	var i int

	for i = range artist {
		globaleAPI.IdAPI = i
		globaleAPI.ImageAPI = artist[i].Image
		globaleAPI.NameAPI = artist[i].Name
		globaleAPI.MembersAPI = artist[i].Members
		globaleAPI.CreationDateAPI = artist[i].CreationDate
		globaleAPI.FirstAlbumAPI = artist[i].FirstAlbum
		globaleAPI.LocationAPI = locationsART.Index[i].Locations
		globaleAPI.DateAPI = datesART.Index[i].Dates
		globaleAPI.RelationAPI = relationsART.Index[i].DatesLocations
		ensemble = append(ensemble, globaleAPI)
		/*func GetArtistbyID(id int) (Api, error) {
			for _, elem := range artistsall {
				if elem.ID == id {
					return elem, nil
				}
			}
			return Api{}, errors.New("UNABLE TO RETRIEVE DATA FOR THIS ARTIST")
		}*/

	}
	return ensemble
}

func getArtist(artist []Artiste) []Artiste {
	x, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer x.Body.Close()

	body2, err := ioutil.ReadAll(x.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body2, &artist)
	if err != nil {
		fmt.Println(err)
	}

	return artist
}

func getLocation() Location {
	var lieu Location

	a, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer a.Body.Close()

	body2, err := ioutil.ReadAll(a.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body2, &lieu)
	if err != nil {
		fmt.Println(err)
	}

	return lieu
}

func getDate() Date {
	var date Date

	z, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer z.Body.Close()

	body2, err := ioutil.ReadAll(z.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body2, &date)
	if err != nil {
		fmt.Println(err)
	}

	return date
}

func getRelation() Relation {
	var rel Relation

	y, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer y.Body.Close()

	body2, err := ioutil.ReadAll(y.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body2, &rel)
	if err != nil {
		fmt.Println(err)
	}

	return rel
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		fmt.Fprintf(w, "Status 404: Page Not Found")
		return
	}

	if err := templates.ExecuteTemplate(w, "test.html", getArtist(artists)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func geo(w http.ResponseWriter, r *http.Request) {
	var n int
	if r.URL.Path != "/geo" {
		http.NotFound(w, r)
		fmt.Fprintf(w, "Status 404: Page Not Found")
		return
	}
	// var ensemble []globale
	n, _ = strconv.Atoi(r.URL.Query().Get("id"))
	if err := templates1.ExecuteTemplate(w, "geo.html", joindre(ensemble)[n-1]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleProxy(w http.ResponseWriter, r *http.Request) {
	remote := r.URL.Query().Get("url")
	if remote == "" {
		http.NotFound(w, r)
		return
	}

	response, _ := http.Get(remote)
	if response.StatusCode != 200 {

		http.NotFound(w, r)
		return
	}
	io.Copy(w, response.Body)
}

func main() {
	// fmt.Print(getRelation())

	// joindre()
	// fmt.Print(joindre(ensemble)[0])
	http.HandleFunc("/", home)
	http.HandleFunc("/geo", geo)
	http.HandleFunc("/proxy", HandleProxy)
	fmt.Println("http://localhost:5555")
	// http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("templates/js/"))))
	http.ListenAndServe(port, nil)
}

/*func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}*/
/*{{range $key, $valeur := .}}
{{$key}}
{{$valeur}}
{{end}}*/
