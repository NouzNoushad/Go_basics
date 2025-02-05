package models

import "time"

type DefaultSettings struct {
	MaxExpire time.Duration
	MinExpire time.Duration
}

type Settings struct {
	CustomExpire time.Duration
	UseCache     bool
}
