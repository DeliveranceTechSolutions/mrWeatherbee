package service

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

	
}

func generateCoordinates(center Bearings, zoomDistance int16) LatLong {
	/*	Do some calculations here	*/

	return LatLong{
		Coordinates{
			Bearings{
				latitude: "91.91",
				longitude: "91.91",
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

func assembleStatesMap(state string) View {
	m := map[string]View{
		"AL": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"AK": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"AZ": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"AR": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"CA": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"CO": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"CT": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"DE": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"FL": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"GA": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"HI": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"ID": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"IL": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"IN": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"IA": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"KS": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"KY": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"LA": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"ME": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"MD": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"MA": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"MI": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"MN": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"MT": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"NE": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"NV": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"NH": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"NJ": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"NM": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"NY": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"NC": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"ND": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"OH": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"OK": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"OR": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"PA": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"RI": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"SC": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"SD": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"TN": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"TX": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"UT": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"VT": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"VA": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"WA": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"WV": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"WI": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
		"WY": {
			center: Bearings{
				latitude: "91.91",
				longitude: "91.91",
			},
		},
	}

	return m[state]
}
