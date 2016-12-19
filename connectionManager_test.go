package main

import (
	"testing"
)

func TestConnectionManager_Contains(t *testing.T) {
	testCM := newConnectionManager()
	if testCM.containsIP("newIP") {
		t.Error("connection manager returned true when it did not contain that entry")
	}

	testCM.connections = append(testCM.connections, "newIP")
	if !testCM.containsIP("newIP") {
		t.Error("connection manager returned false when it did contain that entry")
	}
}

func TestConnectionManager_Add(t *testing.T) {
	testCM := newConnectionManager()

	testCM.addIP("newIP")
	if !testCM.containsIP("newIP") {
		t.Error("connection manager does not add unique IPs")
	}

	testCM.addIP("newIP")
	if !testCM.containsIP("newIP") {
		t.Error("connection manager removed a pre existing IP on add")
	}
}

func TestConnectionManager_Remove(t *testing.T) {
	testCM := newConnectionManager()

	testCM.addIP("newIP")
	testCM.removeIP("newIP")
	if testCM.containsIP("newIP") {
		t.Error("connection manager did not remove IP")
	}

	//test on an empty connection manager
	testCM.removeIP("newIP")
	if testCM.containsIP("newIP") {
		t.Error("connection manager did not remove IP")
	}
}
