package help

func ExampleDefaultCatalog() {
	_ = DefaultCatalog
}

func ExampleCatalog_Add() {
	_ = (*Catalog).Add
}

func ExampleCatalog_List() {
	_ = (*Catalog).List
}

func ExampleCatalog_All() {
	_ = (*Catalog).All
}

func ExampleCatalog_Search() {
	_ = (*Catalog).Search
}

func ExampleCatalog_SearchResults() {
	_ = (*Catalog).SearchResults
}

func ExampleLoadContentDir() {
	_ = LoadContentDir
}

func ExampleCatalog_Get() {
	_ = (*Catalog).Get
}
