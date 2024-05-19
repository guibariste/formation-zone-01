let autorise = document.getElementById("autorise");
let idMess = document.getElementById("idMess");

if (
  autorise.innerText == "Vous n'avez pas poste de messages encore" ||
  autorise.innerText == "Veuillez vous connecter pour lire vos posts likés" ||
  autorise.innerText == "Veuillez vous connecter pour lire vos posts"
) {
  idMess.style.visibility = "hidden";
}
