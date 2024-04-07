package main

import "time"

// FilterAds filters the advertisements based on given criteria.
func FilterAds(ads []Advertisement, offset int, limit int, age int, gender, country, platform string) []ResAd {
	var filteredAds []ResAd
    var responseAd ResAd
    now := time.Now()

	for _, ad := range ads {

        // Check if the advertisement is active
        if now.Before(ad.StartAt) || now.After(ad.EndAt) {
			continue
		}

        // Check if the advertisement meets the age criteria
        if age != 0 {
            if ad.Conditions.AgeStart > age || age > ad.Conditions.AgeEnd && (ad.Conditions.AgeStart != 0 && ad.Conditions.AgeEnd != 0) {
                continue
            }
        }

        // Check if the advertisement meets the gender criteria
        if gender != "" {
            if string(ad.Conditions.Gender) != gender && ad.Conditions.Gender != "" {
                continue
            }
        }

        // check if the advertisement meets the country and platform criteria
        if country != "" {
            flag := false
            for _, c := range ad.Conditions.Country {
                if string(c) == country {
                    flag = true
                    break
                }
            }
            
            if !flag && len(ad.Conditions.Country) > 0 {
                continue
            }
        }

        if platform != "" {
            flag := false
            for _, p := range ad.Conditions.Platform {
                if string(p) == platform {
                    flag = true
                    break
                }
            }
            
            if !flag && len(ad.Conditions.Platform) > 0{
                continue
            }
        }

        responseAd.Title = ad.Title
        responseAd.EndAt = ad.EndAt

		filteredAds = append(filteredAds, responseAd)
	}

	// Apply offset and limit
	end := offset + limit
	if end > len(filteredAds) {
		end = len(filteredAds)
	}
	filteredAds = filteredAds[offset:end]

	return filteredAds
}