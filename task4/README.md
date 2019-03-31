<h3>Console-running application

**Purpose: to estimate http-response duration.** 

* To run the application on your system:
  * Run `go build` to create the binary (`main`)
  * Run the binary `./main` with necessary parameters via flags.
  * To view allowable flags run the binary without any.

* e.g. `./main -url epam.com -url stackoverflow.com -num 3 -timeout 1.5
`