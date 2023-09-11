# Programmerbar Status API

Used for getting and updating the status of echo Programmerbar.

## Routes

### `GET /`

Does nothing. Is only used to check if the server is up.

### `GET /status`

Returns the current status of echo Programmerbar.

`OPEN` - The bar is open and you should come and have a beer.

`CLOSED` - The bar is closed and you should go home.

### `POST /status`

Updates the status of echo Programmerbar, and returns the new status.

## TODO

- [ ] Add tests to test the API
