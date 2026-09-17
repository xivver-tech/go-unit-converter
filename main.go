package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Conversion factors to a base unit
var lengthToMeter = map[string]float64{
	"mm": 0.001,
	"cm": 0.01,
	"m":  1,
	"km": 1000,
	"in": 0.0254,
	"ft": 0.3048,
	"yd": 0.9144,
	"mi": 1609.344,
}

var weightToKg = map[string]float64{
	"mg": 0.000001,
	"g":  0.001,
	"kg": 1,
	"t":  1000,
	"oz": 0.0283495,
	"lb": 0.453592,
}

func convertLength(value float64, from, to string) (float64, error) {
	fromF, ok1 := lengthToMeter[from]
	toF, ok2 := lengthToMeter[to]
	if !ok1 || !ok2 {
		return 0, fmt.Errorf("unsupported length unit. Supported: mm, cm, m, km, in, ft, yd, mi")
	}
	meters := value * fromF
	return meters / toF, nil
}

func convertWeight(value float64, from, to string) (float64, error) {
	fromF, ok1 := weightToKg[from]
	toF, ok2 := weightToKg[to]
	if !ok1 || !ok2 {
		return 0, fmt.Errorf("unsupported weight unit. Supported: mg, g, kg, t, oz, lb")
	}
	kg := value * fromF
	return kg / toF, nil
}

func convertTemp(value float64, from, to string) (float64, error) {
	from = strings.ToLower(from)
	to = strings.ToLower(to)

	var celsius float64
	switch from {
	case "c", "celsius":
		celsius = value
	case "f", "fahrenheit":
		celsius = (value - 32) * 5 / 9
	case "k", "kelvin":
		celsius = value - 273.15
	default:
		return 0, fmt.Errorf("unsupported temperature unit. Use C, F or K")
	}

	switch to {
	case "c", "celsius":
		return celsius, nil
	case "f", "fahrenheit":
		return celsius*9/5 + 32, nil
	case "k", "kelvin":
		return celsius + 273.15, nil
	default:
		return 0, fmt.Errorf("unsupported temperature unit. Use C, F or K")
	}
}

func printHelp() {
	fmt.Println(`Go Unit Converter
=================

Usage:
  converter <value> <from_unit> <to_unit>

Examples:
  converter 100 cm m
  converter 32 f c
  converter 5.5 lb kg
  converter 10 km mi

Supported units:
  Length : mm, cm, m, km, in, ft, yd, mi
  Weight : mg, g, kg, t, oz, lb
  Temp   : C, F, K`)
}

func main() {
	if len(os.Args) != 4 {
		printHelp()
		os.Exit(1)
	}

	value, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Error: first argument must be a number")
		os.Exit(1)
	}

	from := strings.ToLower(os.Args[2])
	to := strings.ToLower(os.Args[3])

	var result float64

	// Try temperature first (single letter units)
	if (from == "c" || from == "f" || from == "k") && (to == "c" || to == "f" || to == "k") {
		result, err = convertTemp(value, from, to)
	} else if _, ok := lengthToMeter[from]; ok {
		result, err = convertLength(value, from, to)
	} else if _, ok := weightToKg[from]; ok {
		result, err = convertWeight(value, from, to)
	} else {
		fmt.Println("Unknown unit. Run without arguments to see help.")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("%.6g %s = %.6g %s\n", value, os.Args[2], result, os.Args[3])
}
