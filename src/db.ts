import Database from "bun:sqlite";

declare global {
  var db: Database | undefined;
}

const dbName = "db.sqlite";

let db: Database;

if (process.env.NODE_ENV === "production") {
  db = new Database();
} else {
  if (!global.db) {
    global.db = new Database(dbName);
  }
  db = global.db;
}

/**
 * Sets up the database
 */
export function setUpDatabase() {
  db.query(
    "CREATE TABLE IF NOT EXISTS status (value  INTEGER NOT NULL DEFAULT 0)"
  ).run();
}

export default db;
