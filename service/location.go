package service

import (
	"fmt"
	"strconv"
)

type Location struct {
	city		string
	state		string		
	country 	string
	coordinates	[2]Coordinates
}

type Bearings struct {
	latitude	string
	longitude	string
}

type LatLng [2]Coordinates
type Coordinates [2]Bearings
func NewCore() Location {
	view := assembleView("WA")
	return Location{ 
		coordinates: generateCoordinates(view),
	}
}

func assembleView(abbreviation string) View {
	return assembleStateView(abbreviation)
}

func generateCoordinates(state View) LatLng {
	/*	Do some calculations here	*/

	sqm := state.sqmiles
	rad := ((sqm / 2) / 2) / 54
	lat, err := strconv.ParseFloat(state.bearings.latitude, 64)
	if err != nil {
		fmt.Println(err)
	}

	lng, err := strconv.ParseFloat(state.bearings.longitude, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("lat: %f, lng: %f", lat + float64(rad), lng + float64(rad))
	eastLat := strconv.Itoa(int(lat + float64(rad)))
	eastLng := strconv.Itoa(int(lng + float64(rad)))
	westLat := strconv.Itoa(int(float64(rad) - lat))
	westLng := strconv.Itoa(int(float64(rad) - lng))
	return LatLng{ 
		Coordinates{
			Bearings{
				latitude: westLat,
				longitude: westLng,
			},
			Bearings{
				latitude: eastLat,
				longitude: eastLng,
			},
		},
	}
}

type View struct {
	bearings Bearings
	sqmiles int32
}	
type States map[string]View
func assembleStateView(abbreviation string) View {
	states := map[string]View{
		"AL": {
			bearings: Bearings{
				latitude: "32.7794",
				longitude: "86.8287",
			},
			sqmiles: 52420,
		},
		
		"AK": {
			bearings: Bearings{
				latitude: "64.0685",
				longitude: "152.2782",
			},
			sqmiles: 665384,
		},
		
		"AZ": {
			bearings: Bearings{
				latitude: "34.2744",
				longitude: "111.6602",
			},
			sqmiles: 113990,
		},
		
		"AR": {
			bearings: Bearings{
				latitude: "34.8938",
				longitude: "92.4426",
			},
			sqmiles: 53179,
		},
		
		"CA": {
			bearings: Bearings{
				latitude: "37.1841",
				longitude: "119.4696",
			},
			sqmiles: 163695,
		},
		
		"CO": {
			bearings: Bearings{
				latitude: "38.9972",
				longitude: "105.5478",
			},
			sqmiles: 104094,
		},
		
		"CT": {
			bearings: Bearings{
				latitude: "41.6219",
				longitude: "72.7273",
			},
			sqmiles: 5543,
		},
		
		"DE": {
			bearings: Bearings{
				latitude: "38.9896",
				longitude: "75.5050",
			},
			sqmiles: 2489,
		},
		
		"FL": {
			bearings: Bearings{
				latitude: "28.6305",
				longitude: "82.4497",
			},
			sqmiles: 65758,
		},
		
		"GA": {
			bearings: Bearings{
				latitude: "32.6415",
				longitude: "83.4426",
			},
			sqmiles: 59425,
		},
		
		"HI": {
			bearings: Bearings{
				latitude: "20.2927",
				longitude: "156.3737",
			},
			sqmiles: 10932,
		},
		
		"ID": {
			bearings: Bearings{
				latitude: "44.3509",
				longitude: "114.6130",
			},
			sqmiles: 83569,
		},
		
		"IL": {
			bearings: Bearings{
				latitude: "40.0417",
				longitude: "89.1965",
			},
			sqmiles: 57914,
		},
		
		"IN": {
			bearings: Bearings{
				latitude: "39.8942",
				longitude: "86.2816",
			},
			sqmiles: 36420,
		},
		
		"IA": {
			bearings: Bearings{
				latitude: "42.0751",
				longitude: "93.4960",
			},
			sqmiles: 56273,
		},
		
		"KS": {
			bearings: Bearings{
				latitude: "38.4937",
				longitude: "98.3804",
			},
			sqmiles: 82278,
		},
		
		"KY": {
			bearings: Bearings{
				latitude: "37.5347",
				longitude: "85.3021",
			},
			sqmiles: 40408,
		},
		
		"LA": {
			bearings: Bearings{
				latitude: "31.0689",
				longitude: "91.9968",
			},
			sqmiles: 52378,
		},
		
		"ME": {
			bearings: Bearings{
				latitude: "45.3695",
				longitude: "69.2428",
			},
			sqmiles: 35380,
		},
		
		"MD": {
			bearings: Bearings{
				latitude: "39.0550",
				longitude: "76.7909",
			},
			sqmiles: 12406,
		},
		
		"MA": {
			bearings: Bearings{
				latitude: "42.2596",
				longitude: "71.8083",
			},
			sqmiles: 10554,
		},
		
		"MI": {
			bearings: Bearings{
				latitude: "44.3467",
				longitude: "85.4102",
			},
			sqmiles: 96714,
		},
		
		"MN": {
			bearings: Bearings{
				latitude: "46.2807",
				longitude: "94.3053",
			},
			sqmiles: 86936,
		},
		
		"MS": {
			bearings: Bearings{
				latitude: "32.7364",
				longitude: "89.6678",
			},
			sqmiles: 48432,
		},

		"MO": {
			bearings: Bearings{
				latitude: "39.3029",
				longitude: "91.3144",
			},
			sqmiles: 69707,
		},
		
		"MT": {
			bearings: Bearings{
				latitude: "47.0527",
				longitude: "109.6333",
			},
			sqmiles: 147040,
		},
		
		"NE": {
			bearings: Bearings{
				latitude: "41.5378",
				longitude: "99.7951",
			},
			sqmiles: 77348,
		},
		
		"NV": {
			bearings: Bearings{
				latitude: "39.3289",
				longitude: "116.6312",
			},
			sqmiles: 110572,
		},
		
		"NH": {
			bearings: Bearings{
				latitude: "43.6805",
				longitude: "71.5811",
			},
			sqmiles: 9349,
		},
		
		"NJ": {
			bearings: Bearings{
				latitude: "40.1907",
				longitude: "74.6728",
			},
			sqmiles: 8723,
		},
		
		"NM": {
			bearings: Bearings{
				latitude: "34.4071",
				longitude: "106.1126",
			},
			sqmiles: 121590,
		},
		
		"NY": {
			bearings: Bearings{
				latitude: "42.9538",
				longitude: "75.5268",
			},
			sqmiles: 54555,
		},
		
		"NC": {
			bearings: Bearings{
				latitude: "35.5557",
				longitude: "79.3877",
			},
			sqmiles: 53819,
		},
		
		"ND": {
			bearings: Bearings{
				latitude: "47.4501",
				longitude: "100.4659",
			},
			sqmiles: 70698,
		},
		
		"OH": {
			bearings: Bearings{
				latitude: "40.2862",
				longitude: "82.7937",
			},
			sqmiles: 44826,
		},
		
		"OK": {
			bearings: Bearings{
				latitude: "35.5889",
				longitude: "97.4943",
			},
			sqmiles: 69899,
		},
		
		"OR": {
			bearings: Bearings{
				latitude: "43.9336",
				longitude: "120.5583",
			},
			sqmiles: 98379,
		},
		
		"PA": {
			bearings: Bearings{
				latitude: "40.8781",
				longitude: "77.7996",
			},
			sqmiles: 46054,
		},
		
		"RI": {
			bearings: Bearings{
				latitude: "41.6762",
				longitude: "71.5562",
			},
			sqmiles: 1545,
		},
		
		"SC": {
			bearings: Bearings{
				latitude: "33.9169",
				longitude: "80.8964",
			},
			sqmiles: 32020,
		},
		
		"SD": {
			bearings: Bearings{
				latitude: "44.4443",
				longitude: "100.2263",
			},
			sqmiles: 77116,
		},
		
		"TN": {
			bearings: Bearings{
				latitude: "35.8580",
				longitude: "86.3505",
			},
			sqmiles: 42144,
		},
		
		"TX": {
			bearings: Bearings{
				latitude: "31.4757",
				longitude: "99.3312",
			},
			sqmiles: 268596,
		},
		
		"UT": {
			bearings: Bearings{
				latitude: "39.3055",
				longitude: "111.6703",
			},
			sqmiles: 84897,
		},
		
		"VT": {
			bearings: Bearings{
				latitude: "44.0687",
				longitude: "72.6658",
			},
			sqmiles: 9616,
		},
		
		"VA": {
			bearings: Bearings{
				latitude: "37.5215",
				longitude: "78.8537",
			},
			sqmiles: 42775,
		},
		
		"WA": {
			bearings: Bearings{
				latitude: "47.3826",
				longitude: "120.4472",
			},
			sqmiles: 71298,
		},
		
		"WV": {
			bearings: Bearings{
				latitude: "38.6409",
				longitude: "80.6227",
			},
			sqmiles: 24230,
		},
		
		"WI": {
			bearings: Bearings{
				latitude: "44.6243",
				longitude: "89.9941",
			},
			sqmiles: 65496,
		},
		
		"WY": {
			bearings: Bearings{
				latitude: "42.9957",
				longitude: "107.5512",
			},
			sqmiles: 97813,
		},
		
	}

	return states[abbreviation]
}
