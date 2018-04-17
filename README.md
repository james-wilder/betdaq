# betdaq

Betdaq API client to make SOAP requests to Betdaq betting API and parse responses.

See [http://betdaqtraders.com/](http://betdaqtraders.com/) for details of the API.

Includes some example API requests (see main.go)

You will need to create a file `config.json` (see `config_example.json` for the format) with your account details.
Note that currently you have to have an account and register to use the API.

Using code generation to generate the structs needed from the WSDL file and the client methods. I tried a couple of off-the-shelf WSDL to code
packages but none worked so I rolled my own that is just about enough for the task.

TODO:

* pull out acceptable values for parameters into enum type
* get rid of Call... prefix on API methods by changing all the struct names to not clash? Or move package?
* change package name to betdaq
* simplify method calls for no/few parameters? (eg GetOddsLadder())
* why is GetPricesRequest.ThresholdAmount is string?
* go test failing due to redefined main func
* fix the many, many lint issues
* hard-code simpleType PolarityEnum
