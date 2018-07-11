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

	// TEST
	pollutants := getPollutants()
	fmt.Print(pollutants)

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
	results := combineTests(tests)
	// results is a list of:
	// Station, Param, Results

	var theResults []MeasuredLocation
	for _, result := range results {
		var theLocation *MeasuredLocation
		for i := range mls {
			if mls[i].ID == result.Station {
				theLocation = &mls[i]
				break
			}
		}

		if theLocation == nil {
			continue
		}

		ml := MeasuredLocation{
			ID:       result.Station,
			Chemical: result.Param,
			Data:     result.Results,
			Lat:      theLocation.Lat,
			Lon:      theLocation.Lon,
			Name:     theLocation.Name,
			Type:     theLocation.Type,
		}

		theResults = append(theResults, ml)
	}

	// fmt.Print(theResults)

	return theResults
}

// WATER
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

// POLLUTION
// Disposal represents a single disposal for a given facility
// Same as Test but with appropriate business names
type Disposal struct {
	Year     int
	ID       string
	Quantity float64
	Material string
}

// DisposalResult represents all the disposal values for a given station
// Similar to TestResults but with appropriate business names
type DisposalResults struct {
	ID         string
	Material   string
	Quantities []float64
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
		result, _ := strconv.ParseFloat(strings.TrimSpace(fields[2]), 64)
		year, _ := strconv.ParseInt(strings.TrimSpace(fields[3]), 10, 64)
		test := Test{
			Station: strings.TrimSpace(fields[0]),
			Param:   strings.TrimSpace(fields[1]),
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
		// fmt.Printf("%+v\n", test)
		_, ok := tmp[test.Station]
		// make a station if it's not available
		if !ok {
			// fmt.Println("wasn't a station")
			tmp[test.Station] = make(map[string]map[int][]float64)
		}

		// make a param if it's not available
		key := getKey(test.Param)
		// fmt.Printf("key is %s\n", key)
		_, ok = tmp[test.Station][key]
		if !ok {
			// fmt.Println("wasn't a param")
			tmp[test.Station][key] = make(map[int][]float64)
		}

		// fmt.Println(test.Year)
		_, ok = tmp[test.Station][key][test.Year]
		if !ok {
			tmp[test.Station][key][test.Year] = []float64{test.Result}
		} else {
			tmp[test.Station][key][test.Year] = append(tmp[test.Station][key][test.Year], test.Result)
		}
	}

	//fmt.Printf("%+v\n", tmp)

	var testResults []TestResults
	for st, t := range tmp { // t -> station

		for pa, p := range t { // p -> param
			// fmt.Println("Param is %s\n", pa)
			result := TestResults{
				Station: st,
				Param:   pa,
			}
			for i := 2000; i <= 2016; i++ {
				vals, ok := p[i]
				if !ok { // no samples for that year
					result.Results = append(result.Results, 0)
				} else {
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

	}

	// fmt.Printf("%+v\n", testResults)

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

func getPollutants() []Disposal {
	dat, err := ioutil.ReadFile("./data/disposal-on-on-site.csv")
	if err != nil {
		panic(err)
	}

	var disposals []Disposal
	lines := strings.Split(string(dat), "\n")[1:]
	for _, line := range lines {
		fields := strings.Split(line, ",")

		year, _ := strconv.ParseInt(strings.TrimSpace(fields[0]), 10, 64)
		quantity, _ := strconv.ParseFloat(strings.TrimSpace(fields[3]), 64)

		disposal := Disposal{
			Year:     int(year),
			ID:       strings.TrimSpace(fields[1]),
			Quantity: quantity,
			Material: strings.TrimSpace(fields[4]),
		}

		fmt.Print(disposal)

		disposals = append(disposals, disposal)
	}

	fmt.Printf("length of disposals: %d\n", len(disposals))
	return disposals
}
