export function renderContent(pathname, routes) {
    const route = routes.find(route => route.path === pathname);
  
    if (route) {
      history.pushState(null, null, pathname);
      if (route.func) {
        route.func();
      }
    } else {
      console.log("pas trouvé");
    }
  }