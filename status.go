package main

import (
	"strconv"
)

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

func GetOrSetStatus() (int, error) {
	status, err := GetStatus()
	if err != nil {
		if err.Error() == "redis: nil" {
			SetStatus(0)
			return 0, nil
		}

		return 0, err
	}

	return status, nil
}
