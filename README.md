# BirdTrack

This application generates randomly-generated representations of Bird vehicles
and prints the details of vehicles traveling faster than 10 mph to stdout.

# Installation and Running

1. Clone this repository: `go get github.com/robbawebba/birdtrack`
2. Run the application:
```
$ go run github.com/robbawebba/birdtrack
Speeding bird! id: ab2d5943-a887-4ba7-b3d7-60038d2a2b1b, speed: 22, lat: 54.4194, lng: 169.2916
Speeding bird! id: 667a346c-455b-417f-b284-49d9af1f8318, speed: 18, lat: 39.3943, lng: 76.4347
Speeding bird! id: ad471aff-ee53-4e24-b15d-410125e3c423, speed: 11, lat: 27.0821, lng: 92.7383
Speeding bird! id: 02ecfd9c-f505-48d3-a407-aafb85f11b15, speed: 20, lat: 42.2001, lng: 50.9461
```

If $GOPATH/bin is added to your $PATH environment variable, then you may also
run the program with just the `birdtrack` command after running the `go get`
command in step 1.
