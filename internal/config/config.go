package config

import "time"

const (
	ServerPort = ":8082"
	MaxMemory  = 100 * 1024 * 1024
	TimeToLive = time.Minute * 5
)
