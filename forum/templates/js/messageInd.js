let reponse =document.getElementById("reponseDemande")

let pseudoMsg=document.getElementById("pseudoMsg")
// voir du cote de dejalike quand on met des strings vides pour le probleme des 0 noirs
let bouton = document.getElementById("boutCo");
let essai = document.getElementById("divComment");
let user = document.getElementById("user");
let boutLike = document.getElementById("like");
let boutDislike = document.getElementById("disLike");
let refusLike = document.getElementById("refusLike");
// let refusCon = document.getElementById("refusCon");
let affLike = document.getElementById("affLike");
let affdisLike = document.getElementById("affdisLike");

// let countLike = 0;
// let countdisLike = 0;
var countLikePseudo = 0;
var countDisLikePseudo = 0;
var likeeValue = getCookie("like");
let likeValue = parseInt(likeeValue);
var disLikeValue = getCookie("disLike");
var dejaLikeJoin = getCookie("dejaLike");
var dejaDisLikeJoin = getCookie("dejaDisLike");
var dejaLike = dejaLikeJoin.split(" ");
var dejaDisLike = dejaDisLikeJoin.split(" ");
var dejaLikePourGO = false;
var dejaDisLikePourGO = false;
let count = 0;
let countDis = 0;
let dejaDisLike1;
let dejaLike1;
let boutLikeCom = document.getElementById("boutLikeCom");
let boutDislikeCom = document.getElementById("boutDislikeCom");
let refusLikeCom = document.getElementById("refusLikeComment");
var modal = document.getElementById("imageModal");
var closeBtn = document.getElementById("closeImageModal");

// Récupérer l'image en miniature
var thumbnailImage = document.getElementById("imgMess");

// Récupérer l'image en taille réelle
var modalImage = document.getElementById("modalImage");

// Ouvrir la modal lorsque l'image miniature est cliquée
if(thumbnailImage){
thumbnailImage.addEventListener("click", function() {
  modal.style.display = "block";
  modalImage.src = thumbnailImage.querySelector("img").src;
});
}
// Fermer la modal lorsque l'utilisateur clique sur le bouton de fermeture
closeBtn.addEventListener("click", function() {
  modal.style.display = "none";
});

// Fermer la modal si l'utilisateur clique en dehors de la modal
window.addEventListener("click", function(event) {
  if (event.target == modal) {
    modal.style.display = "none";
  }
});
function getCookieValue(name) {
  const cookies = document.cookie.split('; ');
  for (let i = 0; i < cookies.length; i++) {
    const [cookieName, cookieValue] = cookies[i].split('=');
    if (cookieName === name) {
      return cookieValue;
    }
  }
  return null;
}

// Récupérer la valeur des cookies comAffLike et comAffDisLike

// const AFFLIKE = JSON.parse(getCookieValue('comAffLike'));
const AFFLIKE = JSON.parse(decodeURIComponent(getCookieValue('comAffLike')));
const AFFDISLIKE = JSON.parse(decodeURIComponent(getCookieValue('comAffDisLike')));

window.onload = function() {
  var commentaires = document.querySelectorAll('#commentaire');
  var comId = Array.from(commentaires).map(function(commentaire) {
    return commentaire.querySelector('#idCom').textContent;
  });
  for (let i=0; i<comId.length; i++){
    let afflikeComm = document.querySelector("#affLikeComment" + comId[i]);
    let affDislikeCom = document.querySelector("#affDislikeComment" + comId[i]);

    if (AFFLIKE[i] === "true") {
        afflikeComm.style.backgroundColor="green";
    }
    if (AFFDISLIKE[i] === "true") {
    affDislikeCom.style.backgroundColor="red";
  }
}
};








affLike.textContent = parseInt(likeValue);
affdisLike.textContent = parseInt(disLikeValue);
if (user.innerText.length > 0) {
  if (
    dejaLike.includes('"' + user.innerText) ||
    dejaLike.includes(user.innerText) ||
    dejaLike.includes(user.innerText + '"')
  ) {
    affLike.style.backgroundColor = "black";
    dejaLikePourGO = true;
    dejaLike1 = true;
    count = 3;
  } else {
    dejaLikePourGO = false;
  }
  //reste a faire qu'on peut pas liker et disliker en meme temps seulement dans le js
  if (
    dejaDisLike.includes('"' + user.innerText) ||
    dejaDisLike.includes(user.innerText) ||
    dejaDisLike.includes(user.innerText + '"')
  ) {
    affdisLike.style.backgroundColor = "black";
    dejaDisLikePourGO = true;
    dejaDisLike1 = true;
    countDis = 3;
  } else {
    dejaDisLikePourGO = false;
  }
}
// console.log(dejaLike);

// console.log(user.innerText.length);
function likee() {
 
  let pseudo=pseudoMsg.textContent
  if (dejaDisLike1 != true) {
    count++;
    refusLike.textContent = "";
  }
  if (dejaDisLike1 == true) {
    refusLike.textContent = "Vous avez deja Dislikes";
  }

  if (user.textContent.length === 0) {
    refusLike.textContent = "Veuillez vous connecter pour réagir aux posts";
  }
  if (user.textContent.length > 0) {
    if (count % 2 == 0 && dejaLike1 == true) {
      dejaLikePourGO = true;
      dejaLike1 = false;
      likeValue -= parseInt(1);
      affLike.style.backgroundColor = "blue";
    }
    if (count % 2 != 0 && dejaDisLike1 != true) {
      //maintenant c'est le count qui merde
      dejaLikePourGO = false; //peut etre passer par encore une autre booleenee
      dejaLike1 = true;
      likeValue += parseInt(1);
      affLike.style.backgroundColor = "black";
    }
    const data = {
      pseudo: pseudo,
      dejaLikePourGO: dejaLikePourGO
    };

    fetch("/like", {
      method: "POST",
      body: JSON.stringify(data), // Envoyez les données JSON
      headers: {
        'Content-Type': 'application/json'
      }
    })
    .then((response) => {
      // Traitement de la réponse du serveur
    });
  }
 
  affLike.textContent = parseInt(likeValue); //ca decremente qd je raffraichis la page il faut trouver une condition qu'on peut pas decrementer plus de 2 fois

 
}

function dislikee() {
  let pseudo=pseudoMsg.textContent
  if (dejaLike1 != true) {
    countDis++;
    refusLike.textContent = "";
  }
  if (dejaLike1 == true) {
    refusLike.textContent = "Vous avez deja likes";
  }

  if (user.textContent.length === 0) {
    refusLike.textContent = "Veuillez vous connecter pour réagir aux posts";
  }
  if (user.textContent.length > 0) {
    if (countDis % 2 == 0 && dejaDisLike1 == true) {
      dejaDisLikePourGO = true;
      dejaDisLike1 = false;
      disLikeValue -= parseInt(1);
      affdisLike.style.backgroundColor = "blue";
    }
    if (countDis % 2 != 0 && dejaLike1 != true) {
      dejaDisLikePourGO = false;
      dejaDisLike1 = true;
      disLikeValue += parseInt(1);
      affdisLike.style.backgroundColor = "black";
    }

    const data = {
      pseudo: pseudo,
      dejaLikePourGO: dejaDisLikePourGO
    };

    fetch("/disLike", {
      method: "POST",
      body: JSON.stringify(data), // Envoyez les données JSON
      headers: {
        'Content-Type': 'application/json'
      }
    })
    .then((response) => {
      // Traitement de la réponse du serveur
    });
  }
  affdisLike.textContent = parseInt(disLikeValue);
}
//ca decremente qd je raffraichis la page il faut trouver une condition qu'on peut pas decrementer plus de 2 fois

function getCookie(name) {
  var value = "; " + document.cookie;
  var parts = value.split("; " + name + "=");
  if (parts.length == 2) return parts.pop().split(";").shift();
}
bouton.addEventListener("click", function () {
  if (user.textContent.length < 1) {
    refusLike.textContent = "Veuillez vous connecter pour commenter";
  } else {
    essai.style.visibility = "visible";
    // console.log(user.textContent.length,"laaaaaaa");
  }
});

function likeCom(event,comId) {
  // let pseudos=pseudoMsg.textContent
  // fo recuperer id de ladresse
  const url = new URL(window.location.href);

  // Obtenez la valeur du paramètre "id" de l'URL
  const id = url.searchParams.get("id");
  
  // Vérifiez si "id" existe et n'est pas vide
  // if (id) {
  //   // Utilisez la valeur de "id" comme vous le souhaitez
  //   console.log("ID récupéré :", id);
  // } else {
  //   console.log("Le paramètre 'id' est absent de l'URL.");
  // }
  let affLikeCom = document.querySelector("#affLikeComment" + comId);
  let valueLike = parseInt(affLikeCom.innerText)
  const dataCom = {
   idCom:comId,
   idPost:id
  };

// fetch("/likeCom", { method: "POST", body: comId}).then(
  fetch("/likeCom", { method: "POST", body: JSON.stringify(dataCom)}).then(
  


  (response) => {
    if (response.ok) {
      response.text().then((text) => {
          console.log(text);
          if((text)=="c bon"){
            valueLike+=1
            affLikeCom.textContent=valueLike
            affLikeCom.style.backgroundColor="green"
            
            }
            if((text)=="dejaLikeCom"&& valueLike>0){
              valueLike-=1
              affLikeCom.textContent=valueLike
              affLikeCom.style.backgroundColor="grey"
              }
         
      });
  } else {
      console.log("Like failed");
  }
  
}
  //a faire quand je clique sur le bouton like je recois une reponse du go si deja like ou pas 
  //si c'est deja like je decremente le contenu de affLike et le chiffre redevient bleu si il est noir(du div correspondant et je supprime dans le go l'entree)
  // si c bon ca insere dans la table et le chiffre se met en noir
);
event.preventDefault()
}

function disLikeCom(event,comId) {
  let affDislikeCom = document.querySelector("#affDislikeComment" + comId);
  let valuedisLike = parseInt(affDislikeCom.innerText)
  const url = new URL(window.location.href);
  const id = url.searchParams.get("id");
  const dataCom = {
    idCom:comId,
    idPost:id
   };
  // Obtenez la valeur du paramètre "id" de l'URL

fetch("/disLikeCom", { method: "POST", body:  JSON.stringify(dataCom)}).then(
  (response) => {
    if (response.ok) {
      response.text().then((text) => {
         
          if((text)=="c bon"){
            valuedisLike+=1
            affDislikeCom.textContent=valuedisLike
            affDislikeCom.style.backgroundColor="red"
            
            }
            if((text)=="dejadisLikeCom"&& valuedisLike>0){
              valuedisLike-=1
              affDislikeCom.textContent=valuedisLike
              affDislikeCom.style.backgroundColor="grey"
              }
         
      });
  } else {
      console.log("Like failed");
  }
  
}
  //a faire quand je clique sur le bouton like je recois une reponse du go si deja like ou pas 
  //si c'est deja like je decremente le contenu de affLike et le chiffre redevient bleu si il est noir(du div correspondant et je supprime dans le go l'entree)
  // si c bon ca insere dans la table et le chiffre se met en noir
);
event.preventDefault()
}
document.addEventListener("DOMContentLoaded", function() {
  const supButtons = document.querySelectorAll(".supprPost");
  
  supButtons.forEach(button => {
      button.addEventListener("click", function(e) {
          const id = e.target.getAttribute('data-id');
          console.log(id)
          

          // Envoyer la valeur du pseudo au serveur Go via une requête POST
          fetch("/supprOwnPost", {
              method: "POST",
              body: JSON.stringify({id}), // Envoyez le pseudo au format JSON
              headers: {
                  'Content-Type': 'application/json'
              }
          })
          .then(response => {
              if (response.ok) {
                 
  reponse.textContent="Vous avez supprime votre message"

              } else {
                reponse.textContent="Votre requete n'a pas abouti"
              }
          })
          .catch(error => {
              console.error("Erreur lors de la requête AJAX : ", error);
          });
      });
  });
})

document.addEventListener("DOMContentLoaded", function() {
  const suprButtons = document.querySelectorAll(".supprCom");
  
  suprButtons.forEach(button => {
      button.addEventListener("click", function(e) {
        
          const id = e.target.getAttribute('data-id');
          console.log(id)
         
          fetch("/supprOwnCom", {
              method: "POST",
              body: JSON.stringify({ id}), // Envoyez le pseudo au format JSON
              headers: {
                  'Content-Type': 'application/json'
              }
          })
          .then(response => {
              if (response.ok) {
                reponse.textContent="Vous avez supprime votre commentaire"

              } else {
                reponse.textContent="Votre requete n'a pas abouti"
              }
          })
          .catch(error => {
              console.error("Erreur lors de la requête AJAX : ", error);
          });
      });
   });
   
})

// document.addEventListener("DOMContentLoaded", function() {
//   const editButtons = document.querySelectorAll(".editMessage");

//   editButtons.forEach(button => {
//     button.addEventListener("click", function(e) {
//       e.preventDefault(); 
//       const idPost = e.target.getAttribute('data-id');
//       const commentaireElement = document.querySelector(".commentaire");

//       if (commentaireElement) {
//         // const isEditing = commentaireElement.contentEditable === "true";
//         // Rendre le contenu éditable
        
        
//         // if (!isEditing) {
//           // Rendre le contenu éditable
//           commentaireElement.contentEditable = true;
//           commentaireElement.focus();
          
//           // Mettre à jour le bouton "Editer" en "Enregistrer"
//           button.textContent = "Enregistrer";
//           button.classList.remove("editMessage");
//           button.classList.add("saveMessage");
//         // } else {
//         //   // Empêcher la propagation de l'événement "click" pour ne pas désactiver l'édition
//         //   e.preventDefault();
//         // }
    
//         // Ajouter un événement "click" pour enregistrer les modifications
//         button.addEventListener("click", function(e) {
//           const updatedText = commentaireElement.innerText;
//           // Envoyer les modifications au serveur Go via une requête POST
//           fetch("/editMess", {
//             method: "POST",
//             body: JSON.stringify({ idPost, updatedText }),
//             headers: {
//               'Content-Type': 'application/json'
//             }
//           })
//           .then(response => {
//             if (response.ok) {
//               // Désactiver l'édition du contenu
//               commentaireElement.contentEditable = false;
//               // Mettre à jour le bouton "Editer"
//               button.textContent = "Editer";
//               button.classList.remove("saveMessage");
//               button.classList.add("editMessage");
//             } else {
//               // Gérer les erreurs si nécessaire
//             }
//           })
//           .catch(error => {
//             // Gérer les erreurs si nécessaire
//           });
//         });
//       }
//     });
//   });
// });
document.addEventListener("DOMContentLoaded", function() {
  const editButtons = document.querySelectorAll(".editMessage");

  editButtons.forEach(button => {
    button.addEventListener("click", function(e) {
      const idPost = e.target.getAttribute('data-id');
      const commentaireElement = document.querySelector(".commentaire");;

      if (commentaireElement) {
        // Vérifier si le commentaire est actuellement en mode édition
        const isEditing = commentaireElement.contentEditable === "true";

        if (!isEditing) {
          // Commencer l'édition
          commentaireElement.contentEditable = true;
          commentaireElement.focus();

          // Mettre à jour le bouton "Editer" en "Enregistrer"
          button.textContent = "Enregistrer";
          button.classList.remove("editMessage");
          button.classList.add("saveMessage");
        } else {
          // Terminer l'édition
          commentaireElement.contentEditable = false;

          // Mettre à jour le bouton "Enregistrer" en "Editer"
          button.textContent = "Editer";
          button.classList.remove("saveMessage");
          button.classList.add("editMessage");

          // Envoyer les modifications au serveur Go via une requête POST
          const updatedText = commentaireElement.innerText;
          fetch("/editMess", {
            method: "POST",
            body: JSON.stringify({ idPost, updatedText }),
            headers: {
              'Content-Type': 'application/json'
            }
          })
          .then(response => {
            if (response.ok) {
              // Les modifications ont été enregistrées avec succès
            } else {
              // Gérer les erreurs si nécessaire
            }
          })
          .catch(error => {
            // Gérer les erreurs si nécessaire
          });
        }
      }
    });
  });
});




document.addEventListener("DOMContentLoaded", function() {
  const editButtons = document.querySelectorAll(".editComment");

  editButtons.forEach(button => {
    button.addEventListener("click", function(e) {
      const idCom = e.target.getAttribute('data-id');
      const commentaireElement = document.getElementById(`comm${idCom}`);

      if (commentaireElement) {
        // Vérifier si le commentaire est actuellement en mode édition
        const isEditing = commentaireElement.contentEditable === "true";

        if (!isEditing) {
          // Commencer l'édition
          commentaireElement.contentEditable = true;
          commentaireElement.focus();

          // Mettre à jour le bouton "Editer" en "Enregistrer"
          button.textContent = "Enregistrer";
          button.classList.remove("editMessage");
          button.classList.add("saveMessage");
        } else {
          // Terminer l'édition
          commentaireElement.contentEditable = false;

          // Mettre à jour le bouton "Enregistrer" en "Editer"
          button.textContent = "Editer";
          button.classList.remove("saveMessage");
          button.classList.add("editMessage");

          // Envoyer les modifications au serveur Go via une requête POST
          const updatedText = commentaireElement.innerText;
          fetch("/editCom", {
            method: "POST",
            body: JSON.stringify({ idCom, updatedText }),
            headers: {
              'Content-Type': 'application/json'
            }
          })
          .then(response => {
            if (response.ok) {
              // Les modifications ont été enregistrées avec succès
            } else {
              // Gérer les erreurs si nécessaire
            }
          })
          .catch(error => {
            // Gérer les erreurs si nécessaire
          });
        }
      }
    });
  });
});







