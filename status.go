package main

import "strconv"

func SetStatus(status int) error {
	return RDB.Set("status", status, 0).Err()
}

func GetStatus() (int, error) {
	status, err := RDB.Get("status").Result()
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(status)
}
