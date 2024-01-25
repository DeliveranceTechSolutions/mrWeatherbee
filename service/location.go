package service

import (
	"fmt"
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

type Coordinates [2]Bearings
type LatLong [2]Coordinates

func NewCore() Location {
	assembleView()
	return Location{ 
		coordinates: generateCoordinates(
			Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
			19,
		),
	}
}

func assembleView() {
	/*

		Access the API here for osm, mapbox, etc.

	*/
	stateAbbrev := "WA"
	stateMap := assembleStateView(stateAbbrev)
	
	fmt.Println(stateMap)
		
}

func generateCoordinates(center Bearings, zoomDistance int16) LatLong {
	/*	Do some calculations here	*/

	return LatLong{
		Coordinates{
			Bearings{
				latitude: "32.7794",
				longitude: "86.8287",
			},
			Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
	}
}

type View struct {
	center Bearings
}
type States map[string]View

func assembleStateView(state string) View {
	m := map[string]View{
		"AL": {
			center: Bearings{
				latitude: "32.7794",
				longitude: "86.8287",
			},
		},
		"AK": {
			center: Bearings{
				latitude: "64.0685",
				longitude: "152.2782",
			},
		},
		"AZ": {
			center: Bearings{
				latitude: "34.2744",
				longitude: "111.6602",
			},
		},
		"AR": {
			center: Bearings{
				latitude: "34.8938",
				longitude: "92.4426",
			},
		},
		"CA": {
			center: Bearings{
				latitude: "37.1841",
				longitude: "119.4696",
			},
		},
		"CO": {
			center: Bearings{
				latitude: "38.9972",
				longitude: "105.5478",
			},
		},
		"CT": {
			center: Bearings{
				latitude: "41.6219",
				longitude: "72.7273",
			},
		},
		"DE": {
			center: Bearings{
				latitude: "38.9896",
				longitude: "75.5050",
			},
		},
		"FL": {
			center: Bearings{
				latitude: "28.6305",
				longitude: "82.4497",
			},
		},
		"GA": {
			center: Bearings{
				latitude: "32.6415",
				longitude: "83.4426",
			},
		},
		"HI": {
			center: Bearings{
				latitude: "20.2927",
				longitude: "156.3737",
			},
		},
		"ID": {
			center: Bearings{
				latitude: "44.3509",
				longitude: "114.6130",
			},
		},
		"IL": {
			center: Bearings{
				latitude: "40.0417",
				longitude: "89.1965",
			},
		},
		"IN": {
			center: Bearings{
				latitude: "39.8942",
				longitude: "86.2816",
			},
		},
		"IA": {
			center: Bearings{
				latitude: "42.0751",
				longitude: "93.4960",
			},
		},
		"KS": {
			center: Bearings{
				latitude: "38.4937",
				longitude: "98.3804",
			},
		},
		"KY": {
			center: Bearings{
				latitude: "37.5347",
				longitude: "85.3021",
			},
		},
		"LA": {
			center: Bearings{
				latitude: "31.0689",
				longitude: "91.9968",
			},
		},
		"ME": {
			center: Bearings{
				latitude: "45.3695",
				longitude: "69.2428",
			},
		},
		"MD": {
			center: Bearings{
				latitude: "39.0550",
				longitude: "76.7909",
			},
		},
		"MA": {
			center: Bearings{
				latitude: "42.2596",
				longitude: "71.8083",
			},
		},
		"MI": {
			center: Bearings{
				latitude: "44.3467",
				longitude: "85.4102",
			},
		},
		"MN": {
			center: Bearings{
				latitude: "46.2807",
				longitude: "94.3053",
			},
		},
		"MS": {
			center: Bearings{
				latitude: "32.7364",
				longitude: "89.6678",
			},
		},
		"MT": {
			center: Bearings{
				latitude: "47.0527",
				longitude: "109.6333",
			},
		},
		"NE": {
			center: Bearings{
				latitude: "41.5378",
				longitude: "99.7951",
			},
		},
		"NV": {
			center: Bearings{
				latitude: "39.3289",
				longitude: "116.6312",
			},
		},
		"NH": {
			center: Bearings{
				latitude: "43.6805",
				longitude: "71.5811",
			},
		},
		"NJ": {
			center: Bearings{
				latitude: "40.1907",
				longitude: "74.6728",
			},
		},
		"NM": {
			center: Bearings{
				latitude: "34.4071",
				longitude: "106.1126",
			},
		},
		"NY": {
			center: Bearings{
				latitude: "42.9538",
				longitude: "75.5268",
			},
		},
		"NC": {
			center: Bearings{
				latitude: "35.5557",
				longitude: "79.3877",
			},
		},
		"ND": {
			center: Bearings{
				latitude: "47.4501",
				longitude: "100.4659",
			},
		},
		"OH": {
			center: Bearings{
				latitude: "40.2862",
				longitude: "82.7937",
			},
		},
		"OK": {
			center: Bearings{
				latitude: "35.5889",
				longitude: "97.4943",
			},
		},
		"OR": {
			center: Bearings{
				latitude: "43.9336",
				longitude: "120.5583",
			},
		},
		"PA": {
			center: Bearings{
				latitude: "40.8781",
				longitude: "77.7996",
			},
		},
		"RI": {
			center: Bearings{
				latitude: "41.6762",
				longitude: "71.5562",
			},
		},
		"SC": {
			center: Bearings{
				latitude: "33.9169",
				longitude: "80.8964",
			},
		},
		"SD": {
			center: Bearings{
				latitude: "44.4443",
				longitude: "100.2263",
			},
		},
		"TN": {
			center: Bearings{
				latitude: "35.8580",
				longitude: "86.3505",
			},
		},
		"TX": {
			center: Bearings{
				latitude: "31.4757",
				longitude: "99.3312",
			},
		},
		"UT": {
			center: Bearings{
				latitude: "39.3055",
				longitude: "111.6703",
			},
		},
		"VT": {
			center: Bearings{
				latitude: "44.0687",
				longitude: "72.6658",
			},
		},
		"VA": {
			center: Bearings{
				latitude: "37.5215",
				longitude: "78.8537",
			},
		},
		"WA": {
			center: Bearings{
				latitude: "47.3826",
				longitude: "120.4472",
			},
		},
		"WV": {
			center: Bearings{
				latitude: "38.6409",
				longitude: "80.6227",
			},
		},
		"WI": {
			center: Bearings{
				latitude: "44.6243",
				longitude: "89.9941",
			},
		},
		"WY": {
			center: Bearings{
				latitude: "42.9957",
				longitude: "107.5512",
			},
		},
	}

	return m[state]
}
