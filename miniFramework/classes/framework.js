 export class MyFramework {
  constructor() {}

  // Créer un nouvel élément DOM avec les attributs et le contenu donnés
  createElement(tagName, attributes = {}, content = '') {
    const element = document.createElement(tagName);

    for (const [key, value] of Object.entries(attributes)) {
      element.setAttribute(key, value);
    }

    element.textContent = content;

    return element;
  }

  // Ajouter un élément DOM enfant à un parent donné
  addElement(parent, child) {
    parent.appendChild(child);
  }

  // Modifier le contenu d'un élément DOM donné
  modifyElementContent(element, newContent) {
    element.textContent = newContent;
  }

  // Ajouter un gestionnaire d'événement à un élément DOM donné
  addEvent(element, eventType, eventHandler) {
    element.addEventListener(eventType, eventHandler);
  }
  removeElement(parent, child) {
    parent.removeChild(child);
  }
  
  // Appliquer un style CSS à un élément DOM donné
  setStyle(element, styles) {
    for (const [property, value] of Object.entries(styles)) {
      element.style[property] = value;
    }
  }
  modifyAttributes(element, newAttributes) {
    for (const [key, value] of Object.entries(newAttributes)) {
      element.setAttribute(key, value);
    }
  }
  removeAllChildren(parent) {
    while (parent.firstChild) {
      parent.removeChild(parent.firstChild);
    }
  }
  changeClass(element, oldClass, newClass) {
    if (element.classList.contains(oldClass)) {
      element.classList.remove(oldClass);
    }
    element.classList.add(newClass);
  }
  // Et ainsi de suite...


 
}
