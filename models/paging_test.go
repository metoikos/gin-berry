package models

import (
	"testing"
)

func TestBuildPaging(t *testing.T) {
	paging := BuildPaging(1, 20, 44)

	if paging.offset != 0 {
		t.Fatalf("Invalid offset. Expected 0 got %d", paging.offset)
	}

	if paging.nextPage != 2 {
		t.Fatalf("Invalid next page. Expected 2 got %d", paging.nextPage)
	}

	if paging.previousPage != 0 {
		t.Fatalf("Invalid previous page. Expected 2 got %d", paging.previousPage)
	}

	if len(paging.pages) != 3 {
		t.Fatalf("Invalid page list. Expected lenght 3 got %d", len(paging.pages))
	}

	if paging.lastPage != 3 {
		t.Fatalf("Invalid page size. Expected 3 got %d", paging.lastPage)
	}

	if !paging.enabled {
		t.Fatalf("Invalid enabled status. Expected true got %v", paging.enabled)
	}
}

func TestBuildPaging2(t *testing.T) {
	paging := BuildPaging(3, 20, 44)

	if paging.offset != 40 {
		t.Fatalf("Invalid offset. Expected 0 got %d", paging.offset)
	}

	if paging.nextPage != 0 {
		t.Fatalf("Invalid next page. Expected 2 got %d", paging.nextPage)
	}

	if paging.previousPage != 2 {
		t.Fatalf("Invalid previous page. Expected 2 got %d", paging.previousPage)
	}

	if len(paging.pages) != 3 {
		t.Fatalf("Invalid page list. Expected lenght 3 got %d", len(paging.pages))
	}

	if paging.lastPage != 3 {
		t.Fatalf("Invalid page size. Expected 3 got %d", paging.lastPage)
	}

	if !paging.enabled {
		t.Fatalf("Invalid enabled status. Expected true got %v", paging.enabled)
	}
}

func TestBuildPaging3(t *testing.T) {
	paging := BuildPaging(1, 20, 19)

	if paging.offset != 0 {
		t.Fatalf("Invalid offset. Expected 0 got %d", paging.offset)
	}

	if paging.nextPage != 0 {
		t.Fatalf("Invalid next page. Expected 2 got %d", paging.nextPage)
	}

	if paging.previousPage != 0 {
		t.Fatalf("Invalid previous page. Expected 2 got %d", paging.previousPage)
	}

	if len(paging.pages) != 1 {
		t.Fatalf("Invalid page list. Expected lenght 3 got %d", len(paging.pages))
	}

	if paging.lastPage != 1 {
		t.Fatalf("Invalid page size. Expected 3 got %d", paging.lastPage)
	}

	if paging.enabled {
		t.Fatalf("Invalid enabled status. Expected false got %v", paging.enabled)
	}
}

func TestBuildPaging4(t *testing.T) {
	paging := BuildPaging(3, 20, 19)

	if paging.offset != 0 {
		t.Fatalf("Invalid offset. Expected 0 got %d", paging.offset)
	}

	if paging.nextPage != 0 {
		t.Fatalf("Invalid next page. Expected 2 got %d", paging.nextPage)
	}

	if paging.previousPage != 0 {
		t.Fatalf("Invalid previous page. Expected 2 got %d", paging.previousPage)
	}

	if len(paging.pages) != 1 {
		t.Fatalf("Invalid page list. Expected lenght 3 got %d", len(paging.pages))
	}

	if paging.lastPage != 1 {
		t.Fatalf("Invalid page size. Expected 3 got %d", paging.lastPage)
	}

	if paging.enabled {
		t.Fatalf("Invalid enabled status. Expected false got %v", paging.enabled)
	}
}
