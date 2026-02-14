package tempconv

var CommonUnits = []string{
	"meter", "kilometer", "centimeter", "millimeter",
	"inch", "foot", "yard", "mile",
	"gram", "kilogram", "milligram", "ton",
	"ounce", "pound",
	"second", "minute", "hour", "day",
	"liter", "milliliter", "cubic_meter",
	"gallon", "quart", "pint",
	"celsius", "fahrenheit", "kelvin",
	"square_meter", "square_kilometer", "square_foot",
	"acre", "hectare",
	"meters_per_second", "kilometers_per_hour", "miles_per_hour",
	"joule", "kilojoule", "calorie", "kilocalorie",
	"watt", "kilowatt",
	"pascal", "bar", "psi", "atmosphere",
	"bit", "byte", "kilobyte", "megabyte",
	"gigabyte", "terabyte",
	"percent", "basis_point",
}

var unitCategory = map[string]string{
	"meter": "length", "kilometer": "length", "centimeter": "length", "millimeter": "length",
	"inch": "length", "foot": "length", "yard": "length", "mile": "length",
	"gram": "mass", "kilogram": "mass", "milligram": "mass", "ton": "mass",
	"ounce": "mass", "pound": "mass",
	"second": "time", "minute": "time", "hour": "time", "day": "time",
	"liter": "volume", "milliliter": "volume", "cubic_meter": "volume",
	"gallon": "volume", "quart": "volume", "pint": "volume",
	"celsius": "temp", "fahrenheit": "temp", "kelvin": "temp",
	"square_meter": "area", "square_kilometer": "area", "square_foot": "area",
	"acre": "area", "hectare": "area",
	"meters_per_second": "speed", "kilometers_per_hour": "speed", "miles_per_hour": "speed",
	"joule": "energy", "kilojoule": "energy", "calorie": "energy", "kilocalorie": "energy",
	"watt": "power", "kilowatt": "power",
	"pascal": "pressure", "bar": "pressure", "psi": "pressure", "atmosphere": "pressure",
	"bit": "digital", "byte": "digital", "kilobyte": "digital", "megabyte": "digital",
	"gigabyte": "digital", "terabyte": "digital",
	"percent": "ratio", "basis_point": "ratio",
}

var toBase = map[string]float64{
	"meter": 1, "kilometer": 1000, "centimeter": 0.01, "millimeter": 0.001,
	"inch": 0.0254, "foot": 0.3048, "yard": 0.9144, "mile": 1609.344,
	"gram": 1, "kilogram": 1000, "milligram": 0.001, "ton": 1e6,
	"ounce": 28.3495, "pound": 453.592,
	"second": 1, "minute": 60, "hour": 3600, "day": 86400,
	"liter": 1, "milliliter": 0.001, "cubic_meter": 1000,
	"gallon": 3.78541, "quart": 0.946353, "pint": 0.473176,
	"square_meter": 1, "square_kilometer": 1e6, "square_foot": 0.092903,
	"acre": 4046.86, "hectare": 10000,
	"meters_per_second": 1, "kilometers_per_hour": 0.277778, "miles_per_hour": 0.44704,
	"joule": 1, "kilojoule": 1000, "calorie": 4.184, "kilocalorie": 4184,
	"watt": 1, "kilowatt": 1000,
	"pascal": 1, "bar": 100000, "psi": 6894.76, "atmosphere": 101325,
	"byte": 1, "bit": 0.125, "kilobyte": 1024, "megabyte": 1048576,
	"gigabyte": 1073741824, "terabyte": 1099511627776,
	"percent": 1, "basis_point": 0.01,
}

func Convert(from, to string, val float64) (float64, bool) {
	if unitCategory[from] != unitCategory[to] {
		return 0, false
	}
	if unitCategory[from] == "temp" {
		return convertTemp(from, to, val), true
	}
	f, ok1 := toBase[from]
	t, ok2 := toBase[to]
	if !ok1 || !ok2 {
		return 0, false
	}
	return val * f / t, true
}

func convertTemp(from, to string, val float64) float64 {
	var c Celsius
	switch from {
	case "celsius":
		c = Celsius(val)
	case "fahrenheit":
		c = FToC(Fahrenheit(val))
	case "kelvin":
		c = KToC(Kelvin(val))
	}
	switch to {
	case "celsius":
		return float64(c)
	case "fahrenheit":
		return float64(CToF(c))
	case "kelvin":
		return float64(CToK(c))
	}
	return 0
}
