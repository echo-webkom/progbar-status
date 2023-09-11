import { decode } from "https://deno.land/std@0.201.0/encoding/base64.ts";
import { load } from "https://deno.land/std@0.201.0/dotenv/mod.ts";

const env = await load();
const db = await Deno.openKv();

/**
 * Gets the open/closed status of the bar
 * @returns {Promise<boolean>} The open/closed status of the bar
 */
async function getStatus() {
  const status = await db.get<boolean>(["status"]);

  if (status.value === null) {
    await setStatus(false);
    return false;
  }

  return status.value;
}

/**
 * Sets the open/closed status of the bar
 * @param {boolean} value
 */
async function setStatus(value: boolean) {
  await db.set(["status"], value);
}

Deno.serve({
  hostname: "0.0.0.0",
  handler: async (req) => {
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
      const status = await getStatus();

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
      const adminKey = env["ADMIN_KEY"];

      if (adminKey) {
        const auth = req.headers.get("Authorization")?.split(" ")[1];

        if (auth !== adminKey) {
          return new Response("Unauthorized. Provide a valid bearer token.", {
            status: 401,
          });
        }
      }

      const status = await getStatus();

      await setStatus(!status);

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
  onListen({ port, hostname }) {
    console.log(`🚀 Server started at http://${hostname}:${port}`);
  },
});
