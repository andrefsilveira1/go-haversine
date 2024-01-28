# go-Haversine

The haversine formula is a very accurate way of computing distances between two points on the surface of a sphere using the latitude and longitude of the two points.
The haversine formula is a re-formulation of the spherical law of cosines, but the formulation in terms of haversines is more useful for small angles and distances.

# How to use

- Install: `github.com/andrefsilveira1/go-haversine`
- Import `go-haversine` module into your go file


# Usage

    package main

    import (
        "fmt"

        "github.com/andrefsilveira1/go-haversine"
    )

    func main() {
        distance :=  haversine.Calculate(37.7749, -122.4194, 34.0522, -118.2437)
        fmt.Printf("The distance between the two points is %.2f km.\n", distance)
    }

After importing haversine package, use `Calculate` function an provide 4 coordinates parameters as `float64`. The distance will be returned as kilometers. The parameters must be:
(latitude of point 1, longitude of point 1, latitude of point 2, longitude of point 2)


# References

* https://en.wikipedia.org/wiki/Haversine_formula
* https://www.themathdoctors.org/distances-on-earth-2-the-haversine-formula/


# Motivation

I was coding a game that needed check the distance between two coordinate points. And as I like to create my own modules to use my stuffs in my projects, I decide to create this another module that easily calculate
the distance between two coordinate points.
