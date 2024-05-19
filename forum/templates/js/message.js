var divImg = document.getElementById("divImg");
var erreur=document.getElementById("autorise");
document.getElementById('fileToUpload').addEventListener('change', function () {
    var input = this;
    if (input.files && input.files[0]) {
        var file = input.files[0];
    
        if (file.size > 2500000) { // Vérifiez si la taille du fichier dépasse 2.5 Mo
          erreur.textContent="Taille superieur a 20mb veuillez prendre une image plus petite"
            input.value = ''; // Réinitialisez le champ de fichier
            return; // Ne continuez pas le traitement
        }else{
            erreur.textContent=""
        }
       
       
        var reader = new FileReader();

        reader.onload = function (e) {
            // Afficher l'élément divImg
            divImg.style.display = 'block';

            // Mettre à jour le champ d'image caché avec les données de l'image
            document.getElementById('imageData').value = e.target.result;

            // Afficher le bouton de suppression de l'image
            document.getElementById('removeImage').style.display = 'block';

            // Afficher la prévisualisation de l'image
            document.getElementById('preview').src = e.target.result;
        };

        reader.readAsDataURL(input.files[0]);
    } else {
        // Masquer l'élément divImg s'il n'y a pas de fichier choisi
        divImg.style.display = 'none';
    }
});

// Fonction appelée pour supprimer l'image de la prévisualisation
document.getElementById('removeImage').addEventListener('click', function () {
    // Masquer l'élément divImg
    divImg.style.display = 'none';

    // Réinitialiser la prévisualisation de l'image
    document.getElementById('preview').src = '';

    // Réinitialiser le champ d'image caché
    document.getElementById('imageData').value = '';

    // Cacher le bouton de suppression de l'image
    document.getElementById('removeImage').style.display = 'none';

    // Réinitialiser le champ de fichier
    document.getElementById('fileToUpload').value = '';
})


