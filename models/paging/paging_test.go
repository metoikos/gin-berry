package paging

import (
	"testing"
)

func TestBuildPaging(t *testing.T) {
	paging := New(1, 20, 44)

	if paging.Offset != 0 {
		t.Fatalf("Invalid offset. Expected 0 got %d", paging.Offset)
	}

	if paging.NextPage != 2 {
		t.Fatalf("Invalid next page. Expected 2 got %d", paging.NextPage)
	}

	if paging.PreviousPage != 0 {
		t.Fatalf("Invalid previous page. Expected 2 got %d", paging.PreviousPage)
	}

	if len(paging.Pages) != 3 {
		t.Fatalf("Invalid page list. Expected lenght 3 got %d", len(paging.Pages))
	}

	if paging.LastPage != 3 {
		t.Fatalf("Invalid page size. Expected 3 got %d", paging.LastPage)
	}

	if !paging.Enabled {
		t.Fatalf("Invalid enabled status. Expected true got %v", paging.Enabled)
	}
}

func TestBuildPaging2(t *testing.T) {
	paging := New(3, 20, 44)

	if paging.Offset != 40 {
		t.Fatalf("Invalid offset. Expected 0 got %d", paging.Offset)
	}

	if paging.NextPage != 0 {
		t.Fatalf("Invalid next page. Expected 2 got %d", paging.NextPage)
	}

	if paging.PreviousPage != 2 {
		t.Fatalf("Invalid previous page. Expected 2 got %d", paging.PreviousPage)
	}

	if len(paging.Pages) != 3 {
		t.Fatalf("Invalid page list. Expected lenght 3 got %d", len(paging.Pages))
	}

	if paging.LastPage != 3 {
		t.Fatalf("Invalid page size. Expected 3 got %d", paging.LastPage)
	}

	if !paging.Enabled {
		t.Fatalf("Invalid enabled status. Expected true got %v", paging.Enabled)
	}
}

func TestBuildPaging3(t *testing.T) {
	paging := New(1, 20, 19)

	if paging.Offset != 0 {
		t.Fatalf("Invalid offset. Expected 0 got %d", paging.Offset)
	}

	if paging.NextPage != 0 {
		t.Fatalf("Invalid next page. Expected 2 got %d", paging.NextPage)
	}

	if paging.PreviousPage != 0 {
		t.Fatalf("Invalid previous page. Expected 2 got %d", paging.PreviousPage)
	}

	if len(paging.Pages) != 1 {
		t.Fatalf("Invalid page list. Expected lenght 3 got %d", len(paging.Pages))
	}

	if paging.LastPage != 1 {
		t.Fatalf("Invalid page size. Expected 3 got %d", paging.LastPage)
	}

	if paging.Enabled {
		t.Fatalf("Invalid enabled status. Expected false got %v", paging.Enabled)
	}
}

func TestBuildPaging4(t *testing.T) {
	paging := New(3, 20, 19)

	if paging.Offset != 0 {
		t.Fatalf("Invalid offset. Expected 0 got %d", paging.Offset)
	}

	if paging.NextPage != 0 {
		t.Fatalf("Invalid next page. Expected 2 got %d", paging.NextPage)
	}

	if paging.PreviousPage != 0 {
		t.Fatalf("Invalid previous page. Expected 2 got %d", paging.PreviousPage)
	}

	if len(paging.Pages) != 1 {
		t.Fatalf("Invalid page list. Expected lenght 3 got %d", len(paging.Pages))
	}

	if paging.LastPage != 1 {
		t.Fatalf("Invalid page size. Expected 3 got %d", paging.LastPage)
	}

	if paging.Enabled {
		t.Fatalf("Invalid enabled status. Expected false got %v", paging.Enabled)
	}
}
