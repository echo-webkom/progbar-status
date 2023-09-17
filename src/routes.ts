import { getStatus, setStatus } from "./models/status";

/**
 * For health checks
 *
 * GET /
 *
 * Returns a 200 OK response
 */
export function handleRoot() {
  return new Response("STATUS OK", {
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
export function handleStatus() {
  const { status, message } = getStatus();

  return new Response(
    JSON.stringify({
      status,
      message,
    }),
    {
      status: 200,
    }
  );
}

/**
 * For toggling the status of the bar
 *
 * POST /status
 *
 * Returns a 200 OK response with the new status of the bar
 */
export function handleUpdateStatus(req: Request) {
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

  // Get the new status from the search params
  const newStatus = Number(new URL(req.url).searchParams.get("status"));

  if (isNaN(newStatus) || newStatus > 1 || newStatus < 0) {
    return new Response("Invalid status. Provide a number.", {
      status: 400,
    });
  }

  setStatus(newStatus);

  const { status, message } = getStatus();

  return new Response(
    JSON.stringify({
      status,
      message,
    }),
    {
      status: 200,
    }
  );
}
