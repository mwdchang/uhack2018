package main

import (
	"fmt"
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
		lat, _ := strconv.ParseFloat(strings.TrimSpace(fields[11]), 64)
		lon, _ := strconv.ParseFloat(strings.TrimSpace(fields[12]), 64)
		mls = append(mls, MeasuredLocation{
			Lat: lat,
			Lon: lon,
			ID:  fields[1],
			//Name: fmt.Sprintf("%s | %s", fields[2], fields[3]),
			Name: fields[2],
			Type: "facility",
		})
	}
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
		lon, _ := strconv.ParseFloat(strings.TrimSpace(fields[3]), 64)
		lat, _ := strconv.ParseFloat(strings.TrimSpace(fields[4]), 64)

		mls = append(mls, MeasuredLocation{
			Lat:  lat,
			Lon:  lon,
			ID:   fields[0],
			Name: fields[1],
			Type: "water",
		})
	}

	tests := getTests()
	combineTests(tests[0:1])

	//fmt.Print(results)

	return mls
}

// Test represents a single test for a given station
type Test struct {
	Station string
	Param   string
	Result  float64
	Year    int
}

// TestResults represents all the tests for a given station
type TestResults struct {
	Station string
	Param   string
	Results []float64 // one per year
}

func getTests() []Test {
	dat, err := ioutil.ReadFile("./data/water-samples-all.csv")
	if err != nil {
		panic(err)
	}

	var tests []Test
	lines := strings.Split(string(dat), "\n")[1:]
	for _, line := range lines {
		fields := strings.Split(line, ",")
		result, _ := strconv.ParseFloat(fields[2], 64)
		year, _ := strconv.ParseInt(fields[3], 10, 64)
		fmt.Printf("James getTests field is %s year is %d\n", fields[3], year)
		test := Test{
			Station: fields[0],
			Param:   fields[1],
			Result:  result,
			Year:    int(year),
		}

		tests = append(tests, test)
	}

	return tests
}

// combineTests combines a list of tests into a testResults
func combineTests(tests []Test) []TestResults {
	//      station    param      year
	tmp := make(map[string]map[string]map[int][]float64)
	for _, test := range tests {
		_, ok := tmp[test.Station]
		// make a station if it's not available
		if !ok {
			tmp[test.Station] = make(map[string]map[int][]float64)
		}
		s, _ := tmp[test.Station]

		// make a param if it's not available
		key := getKey(test.Param)
		_, ok = s[key]
		if !ok {
			s[key] = make(map[int][]float64)
		}
		p, _ := s[key]

		_, ok = p[test.Year]
		if !ok {
			p[test.Year] = []float64{test.Result}
		} else {
			p[test.Year] = append(p[test.Year], test.Result)
		}

	}

	var testResults []TestResults
	for st, t := range tmp { // t -> station
		result := TestResults{
			Station: st,
		}

		for _, p := range t { // p -> param
			for i := 2000; i <= 2016; i++ {
				vals, ok := p[i]
				if !ok { // no samples for that year
					result.Results = append(result.Results, 0)
				}
				cnt := float64(len(vals))
				tot := 0.0
				for _, v := range vals {
					tot += v
				}
				result.Results = append(result.Results, tot/cnt)
			}
		}

		testResults = append(testResults, result)
	}

	return testResults
}

// getKey gets a type of contaminant for a given code
func getKey(code string) string {
	if code == "PBUT" {
		return "LEAD"
	}
	if code == "CRUT" {
		return "CHROMIUM"
	}
	if code == "ASUT" {
		return "ARSENIC"
	}

	panic(fmt.Sprintf("Unknown code %s\n", code))
}
