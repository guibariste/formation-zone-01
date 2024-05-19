
import { ajoutTache } from "./ajoutTache.js";

import { MyFramework } from "./classes/framework.js";

const maListe = new MyFramework();

// Sélectionner le bouton d'ajout de tâche
const boutdelete = document.querySelector('.delete-button');
const boutAjoutTache = document.querySelector('.ajoutTache');
const boutAll = document.querySelector('.all');
const boutActive = document.querySelector('.active');
const boutCompleted = document.querySelector('.completed');
const boutClearCompleted = document.querySelector('.clear-completed');
const liste = document.querySelector('.Liste');
// const countAfaire = liste.querySelectorAll('li').length;
const countAfaireStr = document.querySelector('#count');
let countAfaire2;
const essai = document.querySelector('#essai');
export function update() {
 
  // console.log(parseInt(countAfaireStr.textContent))
  const checkedActive = liste.querySelectorAll('li.all-active');
  const checkedCompleted = liste.querySelectorAll('li.all-completed');
  if(checkedActive.length===0&&checkedCompleted.length===0){
  essai.style.visibility="hidden"
  }else{
    essai.style.visibility="visible"
  }
}
  
  // }
function countVisibleElements(selector) {
  const elements = document.querySelectorAll(selector);
  const visibleElements = Array.from(elements).filter(element => {
    return element.style.visibility !== 'hidden';
  });
  return visibleElements.length;
}

update()



maListe.addEvent(liste, 'dblclick', function(event) {

  const tache = event.target;
  // console.log(maListe.length)
  // Vérifier si l'élément cliqué est une tâche (par exemple, une balise <li>)
  if (tache.tagName === 'LI') {
    const nouveauDiv = document.createElement('div');
    nouveauDiv.contentEditable = true;
    nouveauDiv.textContent = tache.textContent;
    
    // Écouter l'événement de perte de focus du nouveau div
    nouveauDiv.addEventListener('blur', function() {
      const nouveauContenu = nouveauDiv.textContent.trim();
      
      // Mettre à jour le contenu de la tâche si un nouveau contenu est saisi
      if (nouveauContenu !== '') {
        tache.textContent = nouveauContenu;
      }
      
      // Supprimer le nouveau div
      tache.removeChild(nouveauDiv);
     
    });
    
    // Remplacer le contenu de la tâche par le nouveau div
    tache.innerHTML = '';
    tache.appendChild(nouveauDiv);
    
    // Mettre le focus sur le nouveau div
    nouveauDiv.focus();
  }
});

maListe.addEvent(boutAjoutTache, "click", function() {//la il faut compter que les all active
  
  ajoutTache(maListe);
  let countAfaire2 = countVisibleElements('.all-active')




const countAfaireTotal =  parseInt(countAfaire2);
countAfaireStr.textContent = countAfaireTotal;
update()
});
maListe.addEvent(document.querySelector('.aFaire'), 'keydown', function(event) {

  if (event.key === 'Enter') {
    ajoutTache(maListe);
    let countAfaire2 = countVisibleElements('.all-active')




    const countAfaireTotal =  parseInt(countAfaire2);
    countAfaireStr.textContent = countAfaireTotal;
  }
  update()
});


maListe.addEvent(boutAll,"click",function(){

// recupClasseAll(maListe)
const elementsAll = document.querySelectorAll('[class*="all"]');
  elementsAll.forEach(element => {
    element.style.visibility="visible"
});
// let countAfaire1 = countVisibleElements('.all-completed')
let countAfaire2 = countVisibleElements('.all-active')




const countAfaireTotal =  parseInt(countAfaire2);
countAfaireStr.textContent = countAfaireTotal;
update()

}
)



maListe.addEvent(boutActive, "click", function() {
  // recupClasseActive(maListe);
 
  const aCacher = document.querySelectorAll('.all-completed');
  aCacher.forEach((tache) => {
    tache.style.visibility = 'hidden';
  });
  const aMontrer = document.querySelectorAll('.all-active');
  aMontrer.forEach(element => {
     
  element.style.visibility="visible"
    });
 let countAfaire= countVisibleElements('.all-active')
countAfaireStr.textContent = countAfaire;
update()
});
 



maListe.addEvent(boutCompleted,"click",function(){
 
    // recupClasseCompleted(maListe)

    const aCacher = document.querySelectorAll('.all-active');
  aCacher.forEach(element => {
     
  element.style.visibility="hidden"
    });
    const aMontrer = document.querySelectorAll('.all-completed');
    aMontrer.forEach(element => {
       
    element.style.visibility="visible"
      });
      // const elementsAll = document.querySelectorAll('[class*="all"]');
      // countAfaireStr.textContent=countAfaire
      //  let countAfaire= countVisibleElements(elementsAll)//il faut vraiment compter le total
      // //modifier la classed all completed ca fout la merde
      // countAfaireStr.textContent = countAfaire;
      const checkedElements = liste.querySelectorAll('li.all-active');
      const checkedCount = checkedElements.length;
      console.log("Nombre d'éléments cochés : " + checkedCount);
      countAfaireStr.textContent=checkedCount
      update()
    }
    )

    maListe.addEvent(boutClearCompleted,"click",function(){
     
      const aMontrer = document.querySelectorAll('.all-completed');
      aMontrer.forEach(element => {
         
  //  maListe.removeElement(liste,aMontrer)
element.remove()
        });
       
        let countAfaire1 = countVisibleElements('.all-completed')
         countAfaire2 = countVisibleElements('.all-active')
        
        
        
        
        const countAfaireTotal = parseInt(countAfaire1) + parseInt(countAfaire2);
        countAfaireStr.textContent = countAfaireTotal;
        update()
      }
      )
     
