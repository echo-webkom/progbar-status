import db from "../db";

enum Status {
  Closed = 0,
  Open = 1,
}

const getStatusText = (status: Status) => {
  switch (status) {
    case Status.Closed:
      return "Baren er stengt.";
    case Status.Open:
      return "Baren er åpen.";
    default:
      return "Ukjent status.";
  }
};

/**
 * Gets the open/closed status of the bar
 * @returns {Promise<boolean>} The open/closed status of the bar
 */
export function getStatus(): { status: Status; text: string } {
  const query = db.prepare("SELECT value FROM status");

  const status = query.get() as {
    value: number;
  } | null;

  if (status === null) {
    db.query("INSERT INTO status (value) VALUES (0)").run();

    return {
      status: Status.Closed,
      text: getStatusText(Status.Closed),
    };
  }

  return {
    status: status.value,
    text: getStatusText(status.value),
  };
}

/**
 * Sets the open/closed status of the bar
 * @param {boolean} value
 */
export function setStatus(value: Status) {
  const query = db.prepare("UPDATE status SET value = ?");

  query.run(value);
}
