package connectionManager

import (
	"sync"
)

// connectionManager is a simple map of IP adresses and a mutex for
// threadsafe usage. Use NewConnectionManager() to get an instance and
// Add, Contains and Remove to use the connectionManager
type connectionManager struct {
	lock        sync.Mutex
	connections map[string]bool
}

// NewConnectionManager is a factory function for the connectionManager
func NewConnectionManager() *connectionManager {
	return &connectionManager{
		lock:        sync.Mutex{},
		connections: make(map[string]bool),
	}
}

// AddIP adds a new IP to the connection manager
func (cm *connectionManager) AddIP(IP string) {
	if !cm.ContainsIP(IP) {
		cm.lock.Lock()
		cm.connections[IP] = true
		cm.lock.Unlock()
	}
}

// ContainsIP returns true if IP is present in the connectionManager
func (cm *connectionManager) ContainsIP(IP string) bool {
	cm.lock.Lock()
	_, present := cm.connections[IP]
	cm.lock.Unlock()
	return present
}

// RemoveIP removes the IP from the connectionManager if it is present
func (cm *connectionManager) RemoveIP(IP string) {
	cm.lock.Lock()
	delete(cm.connections, IP)
	cm.lock.Unlock()
}
