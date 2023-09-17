import { setUpDatabase } from "./db";
import { handleRoot, handleStatus, handleUpdateStatus } from "./routes";
import { SimpleRouter } from "./simple-router";

setUpDatabase();

const server = Bun.serve({
  fetch: (req) => {
    const router = new SimpleRouter();

    router.get("/", handleRoot);
    router.get("/status", handleStatus);
    router.post("/status", handleUpdateStatus);

    return router.handleRequest(req);
  },
});

console.log(`🚀 Server started at http://${server.hostname}:${server.port}`);
