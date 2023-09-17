# Programmerbar Status API

Used for getting and updating the status of echo Programmerbar.

## Running

Run the server with:

```sh
bun run start
```

Run the server in development mode with:

```sh
bun run dev
```

## Routes

### `GET /`

Does nothing. Is only used to check if the server is up.

### `GET /status`

Returns the current status of echo Programmerbar, in JSON format.

```sh
curl -X GET "http://localhost:3000/status"
```

```json
{
  "status": 1,
  "message": "Baren er åpen"
}
```

Possible status codes:

- `0`: Closed
- `1`: Open

### `POST /status`

Updates the status of echo Programmerbar, and returns the new status. Requires `status` in the search parameters to update the status.

```sh
curl -X POST "http://localhost:3000/status?status=1"
```

```json
{
  "status": 1,
  "message": "Baren er åpen"
}
```

## TODO

- [ ] Add tests to test the API

## Simple Router

This is a simple router to make it to create routes.
