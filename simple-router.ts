export class SimpleRouter {
  private routes: Map<string, Map<string, Function>>;

  constructor() {
    this.routes = new Map();
  }

  /**
   * Add a route and its corresponding handler function for GET requests.
   * @param route
   * @param handler
   */
  get(route: string, handler: Function) {
    this.addRoute("GET", route, handler);
  }

  /**
   * Add a route and its corresponding handler function for POST requests.
   * @param route
   * @param handler
   */
  post(route: string, handler: Function) {
    this.addRoute("POST", route, handler);
  }

  /**
   * Add a route and its corresponding handler function for PUT requests.
   * @param route
   * @param handler
   */
  put(route: string, handler: Function) {
    this.addRoute("PUT", route, handler);
  }

  /**
   * Add a route and its corresponding handler function for DELETE requests.
   * @param route
   * @param handler
   */
  delete(route: string, handler: Function) {
    this.addRoute("DELETE", route, handler);
  }

  /**
   * Add a route and its corresponding handler function for a given method.
   * @param method
   * @param route
   * @param handler
   */
  private addRoute(method: string, route: string, handler: Function) {
    if (!this.routes.has(route)) {
      this.routes.set(route, new Map());
    }
    this.routes.get(route)?.set(method, handler);
  }

  /**
   * Handle a request by matching its method and path to a route and calling the corresponding handler function.
   * @param request
   * @returns
   */
  handleRequest(request: Request) {
    const url = new URL(request.url);
    const pathname = url.pathname;
    const method = request.method;

    if (this.routes.has(pathname)) {
      const routeHandlers = this.routes.get(pathname);
      const handler = routeHandlers?.get(method);

      if (handler instanceof Function) {
        return handler(request);
      } else {
        throw new Response(
          `Route handler for ${method} ${pathname} is not defined.`,
          {
            status: 405,
          }
        );
      }
    } else {
      throw new Response(`Route not found for ${method} ${pathname}.`, {
        status: 404,
      });
    }
  }
}
