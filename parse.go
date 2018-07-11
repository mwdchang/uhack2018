package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// Pollutants gets a list of measuredLocations for the pollutants
func Pollutants() []MeasuredLocation {
	dat, err := ioutil.ReadFile("./data/facilities.csv")
	if err != nil {
		panic(err)
	}

	var mls []MeasuredLocation
	lines := strings.Split(string(dat), "\n")[1:]
	for _, line := range lines {
		fields := strings.Split(line, ",")
		lat, _ := strconv.ParseFloat(fields[11], 64)
		lon, _ := strconv.ParseFloat(fields[12], 64)
		mls = append(mls, MeasuredLocation{
			Lat: lat,
			Lon: lon,
			ID:  fields[1],
			//Name: fmt.Sprintf("%s | %s", fields[2], fields[3]),
			Name: fields[2],
			Type: "facility",
		})
	}
	//fmt.Print(mls)
	return mls
}

// TestSites gets a list of test sites
func TestSites() []MeasuredLocation {
	dat, err := ioutil.ReadFile("./data/test-sites.csv")
	if err != nil {
		panic(err)
	}

	var mls []MeasuredLocation
	lines := strings.Split(string(dat), "\n")[1:]
	for _, line := range lines {
		fields := strings.Split(line, ",")
		lat, _ := strconv.ParseFloat(fields[3], 64)
		lon, _ := strconv.ParseFloat(fields[4], 64)
		mls = append(mls, MeasuredLocation{
			Lat:  lat,
			Lon:  lon,
			ID:   fields[0],
			Name: fields[1],
			Type: "water",
		})
	}
	return mls
}
