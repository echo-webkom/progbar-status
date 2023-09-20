# Programmerbar Status API

Used for getting and updating the status of echo Programmerbar.

## Running

Run the server with:

```sh
make run
```

Run the server in development mode with:

```sh
make dev
```

## Routes

### `GET /`

Does nothing. Is only used to check if the server is up.

### `GET /status`

Returns the current status of echo Programmerbar, in JSON format.

```sh
curl -X GET "http://localhost:8080/status"
```

```json
{
  "status": number,
  "message": string
}
```

Possible status codes:

- `0`: Closed
- `1`: Open

### `POST /status`

Updates the status of echo Programmerbar, and returns the new status. This requires a JSON body with the following format:

```json
{
  "status": number,
}
```

```sh
curl -X POST "http://localhost:8080/status" \
  -H "Content-Type: application/json" \
  -d '{"status": 1}'
```

```json
{
  "status": number,
  "message": string
}
```

## TODO

- [ ] Add tests to test the API
