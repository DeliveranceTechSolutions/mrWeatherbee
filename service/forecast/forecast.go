package forecast

import (
	"github.com/deliveranceTechSolutions/mrWeatherbee/service/location"
)

// ForecastView represents the bearings for a major city for national forecast views
type Bearings location.Bearings
type ForecastView struct {
	position Bearings
}

// assembleForecastView returns a View struct nessecary to calculate the mapbox bounds for each corner.
func assembleForecastView(abbreviation string) ForecastView {
	states := map[string]ForecastView{
		"Seattle": {
		    position: Bearings{
			Latitude: "47.6062",
			Longitude: "-122.3321",
		    },
		},

		"Cut Bank": {
		    position: Bearings{
			Latitude: "48.6346",
			Longitude: "-112.3256",
		    },
		},

		"Minot": {
		    position: Bearings{
			Latitude: "48.2325",
			Longitude: "-101.2963",
		    },
		},

		"International Falls": {
		    position: Bearings{
			Latitude: "48.6015",
			Longitude: "-93.4105",
		    },
		},

		"Minneapolis": {
		    position: Bearings{
			Latitude: "44.9778",
			Longitude: "-93.2650",
		    },
		},

		"Detroit": {
		    position: Bearings{
			Latitude: "42.3314",
			Longitude: "-83.0458",
		    },
		},

		"Chicago": {
		    position: Bearings{
			Latitude: "41.8781",
			Longitude: "-87.6298",
		    },
		},

		"Cincinnati": {
		    position: Bearings{
			Latitude: "39.1031",
			Longitude: "-84.5120",
		    },
		},

		"Charlotte": {
		    position: Bearings{
			Latitude: "35.2271",
			Longitude: "-80.8431",
		    },
		},

		"Norfolk": {
		    position: Bearings{
			Latitude: "36.8508",
			Longitude: "-76.2859",
		    },
		},

		"New York": {
		    position: Bearings{
			Latitude: "40.7128",
			Longitude: "-74.0060",
		    },
		},

		"Buffalo": {
		    position: Bearings{
			Latitude: "42.8864",
			Longitude: "-78.8784",
		    },
		},

		"St. Louis": {
		    position: Bearings{
			Latitude: "38.6270",
			Longitude: "-90.1994",
		    },
		},

		"Memphis": {
		    position: Bearings{
			Latitude: "35.1495",
			Longitude: "-90.0490",
		    },
		},

		"Atlanta": {
		    position: Bearings{
			Latitude: "33.7490",
			Longitude: "-84.3880",
		    },
		},

		"Jacksonville": {
		    position: Bearings{
			Latitude: "30.3322",
			Longitude: "-81.6557",
		    },
		},

		"Tampa": {
		    position: Bearings{
			Latitude: "27.9506",
			Longitude: "-82.4572",
		    },
		},

		"Miami": {
		    position: Bearings{
			Latitude: "25.7617",
			Longitude: "-80.1918",
		    },
		},

		"New Orleans": {
		    position: Bearings{
			Latitude: "29.9511",
			Longitude: "-90.0715",
		    },
		},

		"Houston": {
		    position: Bearings{
			Latitude: "29.7604",
			Longitude: "-95.3698",
		    },
		},

		"Brownsville": {
		    position: Bearings{
			Latitude: "25.9018",
			Longitude: "-97.4975",
		    },
		},

		"San Antonio": {
		    position: Bearings{
			Latitude: "29.4241",
			Longitude: "-98.4936",
		    },
		},

		"El Paso": {
		    position: Bearings{
			Latitude: "31.7619",
			Longitude: "-106.4850",
		    },
		},

		"Lubbock": {
		    position: Bearings{
			Latitude: "33.5779",
			Longitude: "-101.8552",
		    },
		},

		"Albuquerque": {
		    position: Bearings{
			Latitude: "35.0844",
			Longitude: "-106.6504",
		    },
		},

		"Phoenix": {
		    position: Bearings{
			Latitude: "33.4484",
			Longitude: "-112.0740",
		    },
		},

		"Las Vegas": {
		    position: Bearings{
			Latitude: "36.1699",
			Longitude: "-115.1398",
		    },
		},

		"Los Angeles": {
		    position: Bearings{
			Latitude: "34.0522",
			Longitude: "-118.2437",
		    },
		},

		"San Francisco": {
		    position: Bearings{
			Latitude: "37.7749",
			Longitude: "-122.4194",
		    },
		},

		"Reno": {
		    position: Bearings{
			Latitude: "39.5296",
			Longitude: "-119.8138",
		    },
		},

		"Medford": {
		    position: Bearings{
			Latitude: "42.3265",
			Longitude: "-122.8756",
		    },
		},

		"Boise": {
		    position: Bearings{
			Latitude: "43.6150",
			Longitude: "-116.2023",
		    },
		},

		"Salt Lake City": {
		    position: Bearings{
			Latitude: "40.7608",
			Longitude: "-111.8910",
		    },
		},

		"Denver": {
		    position: Bearings{
			Latitude: "39.7392",
			Longitude: "-104.9903",
		    },
		},

		"Billings": {
		    position: Bearings{
			Latitude: "45.7833",
			Longitude: "-108.5007",
		    },
		},

		"Rapid City": {
		    position: Bearings{
			Latitude: "44.0805",
			Longitude: "-103.2310",
		    },
		},

		"Kansas City": {
		    position: Bearings{
			Latitude: "39.0997",
			Longitude: "-94.5786",
		    },
		},

		"Oklahoma City": {
		    position: Bearings{
			Latitude: "35.4676",
			Longitude: "-97.5164",
		    },
		},

		"Dallas": {
		    position: Bearings{
			Latitude: "32.7767",
			Longitude: "-96.7970",
		    },
		},

		"Boston": {
		    position: Bearings{
			Latitude: "42.3601",
			Longitude: "-71.0589",
		    },
		},
	}

	return states[abbreviation]
}
