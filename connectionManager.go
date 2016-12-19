package main

import (
	"sync"
)

type connectionManager struct {
	lock        sync.Mutex
	connections []string
}

func newConnectionManager() *connectionManager {
	return &connectionManager{
		lock:        sync.Mutex{},
		connections: []string{},
	}
}

func (cm *connectionManager) addIP(newIP string) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	for _, ip := range cm.connections {
		if ip == newIP {
			return
		}
	}
	cm.connections = append(cm.connections, newIP)
}

func (cm *connectionManager) containsIP(newIP string) bool {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	for _, ip := range cm.connections {
		if ip == newIP {
			return true
		}
	}
	return false
}

func (cm *connectionManager) removeIP(removeIP string) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	position := -1
	for i, ip := range cm.connections {
		if ip == removeIP {
			position = i
		}
	}

	if position != -1 {
		cm.connections = append(cm.connections[:position], cm.connections[position+1:]...)
	}
}
