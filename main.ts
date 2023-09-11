const db = await Deno.openKv();

async function getStatus() {
  const status = await db.get<boolean>(["status"]);

  if (status.value === null) {
    await setStatus(false);
    return false;
  }

  return status.value;
}

async function setStatus(value: boolean) {
  await db.set(["status"], value);
}

Deno.serve({
  hostname: "0.0.0.0",
  handler: async (req) => {
    const pathname = new URL(req.url).pathname;

    if (pathname === "/") {
      return new Response(undefined, {
        status: 200,
      });
    }

    if (pathname === "/status") {
      const status = await getStatus();

      return new Response(status ? "OPEN" : "CLOSED", {
        status: 200,
      });
    }

    if (pathname === "/status/toggle") {
      const status = await getStatus();

      await setStatus(!status);

      return new Response(!status ? "OPEN" : "CLOSED", {
        status: 200,
      });
    }

    return new Response("Only /, /status and /status/toggle are supported", {
      status: 404,
    });
  },
  onListen({ port, hostname }) {
    console.log(`Server started at http://${hostname}:${port}`);
  },
});
