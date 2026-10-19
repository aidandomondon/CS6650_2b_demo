package main

type Error struct {
	// Go struct field names in Pascal case
	// so that they are exported/exposed,
	// but have JSON tags in snake case to comply with their
	// names in the OpenAPI docs.
	Error   string `json:"error"`
	Message string `json:"message"`
	Details string `json:"details"`
}
