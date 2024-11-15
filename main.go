package main

import (
	"fmt"
	"strings"
)

// CheckPermission checks if a a specific permission was given, considering wildcards
func CheckPermission(permissions *[]string, node string) bool {
	parts := strings.Split(node, ".")
	for i := len(parts); i >= 0; i-- {
		// Build wildcard node for the current level
		wildcardNode := strings.Join(parts[:i], ".") + ".*"
		if i == 0 {
			wildcardNode = "*" // Global wildcard
		}
		// Check if the node or wildcard is in the list
		if hasPermission(permissions, node) || hasPermission(permissions, wildcardNode) {
			return true
		}
	}
	return false
}

// hasPermission checks if a specific node exists in the user's permission list
func hasPermission(permissions *[]string, node string) bool {
	for _, perm := range *permissions {
		if perm == node {
			return true
		}
	}
	return false
}

// AddPermission adds a permission node to the user
func AddPermission(permissions *[]string, node string) {
	if !hasPermission(permissions, node) {
		*permissions = append(*permissions, node)
	}
}

// RemovePermission removes a permission node from the user
func RemovePermission(permissions *[]string, node string) {
	for i, perm := range *permissions {
		if perm == node {
			*permissions = append((*permissions)[:i], (*permissions)[i+1:]...)
			break
		}
	}
}

func main() {
	permissions := []string{}

	// Add permissions
	//AddPermission(&permissions, "plugin.command.*") // Wildcard for all commands in plugin
	AddPermission(&permissions, "plugin.admin") // Specific permission
	//AddPermission(&permissions, "*")                // Global wildcard

	// Check permissions
	fmt.Println("Has 'plugin.command.use':", CheckPermission(&permissions, "plugin.command.use"))         // true (matches "plugin.command.*")
	fmt.Println("Has 'plugin.command.delete':", CheckPermission(&permissions, "plugin.command.delete"))   // true (matches "plugin.command.*")
	fmt.Println("Has 'plugin.admin':", CheckPermission(&permissions, "plugin.admin"))                     // true (exact match)
	fmt.Println("Has 'unknown.node':", CheckPermission(&permissions, "unknown.node"))                     // true (matches "*")
	fmt.Println("Has 'plugin.command.subnode':", CheckPermission(&permissions, "plugin.command.subnode")) // true (matches "plugin.command.*")
}
