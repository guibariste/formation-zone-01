import { renderContent } from './render.js';

export function setupRoutes(routes) {
  renderContent(window.location.pathname, routes);

  document.querySelectorAll('nav a').forEach(link => {
    link.addEventListener('click', event => {
      event.preventDefault();
      renderContent(link.getAttribute('href'), routes);
    });
  });

  window.addEventListener('popstate', () => {
    renderContent(window.location.pathname, routes);
  });
}