package main

//advanced features, faire une page pour chaque utilisateur qui restitue:

//reste a tt checker

//authentification, ne pas oublier de se connecter des comptes google et github si on veut un nouvel user

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"text/template"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

type LikeData struct {
	Pseudo         string `json:"pseudo"`
	DejaLikePourGO bool   `json:"dejaLikePourGO"`
	Id             string `json:"idCom"`
	IdPost         string `json:"idPost"`
}

type UserInfo struct {
	Pseudo   string
	ID       string `json:"id"`
	Email    string `json:"email"`
	Verified bool   `json:"verified_email"`
	// Ajoutez d'autres champs que vous souhaitez extraire ici
}
type Notif struct {
	IDNotif  int    `json:"idNotif"`
	IDPost   string `json:"idPost"`
	IDCom    string `json:"idCom"`
	Pseudo   string `json:"pseudo"`
	Read     string `json:"read"`
	Createur string `json:"createur"`
	Nature   string `json:"nature"`
}

var (
	oauth2Config = oauth2.Config{
		ClientID:     "975083639506-1ai6q4gilsmr883tja8hgvu1du2novi5.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:5555/logine",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	githubOauthConfig = oauth2.Config{
		ClientID:     "3a26fa1b58be5d05c8f1",
		ClientSecret: "5885eead710f4b8af7a528c4ececd9cd2334074e",
		RedirectURL:  "http://localhost:5555/loginGithub",     // Mettez ici l'URL de rappel appropriée pour GitHub
		Scopes:       []string{"user:email", "user:password"}, // Vous pouvez spécifier les scopes nécessaires
		Endpoint:     github.Endpoint,
	}
)

// var templateHome = template.Must(template.ParseFiles("home.html"))
var tokens *oauth2.Token

var (
	countBoutLike    int
	countBoutDisLike int
)

var (
	utilisateur  string
	idComment    string
	idCommentCom string
)

var (
	dejaLike    bool
	dejaLikeCom bool
	date        string
	dateAct     string
	connecte    bool
	roleGlobal  string
)

type erreurs struct {
	MsgErreur string
}
type comments struct {
	User             string
	PseudoMsg        string
	Post             string
	IdPost           string
	Categorie        string
	Date             string
	Image            *string
	Extension        *string
	Statut           string
	PseudoMsgAttente string
	PostAttente      string
	CategorieAttente string
	DateAttente      string
	ImageAttente     string
	ExtensionAttente string
	Like             int
	DisLike          int
	AffLike          []int
	AffDisLike       []int
	ComId            []int
	ComPseudo        []string
	Comment          []string
	Erreur           string
	Role             string
}

type user struct {
	User           string
	Autorisation   string
	Id             []string
	Posts          []string
	IdAppr         []string
	PostsAppr      []string
	Page           int
	PageTotale     int
	PageAppr       int
	PageTotaleAppr int
	Statut         []string
	Role           string
	Demandes       []string
	Pseudos        []string
	PseudoModo     []string
	Signalement    []string
	CountNotif     int
	Notifs         []string
}
type activity struct {
	User         string
	Erreur       string
	IdPost       []string
	IdCom        []string
	IdLike       []string
	IdLikeCom    []string
	Posts        []string
	PostsLike    []string
	PostsLikeCom []string
	Comments     []string
	Likes        []string
	Role         string
}

type rechercher struct {
	Utilisateur      string
	MessParCategorie []string
	Pseudos          []string
	Categ            []string
	Date             []string
	MessParUser      []string
	MessParLike      []string
	PseudoParLike    []string

	Erreur string
	Id     []string
}

const port = ":5555"

var countPage = 1
var countPageAppr = 1
var indexMess = 0

var (
	templateHome    = template.Must(template.ParseFiles("templates/home.html"))
	templateInscr   = template.Must(template.ParseFiles("templates/inscription.html"))
	templateMess    = template.Must(template.ParseFiles("templates/message.html"))
	templateMessInd = template.Must(template.ParseFiles("templates/messageInd.html"))

	templateMessRech = template.Must(template.ParseFiles("templates/messRecherche.html"))
	templateAdmin    = template.Must(template.ParseFiles("templates/admin.html"))
	templateActivite = template.Must(template.ParseFiles("templates/activite.html"))
)

var (
	dejaLikeDemmarrage    []string
	dejaDislikeDemmarrage []string
	likePseudo            string
	likePseudoCom         string
	disLikePseudo         string
	dislikePseudoCom      string
	autorisationMessage   = "salut"
)

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func dejaLikeUser(utilisateur string, tab []int) []string {
	result := make([]string, len(tab))
	for i := range result {
		result[i] = "false"
	}
	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT idCom,PseudoCom, LikeCom,DislikeCom FROM likesCom WHERE PseudoCom = ?", utilisateur)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var idCom int
		var pseudo string
		var like string
		var dislike string
		err := rows.Scan(&idCom, &pseudo, &like, &dislike)
		if err != nil {
			log.Fatal(err)
		}
		for i, id := range tab {
			if idCom == id && like == "like" {
				result[i] = "true"
			}
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func dejadisLikeUser(utilisateur string, tab []int) []string {
	result := make([]string, len(tab))
	for i := range result {
		result[i] = "false"
	}
	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT idCom,PseudoCom, LikeCom,DislikeCom FROM likesCom WHERE PseudoCom = ?", utilisateur)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var idCom int
		var pseudo string
		var like string
		var dislike string
		err := rows.Scan(&idCom, &pseudo, &like, &dislike)
		if err != nil {
			log.Fatal(err)
		}
		for i, id := range tab {
			if idCom == id && dislike == "disLike" {
				result[i] = "true"
			}
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func countLikesCom(tab []int) []int {
	var result []int = make([]int, len(tab)) // Initialiser un tableau avec des compteurs à zéro
	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT idCom, LikeCom FROM likesCom")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var idCom int
		var like string
		err := rows.Scan(&idCom, &like)
		if err != nil {
			log.Fatal(err)
		}
		if like == "like" {
			for i, v := range tab {
				if v == idCom {
					result[i]++
				}
			}
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func countDisLikesCom(tab []int) []int {
	var result []int = make([]int, len(tab)) // Initialiser un tableau avec des compteurs à zéro
	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT idCom, DislikeCom FROM likesCom")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var idCom int
		var dislike string
		err := rows.Scan(&idCom, &dislike)
		if err != nil {
			log.Fatal(err)
		}
		if dislike == "disLike" {
			for i, v := range tab {
				if v == idCom {
					result[i]++
				}
			}
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func removeDuplicates(s []string) []string {
	// Utilisez un ensemble (set) pour stocker les éléments uniques
	set := make(map[string]bool)
	result := []string{}

	for _, element := range s {
		if set[element] {
			// Si l'élément existe déjà dans l'ensemble, ignorez-le
			continue
		}
		set[element] = true
		result = append(result, element)
	}

	return result
}

func likeCom(w http.ResponseWriter, r *http.Request) {
	// n := r.URL.Query().Get("id")

	// dejaLikeCom = false
	if r.Method == "POST" {

		var likeData LikeData

		// // Décodez le corps de la demande JSON dans la structure LikeData
		err := json.NewDecoder(r.Body).Decode(&likeData)
		if err != nil {
			// Gérez l'erreur de décodage JSON
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// 	w.Write([]byte("c bon"))
		// 	// fmt.Print(n)
		// 	fmt.Print(likeData.IdPost)

		// }

		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	http.Error(w, "Error reading request body",
		// 		http.StatusInternalServerError)
		// }
		// idCommentCom := string(body)
		// iddd := body
		// fmt.Print(iddd)
		// print(idCommentCom)
		idCommentCom = likeData.Id

		idPost := likeData.IdPost

		// fmt.Println(likePseudoCom)
		db, _ := sql.Open("sqlite3", "./dbForum.db")
		parcoursLikesCom, _ := db.Query("SELECT idCom,PseudoCom,LikeCom,DislikeCom FROM likesCom ")
		// 	// parcoursdisLikes, _ := db.Query("SELECT id, Pseudo,DisLike FROM disLikes ")

		var idCom string

		var pseudoLikesCom string
		var likesCom string
		var dislikesCom string

		for parcoursLikesCom.Next() {
			parcoursLikesCom.Scan(&idCom, &pseudoLikesCom, &likesCom, &dislikesCom)

			if idCom == idCommentCom && pseudoLikesCom == utilisateur && (likesCom == "like" || dislikesCom == "disLike") {
				dejaLikeCom = true

				w.Write([]byte("dejaLikeCom"))
			}

		}

		if connecte && dejaLikeCom == false {
			state, _ := db.Prepare("CREATE TABLE IF NOT EXISTS likesCom (idComText,idPost TEXT,PseudoCom TEXT,LikeCom text,DislikeCom text)")
			state.Exec()
			state, _ = db.Prepare("INSERT INTO likesCom (idCom,idPost,PseudoCom,LikeCom,DislikeCom) VALUES (?,?,?,?,?)")
			state.Exec(idCommentCom, idPost, utilisateur, "like", "")
			w.Write([]byte("c bon"))

		}
		// fmt.Print(dejaLikeCom)
		if dejaLikeCom {
			statement, _ := db.Prepare("delete from likesCom where idCom = ? and PseudoCom = ?") // ca ca marche a voir dans le js
			statement.Exec(idCommentCom, utilisateur)

		}

	} else {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

// // il faut afficher des messages d'erreur si on like et si on dislike

func disLikeCom(w http.ResponseWriter, r *http.Request) {
	dejadisLikeCom := false
	if r.Method == "POST" {
		var likeData LikeData

		// // Décodez le corps de la demande JSON dans la structure LikeData
		err := json.NewDecoder(r.Body).Decode(&likeData)
		if err != nil {
			// Gérez l'erreur de décodage JSON
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		idCommentCom = likeData.Id

		idPost := likeData.IdPost
		db, _ := sql.Open("sqlite3", "./dbForum.db")
		parcoursLikesCom, _ := db.Query("SELECT idCom, PseudoCom,LikeCom,DislikeCom FROM likesCom ")
		// 	// parcoursdisLikes, _ := db.Query("SELECT id, Pseudo,DisLike FROM disLikes ")

		var id string
		var pseudoLikes string
		var likes string
		var dislikes string

		for parcoursLikesCom.Next() {
			parcoursLikesCom.Scan(&id, &pseudoLikes, &likes, &dislikes)

			if id == idCommentCom && pseudoLikes == utilisateur && (likes == "like" || dislikes == "disLike") {
				dejadisLikeCom = true
				w.Write([]byte("dejadisLikeCom"))
			}

		}
		// reste a faire pour quon puisse pas lioke et dislike peut etre les mettre like et dislike dans une seule db

		if connecte && dejadisLikeCom == false {
			state, _ := db.Prepare("CREATE TABLE IF NOT EXISTS likesCom (idCom Text,idPost TEXT,PseudoCom TEXT,LikeCom text,DislikeCom text)")
			state.Exec()
			state, _ = db.Prepare("INSERT INTO likesCom (idCom,idPost,PseudoCom,LikeCom,DislikeCom) VALUES (?,?,?,?,?)")
			state.Exec(idCommentCom, idPost, utilisateur, "", "disLike")
			w.Write([]byte("c bon"))

		}
		if dejadisLikeCom {
			statement, _ := db.Prepare("delete from likesCom where idCom = ? and PseudoCom = ?") // ca ca marche a voir dans le js
			statement.Exec(idCommentCom, utilisateur)

		}

		// fmt.Print("Merci pour le like!")
	} else {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
	// fmt.Print(idCommentCom)
}

func commenter(w http.ResponseWriter, r *http.Request) {
	use := user{
		Autorisation: "",
	}
	pseudoMsg := r.FormValue("pseudoMsg")
	receptComment := r.FormValue("comment")
	if len(receptComment) == 0 {
		use.Autorisation = "Veuillez mettre du contenu"
	}
	if len(receptComment) > 0 {
		db, _ := sql.Open("sqlite3", "./dbForum.db")
		statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS comments (idComment INTEGER PRIMARY KEY,id Text, Pseudo TEXT,Commentaires TEXT)")
		statement.Exec()
		statement, _ = db.Prepare("INSERT INTO comments (id,Pseudo,Commentaires) VALUES (?,?,?)")
		statement.Exec(idComment, utilisateur, receptComment)
		if pseudoMsg != utilisateur {
			statemente, _ := db.Prepare("CREATE TABLE IF NOT EXISTS notifs (idNotif INTEGER PRIMARY KEY,idPost Text,idCom TEXT, Pseudo TEXT,read TEXT,Createur TEXT,Nature TEXT)")
			statemente.Exec()
			statemente, _ = db.Prepare("INSERT INTO notifs(idPost,idCom,Pseudo,read,Createur,Nature) VALUES (?,?,?,?,?,?)")
			statemente.Exec(idComment, "", utilisateur, false, pseudoMsg, "comment")

			use.Autorisation = "Votre commentaire a été ajouté"
			// creer un idUnique pour chaque commentaire
		}
	}
	if connecte == false {
		use.Autorisation = "Vous devez etre connecte pour commenter"
	}
	templateHome.ExecuteTemplate(w, "home.html", use)
}

func like(w http.ResponseWriter, r *http.Request) {

	// receptComment := r.FormValue("comment")

	dejaLike = false
	if r.Method == "POST" {
		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	http.Error(w, "Error reading request body",
		// 		http.StatusInternalServerError)
		// }
		// likePseudo = string(body)
		// // fmt.Println(likePseudo)
		db, _ := sql.Open("sqlite3", "./dbForum.db")
		parcoursLikes, _ := db.Query("SELECT id, Pseudo,Like,Dislike FROM likes ")
		// // 	// parcoursdisLikes, _ := db.Query("SELECT id, Pseudo,DisLike FROM disLikes ")
		var likeData LikeData

		// Décodez le corps de la demande JSON dans la structure LikeData
		err := json.NewDecoder(r.Body).Decode(&likeData)
		if err != nil {
			// Gérez l'erreur de décodage JSON
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Vous pouvez maintenant accéder à likeData.Pseudo et likeData.DejaLikePourGO
		pseudoUser := likeData.Pseudo
		likePseudo := likeData.DejaLikePourGO

		// fmt.Print(pseudo, dejaLikePourGO)
		var id string
		var pseudoLikes string
		var likes string
		var dislikes string

		for parcoursLikes.Next() {
			parcoursLikes.Scan(&id, &pseudoLikes, &likes, &dislikes)

			if id == idComment && pseudoLikes == utilisateur && (likes == "like" || dislikes == "disLike") {
				dejaLike = true
			}

		}
		// reste a faire pour quon puisse pas lioke et dislike peut etre les mettre like et dislike dans une seule db

		if connecte && dejaLike == false {
			state, _ := db.Prepare("CREATE TABLE IF NOT EXISTS likes (id Text, Pseudo TEXT,Like text,Dislike text)")
			state.Exec()
			state, _ = db.Prepare("INSERT INTO likes (id,Pseudo,Like,Dislike) VALUES (?,?,?,?)")
			state.Exec(idComment, utilisateur, "like", "")
			if pseudoUser != utilisateur {
				statemente, _ := db.Prepare("CREATE TABLE IF NOT EXISTS notifs (idNotif INTEGER PRIMARY KEY,idPost Text,idCom TEXT, Pseudo TEXT,read TEXT,Createur TEXT,Nature TEXT)")
				statemente.Exec()
				statemente, _ = db.Prepare("INSERT INTO notifs(idPost,idCom,Pseudo,read,Createur,Nature) VALUES (?,?,?,?,?,?)")
				statemente.Exec(idComment, "", utilisateur, false, pseudoUser, "like")
			}
		}
		if likePseudo {
			statement, _ := db.Prepare("delete from likes where id = ? and Pseudo = ?") // ca ca marche a voir dans le js
			statement.Exec(idComment, utilisateur)

			statemente, _ := db.Prepare("delete from notifs where idPost = ? and Pseudo = ?") // ca ca marche a voir dans le js
			statemente.Exec(idComment, utilisateur)
		}

		// fmt.Print("Merci pour le like!")
	} else {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

// il faut afficher des messages d'erreur si on like et si on dislike

func disLike(w http.ResponseWriter, r *http.Request) {
	dejaLike = false
	if r.Method == "POST" {
		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	http.Error(w, "Error reading request body",
		// 		http.StatusInternalServerError)
		// }

		// disLikePseudo = string(body)

		var likeData LikeData

		// Décodez le corps de la demande JSON dans la structure LikeData
		err := json.NewDecoder(r.Body).Decode(&likeData)
		if err != nil {
			// Gérez l'erreur de décodage JSON
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Vous pouvez maintenant accéder à likeData.Pseudo et likeData.DejaLikePourGO
		pseudoUser := likeData.Pseudo

		disLikePseudo := likeData.DejaLikePourGO

		// fmt.Println(likePseudo)
		db, _ := sql.Open("sqlite3", "./dbForum.db")
		parcoursLikes, _ := db.Query("SELECT id, Pseudo,Like,Dislike FROM likes ")
		// 	// parcoursdisLikes, _ := db.Query("SELECT id, Pseudo,DisLike FROM disLikes ")

		var id string
		var pseudoLikes string
		var likes string
		var dislikes string

		for parcoursLikes.Next() {
			parcoursLikes.Scan(&id, &pseudoLikes, &likes, &dislikes)

			if id == idComment && pseudoLikes == utilisateur && (likes == "like" || dislikes == "disLike") {
				dejaLike = true
			}

		}
		// reste a faire pour quon puisse pas lioke et dislike peut etre les mettre like et dislike dans une seule db

		if connecte && dejaLike == false {
			state, _ := db.Prepare("CREATE TABLE IF NOT EXISTS likes (id Text, Pseudo TEXT,Like text,Dislike text)")
			state.Exec()
			state, _ = db.Prepare("INSERT INTO likes (id,Pseudo,Like,Dislike) VALUES (?,?,?,?)")
			state.Exec(idComment, utilisateur, "", "disLike")
			if pseudoUser != utilisateur {
				statemente, _ := db.Prepare("CREATE TABLE IF NOT EXISTS notifs (idNotif INTEGER PRIMARY KEY,idPost Text,idCom TEXT, Pseudo TEXT,read TEXT,Createur TEXT,Nature TEXT)")
				statemente.Exec()
				statemente, _ = db.Prepare("INSERT INTO notifs(idPost,idCom,Pseudo,read,Createur,Nature) VALUES (?,?,?,?,?,?)")
				statemente.Exec(idComment, "", utilisateur, false, pseudoUser, "dislike")

			}
		}
		if disLikePseudo == true {
			statement, _ := db.Prepare("delete from likes where id = ? and Pseudo = ?") // ca ca marche a voir dans le js
			statement.Exec(idComment, utilisateur)
			statemente, _ := db.Prepare("delete from notifs where idPost = ? and Pseudo = ?") // ca ca marche a voir dans le js
			statemente.Exec(idComment, utilisateur)
		}

		// fmt.Print("Merci pour le like!")
	} else {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

// func lienMessage(id string) {
// 	database, _ := sql.Open("sqlite3", "./dbForum.db")
// 	// parcours, _ := database.Query("SELECT Pseudo, Posts,categorie,Date,Image,extension,statut FROM messages")
// 	parcours, _ := database.Query("SELECT Pseudo, Posts, categorie, Date, Image, extension, statut FROM messages WHERE id = ?", id)

// 	var posts string
// 	var pseudoMessage string
// 	var categ string
// 	var date string
// 	var image string
// 	var extension string
// 	var statut string
// 	for parcours.Next() {
// 		parcours.Scan(&pseudoMessage, &posts, &categ, &date, &image, &extension, &statut)
// 		fmt.Print(posts)

// 	}
// }

func lienMessage(id string) (comments, error) {
	var imageNullable sql.NullString
	var extNullable sql.NullString
	database, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		return comments{}, err
	}

	row := database.QueryRow("SELECT Pseudo, Posts, categorie, Date, Image, extension, statut FROM messages WHERE id = ?", id)

	var messageInfo comments
	err = row.Scan(
		&messageInfo.PseudoMsg,
		&messageInfo.Post,
		&messageInfo.Categorie,
		&messageInfo.Date,
		&imageNullable,
		&extNullable,
		&messageInfo.Statut,
	)
	if imageNullable.Valid {
		messageInfo.Image = &imageNullable.String
	} else {
		messageInfo.Image = nil // Si l'image est NULL, affectez un pointeur nul
	}
	if extNullable.Valid {
		messageInfo.Extension = &extNullable.String
	} else {
		messageInfo.Extension = nil // Si l'image est NULL, affectez un pointeur nul
	}
	if err != nil {
		return comments{}, err
	}

	return messageInfo, nil
}
func reagirMess(w http.ResponseWriter, r *http.Request) {
	// trouver un moyen de recuperer l'id unique du commentaire sur lequel je suis
	// fmt.Print(roleGlobal)
	dejaLike = false
	// var trouvMess []string
	// var pseudoMess []string
	// var categMess []string
	// var dateMess []string
	// var imgMess []string
	// var imgExtension []string

	var countLike int
	var countLikeCom int
	var countdisLike int
	var countdisLikeCom int

	com := comments{
		PseudoMsg: "",
		Post:      "",
		Categorie: "",
		Date:      "",

		Like:       0,
		DisLike:    0,
		ComId:      []int{0},
		AffLike:    []int{0},
		AffDisLike: []int{0},
		ComPseudo:  []string{""},
		Comment:    []string{""},
		Erreur:     "",
	}
	// var com comments
	n := r.URL.Query().Get("id")
	com, err := lienMessage(n)
	if err != nil {
		fmt.Print("error")
		// Gérez l'erreur, par exemple en affichant un message d'erreur
	}
	com.IdPost = n
	// fmt.Print(lienMessage(n))
	idComment = r.URL.Query().Get("id")
	// ne, _ := strconv.Atoi(r.URL.Query().Get("ide"))

	// var commentInd string

	// var trouvAll []string
	if r.URL.Path != "/messagesInd" {
		http.NotFound(w, r)
		fmt.Fprintf(w, "Status 404: Page Not Found")
		return
	}
	database, _ := sql.Open("sqlite3", "./dbForum.db")

	parcoursComm, _ := database.Query("SELECT idComment,id, Pseudo,Commentaires FROM comments ") //
	parcoursLikes, _ := database.Query("SELECT id, Pseudo,Like ,Dislike FROM likes ")
	parcoursLikesCom, _ := database.Query("SELECT idCom, PseudoCom,LikeCom ,DislikeCom FROM likesCom ")
	// parcoursdisLikes, _ := database.Query("SELECT id, Pseudo,DisLike FROM disLikes ") //

	// var posts string
	// var pseudoMessage string
	// var categ string
	// var date string
	// var image string
	// var extension string
	// var statut string

	var pseudoDejaLike []string
	var pseudoDejaDisLike []string
	var pseudoDejaLikeCom []string
	var pseudoDejaDisLikeCom []string

	var idLike string
	var pseudoLike string
	var like string
	var dislike string
	// var essai []string

	for parcoursLikes.Next() {
		parcoursLikes.Scan(&idLike, &pseudoLike, &like, &dislike)
		if idComment == idLike && like == "like" {
			countLike++
		}
		if idComment == idLike && (like == "like") {
			pseudoDejaLike = append(pseudoDejaLike, pseudoLike)
		}
		if idComment == idLike && (dislike == "disLike") {
			pseudoDejaDisLike = append(pseudoDejaDisLike, pseudoLike)
		}
		if idComment == idLike && dislike == "disLike" {
			countdisLike++
		}

	}
	var idLikeCom string
	var pseudoLikeCom string
	var likeCom string
	var dislikeCom string
	for parcoursLikesCom.Next() {
		parcoursLikesCom.Scan(&idLikeCom, &pseudoLikeCom, &likeCom, &dislikeCom)
		if idCommentCom == idLikeCom && likeCom == "like" {
			countLikeCom++
		}
		if idCommentCom == idLikeCom && (likeCom == "like") {
			pseudoDejaLikeCom = append(pseudoDejaLikeCom, pseudoLikeCom)
		}
		if idCommentCom == idLikeCom && (dislikeCom == "disLike") {
			pseudoDejaDisLikeCom = append(pseudoDejaDisLikeCom, pseudoLikeCom)
		}
		if idCommentCom == idLikeCom && dislikeCom == "disLike" {
			countdisLikeCom++
		}

	}
	// fmt.Print(pseudoDejaLike)
	var idCommentaire int
	var idCom string
	var pseudoComment string
	var comComment string
	var affIdCommentaire []int
	var affPseudoComment []string
	var affComment []string

	for parcoursComm.Next() {
		parcoursComm.Scan(&idCommentaire, &idCom, &pseudoComment, &comComment)
		if idComment == idCom {
			affIdCommentaire = append(affIdCommentaire, idCommentaire)
			affPseudoComment = append(affPseudoComment, pseudoComment)
			affComment = append(affComment, comComment)

		}

	}

	if connecte {
		com.User = utilisateur
	}
	dejaLikeDemmarrage = (dejaLikeUser(utilisateur, affIdCommentaire))
	dejaDislikeDemmarrage = (dejadisLikeUser(utilisateur, affIdCommentaire))
	// fmt.Print(dejaLikeDemmarrage, dejaDislikeDemmarrage)

	comAffLikeJSON, _ := json.Marshal(dejaLikeDemmarrage)
	comAffDisLikeJSON, _ := json.Marshal(dejaDislikeDemmarrage)

	http.SetCookie(w, &http.Cookie{
		Name:  "comAffLike",
		Value: url.PathEscape(string(comAffLikeJSON)),
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "comAffDisLike",
		Value: url.PathEscape(string(comAffDisLikeJSON)),
	})

	com.AffDisLike = countDisLikesCom(affIdCommentaire)
	com.AffLike = countLikesCom(affIdCommentaire)
	com.ComId = affIdCommentaire
	com.ComPseudo = affPseudoComment
	com.Comment = affComment

	com.Role = roleGlobal

	com.Like = countLike
	com.DisLike = countdisLike

	cookieLike := &http.Cookie{
		Name:    "like",
		Value:   strconv.Itoa(countLike),
		Expires: time.Now().Add(365 * 24 * time.Hour), // envoyer directement au javascript les likes de la db avec cookies
	}

	http.SetCookie(w, cookieLike)
	cookieDislike := &http.Cookie{
		Name:    "disLike",
		Value:   strconv.Itoa(countdisLike),
		Expires: time.Now().Add(365 * 24 * time.Hour),
	}
	http.SetCookie(w, cookieDislike)
	cookiePseudolike := &http.Cookie{
		Name:    "dejaLike",
		Value:   strings.Join(pseudoDejaLike, " "),
		Expires: time.Now().Add(365 * 24 * time.Hour),
	}
	http.SetCookie(w, cookiePseudolike)
	cookiePseudoDislike := &http.Cookie{
		Name:    "dejaDisLike",
		Value:   strings.Join(pseudoDejaDisLike, " "),
		Expires: time.Now().Add(365 * 24 * time.Hour),
	}
	http.SetCookie(w, cookiePseudoDislike)

	// fmt.Print(countLikesCom(affIdCommentaire))

	templateMessInd.ExecuteTemplate(w, "messageInd.html", com)

	// fmt.Fprint(w, nomImgMess[n-1])
}

func formMessage(w http.ResponseWriter, r *http.Request) {
	use := user{
		User:         "",
		Autorisation: "",
	}
	if connecte {

		use.User = utilisateur

		templateMess.ExecuteTemplate(w, "message.html", use)
	} else {
		use.Autorisation = "Veuillez vous connecter"
		templateHome.ExecuteTemplate(w, "home.html", use)

	}
}

func ecrireMessage(w http.ResponseWriter, r *http.Request) {
	var fileName = ""
	var encodedImage = ""
	use := user{
		User:         "",
		Autorisation: "",
	}
	r.ParseMultipartForm(10 << 20) // Taille maximale de fichier (10 Mo)

	file, fileHeader, err := r.FormFile("fileToUpload")
	if err != nil {
		// Aucun fichier n'a été téléchargé, gérer cette situation en conséquence
		// Par exemple, vous pouvez définir une valeur par défaut pour fileName
		fileName = "default.jpg" // Remplacez "default.jpg" par le nom de votre choix
	} else {
		defer file.Close()
		// Récupérer le nom du fichier à partir de l'en-tête

		defer file.Close()
		fileName = fileHeader.Filename
		imageBytes, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		encodedImage = base64.StdEncoding.EncodeToString(imageBytes)
		// Le reste de votre code pour traiter le fichier téléchargé
		// ...
	}

	// Récupérer le nom du fichier à partir de l'en-tête

	// Diviser le nom du fichier pour obtenir l'extension
	parts := strings.Split(fileName, ".")
	extension := parts[len(parts)-1] // Obtient la dernière partie du nom du fichier (l'extension)
	// fmt.Print(extension)
	// Convertissez le contenu de l'image en base64

	// Récupérez les autres données du formulaire (message, catégorie, etc.)
	receptMess := strings.TrimSpace(r.FormValue("message"))
	categories := r.Form["categorie[]"]
	categorieStr := strings.Join(categories, ",")

	// Insérez les données (y compris l'image encodée) dans la base de données
	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS messages (id INTEGER PRIMARY KEY, Pseudo TEXT, Posts TEXT, Date TEXT, categorie TEXT, Image TEXT,extension TEXT,statutTEXT,signalement TEXT)")
	statement.Exec()
	statement, err = db.Prepare("INSERT INTO messages (Pseudo, Posts, Date, categorie, Image,extension,statut,signalement) VALUES (?, ?, ?, ?, ?,?,?,?)")
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	t := time.Now()
	date := t.Format("2006-01-02 15:04:05")
	if len(receptMess) > 0 && (roleGlobal == "admin" || roleGlobal == "modo") {
		_, err = statement.Exec(utilisateur, receptMess, date, categorieStr, encodedImage, extension, "approuve", false)
		if err != nil {
			panic(err)
		}
		use.Autorisation = "Votre message a été ajouté"
	} else {
		use.Autorisation = "Veuillez mettre du contenu dans le message"
	}
	if len(receptMess) > 0 && roleGlobal == "user" {
		_, err = statement.Exec(utilisateur, receptMess, date, categorieStr, encodedImage, extension, "attente", false)
		if err != nil {
			panic(err)
		}
		use.Autorisation = "Votre message a été ajouté"
	} else {
		use.Autorisation = "Veuillez mettre du contenu dans le message"
	}

	// Réponse de succès

	templateHome.ExecuteTemplate(w, "home.html", use)
}

func connexion(w http.ResponseWriter, r *http.Request) {
	// cookies := r.Cookies()
	userBrowser := r.UserAgent()
	// fmt.Print(userBrowser)
	connecte = false

	t := time.Now()

	date := t.Format("2006-01-02 15:04:05")

	// t := time.Now()

	// date = t.Format("2006-01-02 15:04:05")

	use := user{
		User:         "",
		Autorisation: "",
	}

	use.Autorisation = "La connexion a reussi"

	receptPseudo := r.FormValue("pseudo")
	receptMdp := r.FormValue("mdp")

	database, _ := sql.Open("sqlite3", "./dbForum.db")
	parcours, _ := database.Query("SELECT id, Pseudo, Mdp,role FROM people")
	var id int
	var pseudo string
	var mdp string
	var role string
	for parcours.Next() {
		parcours.Scan(&id, &pseudo, &mdp, &role)

		err := bcrypt.CompareHashAndPassword([]byte(mdp), []byte(receptMdp))
		if err == nil && receptPseudo == pseudo {
			connecte = true
			roleGlobal = role
		}
		if err == nil && receptPseudo == "admin" {

		}

	}

	if connecte {
		utilisateur = receptPseudo
		cookie := &http.Cookie{
			Name:   utilisateur,
			Value:  date, // faire avec la date
			MaxAge: 10 * 365 * 24 * 60 * 60,
		}
		http.SetCookie(w, cookie)

		statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS cookies (id INTEGER PRIMARY KEY, CookieValue TEXT,Pseudo TEXT,NavigateurAct TEXT,Role TEXT)")
		statement.Exec()

		statement, _ = database.Prepare("INSERT INTO cookies (CookieValue,Pseudo,NavigateurAct,role) VALUES (?,?,?,?)")
		statement.Exec(date, utilisateur, userBrowser, roleGlobal)

		use.User = "Bonjour " + utilisateur
		use.Autorisation = "La connexion a reussi"
		use.Role = roleGlobal

		templateHome.ExecuteTemplate(w, "home.html", use)

	} else {
		connecte = false
		use.User = "invite"
		use.Autorisation = "pseudo ou MDP inconnu"
		if len(receptPseudo) == 0 && len(receptMdp) == 0 {
			use.Autorisation = "Veuillez remplir les champs vides:"
		}

		templateHome.ExecuteTemplate(w, "home.html", use)

	}
}

func deconnexion(w http.ResponseWriter, r *http.Request) {
	connecte = false
	utilisateur = "deconnecte"
	c := http.Cookie{
		Name:   utilisateur,
		Value:  "",
		MaxAge: 10 * 365 * 24 * 60 * 60,
	}
	http.SetCookie(w, &c)

	var messageSlice []string
	var idSlice []string

	use := user{
		User: "",
	}

	database, _ := sql.Open("sqlite3", "./dbForum.db")

	parcours, _ := database.Query("SELECT id, Posts  FROM messages")
	var id string
	var messages string

	for parcours.Next() {
		parcours.Scan(&id, &messages)

		messageSlice = append(messageSlice, messages)
		idSlice = append(idSlice, id)

	}

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS cookies (id INTEGER PRIMARY KEY, CookieValue TEXT,Pseudo TEXT)")
	statement.Exec()

	statement, _ = database.Prepare("INSERT INTO cookies (CookieValue,Pseudo) VALUES (?,?)")
	statement.Exec(date, utilisateur)
	use.Id = idSlice
	use.Posts = messageSlice
	templateHome.ExecuteTemplate(w, "home.html", use)

	// mettre qd meme acces aux messages
}

func insertInscription(w http.ResponseWriter, r *http.Request) {
	erreure := false
	receptPseudoEnr := r.FormValue("pseudoEnr")
	receptMdpEnr := r.FormValue("mdpEnr")
	receptMailEnr := r.FormValue("mailEnr")

	database, _ := sql.Open("sqlite3", "./dbForum.db")

	erreur := erreurs{
		MsgErreur: "",
	}

	password := []byte(receptPseudoEnr)

	// Générer le hash du mot de passe. La valeur de retour est un slice de bytes
	// qui contient le hash. La fonction prend en paramètre le niveau de complexité
	// souhaité (ici, 12). Plus ce niveau est élevé, plus le hash sera sécurisé,
	// mais le temps de calcul sera plus long.
	hash, err := bcrypt.GenerateFromPassword(password, 12)
	if err != nil {
		panic(err)
	}

	// fmt.Fprint(w, erreur.MsgErreur)
	if len(receptMdpEnr) <= 1 && len(receptMailEnr) <= 1 && len(receptMdpEnr) <= 1 {
		erreur.MsgErreur = "Veuillez ecrire plus de 1 caractere pour tous les champs"
		erreure = true

	} else if !strings.Contains(receptMailEnr, "@") {
		erreur.MsgErreur = "Veuillez entrer un mail au bon format(exemple@exemple.com)"
		erreure = true

	}
	parcourse, _ := database.Query("SELECT  Pseudo, Mail FROM people")

	var pseudos string
	var mail string
	for parcourse.Next() {
		parcourse.Scan(&pseudos, &mail)
		if receptPseudoEnr == pseudos || receptMailEnr == mail {
			erreur.MsgErreur = "Ces identifiants ou mails sont deja pris, veuillez en choisir un autre"
			erreure = true
		}

	}
	if erreure == true {
		templateInscr.ExecuteTemplate(w, "inscription.html", erreur)
	}

	if erreure == false {
		erreur.MsgErreur = "Votre inscription a reussi"
		statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, Pseudo TEXT, Mdp TEXT,Mail TEXT,role TEXT)")
		statement.Exec()
		statement, _ = database.Prepare("INSERT INTO people (Pseudo, Mdp,Mail,role) VALUES (?,?,?,?)")
		statement.Exec(receptPseudoEnr, hash, receptMailEnr, "user") // si on rajoute une nouvelle rubrique il faut suppr le fichier ou creer une nvelle table
		templateHome.ExecuteTemplate(w, "home.html", erreur)         // faire une page qui affiche les derniers messages

	}

	//
	// 	str = "pas de mots vides"

	// 	fmt.Fprint(w, str)
	// }
}
func splitMots(message string) string {
	words := strings.Split(message, " ")
	var truncatedMessage string
	// Vérifiez que le message a au moins un mot
	if len(words) > 0 {
		// Récupérez le premier mot
		firstWord := words[0]

		// Récupérez le dernier mot
		lastWord := words[len(words)-1]

		// Créez une nouvelle chaîne avec le premier et le dernier mot
		truncatedMessage = firstWord + "..." + lastWord

	}
	return truncatedMessage

}

func home(w http.ResponseWriter, r *http.Request) {

	navigateur := r.UserAgent()

	// var messSliceInverse []string
	// var idSliceInverse []string
	// var statutSliceInverse []string
	var pageTotale int
	var pageTotaleAppr int

	boutSuivant := r.FormValue("pageSuivante")
	boutPrecedent := r.FormValue("pagePrecedente")
	boutSuivantAppr := r.FormValue("pageSuivanteAppr")
	boutPrecedentAppr := r.FormValue("pagePrecedenteAppr")
	use := user{
		User:         "",
		Autorisation: "",

		Page:       0,
		PageTotale: 0,
	}

	if r.URL.Path != "/home" {
		http.NotFound(w, r)
		fmt.Fprintf(w, "Status 404: Page Not Found")
		return
	}
	cookies := r.Cookies()
	var valSlice []string
	var pseudoSlice []string
	var navigator []string
	var roleStatut []string
	var signalement []string
	// var statutSlice []string
	// var connecte bool

	db, _ := sql.Open("sqlite3", "./dbForum.db")
	parcours, _ := db.Query("SELECT CookieValue, Pseudo,NavigateurAct,role FROM cookies ")

	var pseudo string
	var valeur string
	var navig string
	var role string
	for parcours.Next() {
		parcours.Scan(&valeur, &pseudo, &navig, &role)
		valSlice = append(valSlice, valeur)
		pseudoSlice = append(pseudoSlice, pseudo)
		navigator = append(navigator, navig)
		roleStatut = append(roleStatut, role)
		// si la valeur du cookie est egale a la derniere valeur du cookie de la db et du pseudo il peut etre connecte sinon non
	}
	for _, cookie := range cookies {
		if pseudoSlice[len(pseudoSlice)-1] == "deconnecte" || cookie.Name == "deconnecte" || navigator[len(navigator)-1] != navigateur {
			connecte = false
			use.User = "invite"
			use.Autorisation = "Vous n'etes pas connecte"
		} else {

			utilisateur = pseudoSlice[len(pseudoSlice)-1]
			connecte = true
			use.User = utilisateur
			roleGlobal = roleStatut[len(roleStatut)-1]
			fmt.Print(use.Role)
			use.Autorisation = ""
			// fmt.Print(use.Role)
		}

	}
	var pseudoOnly []string
	var messageSlice []string
	var messageSliceAppr []string
	var idSlice []string
	var idSliceAppr []string
	// var statutSlice []string

	database, _ := sql.Open("sqlite3", "./dbForum.db")

	parcourse, _ := database.Query("SELECT  Pseudo FROM people")
	var pseudon string
	for parcourse.Next() {
		parcourse.Scan(&pseudon)

		pseudoOnly = append(pseudoOnly, pseudon)

	}

	parcoursette, _ := database.Query("SELECT id, Posts,statut,signalement  FROM messages")
	var id string
	var messages string
	var statut string
	var signal string
	// Déclarer une variable pour garder la trace de l'index

	for parcoursette.Next() {
		parcoursette.Scan(&id, &messages, &statut, &signal)

		if (roleGlobal == "admin" || roleGlobal == "modo") && statut == "attente" {

			messageSlice = append(messageSlice, splitMots(messages))

			signalement = append(signalement, signal)
			idSlice = append(idSlice, id)

		}
		if (roleGlobal == "admin" || roleGlobal == "modo") && statut == "approuve" {
			messageSliceAppr = append(messageSliceAppr, splitMots(messages))
			idSliceAppr = append(idSliceAppr, id)

		}

		if statut == "approuve" && roleGlobal == "user" {

			messageSlice = append(messageSlice, splitMots(messages))

			idSlice = append(idSlice, id)

		}

	}

	pageTotale = (len(messageSlice) + 9) / 10
	debut := (countPage - 1) * 10
	fin := countPage * 10
	if fin > len(messageSlice) {
		fin = len(messageSlice)
	}

	pageTotaleAppr = (len(messageSliceAppr) + 9) / 10
	debutAppr := (countPageAppr - 1) * 10
	finAppr := countPageAppr * 10
	if finAppr > len(messageSliceAppr) {
		finAppr = len(messageSliceAppr)
	}
	use.Signalement = signalement
	use.Id = idSlice[debut:fin]
	use.Posts = messageSlice[debut:fin]
	use.IdAppr = idSliceAppr[debutAppr:finAppr]
	use.PostsAppr = messageSliceAppr[debutAppr:finAppr]
	use.CountNotif = countNotifs(utilisateur)

	if boutPrecedent == "precedent" && countPage > 1 {
		countPage--
	}

	if boutSuivant == "suivant" && countPage < pageTotale {
		countPage++
	}

	if boutPrecedentAppr == "precedent" && countPageAppr > 1 {
		countPageAppr--
	}

	if boutSuivantAppr == "suivant" && countPageAppr < pageTotaleAppr {
		countPageAppr++
	}

	use.Page = countPage
	use.PageTotale = pageTotale
	use.PageAppr = countPage
	use.PageTotaleAppr = pageTotale
	use.Role = roleGlobal
	// fmt.Print(messageSlice)
	templateHome.ExecuteTemplate(w, "home.html", use)

}

func inscription(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/inscription" {
		http.NotFound(w, r)
		fmt.Fprintf(w, "Status 404: Page Not Found")
		return
	}

	templateInscr.ExecuteTemplate(w, "inscription.html", "")
}

func recherche(w http.ResponseWriter, r *http.Request) {
	// il faut separer mesmessages des categories

	var messagesCateg []string
	var pseudoCateg []string
	var categCateg []string
	var dateCateg []string
	var idCateg []string
	structRecherche := rechercher{
		Utilisateur:      "",
		MessParCategorie: []string{""},
		Pseudos:          []string{""},
		Categ:            []string{""},
		Date:             []string{""},
		MessParUser:      []string{""},
		MessParLike:      []string{""},
		PseudoParLike:    []string{""},

		Erreur: "",
		Id:     []string{""},
	}
	if connecte {
		structRecherche.Utilisateur = utilisateur
		structRecherche.Erreur = ""
	}
	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r.ParseForm()
	categories, ok := r.Form["filtreCategorie[]"]
	if connecte == false && categories != nil && contains(categories, "mesLikes") {
		structRecherche.Erreur = "Veuillez vous connecter pour voir vos likes"
	}
	if connecte == false && categories != nil && contains(categories, "mesMessages") {
		structRecherche.Erreur = "Veuillez vous connecter pour voir vos messages"
	}
	if connecte && categories != nil && contains(categories, "mesLikes") {
		rows, err := db.Query("SELECT id FROM likes WHERE Pseudo=? AND Like='like'", utilisateur)
		if err != nil {
			fmt.Print(err)
		}
		defer rows.Close()

		var ids []string
		// var idCom []string
		for rows.Next() {
			var id string
			err := rows.Scan(&id)
			if err != nil {
				fmt.Print(err)
			}
			ids = append(ids, id)
		}
		// rowse, err := db.Query("SELECT idCom-mess FROM likesCom WHERE PseudoCom=? AND LikeCom='like'", utilisateur)
		// if err != nil {
		// 	fmt.Print(err)
		// }
		// defer rowse.Close()

		// for rowse.Next() {
		// 	var ide string
		// 	err := rowse.Scan(&ide)
		// 	if err != nil {
		// 		fmt.Print(err)
		// 	}
		// 	ids = append(ids, ide)
		// }

		if len(ids) == 0 {
			structRecherche.Erreur = "Pas de likes"
		} else {

			query := "SELECT id,Pseudo,Posts,Date,categorie FROM messages WHERE "
			for i, id := range ids {
				if i > 0 {
					query += "OR "
				}
				query += "id=" + id + " "
			}

			rows, err = db.Query(query)

			defer rows.Close()
		}
		var messagesLike []string
		var pseudoLike []string
		var categorieLike []string
		var idLike []string
		for rows.Next() {
			var id int
			var pseudo string
			var posts string
			var date string
			var categorie string

			err := rows.Scan(&id, &pseudo, &posts, &date, &categorie)
			if err != nil {
				fmt.Print(err)
			}

			messagesLike = append(messagesLike, posts)
			pseudoLike = append(pseudoLike, pseudo)
			categorieLike = append(categorieLike, categorie)
			idLike = append(idLike, strconv.Itoa(id))

		}

		structRecherche.MessParCategorie = messagesLike
		structRecherche.Categ = categorieLike
		structRecherche.Id = idLike

	} else if connecte && categories != nil && contains(categories, "mesMessages") {
		rows, err := db.Query("SELECT id,Pseudo,Posts,Date,categorie FROM messages WHERE Pseudo=? ", utilisateur)
		if err != nil {
			fmt.Print(err)
		}
		defer rows.Close()
		var messMess []string

		var categMess []string
		var idMess []string

		for rows.Next() {
			var id int
			var pseudo string
			var posts string
			var date string
			var categorie string

			err := rows.Scan(&id, &pseudo, &posts, &date, &categorie)
			if err != nil {
				fmt.Print(err)
			}

			messMess = append(messMess, splitMots(posts))

			categMess = append(categMess, categorie)
			idMess = append(idMess, strconv.Itoa(id))

		}
		if len(messagesCateg) == 0 {
			structRecherche.Erreur = "Aucun messages trouves"
		}
		structRecherche.MessParCategorie = messMess
		structRecherche.Categ = categMess
		structRecherche.Id = idMess
	} else {

		if ok {
			query := "SELECT id,Pseudo,Posts,Date,categorie,statut FROM messages WHERE "
			for i, cat := range categories {
				if i > 0 {
					query += "OR "
				}
				query += "categorie LIKE '%" + cat + "%' "
			}

			rows, err := db.Query(query)
			if err != nil {
				panic(err.Error())
			}
			defer rows.Close()

			// fmt.Println("Messages correspondants :")
			for rows.Next() {
				var id int
				var pseudo string
				var posts string
				var date string
				var categorie string
				var statut string

				err := rows.Scan(&id, &pseudo, &posts, &date, &categorie, &statut)
				if err != nil {
					panic(err.Error())
				}
				if statut == "approuve" {
					messagesCateg = append(messagesCateg, splitMots(posts))
					pseudoCateg = append(pseudoCateg, pseudo)
					categCateg = append(categCateg, categorie)
					dateCateg = append(dateCateg, date)
					idCateg = append(idCateg, strconv.Itoa(id))
				}
			}

		} else {
			fmt.Println("Aucune catégorie sélectionnée.")
		}

		structRecherche.MessParCategorie = messagesCateg
		structRecherche.Id = idCateg
		structRecherche.Categ = categCateg

	}
	templateMessRech.ExecuteTemplate(w, "messRecherche.html", structRecherche)
}
func getStatut() string {

	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT role FROM people WHERE Pseudo=? ", utilisateur)
	if err != nil {
		fmt.Print(err)
	}
	defer rows.Close()
	for rows.Next() {
		var role string

		err := rows.Scan(&role)
		if err != nil {
			panic(err.Error())
		}

		roleGlobal = role
	}
	return roleGlobal

}
func dejaDemandeModo() bool {

	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT pseudo,demande FROM contactAdmin")
	if err != nil {
		fmt.Print(err)
	}
	defer rows.Close()
	for rows.Next() {
		var pseudo string
		var demande string

		err := rows.Scan(&pseudo, &demande)
		if err != nil {
			panic(err.Error())
		}

		if pseudo == utilisateur && demande == "attente" {
			return true
		}
	}
	return false

}
func admin(w http.ResponseWriter, r *http.Request) {
	useRole := user{
		Role: "",
		User: utilisateur,
	}
	useRole.Role = getStatut()
	useRole.User = utilisateur

	fmt.Print(autorisationMessage)

	templateAdmin.ExecuteTemplate(w, "admin.html", useRole)
}
func maintenance(w http.ResponseWriter, r *http.Request) {
	useRol := user{
		Demandes: []string{""},
		Pseudos:  []string{""},
		Role:     "",
		User:     utilisateur,
	}

	var demandesModo []string
	var pseudoModos []string
	var pseudoDemandes []string
	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rowse, err := db.Query("SELECT Pseudo,role FROM people")
	if err != nil {
		fmt.Print(err)
	}
	defer rowse.Close()
	for rowse.Next() {
		var pseudo string
		var role string

		err := rowse.Scan(&pseudo, &role)
		if err != nil {
			panic(err.Error())
		}
		if role == "modo" {
			pseudoModos = append(pseudoModos, pseudo)

		}

	}
	rows, err := db.Query("SELECT pseudo,demande FROM contactAdmin")
	if err != nil {
		fmt.Print(err)
	}
	defer rows.Close()
	for rows.Next() {
		var pseudo string
		var demande string

		err := rows.Scan(&pseudo, &demande)
		if err != nil {
			panic(err.Error())
		}
		if demande == "attente" {
			pseudoDemandes = append(pseudoDemandes, pseudo)
			demandesModo = append(demandesModo, demande)
		}

	}
	useRol.Demandes = demandesModo
	useRol.Pseudos = pseudoDemandes
	useRol.PseudoModo = pseudoModos
	useRol.Role = getStatut()
	fmt.Print(useRol.Demandes, useRol.Pseudos)
	templateAdmin.ExecuteTemplate(w, "admin.html", useRol)
}
func upgrade(w http.ResponseWriter, r *http.Request) {

	// fmt.Print(dejaDemandeModo())
	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if r.Method == http.MethodPost {

		state, _ := db.Prepare("CREATE TABLE IF NOT EXISTS contactAdmin (pseudo,demande,role,idMessSignal)")
		state.Exec()
		if roleGlobal != "modo" && !dejaDemandeModo() {
			state, _ = db.Prepare("INSERT INTO contactAdmin(pseudo,demande,role) VALUES (?,?,?)")
			state.Exec(utilisateur, "attente", roleGlobal)

			// w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusConflict)
		}

		return
	}

}

func acceptModo(pseudo string) bool {
	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Vérifiez d'abord si la demande pour ce pseudo est en attente
	rows, err := tx.Query("SELECT pseudo, demande FROM contactAdmin WHERE pseudo = ? AND demande = 'attente'", pseudo)
	if err != nil {
		fmt.Print(err)
		tx.Rollback() // Annuler la transaction en cas d'erreur
		return false
	}
	defer rows.Close()

	if rows.Next() {
		// La demande est en attente, mettez à jour le statut en "approuve"
		_, err := tx.Exec("UPDATE contactAdmin SET demande = 'approuve' WHERE pseudo = ?", pseudo)
		if err != nil {
			fmt.Print(err)
			tx.Rollback() // Annuler la transaction en cas d'erreur
			return false
		}
	} else {
		tx.Rollback() // Annuler la transaction si la demande n'est pas en attente
		return false
	}

	// Vérifiez si le pseudo existe dans la table "people" et mettez à jour le rôle en "modo"
	rowse, err := tx.Query("SELECT pseudo FROM people WHERE pseudo = ?", pseudo)
	if err != nil {
		fmt.Print(err)
		tx.Rollback() // Annuler la transaction en cas d'erreur
		return false
	}
	defer rowse.Close()

	if rowse.Next() {
		// Mettez à jour le rôle en "modo"
		_, err := tx.Exec("UPDATE people SET role = 'modo' WHERE pseudo = ?", pseudo)
		if err != nil {
			fmt.Print(err)
			tx.Rollback() // Annuler la transaction en cas d'erreur
			return false
		}
	} else {
		tx.Rollback() // Annuler la transaction si le pseudo n'existe pas dans "people"
		return false
	}

	// Toutes les mises à jour se sont bien déroulées, validez la transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return true
}
func refusModo(pseudo string) bool {
	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Vérifiez d'abord si la demande pour ce pseudo est en attente
	rows, err := tx.Query("SELECT pseudo, demande FROM contactAdmin WHERE pseudo = ? AND demande = 'attente'", pseudo)
	if err != nil {
		fmt.Print(err)
		tx.Rollback() // Annuler la transaction en cas d'erreur
		return false
	}
	defer rows.Close()

	if rows.Next() {
		// La demande est en attente, mettez à jour le statut en "approuve"
		_, err := tx.Exec("UPDATE contactAdmin SET demande = 'refuse' WHERE pseudo = ?", pseudo)
		if err != nil {
			fmt.Print(err)
			tx.Rollback() // Annuler la transaction en cas d'erreur
			return false
		}
	} else {
		tx.Rollback() // Annuler la transaction si la demande n'est pas en attente
		return false
	}

	// Toutes les mises à jour se sont bien déroulées, validez la transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return true
}

func acceptDemande(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Lire le corps de la requête JSON
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erreur de lecture du corps de la requête", http.StatusBadRequest)
		return
	}

	// Structure pour stocker la valeur du pseudo
	var requestData struct {
		Pseudo string `json:"pseudo"`
	}

	// Analyser le corps JSON
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Erreur d'analyse JSON", http.StatusBadRequest)
		return
	}

	// Maintenant, requestData.Pseudo contient la valeur du pseudo
	pseudo := requestData.Pseudo
	fmt.Print(pseudo)
	acceptModo(pseudo)

	w.WriteHeader(http.StatusOK)
}
func refusDemande(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Lire le corps de la requête JSON
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erreur de lecture du corps de la requête", http.StatusBadRequest)
		return
	}

	// Structure pour stocker la valeur du pseudo
	var requestData struct {
		Pseudo string `json:"pseudo"`
	}

	// Analyser le corps JSON
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Erreur d'analyse JSON", http.StatusBadRequest)
		return
	}

	// Maintenant, requestData.Pseudo contient la valeur du pseudo
	pseudo := requestData.Pseudo
	// fmt.Print(pseudo)
	refusModo(pseudo)

	w.WriteHeader(http.StatusOK)
}

// func getIdMsg(id string) {
// 	db, _ := sql.Open("sqlite3", "./dbForum.db")

// 	rows, err := db.Query("SELECT idCom FROM likesCom")
// 	if err != nil {
// 		fmt.Print(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var id string

// 		err := rows.Scan(&id)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		fmt.Print(id)

// 	}
// }

func supprMessage(id string) bool {
	// Convertir l'ID en int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Erreur de conversion de l'ID en int:", err)
		return false
	}

	// Ouvrir la connexion à la base de données
	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Commencer une transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Supprimer la ligne avec l'ID correspondant
	_, _ = tx.Exec("DELETE FROM messages WHERE id = ?", idInt)
	_, _ = tx.Exec("DELETE FROM comments WHERE id = ?", idInt)
	_, _ = tx.Exec("DELETE FROM likes WHERE id = ?", idInt)
	_, _ = tx.Exec("DELETE FROM notifs WHERE idPost = ?", idInt)
	_, _ = tx.Exec("DELETE FROM likesCom WHERE idPost = ?", idInt)
	// getIdMsg(id)
	//voir pr recup l'id du message ds likeCom

	// if err != nil {
	// 	log.Println("Erreur lors de la suppression du message:", err)
	// 	tx.Rollback() // Annuler la transaction en cas d'erreur
	// 	return false
	// }

	// Toutes les opérations se sont bien déroulées, valider la transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return true
}
func supprMess(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Lire le corps de la requête JSON
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erreur de lecture du corps de la requête", http.StatusBadRequest)
		return
	}

	// Structure pour stocker la valeur du pseudo
	var requestData struct {
		Id string `json:"id"`
	}

	// Analyser le corps JSON
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Erreur d'analyse JSON", http.StatusBadRequest)
		return
	}

	id := requestData.Id
	// fmt.Print(pseudo)
	supprMessage(id)

	w.WriteHeader(http.StatusOK)
}
func apprMessage(id string) bool {
	// Convertir l'ID en int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Erreur de conversion de l'ID en int:", err)
		return false
	}

	// Ouvrir la connexion à la base de données
	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Commencer une transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Supprimer la ligne avec l'ID correspondant
	rows, err := tx.Query("SELECT statut FROM messages WHERE id = ? AND statut = 'attente'", idInt)
	if err != nil {
		fmt.Print(err)
		tx.Rollback() // Annuler la transaction en cas d'erreur
		return false
	}
	defer rows.Close()

	if rows.Next() {
		// La demande est en attente, mettez à jour le statut en "approuve"
		_, err := tx.Exec("UPDATE messages SET statut = 'approuve' WHERE id = ?", idInt)
		if err != nil {
			fmt.Print(err)
			tx.Rollback() // Annuler la transaction en cas d'erreur
			return false
		}
	}
	// Toutes les opérations se sont bien déroulées, valider la transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return true
}
func approuveMess(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Lire le corps de la requête JSON
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erreur de lecture du corps de la requête", http.StatusBadRequest)
		return
	}

	// Structure pour stocker la valeur du pseudo
	var requestData struct {
		Id string `json:"id"`
	}

	// Analyser le corps JSON
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Erreur d'analyse JSON", http.StatusBadRequest)
		return
	}

	id := requestData.Id
	// fmt.Print(pseudo)
	apprMessage(id)

	w.WriteHeader(http.StatusOK)
}
func signMessage(id string) bool {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Erreur de conversion de l'ID en int:", err)
		return false
	}

	// Ouvrir la connexion à la base de données
	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Commencer une transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Supprimer la ligne avec l'ID correspondant
	rows, err := tx.Query("SELECT signalement FROM messages WHERE id = ? ", idInt)
	if err != nil {
		fmt.Print(err)
		tx.Rollback() // Annuler la transaction en cas d'erreur
		return false
	}
	defer rows.Close()

	if rows.Next() {
		// La demande est en attente, mettez à jour le statut en "approuve"
		_, err := tx.Exec("UPDATE messages SET signalement = 'true' WHERE id = ? ", idInt)
		if err != nil {
			fmt.Print(err)
			tx.Rollback() // Annuler la transaction en cas d'erreur
			return false
		}
	}
	// Toutes les opérations se sont bien déroulées, valider la transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return true
}
func retrogade(pseudo string) bool {

	// Ouvrir la connexion à la base de données
	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Commencer une transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	rowse, err := tx.Query("SELECT demande FROM contactAdmin WHERE Pseudo = ? ", pseudo)
	if err != nil {
		fmt.Print(err)
		tx.Rollback() // Annuler la transaction en cas d'erreur
		return false
	}
	defer rowse.Close()

	if rowse.Next() {

		_, err := tx.Exec("UPDATE contactAdmin SET demande = 'refuse' WHERE pseudo = ? ", pseudo)
		if err != nil {
			fmt.Print(err)
			tx.Rollback() // Annuler la transaction en cas d'erreur
			return false
		}
	}
	rows, err := tx.Query("SELECT role FROM people WHERE Pseudo = ? ", pseudo)
	if err != nil {
		fmt.Print(err)
		tx.Rollback() // Annuler la transaction en cas d'erreur
		return false
	}
	defer rows.Close()

	if rows.Next() {

		_, err := tx.Exec("UPDATE people SET role = 'user' WHERE pseudo = ? ", pseudo)
		if err != nil {
			fmt.Print(err)
			tx.Rollback() // Annuler la transaction en cas d'erreur
			return false
		}
	}
	// Toutes les opérations se sont bien déroulées, valider la transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return true
}
func signalerMess(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Lire le corps de la requête JSON
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erreur de lecture du corps de la requête", http.StatusBadRequest)
		return
	}

	// Structure pour stocker la valeur du pseudo
	var requestData struct {
		Id string `json:"id"`
	}

	// Analyser le corps JSON
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Erreur d'analyse JSON", http.StatusBadRequest)
		return
	}

	id := requestData.Id
	// fmt.Print(pseudo)
	signMessage(id)

	w.WriteHeader(http.StatusOK)
}
func retrogader(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Lire le corps de la requête JSON
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erreur de lecture du corps de la requête", http.StatusBadRequest)
		return
	}

	// Structure pour stocker la valeur du pseudo
	var requestData struct {
		Pseudo string `json:"pseudo"`
	}

	// Analyser le corps JSON
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Erreur d'analyse JSON", http.StatusBadRequest)
		return
	}

	pseudos := requestData.Pseudo
	fmt.Print(pseudos)
	retrogade(pseudos)

	w.WriteHeader(http.StatusOK)
}

func dejaUser(email string) bool {
	database, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer database.Close()

	// Utilisation d'une requête SQL préparée pour éviter les problèmes de casse
	query := "SELECT EXISTS (SELECT 1 FROM people WHERE LOWER(Mail) = LOWER(?))"
	var exists bool
	err = database.QueryRow(query, email).Scan(&exists)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return exists
}
func dejaPseudo(pseudo string) bool {
	database, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer database.Close()

	// Utilisation d'une requête SQL préparée pour éviter les problèmes de casse
	query := "SELECT EXISTS (SELECT 1 FROM people WHERE LOWER(Pseudo) = LOWER(?))"
	var exists bool
	err = database.QueryRow(query, pseudo).Scan(&exists)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return exists
}
func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Redirigez vers la page d'authentification Google
	http.Redirect(w, r, oauth2Config.AuthCodeURL("", oauth2.AccessTypeOffline), http.StatusFound)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	use := user{}
	// cookies := r.Cookies()
	userBrowser := r.UserAgent()
	// fmt.Print(userBrowser)
	connecte = false

	t := time.Now()

	date := t.Format("2006-01-02 15:04:05")

	// t := time.Now()

	// date = t.Format("2006-01-02 15:04:05")
	database, _ := sql.Open("sqlite3", "./dbForum.db")

	// Échangez le code d'authentification contre un jeton d'accès
	code := r.URL.Query().Get("code")
	token, err := oauth2Config.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userInfo := UserInfo{}

	// Utilisez un décodeur JSON pour extraire les données de resp.Body

	// Maintenant, userInfo contient les données de l'utilisateur extraites de resp.Body
	// Vous pouvez y accéder comme ceci :

	// Utilisez le token d'accès pour accéder aux ressources protégées
	client := oauth2Config.Client(r.Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if err != nil {
		// Gérez les erreurs
		fmt.Println("Erreur lors de l'échange du code d'autorisation : ", err)
		return
	}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&userInfo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s := strings.Split(userInfo.Email, "@")

	userInfo.Pseudo = s[0]
	fmt.Print(dejaUser(userInfo.Email))
	// fmt.Print(userInfo.Email, userInfo.Pseudo)
	if !dejaUser(userInfo.Email) { //si ya pas le mail deja dans la table on inscrit, sinon on le connecte

		password := []byte(userInfo.ID)

		hash, err := bcrypt.GenerateFromPassword(password, 12)
		if err != nil {
			panic(err)
		}

		// fmt.Fprint(w, erreur.MsgErreur)

		statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, Pseudo TEXT, Mdp TEXT,Mail TEXT,role TEXT)")
		statement.Exec()
		statement, _ = database.Prepare("INSERT INTO people (Pseudo, Mdp,Mail,role) VALUES (?,?,?,?)")
		statement.Exec(userInfo.Pseudo, hash, userInfo.Email, "user") // si on rajoute une nouvelle rubrique il faut suppr le fichier ou creer une nvelle table
		connecte = true
		roleGlobal = "user"
	} else if dejaUser(userInfo.Email) {

		connecte = true

	}
	if connecte {
		utilisateur = userInfo.Pseudo
		cookie := &http.Cookie{
			Name:   utilisateur,
			Value:  date, // faire avec la date
			MaxAge: 10 * 365 * 24 * 60 * 60,
		}
		http.SetCookie(w, cookie)

		statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS cookies (id INTEGER PRIMARY KEY, CookieValue TEXT,Pseudo TEXT,NavigateurAct TEXT,Role TEXT)")
		statement.Exec()

		statement, _ = database.Prepare("INSERT INTO cookies (CookieValue,Pseudo,NavigateurAct,role) VALUES (?,?,?,?)")
		statement.Exec(date, utilisateur, userBrowser, roleGlobal)

		use.User = "Bonjour " + utilisateur
		use.Autorisation = "La connexion a reussi"
		use.Role = roleGlobal

		templateHome.ExecuteTemplate(w, "home.html", use)

	} else {
		connecte = false
		use.User = "invite"
		use.Autorisation = "pseudo ou MDP inconnu"

		templateHome.ExecuteTemplate(w, "home.html", use)

	}
}

func handleLoginGithub(w http.ResponseWriter, r *http.Request) {

	// Créez une URL d'authentification OAuth2
	url := githubOauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)

	// Redirigez l'utilisateur vers cette URL
	http.Redirect(w, r, url, http.StatusFound)
}
func handleCallbackGithub(w http.ResponseWriter, r *http.Request) {
	use := user{}
	// cookies := r.Cookies()
	userBrowser := r.UserAgent()
	// fmt.Print(userBrowser)
	connecte = false

	t := time.Now()

	date := t.Format("2006-01-02 15:04:05")
	var mailGit string
	// // t := time.Now()

	// // date = t.Format("2006-01-02 15:04:05")
	// userInfo := UserInfo{}
	database, _ := sql.Open("sqlite3", "./dbForum.db")
	code := r.URL.Query().Get("code")
	token, err := githubOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := githubOauthConfig.Client(r.Context(), token)

	// Récupérer les informations de l'utilisateur depuis GitHub
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Print("log GitHub: ", userInfo["login"])
	fmt.Print("mail GitHub: ", userInfo["email"])
	pseudoGit, _ := userInfo["login"].(string)
	idGit, _ := userInfo["id"].(string)
	if userInfo["email"] != nil {
		mailGit, _ = userInfo["email"].(string)
	} else {
		mailGit = "nonCommunique"

	}
	if !dejaPseudo(pseudoGit) {

		password := []byte(idGit)

		hash, err := bcrypt.GenerateFromPassword(password, 12)
		if err != nil {
			panic(err)
		}

		// fmt.Fprint(w, erreur.MsgErreur)

		statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, Pseudo TEXT, Mdp TEXT,Mail TEXT,role TEXT)")
		statement.Exec()
		statement, _ = database.Prepare("INSERT INTO people (Pseudo, Mdp,Mail,role) VALUES (?,?,?,?)")
		statement.Exec(pseudoGit, hash, mailGit, "user") // si on rajoute une nouvelle rubrique il faut suppr le fichier ou creer une nvelle table
		connecte = true
		roleGlobal = "user"
	} else if dejaPseudo(pseudoGit) {

		connecte = true

	}
	if connecte {
		utilisateur = pseudoGit
		cookie := &http.Cookie{
			Name:   utilisateur,
			Value:  date, // faire avec la date
			MaxAge: 10 * 365 * 24 * 60 * 60,
		}
		http.SetCookie(w, cookie)

		statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS cookies (id INTEGER PRIMARY KEY, CookieValue TEXT,Pseudo TEXT,NavigateurAct TEXT,Role TEXT)")
		statement.Exec()

		statement, _ = database.Prepare("INSERT INTO cookies (CookieValue,Pseudo,NavigateurAct,role) VALUES (?,?,?,?)")
		statement.Exec(date, utilisateur, userBrowser, roleGlobal)

		use.User = "Bonjour " + utilisateur
		use.Autorisation = "La connexion a reussi"
		use.Role = roleGlobal

		templateHome.ExecuteTemplate(w, "home.html", use)

	} else {
		connecte = false
		use.User = "invite"
		use.Autorisation = "pseudo ou MDP inconnu"

		templateHome.ExecuteTemplate(w, "home.html", use)

	}

	// fmt.Fprintf(w, "Nom d'utilisateur GitHub: %v\n", userInfo["login"])
	// fmt.Fprintf(w, "Email GitHub: %v\n", userInfo["email"])
}

func getOwnMess(utilisateur string) (activity, error) {
	var userMessages activity

	database, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		return userMessages, err
	}
	defer database.Close()

	query := "SELECT id, posts FROM messages WHERE Pseudo = ?"

	rows, err := database.Query(query, utilisateur)
	if err != nil {
		return userMessages, err
	}
	defer rows.Close()

	userMessages.IdPost = nil // Initialiser avec une valeur par défaut

	for rows.Next() {
		var id int
		var message string
		if err := rows.Scan(&id, &message); err != nil {
			return userMessages, err
		}

		userMessages.IdPost = append(userMessages.IdPost, strconv.Itoa(id))

		userMessages.Posts = append(userMessages.Posts, splitMots(message))
	}

	if err := rows.Err(); err != nil {
		return userMessages, err
	}

	return userMessages, nil
}

// func getPostsForIDs(db *sql.DB, idCom []string) ([]string, error) {
// 	// Convertir les IDs de type string en int
// 	var idComInt []int
// 	for _, id := range idCom {
// 		idInt, err := strconv.Atoi(id)
// 		if err != nil {
// 			return nil, err
// 		}
// 		idComInt = append(idComInt, idInt)
// 	}

// 	// Construire une chaîne de placeholders pour les IDs
// 	placeholders := make([]string, len(idComInt))
// 	values := make([]interface{}, len(idComInt))
// 	for i, id := range idComInt {
// 		placeholders[i] = "?"
// 		values[i] = id
// 	}

// 	// Construire la requête SQL avec les placeholders
// 	query := fmt.Sprintf("SELECT Posts FROM messages WHERE id IN (%s)", strings.Join(placeholders, ","))

// 	// Exécuter la requête SQL avec les IDs
// 	rows, err := db.Query(query, values...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var posts []string

// 	for rows.Next() {
// 		var post string

// 		if err := rows.Scan(&post); err != nil {
// 			return nil, err
// 		}

// 		posts = append(posts, post)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

//		return posts, nil
//	}
func getPostsForIDs(db *sql.DB, idCom []string) ([]string, error) {
	// Convertir les IDs de type string en int
	var postes []string
	var idComInt []int
	for _, id := range idCom {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		idComInt = append(idComInt, idInt)
	}
	database, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		return nil, err
	}
	defer database.Close()

	query := "SELECT Posts FROM messages WHERE id = ?"

	for i := 0; i < len(idComInt); i++ {

		rows, err := database.Query(query, idComInt[i])
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var posts string

			if err := rows.Scan(&posts); err != nil {
				return nil, err
			}

			postes = append(postes, splitMots(posts))

		}
	}
	return postes, nil
}

func getOwnComment(utilisateur string) (activity, error) {
	var userMessages activity

	database, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		return userMessages, err
	}
	defer database.Close()

	query := "SELECT id FROM comments WHERE Pseudo = ?"

	rows, err := database.Query(query, utilisateur)
	if err != nil {
		return userMessages, err
	}
	defer rows.Close()

	userMessages.IdCom = nil // Initialiser avec une valeur par défaut

	for rows.Next() {
		var id string

		if err := rows.Scan(&id); err != nil {
			return userMessages, err
		}

		userMessages.IdCom = append(userMessages.IdCom, id)

	}

	if err := rows.Err(); err != nil {
		return userMessages, err
	}

	return userMessages, nil
}
func getOwnLike(utilisateur string) (activity, error) {
	var userMessages activity

	database, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		return userMessages, err
	}
	defer database.Close()

	query := "SELECT id FROM likes WHERE Pseudo = ?"

	rows, err := database.Query(query, utilisateur)
	if err != nil {
		return userMessages, err
	}
	defer rows.Close()

	userMessages.IdLike = nil // Initialiser avec une valeur par défaut

	for rows.Next() {
		var id string

		if err := rows.Scan(&id); err != nil {
			return userMessages, err
		}

		userMessages.IdLike = append(userMessages.IdLike, id)

	}

	if err := rows.Err(); err != nil {
		return userMessages, err
	}

	return userMessages, nil
}
func getOwnLikeCom(utilisateur string) (activity, error) {
	var userMessages activity

	database, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		return userMessages, err
	}
	defer database.Close()

	query := "SELECT idCom FROM likesCom WHERE PseudoCom = ?"

	rows, err := database.Query(query, utilisateur)
	if err != nil {
		return userMessages, err
	}
	defer rows.Close()

	userMessages.IdLikeCom = nil // Initialiser avec une valeur par défaut

	for rows.Next() {
		var id string

		if err := rows.Scan(&id); err != nil {
			return userMessages, err
		}

		userMessages.IdLikeCom = append(userMessages.IdLikeCom, id)

		// fmt.Print(id)
	}
	// fmt.Print(userMessages.IdLikeCom)
	if err := rows.Err(); err != nil {
		return userMessages, err
	}

	return userMessages, nil
}
func activite(w http.ResponseWriter, r *http.Request) {
	database, _ := sql.Open("sqlite3", "./dbForum.db")
	activites := activity{}
	userMessages, _ := getOwnMess(utilisateur)
	userCom, _ := getOwnComment(utilisateur)
	userLike, _ := getOwnLike(utilisateur)
	userLikeCom, _ := getOwnLikeCom(utilisateur)

	activites.User = utilisateur
	activites.Role = roleGlobal

	activites.Posts = userMessages.Posts
	activites.IdPost = userMessages.IdPost
	activites.IdCom = userCom.IdCom
	activites.Comments, _ = getPostsForIDs(database, userCom.IdCom)
	activites.IdLike = userLike.IdLike
	activites.IdLikeCom = userLikeCom.IdLikeCom
	activites.PostsLike, _ = getPostsForIDs(database, userLike.IdLike)
	activites.PostsLikeCom, _ = getPostsForIDs(database, userLikeCom.IdLikeCom)

	// fmt.Print(userLikeCom)
	// fmt.Print(activites.Comments[1])

	// fmt.Print(activites.Posts, activites.Id)
	templateActivite.ExecuteTemplate(w, "activite.html", activites)
}

func supprOwnComment(id string) bool {
	// Convertir l'ID en int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Erreur de conversion de l'ID en int:", err)
		return false
	}

	// Ouvrir la connexion à la base de données
	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Commencer une transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Supprimer la ligne avec l'ID correspondant
	_, err = tx.Exec("DELETE FROM comments WHERE idComment = ?", idInt)
	if err != nil {
		log.Println("Erreur lors de la suppression du message:", err)
		tx.Rollback() // Annuler la transaction en cas d'erreur
		return false
	}

	// Toutes les opérations se sont bien déroulées, valider la transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return true
}
func supprOwnPost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Lire le corps de la requête JSON
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erreur de lecture du corps de la requête", http.StatusBadRequest)
		return
	}

	// Structure pour stocker la valeur du pseudo
	var requestData struct {
		Id string `json:"id"`
	}

	// Analyser le corps JSON
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Erreur d'analyse JSON", http.StatusBadRequest)
		return
	}

	// Maintenant, requestData.Pseudo contient la valeur du pseudo
	id := requestData.Id
	supprMessage(id)
	// acceptModo(pseudo)

	w.WriteHeader(http.StatusOK)
}
func supprOwnCom(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Lire le corps de la requête JSON
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erreur de lecture du corps de la requête", http.StatusBadRequest)
		return
	}

	// Structure pour stocker la valeur du pseudo
	var requestData struct {
		Id string `json:"id"`
	}

	// Analyser le corps JSON
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Erreur d'analyse JSON", http.StatusBadRequest)
		return
	}

	// Maintenant, requestData.Pseudo contient la valeur du pseudo
	id := requestData.Id

	fmt.Print(id)
	supprOwnComment(id)

	w.WriteHeader(http.StatusOK)
}
func countNotifs(utilisateur string) int {
	// Vous devrez utiliser votre bibliothèque/ORM de base de données ici pour interagir avec la table "notifs".
	// L'exemple suivant suppose que vous utilisez SQL avec la bibliothèque "database/sql".

	// Ouvrez une connexion à votre base de données (à adapter à votre configuration).
	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Préparez une requête SQL pour compter les notifications non lues.
	query := "SELECT COUNT(*) FROM notifs WHERE read = false AND Createur = ? AND Createur != Pseudo"

	// Exécutez la requête SQL en utilisant l'utilisateur spécifié comme paramètre.
	var count int
	err = db.QueryRow(query, utilisateur).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	// Retournez le nombre de notifications non lues.
	return count
}

func editMessage(id string, texte string) bool {
	idInt, _ := strconv.Atoi(id)

	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Commencer une transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback() // Rollback en cas d'erreur

	// Exécuter la requête de mise à jour
	_, err = tx.Exec("UPDATE messages SET Posts = ? WHERE id = ?", texte, idInt)
	if err != nil {
		log.Print(err)
		return false
	}

	// Toutes les opérations se sont bien déroulées, valider la transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return true
}
func editMess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Lire le corps de la requête JSON
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erreur de lecture du corps de la requête", http.StatusBadRequest)
		return
	}

	// Structure pour stocker la valeur du pseudo
	var requestData struct {
		Id   string `json:"idPost"`
		Text string `json:"updatedText"` //idCom, updatedText
	}

	// Analyser le corps JSON
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Erreur d'analyse JSON", http.StatusBadRequest)
		return
	}

	// Maintenant, requestData.Pseudo contient la valeur du pseudo
	id := requestData.Id
	texte := requestData.Text
	editMessage(id, texte)
	fmt.Print(id, texte)

	w.WriteHeader(http.StatusOK)
}
func editComment(id string, texte string) bool {
	idInt, _ := strconv.Atoi(id)

	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Commencer une transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback() // Rollback en cas d'erreur

	// Exécuter la requête de mise à jour
	_, err = tx.Exec("UPDATE comments SET Commentaires = ? WHERE idComment = ?", texte, idInt)
	if err != nil {
		log.Print(err)
		return false
	}

	// Toutes les opérations se sont bien déroulées, valider la transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return true
}
func editCom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Lire le corps de la requête JSON
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erreur de lecture du corps de la requête", http.StatusBadRequest)
		return
	}

	// Structure pour stocker la valeur du pseudo
	var requestData struct {
		Id   string `json:"idCom"`
		Text string `json:"updatedText"` //idCom, updatedText
	}

	// Analyser le corps JSON
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Erreur d'analyse JSON", http.StatusBadRequest)
		return
	}

	// Maintenant, requestData.Pseudo contient la valeur du pseudo
	id := requestData.Id
	texte := requestData.Text
	editComment(id, texte)
	fmt.Print(id, texte)

	w.WriteHeader(http.StatusOK)
}
func getNotifs(utilisateur string) ([]Notif, error) {
	database, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		return nil, err
	}
	defer database.Close()

	query := "SELECT idNotif, idPost, idCom, Pseudo, read, Createur, Nature FROM notifs WHERE Createur = ?"
	rows, err := database.Query(query, utilisateur)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifs []Notif

	for rows.Next() {
		var notif Notif
		err := rows.Scan(&notif.IDNotif, &notif.IDPost, &notif.IDCom, &notif.Pseudo, &notif.Read, &notif.Createur, &notif.Nature)
		if err != nil {
			return nil, err
		}
		notifs = append(notifs, notif)
	}

	return notifs, nil
}

func affNotifs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	notifs, err := getNotifs(utilisateur)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des notifications", http.StatusInternalServerError)
		return
	}

	// Convertissez les données en JSON
	notificationsJSON, err := json.Marshal(notifs)
	if err != nil {
		http.Error(w, "Erreur de sérialisation JSON", http.StatusInternalServerError)
		return
	}

	// Définissez le type de contenu de la réponse comme JSON
	w.Header().Set("Content-Type", "application/json")

	// Écrivez les données JSON dans la réponse
	_, err = w.Write(notificationsJSON)
	if err != nil {
		http.Error(w, "Erreur lors de l'écriture de la réponse", http.StatusInternalServerError)
		return
	}
}
func supprNotification(idNotif int) bool {

	// Ouvrir la connexion à la base de données
	db, err := sql.Open("sqlite3", "./dbForum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Commencer une transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Supprimer la ligne avec l'ID correspondant
	_, _ = tx.Exec("DELETE FROM  notifs WHERE idNotif = ?", idNotif)

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return true

}
func supprNotifs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Lire le corps de la requête JSON
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erreur de lecture du corps de la requête", http.StatusBadRequest)
		return
	}

	// Structure pour stocker l'ID de la notification
	var requestData struct {
		ID int `json:"id"`
	}

	// Analyser le corps JSON
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Erreur d'analyse JSON", http.StatusBadRequest)
		return
	}

	// Récupérez l'ID de la notification
	idNotif := requestData.ID
	fmt.Print(idNotif)
	supprNotification(idNotif)

	w.WriteHeader(http.StatusOK)

}

func main() {

	http.HandleFunc("/home", home)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/maintenance", maintenance)
	http.HandleFunc("/activite", activite)
	http.HandleFunc("/upgrade", upgrade)
	http.HandleFunc("/accepterDemande", acceptDemande)
	http.HandleFunc("/refuserDemande", refusDemande)
	http.HandleFunc("/supprMess", supprMess)
	http.HandleFunc("/supprOwnPost", supprOwnPost)
	http.HandleFunc("/supprOwnCom", supprOwnCom)
	http.HandleFunc("/editMess", editMess)
	http.HandleFunc("/editCom", editCom)
	http.HandleFunc("/supprNotif", supprNotifs)

	http.HandleFunc("/affNotifs", affNotifs)

	http.HandleFunc("/approuveMess", approuveMess)
	http.HandleFunc("/signalerMess", signalerMess)
	http.HandleFunc("/retrogader", retrogader)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/logine", handleCallback)
	http.HandleFunc("/loginGit", handleLoginGithub)
	http.HandleFunc("/loginGithub", handleCallbackGithub)

	http.HandleFunc("/inscription", inscription)
	http.HandleFunc("/inscrit", insertInscription)
	http.HandleFunc("/connexion", connexion)
	http.HandleFunc("/deconnexion", deconnexion)

	http.HandleFunc("/formMessage", formMessage)

	http.HandleFunc("/message", ecrireMessage)
	http.HandleFunc("/messagesInd", reagirMess)
	http.HandleFunc("/commenter", commenter)
	http.HandleFunc("/like", like)
	http.HandleFunc("/likeCom", likeCom)

	http.HandleFunc("/disLike", disLike)
	http.HandleFunc("/disLikeCom", disLikeCom)
	http.HandleFunc("/recherche", recherche)
	// http.HandleFunc("/page", page)

	// http.HandleFunc("/readcookie", ReadCookie)
	// http.HandleFunc("/deletecookie", DeleteCookie)
	// http.HandleFunc("/createCookie", CreateCookie)
	fmt.Println("http://localhost:5555/home")
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css/"))))
	// http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("templates/js/"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("templates/images/"))))
	http.Handle("/", http.FileServer(http.Dir("./templates")))
	// http.HandleFunc("/js/messageInd.js", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./templates/js/messageInd.js")
	// })

	http.HandleFunc("/js/message.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/js/message.js")
	})

	http.ListenAndServe(port, nil)
	log.Fatalln(http.ListenAndServe(port, nil))
}
