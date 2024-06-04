package main

import (
	"flag"
)

type Config struct {
	Host        string
	Port        int
	CookieName  string
	CookieValue string
	CookieTTL   int
	TargetURL   string
	Password    string
	Prefix      string
}

var configInstance Config

func init() {
	flag.StringVar(&configInstance.Host, "host", "0.0.0.0", "The host to listen on")
	flag.IntVar(&configInstance.Port, "port", 8080, "The port to listen on")
	flag.StringVar(&configInstance.CookieName, "cookie-name", "", "The cookie name to use for authentication")
	flag.StringVar(&configInstance.CookieValue, "cookie-value", "", "The cookie value to use for authentication")
	flag.IntVar(&configInstance.CookieTTL, "ttl", 300, "The cookie TTL in seconds")
	flag.StringVar(&configInstance.TargetURL, "target", "https://www.baidu.com", "The target URL to proxy")
	flag.StringVar(&configInstance.Password, "password", "", "The secret password")
	flag.StringVar(&configInstance.Prefix, "prefix", "", "The prefix to use for the proxy")
	flag.Parse()

	if configInstance.Port < 1 || configInstance.Port > 65535 {
		panic("Port must be between 1 and 65535")
	}
	if configInstance.CookieTTL < 1 {
		panic("Cookie TTL must be greater than 0")
	}
	if configInstance.Password == "" {
		panic("Password is required")
	}
	if configInstance.CookieName == "" {
		panic("Cookie Name is required")
	}
	if configInstance.CookieValue == "" {
		panic("Cookie Value is required")
	}
}
