package main

import (
	"fmt"
	"net/http"
)

// handle request methods (account)
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccounts(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

// handle get accounts
func (s *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// handle create account
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}