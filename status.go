package main

func SetStatus(status string) error {
	return RDB.Set("status", status, 0).Err()
}

func GetStatus() (string, error) {
	return RDB.Get("status").Result()
}
