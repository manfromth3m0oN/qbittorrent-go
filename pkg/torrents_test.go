package pkg

import "testing"

func SearchCriteriaMarshalTest(t *testing.T) {
	expected := "?filter=downloading&category=MyCategory&tag=MyTag&sort=downloaded&limit=10&offset=1&hashes=12398546dfg|sdkjfgy278"

	filter := Downloading
	category := "MyCategory"
	tag := "MyTag"
	sort := "downloaded"
	limit := 10
	offset := 1
	hashes := []string{"12398546dfg", "sdkjfgy278"}
	searchCriteria := &SearchCriteria{
		filter:   &filter,
		Category: &category,
		tag:      &tag,
		sort:     &sort,
		reverse:  false,
		limit:    &limit,
		offset:   &offset,
		hashes:   &hashes,
	}

	marshaledText := searchCriteria.Marshal()

	if marshaledText != expected {
		t.Fatalf("Marshaled text does not match expected: %s", marshaledText)
	}

}
