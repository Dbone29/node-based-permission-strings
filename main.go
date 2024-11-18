package main

import (
	"fmt"

	"github.com/Dbone29/node-based-permission-strings/perms"
)

func main() {
	permissionStrings := []string{}

	// Add permissions
	//AddPermission(&permissions, "plugin.command.*") // Wildcard for all commands in plugin
	perms.AddPermission(&permissionStrings, "plugin.admin") // Specific permission
	//AddPermission(&permissions, "*")                // Global wildcard

	// Check permissions
	fmt.Println("Has 'plugin.command.use':", perms.CheckPermission(&permissionStrings, "plugin.command.use"))         // true (matches "plugin.command.*")
	fmt.Println("Has 'plugin.command.delete':", perms.CheckPermission(&permissionStrings, "plugin.command.delete"))   // true (matches "plugin.command.*")
	fmt.Println("Has 'plugin.admin':", perms.CheckPermission(&permissionStrings, "plugin.admin"))                     // true (exact match)
	fmt.Println("Has 'unknown.node':", perms.CheckPermission(&permissionStrings, "unknown.node"))                     // true (matches "*")
	fmt.Println("Has 'plugin.command.subnode':", perms.CheckPermission(&permissionStrings, "plugin.command.subnode")) // true (matches "plugin.command.*")
}
