package testData

// Generate file for easy access to testing data.
//go:generate go-bindata -pkg $GOPACKAGE -nometadata -o assets.go data/

// This file is intentionally blank.
//
// To answer your question as to why an entire package for some testing data?
// The namespace in the dbTest package was very polluted and my editor
// didn't like that. Hence, this.
