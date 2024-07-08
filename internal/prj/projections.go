package prj

// https://epsg.io/31982
var BrazilCs = CoordinateSystem{
	Name: "SIRGAS_2000_UTM_Zone_22S",
	GeoCoordinateSystem: GeoCoordinateSystem{
		Name: "GCS_SIRGAS_2000",
		Datum: Datum{
			Name: "D_SIRGAS_2000",
			Spheroid: Spheroid{
				Name:              "GRS_1980",
				SemiMajorAxis:     6378137,
				InverseFlattening: 298.257222101,
			},
		},
		PrimeMeridian: Primem{
			Name:  "Greenwich",
			Value: 0,
		},
		Unit: Unit{
			Name:  "Degree",
			Value: 0.0174532925199433,
		},
	},
	Projection:       "Transverse_Mercator",
	FalseEasting:     500000,
	FalseNorthing:    10000000,
	CentralMeridian:  -51,
	ScaleFactor:      0.9996,
	LatitudeOfOrigin: 0,
	Unit: Unit{
		Name:  "Meter",
		Value: 1,
	},
}

// https://epsg.io/28354
var AustraliaCs = CoordinateSystem{
	Name: "GDA94 / MGA zone 54",
	GeoCoordinateSystem: GeoCoordinateSystem{
		Name: "GDA94",
		Datum: Datum{
			Name: "Geocentric_Datum_of_Australia_1994",
			Spheroid: Spheroid{
				Name:              "GRS_1980",
				SemiMajorAxis:     6378137,
				InverseFlattening: 298.257222101,
			},
		},
		PrimeMeridian: Primem{
			Name:  "Greenwich",
			Value: 0,
		},
		Unit: Unit{
			Name:  "Degree",
			Value: 0.0174532925199433,
		},
	},
	Projection:       "Transverse_Mercator",
	FalseEasting:     500000,
	FalseNorthing:    10000000,
	CentralMeridian:  141,
	ScaleFactor:      0.9996,
	LatitudeOfOrigin: 0,
	Unit: Unit{
		Name:  "Meter",
		Value: 1,
	},
}

// https://epsg.io/27700
var BritainCs = CoordinateSystem{
	Name: "OSGB36 / British National Grid",
	GeoCoordinateSystem: GeoCoordinateSystem{
		Name: "OSGB36",
		Datum: Datum{
			Name: "Ordnance_Survey_of_Great_Britain_1936",
			Spheroid: Spheroid{
				Name:              "Airy 1830",
				SemiMajorAxis:     6377563.396,
				InverseFlattening: 299.3249646,
			},
		},
		PrimeMeridian: Primem{
			Name:  "Greenwich",
			Value: 0,
		},
		Unit: Unit{
			Name:  "Degree",
			Value: 0.0174532925199433,
		},
	},
	Projection:       "Transverse_Mercator",
	FalseEasting:     400000,
	FalseNorthing:    -100000,
	CentralMeridian:  -2,
	ScaleFactor:      0.9996012717,
	LatitudeOfOrigin: 49,
	Unit: Unit{
		Name:  "Meter",
		Value: 1,
	},
}

var WgsSphere = Spheroid{
	Name:              "WGS 84",
	SemiMajorAxis:     6378137,
	InverseFlattening: 298.257223563,
}
