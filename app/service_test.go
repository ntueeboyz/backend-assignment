package main

import (
	"testing"
)

// mockAdvertisement creates a mock advertisement for testing purposes.
func mockAdvertisement(title string, ageStart, ageEnd int, gender GenderType, country []ISO3166, platform []PlatformType) Advertisement {
	return Advertisement{
		Title: title,
		Conditions: Conditions{
			AgeStart: ageStart,
			AgeEnd:   ageEnd,
			Gender:   gender,
			Country:  country,
			Platform: platform,
		},
	}
}

func TestFilterAds(t *testing.T) {
	ads := []Advertisement{
		mockAdvertisement("Ad 1", 15, 30, "", []ISO3166{"JP", "KR"}, []PlatformType{"web"}),
		mockAdvertisement("Ad 2", 40, 50, "M", []ISO3166{"TW", "JP"}, []PlatformType{"android", "ios", "web"}),
		mockAdvertisement("Ad 3", 25, 35, "F", []ISO3166{}, []PlatformType{"ios"}),
		mockAdvertisement("Ad 4", 10, 20, "F", []ISO3166{"JP", "CA", "US"}, []PlatformType{"android"}),
		mockAdvertisement("Ad 5", 45, 60, "M", []ISO3166{"US", "CA", "TW"}, []PlatformType{"web", "android"}),
		mockAdvertisement("Ad 6", 35, 45, "F", []ISO3166{"JP", "US"}, []PlatformType{"ios", "android"}),
		mockAdvertisement("Ad 7", 15, 65, "M", []ISO3166{"KR", "CA", "TW"}, []PlatformType{}),
		mockAdvertisement("Ad 8", 10, 18, "F", []ISO3166{}, []PlatformType{"web", "android", "ios"}),
		mockAdvertisement("Ad 9", 0, 0, "", []ISO3166{"JP", "CA"}, []PlatformType{"web", "ios"}),
		mockAdvertisement("Ad 10", 40, 100, "", []ISO3166{"JP", "CA", "KR", "US"}, []PlatformType{}),
		// Add more mock ads as necessary for thorough testing
	}

	tests := []struct {
		name           string
		offset, limit  int
		age            int
		gender         string
		country        string
		platform       string
		expectedTitles []string // Titles of ads we expect to pass the filter
	}{
		{
			name:           "Age filter",
			limit:		  	100,
			offset: 		0,
			age:            27,
			country: 	    "",
			platform:       "",
			gender: 	   	"",
			expectedTitles: []string{"Ad 1", "Ad 3", "Ad 7", "Ad 9"},
		},
		{
			name:           "Gender filter",
			limit:		  	100,
			offset: 		0,
			age:            0,
			gender:         "F",
			country: 	    "",
			platform:       "",
			expectedTitles: []string{"Ad 1", "Ad 3", "Ad 4", "Ad 6", "Ad 8", "Ad 9", "Ad 10"},
		},
		{
			name:           "Country filter",
			limit:		  	100,
			offset: 		0,
			age:            0,
			gender:         "",
			country:        "TW",
			platform:       "",
			expectedTitles: []string{"Ad 2", "Ad 3", "Ad 5", "Ad 7", "Ad 8"},
		},
		{
			name:           "Platform filter",
			limit:		  	100,
			offset: 		0,
			age:            0,
			gender:         "",
			country:        "",
			platform:       "ios",
			expectedTitles: []string{"Ad 2", "Ad 3", "Ad 6", "Ad 7", "Ad 8", "Ad 9", "Ad 10"},
		},
		{
			name:           "Combined filters",
			limit:		  	3,
			offset: 		0,
			age:            20,
			gender:         "F",
			country:        "CA",
			platform:       "android",
			expectedTitles: []string{"Ad 4"},
		},
		{
			name:           "Offset and limit",
			offset:         1,
			limit:          3,
			age:            0,
			gender:         "",
			country:        "",
			platform:       "",
			expectedTitles: []string{"Ad 2", "Ad 3", "Ad 4"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			filteredAds := FilterAds(ads, tc.offset, tc.limit, tc.age, tc.gender, tc.country, tc.platform)
			if len(filteredAds) != len(tc.expectedTitles) {
				t.Fatalf("Expected %d ads, got %d", len(tc.expectedTitles), len(filteredAds))
			}
			for i, ad := range filteredAds {
				if ad.Title != tc.expectedTitles[i] {
					t.Errorf("Expected %s, got %s", tc.expectedTitles[i], ad.Title)
				}
			}
		})
	}
}
