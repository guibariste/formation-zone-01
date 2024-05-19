import { setupRoutes } from './setupRoute.js';
import { essai } from './essaiExport.js';
import { renderContent } from './render.js';

const routes = [
  {
    path: '/',
    func: () => console.log('Vous êtes sur la page d\'accueil'),
  },
  {
    path: '/about',
    func: () => essai(),
  },
];

setupRoutes(routes);
