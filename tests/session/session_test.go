package main

import (
	"fmt"
	"testing"

	"../../helpers/session"
)

const (
	USERNAME = "TEST"
)

func TestSessionId(t *testing.T) {
	sessionId, err := sessionHelpers.GetSessionId()
	if err != nil {
		t.Error("Error occured while creating session id: " + err.Error())
	} else {
		fmt.Print("Session Id is: " + sessionId)
	}
}

func TestSessionConnection(t *testing.T) {
	connection, err := sessionHelpers.CreateRedisConnection()
	if err != nil {
		t.Error("Error Occured while creating a connection " + err.Error())
	} else {
		print("Connection Created Successfully ")
		connection.Close()
	}
}

func TestSetAndGetSession(t *testing.T) {
	sessionId, err := sessionHelpers.GetSessionId()
	if err != nil {
		t.Error("Error occured while creating session id: " + err.Error())
	}
	connection, err := sessionHelpers.CreateRedisConnection()
	if err != nil {
		t.Error("Error Occured while creating a connection " + err.Error())
	} else {
		err := sessionHelpers.AddSessionTokenToRedis(connection, sessionId, USERNAME)
		if err != nil {
			t.Error("Error Occured while setting a token " + err.Error())
		} else {
			username, err := sessionHelpers.GetSessionToken(connection, sessionId)
			if err != nil {
				t.Error("Error Occured while getting Session Id " + err.Error())
			} else {
				fmt.Println(username)
			}
		}
	}
}
