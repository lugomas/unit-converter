# Unit Converter Web App
`unit-converter` is a simple web application designed to help users convert between different units of measurement. It supports conversions for units of length, weight, temperature, and more. The app allows users to input a value, select the units to convert from and to, and displays the converted result.

## Features
- Convert between various units of measurement, including:
    - Length: millimeter, centimeter, meter, kilometer, inch, foot, yard, mile.
    - Weight: gram, kilogram, ounce, pound.
    - Temperature: Celsius, Fahrenheit
- Simple web interface to input values, select conversion units, and view the result.
- Responsive design that works across different devices.

## Installation
To install and run the `unit-converter` app locally, clone the repository and build the Go binary, following the steps below:
```
git clone https://github.com/lugomas/unit-converter.git
cd unit-converter
go build -o unit-converter
```

## Usage
1. Start the web app by running the following command in your terminal:
   ```./unit-converter```
2. Open your web browser and enter the following url address:
   ```localhost:8080```
3. Choose the desired category of units (e.g., Length, Weight, Temperature).
4. Enter the value to be converted.
5. Select the units to convert from and to.
6. Click the convert button to view the converted value.

For example, converting 1 kilometer to miles will display the equivalent in miles.

## License
This project is licensed under the MIT License.

## Project Inspiration
This project was developed based on the guidelines provided by [roadmap.sh's Unit Converter project](https://roadmap.sh/projects/unit-converter)