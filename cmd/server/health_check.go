package main

import (
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func startHealthCheck(logger log.Logger, port int) {
	check := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router := http.NewServeMux()
	router.Handle("/health_check", check)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		level.Error(logger).Log("event", "health_check.failed", "err", err)
	}
}
