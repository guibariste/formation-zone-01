const creaCurseurMoins = document.getElementById("creaRange-");
const creaCurseurPlus = document.getElementById("creaRange+");
const creaAffCurseurPlus = document.getElementById("creaAffRange+");
const creaAffCurseurMoins = document.getElementById("creaAffRange-");
const firstCurseurMoins = document.getElementById("firstRange-");
const firstCurseurPlus = document.getElementById("firstRange+");
const firstAffCurseurPlus = document.getElementById("firstAffRange+");
const firstAffCurseurMoins = document.getElementById("firstAffRange-");
const bypassCrea = document.getElementById("bypassCrea");
const bypassVille = document.getElementById("bypassVille");
const button = document.getElementById("bouton");
const reset = document.getElementById("reset");
const bypassFirst = document.getElementById("bypassFirst");
const bypassMembres = document.getElementById("bypassMembres");
const aff = document.getElementById("aff");
var mbr1 = document.getElementById("mbr1");
var mbr2 = document.getElementById("mbr2");
var mbr3 = document.getElementById("mbr3");
var mbr4 = document.getElementById("mbr4");
var mbr5 = document.getElementById("mbr5");
var mbr6 = document.getElementById("mbr6");
var mbr7 = document.getElementById("mbr7");
var mbr8 = document.getElementById("mbr8");

var artiste = [];
var ville = [];

function executerRequeteArtiste(callback) {
  if (artiste.length === 0) {
    var xhr = getXMLHttpRequest();

    xhr.onreadystatechange = function () {
      if (xhr.readyState == 4 && (xhr.status == 200 || xhr.status == 0)) {
        artiste = JSON.parse(xhr.responseText);

        callback();
      }
    };

    xhr.open(
      "GET",
      "/proxy?url=" +
        encodeURIComponent("https://groupietrackers.herokuapp.com/api/artists"),
      true
    );
    xhr.send();
  } else {
    callback();
  }
}
function executerRequeteVille(callbac) {
  if (ville.length === 0) {
    var xhr1 = getXMLHttpRequest();

    xhr1.onreadystatechange = function () {
      if (xhr1.readyState == 4 && (xhr1.status == 200 || xhr1.status == 0)) {
        ville = JSON.parse(xhr1.responseText);

        callbac();
      }
    };

    xhr1.open(
      "GET",
      "/proxy?url=" +
        encodeURIComponent(
          "https://groupietrackers.herokuapp.com/api/locations"
        ),
      true
    );
    xhr1.send();
  } else {
    callbac();
  }
}
function curseurCrea() {
  creaAffCurseurMoins.innerHTML = creaCurseurMoins.value;
  creaAffCurseurPlus.innerHTML = creaCurseurPlus.value;
  creaCurseurMoins.oninput = function () {
    creaAffCurseurMoins.innerHTML = this.value;
  };

  creaCurseurPlus.oninput = function () {
    creaAffCurseurPlus.innerHTML = this.value;
  };
}
function curseurFirst() {
  firstAffCurseurMoins.innerHTML = firstCurseurMoins.value;
  firstAffCurseurPlus.innerHTML = firstCurseurPlus.value;
  firstCurseurMoins.oninput = function () {
    firstAffCurseurMoins.innerHTML = this.value;
  };

  firstCurseurPlus.oninput = function () {
    firstAffCurseurPlus.innerHTML = this.value;
  };
}
function affiche() {
  var select = document.getElementById("selectVille");
  curseurCrea();
  curseurFirst();
  // if (bypassVille.checked === true) {
  bypassVille.onchange = function () {
    if (bypassVille.checked === true) {
      for (i = 0; i < ville.index.length; i++) {
        for (j = 0; j <= ville.index[i].locations.length; j++) {
          var opt = ville.index[i].locations[j];

          select.innerHTML +=
            '<option value="' + opt + '">' + opt + "</option>";
        }
      }
    }
  };
  button.addEventListener("click", () => {
    if (
      bypassCrea.checked === true &&
      bypassMembres.checked === false &&
      bypassFirst.checked === false &&
      bypassVille.checked === false
    ) {
      triCreationDate();
    }
    if (
      bypassMembres.checked === true &&
      bypassCrea.checked === false &&
      bypassFirst.checked === false &&
      bypassVille.checked === false
    ) {
      selectMembres();
    }

    if (
      bypassFirst.checked === true &&
      bypassMembres.checked === false &&
      bypassCrea.checked === false &&
      bypassVille.checked === false
    ) {
      triFirstAlbum();
    }
    if (bypassCrea.checked === true && bypassMembres.checked === true) {
      combCreaMemb();
    }
    if (
      bypassVille.checked === true &&
      bypassCrea.checked === false &&
      bypassMembres.checked === false &&
      bypassFirst.checked === false
    ) {
    }
    if (
      bypassVille.checked === true &&
      bypassCrea.checked === false &&
      bypassMembres.checked === false &&
      bypassFirst.checked === false
    ) {
      concert();
    }
    if (
      bypassVille.checked === true &&
      bypassCrea.checked === false &&
      bypassMembres.checked === true &&
      bypassFirst.checked === false
    ) {
      combConcMemb();
    }
    if (
      bypassCrea.checked === true &&
      bypassFirst.checked === true &&
      bypassMembres.checked === false &&
      bypassVille.checked === false
    ) {
      combCreaFirst();
    }
    if (
      bypassFirst.checked === true &&
      bypassMembres.checked === true &&
      bypassCrea.checked === false &&
      bypassVille.checked === false
    ) {
      combFirstMemb();
    }
  });

  reset.onclick = function () {
    aff.innerHTML = "";
  };
}
function triCreationDate() {
  for (i = 0; i < artiste.length; i++) {
    if (
      artiste[i].creationDate >= creaCurseurMoins.value &&
      artiste[i].creationDate <= creaCurseurPlus.value
    ) {
      // console.log(artiste[i].name);
      aff.innerHTML += artiste[i].name + "<br>";
    }
  }
}
function concert() {
  var essai = [];
  var select = document.getElementById("selectVille");
  var valeur = select.options[select.selectedIndex].value;
  for (i = 0; i < ville.index.length; i++) {
    for (i = 0; i < artiste.length; i++) {
      if (ville.index[i].locations.includes(valeur)) {
        essai.push(artiste[i].name);

        triEssai = [...new Set(essai)];

        aff.innerHTML = triEssai;

        select.onchange = function () {
          essai = [];
        };
      }
    }
  }
}

function triFirstAlbum() {
  for (i = 0; i < artiste.length; i++) {
    var art = artiste[i].firstAlbum.split("-");
    if (
      art[2] >= firstCurseurMoins.value && // convertir firstalbum en annee
      art[2] <= firstCurseurPlus.value
    ) {
      aff.innerHTML += artiste[i].name + "<br>";
    }
  }

  // console.log(art[2]);
}

function selectMembres() {
  for (i = 0; i < artiste.length; i++) {
    long1 = artiste[i].members.length === 1;
    long2 = artiste[i].members.length === 2;
    long3 = artiste[i].members.length == 3;
    long4 = artiste[i].members.length == 4;
    long5 = artiste[i].members.length == 5;
    long6 = artiste[i].members.length == 6;
    long7 = artiste[i].members.length == 7;
    long8 = artiste[i].members.length == 8;
    if (mbr1.checked && long1) {
      aff.innerHTML += artiste[i].name + "<br>";
    }
    if (mbr2.checked && long2) {
      aff.innerHTML += artiste[i].name + "<br>";
    }
    if (mbr3.checked && long3) {
      aff.innerHTML += artiste[i].name + "<br>";
    }
    if (mbr4.checked && long4) {
      aff.innerHTML += artiste[i].name + "<br>";
    }

    if (mbr5.checked && long5) {
      aff.innerHTML += artiste[i].name + "<br>";
    }
    if (mbr6.checked && long6) {
      aff.innerHTML += artiste[i].name + "<br>";
    }
    if (mbr7.checked && long7) {
      aff.innerHTML += artiste[i].name + "<br>";
    }
    if (mbr8.checked && long8) {
      aff.innerHTML += artiste[i].name + "<br>";
    }
  }
}
function combCreaMemb() {
  for (i = 0; i < artiste.length; i++) {
    long1 = artiste[i].members.length === 1;
    long2 = artiste[i].members.length === 2;
    long3 = artiste[i].members.length === 3;
    long4 = artiste[i].members.length === 4;
    long5 = artiste[i].members.length === 5;
    long6 = artiste[i].members.length === 6;
    long7 = artiste[i].members.length === 7;
    long8 = artiste[i].members.length === 8;
    if (
      artiste[i].creationDate >= creaCurseurMoins.value &&
      artiste[i].creationDate <= creaCurseurPlus.value &&
      long1 &&
      mbr1.checked
    ) {
      aff.innerHTML += artiste[i].name + "<br>";
    }
    if (
      artiste[i].creationDate >= creaCurseurMoins.value &&
      artiste[i].creationDate <= creaCurseurPlus.value &&
      long2 &&
      mbr2.checked
    ) {
      aff.innerHTML += artiste[i].name + "<br>";
    }
    if (
      artiste[i].creationDate >= creaCurseurMoins.value &&
      artiste[i].creationDate <= creaCurseurPlus.value &&
      long3 &&
      mbr3.checked
    ) {
      aff.innerHTML += artiste[i].name + "<br>";
    }
    if (
      artiste[i].creationDate >= creaCurseurMoins.value &&
      artiste[i].creationDate <= creaCurseurPlus.value &&
      long4 &&
      mbr4.checked
    ) {
      aff.innerHTML += artiste[i].name + "<br>";
    }
    if (
      artiste[i].creationDate >= creaCurseurMoins.value &&
      artiste[i].creationDate <= creaCurseurPlus.value &&
      long5 &&
      mbr5.checked
    ) {
      aff.innerHTML += artiste[i].name + "<br>";
    }
    if (
      artiste[i].creationDate >= creaCurseurMoins.value &&
      artiste[i].creationDate <= creaCurseurPlus.value &&
      long6 &&
      mbr6.checked
    ) {
      aff.innerHTML += artiste[i].name + "<br>";
    }
    if (
      artiste[i].creationDate >= creaCurseurMoins.value &&
      artiste[i].creationDate <= creaCurseurPlus.value &&
      long7 &&
      mbr7.checked
    ) {
      aff.innerHTML += artiste[i].name + "<br>";
    }
    if (
      artiste[i].creationDate >= creaCurseurMoins.value &&
      artiste[i].creationDate <= creaCurseurPlus.value &&
      long8 &&
      mbr8.checked
    ) {
      aff.innerHTML += artiste[i].name + "<br>";
    }
  }
}
function combConcMemb() {
  var select = document.getElementById("selectVille");

  var essai = [];

  var valeur = select.options[select.selectedIndex].value;
  for (i = 0; i < artiste.length; i++) {
    long1 = artiste[i].members.length === 1;
    long2 = artiste[i].members.length === 2;
    long3 = artiste[i].members.length === 3;
    long4 = artiste[i].members.length === 4;
    long5 = artiste[i].members.length === 5;
    long6 = artiste[i].members.length === 6;
    long7 = artiste[i].members.length === 7;
    long8 = artiste[i].members.length === 8;
    if (
      (ville.index[i].locations.includes(valeur) &&
        long1 &&
        mbr1.checked === true) ||
      (ville.index[i].locations.includes(valeur) &&
        long2 &&
        mbr2.checked === true) ||
      (ville.index[i].locations.includes(valeur) &&
        long3 &&
        mbr3.checked === true) ||
      (ville.index[i].locations.includes(valeur) &&
        long4 &&
        mbr4.checked === true) ||
      (ville.index[i].locations.includes(valeur) &&
        long5 &&
        mbr5.checked === true) ||
      (ville.index[i].locations.includes(valeur) &&
        long6 &&
        mbr6.checked === true) ||
      (ville.index[i].locations.includes(valeur) &&
        long7 &&
        mbr7.checked === true) ||
      (ville.index[i].locations.includes(valeur) &&
        long8 &&
        mbr8.checked === true)
    ) {
      essai.push(artiste[i].name);

      triEssai = [...new Set(essai)];

      aff.innerHTML = triEssai;

      select.onchange = function () {
        essai = [];
      };
    }
  }
}
function combCreaFirst() {
  for (i = 0; i < artiste.length; i++) {
    var art = artiste[i].firstAlbum.split("-");
    if (
      artiste[i].creationDate >= creaCurseurMoins.value &&
      artiste[i].creationDate <= creaCurseurPlus.value &&
      art[2] >= firstCurseurMoins.value &&
      art[2] <= firstCurseurPlus.value
    ) {
      aff.innerHTML += artiste[i].name + "<br>";
    }
  }
}
function combFirstMemb() {
  for (i = 0; i < artiste.length; i++) {
    var art = artiste[i].firstAlbum.split("-");
    long1 = artiste[i].members.length === 1;
    long2 = artiste[i].members.length === 2;
    long3 = artiste[i].members.length === 3;
    long4 = artiste[i].members.length === 4;
    long5 = artiste[i].members.length === 5;
    long6 = artiste[i].members.length === 6;
    long7 = artiste[i].members.length === 7;
    long8 = artiste[i].members.length === 8;

    if (
      (art[2] >= firstCurseurMoins.value &&
        art[2] <= firstCurseurPlus.value &&
        long1 &&
        mbr1.checked) ||
      (art[2] >= firstCurseurMoins.value &&
        art[2] <= firstCurseurPlus.value &&
        long2 &&
        mbr2.checked) ||
      (art[2] >= firstCurseurMoins.value &&
        art[2] <= firstCurseurPlus.value &&
        long3 &&
        mbr3.checked) ||
      (art[2] >= firstCurseurMoins.value &&
        art[2] <= firstCurseurPlus.value &&
        long4 &&
        mbr4.checked) ||
      (art[2] >= firstCurseurMoins.value &&
        art[2] <= firstCurseurPlus.value &&
        long5 &&
        mbr5.checked) ||
      (art[2] >= firstCurseurMoins.value &&
        art[2] <= firstCurseurPlus.value &&
        long6 &&
        mbr6.checked) ||
      (art[2] >= firstCurseurMoins.value &&
        art[2] <= firstCurseurPlus.value &&
        long7 &&
        mbr7.checked) ||
      (art[2] >= firstCurseurMoins.value &&
        art[2] <= firstCurseurPlus.value &&
        long8 &&
        mbr8.checked)
    ) {
      aff.innerHTML += artiste[i].name + "<br>";
    }
  }
}
//faire combine creation date first album,ville et membres,first album et membres,
executerRequeteArtiste(affiche);
executerRequeteVille(null);
