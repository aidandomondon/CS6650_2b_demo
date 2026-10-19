package main

type Product struct {
	// Go struct field names in Pascal case
	// so that they are exported/exposed,
	// but have JSON tags in snake case to comply with their
	// names in the OpenAPI docs.
	ProductID    int32  `json:"product_id"`
	SKU          string `json:"sku"`
	Manufacturer string `json:"manufacturer"`
	CategoryID   int32  `json:"category_id"`
	Weight       int32  `json:"weight"`
	SomeOtherID  int32  `json:"some_other_id"`
}
