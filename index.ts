import { Database } from "bun:sqlite";

const db = new Database("db.sqlite");

// Create the table if it doesn't exist
db.query(
  "CREATE TABLE IF NOT EXISTS status (value BOOLEAN NOT NULL DEFAULT FALSE)"
).run();

/**
 * Gets the open/closed status of the bar
 * @returns {Promise<boolean>} The open/closed status of the bar
 */
function getStatus() {
  const query = db.prepare("SELECT value FROM status");

  const status = query.get() as {
    value: number;
  };

  if (status === null) {
    db.query("INSERT INTO status (value) VALUES (FALSE)").run();

    return false;
  }

  return status.value === 1;
}

/**
 * Sets the open/closed status of the bar
 * @param {boolean} value
 */
function setStatus(value: boolean) {
  const query = db.prepare("UPDATE status SET value = ?");

  query.run(value ? 1 : 0);

  return !value;
}

const server = Bun.serve({
  fetch: (req) => {
    const pathname = new URL(req.url).pathname;

    /**
     * For health checks
     *
     * GET /
     *
     * Returns a 200 OK response
     */
    if (pathname === "/" && req.method === "GET") {
      return new Response(undefined, {
        status: 200,
      });
    }

    /**
     * For getting the status of the bar
     *
     * GET /status
     *
     * Returns a 200 OK response with the status of the bar
     */
    if (pathname === "/status" && req.method === "GET") {
      const status = getStatus();

      return new Response(status ? "OPEN" : "CLOSED", {
        status: 200,
      });
    }

    /**
     * For toggling the status of the bar
     *
     * POST /status
     *
     * Returns a 200 OK response with the new status of the bar
     */
    if (pathname === "/status" && req.method === "POST") {
      const adminKey = Bun.env.API_KEY;

      /**
       * If there is an admin key set, check if the request has a valid bearer token.
       */
      if (adminKey) {
        const auth = req.headers.get("Authorization")?.split(" ")[1];

        if (auth !== adminKey) {
          return new Response("Unauthorized. Provide a valid bearer token.", {
            status: 401,
          });
        }
      }

      const status = getStatus();

      setStatus(!status);

      return new Response(!status ? "OPEN" : "CLOSED", {
        status: 200,
      });
    }

    /**
     * If no valid route is found send a 404 response with a message
     */
    return new Response("Only /, /status and /status/toggle are supported", {
      status: 404,
    });
  },
});

console.log(`🚀 Server started at http://${server.hostname}:${server.port}`);
