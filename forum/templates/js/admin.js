let reponse =document.getElementById("reponseDemande")
document.addEventListener("DOMContentLoaded", function() {
    document.getElementById("demandeModo").addEventListener("click", function() {
        fetch("/upgrade", {
            method: "POST"
        })
        .then(response => {
            if (response.ok) {
                reponse.textContent="Votre demande a été enregistrée avec succès.";
            } else if (response.status === 409) {
       reponse.textContent="Vous avez déjà fait une demande.";
            } else {
                reponse.textContent="Votre requete n'a pas abouti";
            }
        })
        .catch(error => {
            console.error("Erreur lors de la requête AJAX : ", error);
        });
    });
});
document.addEventListener("DOMContentLoaded", function() {
    const accepterDemandeButtons = document.querySelectorAll(".accepterDemande");
    
    accepterDemandeButtons.forEach(button => {
        button.addEventListener("click", function(e) {
            const pseudo = e.target.getAttribute('data-pseudo');
            
            // Envoyer la valeur du pseudo au serveur Go via une requête POST
            fetch("/accepterDemande", {
                method: "POST",
                body: JSON.stringify({ pseudo }), // Envoyez le pseudo au format JSON
                headers: {
                    'Content-Type': 'application/json'
                }
            })
            .then(response => {
                if (response.ok) {
                    // console.log("vous avez accepte la demande")
                    reponse.textContent="vous avez accepte la demande";

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
    const refuserDemandeButtons = document.querySelectorAll(".refuserDemande");
    
    refuserDemandeButtons.forEach(button => {
        button.addEventListener("click", function(e) {
            const pseudo = e.target.getAttribute('data-pseudo');
            console.log(pseudo)
            // Envoyer la valeur du pseudo au serveur Go via une requête POST
            fetch("/refuserDemande", {
                method: "POST",
                body: JSON.stringify({ pseudo }), // Envoyez le pseudo au format JSON
                headers: {
                    'Content-Type': 'application/json'
                }
            })
            .then(response => {
                if (response.ok) {
                    // console.log("vous avez accepte la demande")
                    reponse.textContent="vous avez refuse la demande";

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
    const retrogaderButtons = document.querySelectorAll(".retrogader");
    
   retrogaderButtons.forEach(button => {
        button.addEventListener("click", function(e) {
            const pseudo = e.target.getAttribute('data-pseudo');
            
            // Envoyer la valeur du pseudo au serveur Go via une requête POST
            fetch("/retrogader", {
                method: "POST",
                body: JSON.stringify({ pseudo }), // Envoyez le pseudo au format JSON
                headers: {
                    'Content-Type': 'application/json'
                }
            })
            .then(response => {
                if (response.ok) {
                    // console.log("vous avez accepte la demande")
                    reponse.textContent="vous avez retrogade un modo";

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