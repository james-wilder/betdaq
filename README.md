# betdaq

Betdaq API client to make SOAP requests to Betdaq betting API and parse responses. See
[http://betdaqtraders.com/](http://betdaqtraders.com/) for details of the API.

Includes some example API requests (see examples/main.go). Run with:
`go run examples/betdaq_api_examples.go`
(run from the root project directory)

Run all the unit tests with `go test ./...`

Run the code generation from the project root with `go generate` (you should never need to do this but...)

You will need to create a file `config.json` (see `config_example.json` for the format)
with your account details. Note that currently you have to have an account and register
to use the API.

Using code generation to generate the structs needed from the WSDL file and the client
methods. I tried a couple of off-the-shelf WSDL to code packages but none worked so I
rolled my own that is just about enough for the task.

TODO:
* why is GetPricesRequest.ThresholdAmount is string?
* hard-code simpleType PolarityEnum
