package location

import (
	"fmt"
	"math"
	"strconv"
)

// Location provides metadata for the frontend along with map configs
type Location struct {
	City		string
	State		string		
	Country 	string
	Coordinates	*Coordinates
}

// Bearings carry lat and lng qualities for the front end
type Bearings struct {
	Latitude	string
	Longitude	string
}

// Coordinates provides points for each corner of a mapbox view
type Coordinates struct {
	NW Bearings
	NE Bearings
	SW Bearings
	SE Bearings
}

// NewCore is a coloquialism used within each package
// To hide implementation details from func names
// Making it more difficult to replicate or spoof
func NewCore(abbreviation string) *Location {
	// Retrieve mapping of state square miles and center coordinate
	view := assembleStateView(abbreviation)
	coords := generateCoordinates(view)
	return &Location {
		Coordinates: coords,
	}	
}

// convertMilesToDegrees provides a straight conversion of miles into map degree units
func convertMilesToDegrees(sqm int32) float64 {
	// A quarter of the square miles, gives us a quadrant.
	side := math.Sqrt(float64(sqm))

	// The diagonal of a quadrant*sqrt2 is the radius of a square.
	diagonal := side * math.Sqrt2

	// Results vary as the curvature of the earth changes; however,
	// 69 is a mathematical "constant" to convert from miles to degrees.
	degrees := (diagonal / float64(2)) / float64(69)

	return degrees
}

// generateCoordinates provides each corner of a state's "view" in the mapbox
func generateCoordinates(state View) *Coordinates {
	// Get the offset degree range to form the map box.
	sqm := state.sqm
	rad := convertMilesToDegrees(sqm)
	lat, err := strconv.ParseFloat(state.center.Latitude, 64)
	if err != nil {
		fmt.Println(err)
	}

	lng, err := strconv.ParseFloat(state.center.Longitude, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(lat, lng)
	fmt.Println(rad)
	// Quadrant 1
	nwLat := strconv.Itoa(int(float64(rad) + lat))
	nwLng := strconv.Itoa(int(lng - float64(rad)))

	// Quadrant 2
	neLat := strconv.Itoa(int(float64(rad) + lat))
	neLng := strconv.Itoa(int(float64(rad) + lng))

	// Quadrant 3
	seLat := strconv.Itoa(int(lat - float64(rad)))
	seLng := strconv.Itoa(int(float64(rad) + lng))

	// Quadrant 4
	swLat := strconv.Itoa(int(lat - float64(rad)))
	swLng := strconv.Itoa(int(lng - float64(rad)))

	return &Coordinates{
		NW: Bearings{
			Latitude: nwLat,
			Longitude: nwLng,
		},
		NE: Bearings{
			Latitude: neLat,
			Longitude: neLng,
		},
		SE: Bearings{
			Latitude: seLat,
			Longitude: seLng,
		},
		SW: Bearings{
			Latitude: swLat,
			Longitude: swLng,
		},
	}
}

type States map[string]View
type View struct {
	center Bearings
	sqm int32
}	

// assembleStateView returns a View struct nessecary to calculate the mapbox bounds for each corner.
func assembleStateView(abbreviation string) View {
	states := map[string]View{
		"AL": {
			center: Bearings{
				Latitude: "32.7794",
				Longitude: "86.8287",
			},
			sqm: 52420,
		},
		
		"AK": {
			center: Bearings{
				Latitude: "64.0685",
				Longitude: "152.2782",
			},
			sqm: 665384,
		},
		
		"AZ": {
			center: Bearings{
				Latitude: "34.2744",
				Longitude: "111.6602",
			},
			sqm: 113990,
		},
		
		"AR": {
			center: Bearings{
				Latitude: "34.8938",
				Longitude: "92.4426",
			},
			sqm: 53179,
		},
		
		"CA": {
			center: Bearings{
				Latitude: "37.1841",
				Longitude: "119.4696",
			},
			sqm: 163695,
		},
		
		"CO": {
			center: Bearings{
				Latitude: "38.9972",
				Longitude: "105.5478",
			},
			sqm: 104094,
		},
		
		"CT": {
			center: Bearings{
				Latitude: "41.6219",
				Longitude: "72.7273",
			},
			sqm: 5543,
		},
		
		"DE": {
			center: Bearings{
				Latitude: "38.9896",
				Longitude: "75.5050",
			},
			sqm: 2489,
		},
		
		"FL": {
			center: Bearings{
				Latitude: "28.6305",
				Longitude: "82.4497",
			},
			sqm: 65758,
		},
		
		"GA": {
			center: Bearings{
				Latitude: "32.6415",
				Longitude: "83.4426",
			},
			sqm: 59425,
		},
		
		"HI": {
			center: Bearings{
				Latitude: "20.2927",
				Longitude: "156.3737",
			},
			sqm: 10932,
		},
		
		"ID": {
			center: Bearings{
				Latitude: "44.3509",
				Longitude: "114.6130",
			},
			sqm: 83569,
		},
		
		"IL": {
			center: Bearings{
				Latitude: "40.0417",
				Longitude: "89.1965",
			},
			sqm: 57914,
		},
		
		"IN": {
			center: Bearings{
				Latitude: "39.8942",
				Longitude: "86.2816",
			},
			sqm: 36420,
		},
		
		"IA": {
			center: Bearings{
				Latitude: "42.0751",
				Longitude: "93.4960",
			},
			sqm: 56273,
		},
		
		"KS": {
			center: Bearings{
				Latitude: "38.4937",
				Longitude: "98.3804",
			},
			sqm: 82278,
		},
		
		"KY": {
			center: Bearings{
				Latitude: "37.5347",
				Longitude: "85.3021",
			},
			sqm: 40408,
		},
		
		"LA": {
			center: Bearings{
				Latitude: "31.0689",
				Longitude: "91.9968",
			},
			sqm: 52378,
		},
		
		"ME": {
			center: Bearings{
				Latitude: "45.3695",
				Longitude: "69.2428",
			},
			sqm: 35380,
		},
		
		"MD": {
			center: Bearings{
				Latitude: "39.0550",
				Longitude: "76.7909",
			},
			sqm: 12406,
		},
		
		"MA": {
			center: Bearings{
				Latitude: "42.2596",
				Longitude: "71.8083",
			},
			sqm: 10554,
		},
		
		"MI": {
			center: Bearings{
				Latitude: "44.3467",
				Longitude: "85.4102",
			},
			sqm: 96714,
		},
		
		"MN": {
			center: Bearings{
				Latitude: "46.2807",
				Longitude: "94.3053",
			},
			sqm: 86936,
		},
		
		"MS": {
			center: Bearings{
				Latitude: "32.7364",
				Longitude: "89.6678",
			},
			sqm: 48432,
		},

		"MO": {
			center: Bearings{
				Latitude: "39.3029",
				Longitude: "91.3144",
			},
			sqm: 69707,
		},
		
		"MT": {
			center: Bearings{
				Latitude: "47.0527",
				Longitude: "109.6333",
			},
			sqm: 147040,
		},
		
		"NE": {
			center: Bearings{
				Latitude: "41.5378",
				Longitude: "99.7951",
			},
			sqm: 77348,
		},
		
		"NV": {
			center: Bearings{
				Latitude: "39.3289",
				Longitude: "116.6312",
			},
			sqm: 110572,
		},
		
		"NH": {
			center: Bearings{
				Latitude: "43.6805",
				Longitude: "71.5811",
			},
			sqm: 9349,
		},
		
		"NJ": {
			center: Bearings{
				Latitude: "40.1907",
				Longitude: "74.6728",
			},
			sqm: 8723,
		},
		
		"NM": {
			center: Bearings{
				Latitude: "34.4071",
				Longitude: "106.1126",
			},
			sqm: 121590,
		},
		
		"NY": {
			center: Bearings{
				Latitude: "42.9538",
				Longitude: "75.5268",
			},
			sqm: 54555,
		},
		
		"NC": {
			center: Bearings{
				Latitude: "35.5557",
				Longitude: "79.3877",
			},
			sqm: 53819,
		},
		
		"ND": {
			center: Bearings{
				Latitude: "47.4501",
				Longitude: "100.4659",
			},
			sqm: 70698,
		},
		
		"OH": {
			center: Bearings{
				Latitude: "40.2862",
				Longitude: "82.7937",
			},
			sqm: 44826,
		},
		
		"OK": {
			center: Bearings{
				Latitude: "35.5889",
				Longitude: "97.4943",
			},
			sqm: 69899,
		},
		
		"OR": {
			center: Bearings{
				Latitude: "43.9336",
				Longitude: "120.5583",
			},
			sqm: 98379,
		},
		
		"PA": {
			center: Bearings{
				Latitude: "40.8781",
				Longitude: "77.7996",
			},
			sqm: 46054,
		},
		
		"RI": {
			center: Bearings{
				Latitude: "41.6762",
				Longitude: "71.5562",
			},
			sqm: 1545,
		},
		
		"SC": {
			center: Bearings{
				Latitude: "33.9169",
				Longitude: "80.8964",
			},
			sqm: 32020,
		},
		
		"SD": {
			center: Bearings{
				Latitude: "44.4443",
				Longitude: "100.2263",
			},
			sqm: 77116,
		},
		
		"TN": {
			center: Bearings{
				Latitude: "35.8580",
				Longitude: "86.3505",
			},
			sqm: 42144,
		},
		
		"TX": {
			center: Bearings{
				Latitude: "31.4757",
				Longitude: "99.3312",
			},
			sqm: 268596,
		},
		
		"UT": {
			center: Bearings{
				Latitude: "39.3055",
				Longitude: "111.6703",
			},
			sqm: 84897,
		},
		
		"VT": {
			center: Bearings{
				Latitude: "44.0687",
				Longitude: "72.6658",
			},
			sqm: 9616,
		},
		
		"VA": {
			center: Bearings{
				Latitude: "37.5215",
				Longitude: "78.8537",
			},
			sqm: 42775,
		},
		
		"WA": {
			center: Bearings{
				Latitude: "47.3826",
				Longitude: "120.4472",
			},
			sqm: 71298,
		},
		
		"WV": {
			center: Bearings{
				Latitude: "38.6409",
				Longitude: "80.6227",
			},
			sqm: 24230,
		},
		
		"WI": {
			center: Bearings{
				Latitude: "44.6243",
				Longitude: "89.9941",
			},
			sqm: 65496,
		},
		
		"WY": {
			center: Bearings{
				Latitude: "42.9957",
				Longitude: "107.5512",
			},
			sqm: 97813,
		},
		
	}

	return states[abbreviation]
}
