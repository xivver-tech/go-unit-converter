# Go Unit Converter

A practical command-line unit converter supporting length, weight and temperature.

## Features
- Length: mm, cm, m, km, in, ft, yd, mi
- Weight: mg, g, kg, t, oz, lb
- Temperature: Celsius, Fahrenheit, Kelvin

## How to run
```bash
go run main.go 100 cm m
go run main.go 32 f c
go run main.go 5.5 lb kg
go run main.go 10 km mi
```

Or build it:
```bash
go build -o converter
./converter 100 cm m
```
