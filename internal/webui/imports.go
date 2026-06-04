package webui

// the only job of this package is to import all webui packages so they're actually registered
// Call Register() from main() instead of blank-importing this package.

import (
	_ "elichika/internal/webui/user"
	//_ "elichika/internal/webui/admin"
	_ "elichika/internal/webui/agnostic"
)

// Register triggers all webui handler registrations by importing this package.
func Register() {}
