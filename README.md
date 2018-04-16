# betdaq

Betdaq API client to make SOAP requests to Betdaq betting API and parse responses.

See [http://betdaqtraders.com/](http://betdaqtraders.com/) for details of the API.

Includes an example API request (see main.go)

You will need to create a file `config.json` (see `config_example.json` for the format) with your account details.
Note that currently you have to have an account and register to use the API.

Using code generation to generate the structs needed from the WSDL file. I tried a couple of off-the-shelf WSDL to code
packages but none worked so I rolled my own that is just about enough for the task.

Requests implemented:
* GetOddsLadder
* GetAccountBalances

TODO:
auto-generate the functions for API calls rather than coding by hand.
