const lieux = document.getElementById("lieu");
const groupeLieux = document.getElementById("groupeLieu");
const groupeMembre = document.getElementById("groupeMembre");
const membres = document.getElementById("membre");
const groupeNoms = document.getElementById("groupeNom");
const dateCreation = document.getElementById("dateCreation");
const first = document.getElementById("first");
const compo = document.getElementById("compo");
const searchUser = document.getElementById("search");

var artiste = [];
var date = [];
var lieu = [];

function executerRequeteArtiste(callback) {
  // on vérifie si le artiste a déjà été chargé pour n'exécuter la requête AJAX
  // qu'une seule fois<form id = "" action="/search" method="post" >

  if (artiste.length === 0) {
    // on récupère un objet XMLHttpRequest
    var xhr = getXMLHttpRequest();
    // on réagit à l'événement onreadystatechange
    xhr.onreadystatechange = function () {
      // test du statut de retour de la requête AJAX
      if (xhr.readyState == 4 && (xhr.status == 200 || xhr.status == 0)) {
        // on désérialise le artiste et on le sauvegarde dans une variable
        artiste = JSON.parse(xhr.responseText);
        // on lance la fonction de callback avec le artiste récupéré
        callback();
      }
    };
    // la requête AJAX : lecture de data.json
    xhr.open(
      "GET",
      "/proxy?url=" +
        encodeURIComponent("https://groupietrackers.herokuapp.com/api/artists"),
      true
    );
    xhr.send();
  } else {
    // on lance la fonction de callback avec le artiste déjà récupéré précédemment
    callback();
  }
}
function executerRequeteDate(callback) {
  // on vérifie si le artiste a déjà été chargé pour n'exécuter la requête AJAX
  // qu'une seule fois
  if (date.length === 0) {
    // on récupère un objet XMLHttpRequest
    var xhr1 = getXMLHttpRequest();
    // on réagit à l'événement onreadystatechange
    xhr1.onreadystatechange = function () {
      // test du statut de retour de la requête AJAX
      if (xhr1.readyState == 4 && (xhr1.status == 200 || xhr1.status == 0)) {
        // on désérialise le artiste et on le sauvegarde dans une variable
        date = JSON.parse(xhr1.responseText);
        // on lance la fonction de callback avec le artiste récupéré
        callback();
      }
    };
    // la requête AJAX : lecture de data.json

    xhr1.open(
      "GET",
      "/proxy?url=" +
        encodeURIComponent("https://groupietrackers.herokuapp.com/api/dates"),
      true
    );
    xhr1.send();
  } else {
    // on lance la fonction de callback avec le artiste déjà récupéré précédemment
    callback();
  }
}
function executerRequeteLieu(callback) {
  // on vérifie si le artiste a déjà été chargé pour n'exécuter la requête AJAX
  // qu'une seule fois
  if (lieu.length === 0) {
    // on récupère un objet XMLHttpRequest
    var xhr2 = getXMLHttpRequest();
    // on réagit à l'événement onreadystatechange
    xhr2.onreadystatechange = function () {
      // test du statut de retour de la requête AJAX
      if (xhr2.readyState == 4 && (xhr2.status == 200 || xhr2.status == 0)) {
        // on désérialise le artiste et on le sauvegarde dans une variable
        lieu = JSON.parse(xhr2.responseText);
        // on lance la fonction de callback avec le artiste récupéré
        callback();
      }
    };
    // la requête AJAX : lecture de data.json
    xhr2.open(
      "GET",
      "/proxy?url=" +
        encodeURIComponent(
          "https://groupietrackers.herokuapp.com/api/locations"
        ),
      true
    );
    xhr2.send();
  } else {
    // on lance la fonction de callback avec le artiste déjà récupéré précédemment
    callback();
  }
}
function afficheArtiste() {
  groupeNom = [];
  composition = [];
  searchUser.addEventListener("input", (e) => {
    element = e.target.value.toLowerCase();
    taille = e.target.value.length;
    for (i = 0; i < artiste.length; i++) {
      inclArtist = artiste[i].name.toLowerCase().includes(element);

      if (inclArtist) {
        groupeNom.push(artiste[i].name);
        composition.push(artiste[i].members);
        var trigroupeNom = [...new Set(groupeNom)];
        var tricompo = [...new Set(composition)];
        groupeNoms.innerHTML =
          " GROUPE PRINCIPAL (ou secondaire) : " + trigroupeNom;
        compo.innerHTML = "COMPOSITION : " + tricompo;
        if (taille % 2 == 0) {
          groupeNom = [];
          composition = [];
        }
        // groupeMembre.innerHTML = "";
        membres.innerHTML = "";
        lieux.innerHTML = "";
        groupeLieux.innerHTML = "";
      }
      if (element === "") {
        groupeMembre.innerHTML = "";
        membres.innerHTML = "";
        groupeNoms.innerHTML = "";
        compo.innerHTML = "";
        groupeNom = [];
        composition = [];
      }
    }
  });
}
function afficheMembre() {
  membreGroupe = [];
  membre = [];
  searchUser.addEventListener("input", (e) => {
    element = e.target.value.toLowerCase();
    taille = e.target.value.length;
    for (i = 0; i < artiste.length; i++) {
      for (j = 0; j < artiste[i].members.length; j++) {
        inclArtist = artiste[i].members[j].toLowerCase().includes(element);

        if (inclArtist) {
          membreGroupe.push(artiste[i].name);
          membre.push(artiste[i].members[j]);

          var triMembreGroupe = [...new Set(membreGroupe)];
          var triMembre = [...new Set(membre)];
          groupeMembre.innerHTML = "MEMBRE DU GROUPE : " + triMembreGroupe;
          membres.innerHTML = "NOM DU MEMBRE : " + triMembre;
          if (taille == 4 || taille == 6 || taille == 8) {
            membreGroupe = [];
            membre = [];
          }
          lieux.innerHTML = "";
          groupeLieux.innerHTML = "";
          compo.innerHTML = "";
        }
        if (element === "") {
          groupeMembre.innerHTML = "";
          membres.innerHTML = "";
          membre = [];
          membreGroupe = [];
        }
      }
    }
  });
}

function afficheLieu() {
  lieuConcert = [];
  groupeLieu = [];

  searchUser.addEventListener("input", (e) => {
    element = e.target.value.toLowerCase();
    taille = e.target.value.length;
    for (i = 0; i < lieu.index.length; i++) {
      for (j = 0; j < lieu.index[i].locations.length; j++) {
        inclLieu = lieu.index[i].locations[j].toLowerCase().includes(element);
        if (inclLieu) {
          lieuConcert.push(lieu.index[i].locations[j]);

          groupeLieu.push(artiste[i].name);

          var triGroupeLieu = [...new Set(groupeLieu)];
          var triLieu = [...new Set(lieuConcert)];
          groupeLieux.innerHTML =
            " GROUPE / LIEUX DE CONCERTS : " + triGroupeLieu;
          lieux.innerHTML = triLieu + " (lieu) ";
          if (taille % 2 == 0) {
            lieuConcert = [];
            groupeLieu = [];
          }
          membres.innerHTML = "";
          groupeNoms.innerHTML = "";
          groupeMembre.innerHTML = "";
          compo.innerHTML = "";
        }
        if (element === "") {
          lieuConcert = [];
          groupeLieu = [];
          lieux.innerHTML = "";
          groupeLieux.innerHTML = "";
        }
      }
    }
  });
}
function afficheCreationDate() {
  groupeDate = [];
  searchUser.addEventListener("input", (e) => {
    const element = e.target.value;
    const taille = e.target.value.length;
    for (i = 0; i < artiste.length; i++) {
      //inclArtist = artiste[i].creationDate.includes(element);

      if (element == artiste[i].creationDate) {
        groupeDate.push(artiste[i].name);

        var trigroupeDate = [...new Set(groupeDate)];

        dateCreation.innerHTML = "DATE DE CREATION : " + trigroupeDate;

        lieux.innerHTML = "";
        groupeLieux.innerHTML = "";
        groupeMembre.innerHTML = "";
        membres.innerHTML = "";
        groupeNoms.innerHTML = "";
      }
      if (element === "") {
        dateCreation.innerHTML = "";
        groupeDate = [];
      }
    }
  });
}
function afficheFirstAlbum() {
  Album = [];
  searchUser.addEventListener("input", (e) => {
    const element = e.target.value.toLowerCase();
    const taille = e.target.value.length;
    for (i = 0; i < artiste.length; i++) {
      inclfirst = artiste[i].firstAlbum.toLowerCase().includes(element);

      if (inclfirst && taille == 10) {
        Album.push(artiste[i].name);

        var trigroupeAlbum = [...new Set(Album)];

        first.innerHTML = "DATE 1 er ALBUM DE : " + trigroupeAlbum;

        lieux.innerHTML = "";
        groupeLieux.innerHTML = "";
        groupeMembre.innerHTML = "";
        membres.innerHTML = "";
        groupeNoms.innerHTML = "";
        dateCreation.innerHTML = "";
      }
      if (element === "") {
        first.innerHTML = "";
        Album = [];
      }
    }
  });
}

executerRequeteArtiste(afficheMembre);
executerRequeteLieu(afficheLieu);
executerRequeteArtiste(afficheArtiste);
executerRequeteArtiste(afficheCreationDate);
executerRequeteArtiste(afficheFirstAlbum);
