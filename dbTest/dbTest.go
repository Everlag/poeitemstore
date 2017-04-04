// Package dbTest holds tests for externally
// available functionality of the db package
//
// This avoids clutter and keeps things pretty.
package dbTest

// Generate file for easy access to testing data.
//go:generate go-bindata -pkg $GOPACKAGE -nometadata -o assets.go testData/

// The remainder of this file is left intentionally blank to avoid
// polluting the non-testing environment
