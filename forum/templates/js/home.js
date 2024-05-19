let reponse =document.getElementById("reponseDemande")
document.addEventListener("DOMContentLoaded", function() {
    const fermerNotifButton = document.getElementById("fermerNotif");

    if (fermerNotifButton) {
        fermerNotifButton.addEventListener("click", function() {
            const container = document.getElementById("divNotifs");
            const dataDiv = document.querySelector(".data-list");
            
    
            if (dataDiv) {
                container.removeChild(dataDiv);
            }
        });
    }
});


document.addEventListener("DOMContentLoaded", function() {
    const supprButtons = document.querySelectorAll(".supprMess");
    
    supprButtons.forEach(button => {
        button.addEventListener("click", function(e) {
            const id = e.target.getAttribute('data-id');
            
            // Envoyer la valeur du pseudo au serveur Go via une requête POST
            fetch("/supprMess", {
                method: "POST",
                body: JSON.stringify({ id}), // Envoyez le pseudo au format JSON
                headers: {
                    'Content-Type': 'application/json'
                }
            })
            .then(response => {
                if (response.ok) {
                    // console.log("vous avez accepte la demande")
                    reponse.textContent="vous avez supprime le message";

                } else {
                    reponse.textContent="la requete n'a pas abouti";
                }
            })
            .catch(error => {
                console.error("Erreur lors de la requête AJAX : ", error);
            });
        });
    });
});
document.addEventListener("DOMContentLoaded", function() {
    const apprButtons = document.querySelectorAll(".approuverMess");
    
    apprButtons.forEach(button => {
        button.addEventListener("click", function(e) {
            const id = e.target.getAttribute('data-id');
            
            // Envoyer la valeur du pseudo au serveur Go via une requête POST
            fetch("/approuveMess", {
                method: "POST",
                body: JSON.stringify({ id}), // Envoyez le pseudo au format JSON
                headers: {
                    'Content-Type': 'application/json'
                }
            })
            .then(response => {
                if (response.ok) {
                    // console.log("vous avez accepte la demande")
                    reponse.textContent="vous avez approuve le message";

                } else {
                    reponse.textContent="la requete n'a pas abouti";
                }
            })
            .catch(error => {
                console.error("Erreur lors de la requête AJAX : ", error);
            });
        });
    });
});
document.addEventListener("DOMContentLoaded", function() {
    const signButtons = document.querySelectorAll(".signalerMess");
    
    signButtons.forEach(button => {
        button.addEventListener("click", function(e) {
            const id = e.target.getAttribute('data-id');
            
            // Envoyer la valeur du pseudo au serveur Go via une requête POST
            fetch("/signalerMess", {
                method: "POST",
                body: JSON.stringify({ id}), // Envoyez le pseudo au format JSON
                headers: {
                    'Content-Type': 'application/json'
                }
            })
            .then(response => {
                if (response.ok) {
                    // console.log("vous avez accepte la demande")
                    reponse.textContent="vous avez signale le message";

                } else {
                    reponse.textContent="la requete n'a pas abouti";
                }
            })
            .catch(error => {
                console.error("Erreur lors de la requête AJAX : ", error);
            });
        });
    });
});


document.addEventListener("DOMContentLoaded", function() {
    const Buttons = document.querySelectorAll(".notifs");
    
    Buttons.forEach(button => {
        button.addEventListener("click", function(e) {
            e.preventDefault()
            console.log("ok")
            // Envoyer une requête POST vide
            fetch("/affNotifs", {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json'
                }
            })
            .then(response => {
                if (response.ok) {
                    return response.json(); // Convertir la réponse en JSON
                } else {
                    reponse.textContent = "La requête n'a pas abouti";
                }
            })
            .then(data => {
                // Utilisez les données JSON dans 'data' dans votre code JavaScript
                const container = document.getElementById("divNotifs"); // Remplacez "container" par l'ID de l'endroit où vous souhaitez afficher la liste

                // Créez une div dynamiquement
                const dataDiv = document.createElement("div");
                
                // Donnez-lui une classe CSS pour la mise en forme ou utilisez d'autres attributs comme l'ID au besoin
                dataDiv.classList.add("data-list"); // Ajoutez une classe CSS si nécessaire
                
                data.forEach((item, index) => {
                    const paragraph = document.createElement("p");
                    paragraph.textContent = `${item.pseudo} vous a envoyé un ${item.nature}:`;
                
                    // Créez un lien "Voir" pour chaque notification
                    const voirLink = document.createElement("a");
                    voirLink.textContent = "Voir";
                    voirLink.href = `http://localhost:5555/messagesInd?id=${item.idPost}`;


                    const supprimerButton = document.createElement("button");
                    supprimerButton.textContent = "Supprimer";
                
                    // Ajoutez un gestionnaire d'événements pour le bouton "X"
                    supprimerButton.addEventListener("click", () => {
                        // Créez un objet JSON avec l'ID de la notification
                        let id= item.idNotif ;
                
                        // Envoyez une requête POST avec l'objet JSON
                        fetch("/supprNotif", {
                            method: "POST",
                            body: JSON.stringify({ id}),
                            headers: {
                                'Content-Type': 'application/json'
                            }
                        })
                        .then(response => {
                            if (response.ok) {
                                // La notification a été marquée pour suppression
                                // Vous pouvez gérer la mise à jour de l'affichage si nécessaire
                            } else {
                                // Gérer les erreurs si la suppression a échoué
                                console.error("Échec de la suppression de la notification.");
                            }
                        })
                        .catch(error => {
                            console.error("Erreur lors de la requête AJAX : ", error);
                        });
                    });
                
                    // Ajoutez le paragraphe et le lien à la div de données
                    dataDiv.appendChild(paragraph);
                    dataDiv.appendChild(voirLink);
                    dataDiv.appendChild(supprimerButton);
                });
                // Ajoutez la div avec les données à l'endroit spécifié dans le DOM
                container.appendChild(dataDiv);
            })
            .catch(error => {
                console.error("Erreur lors de la requête AJAX : ", error);
            });
        })
    })
})
