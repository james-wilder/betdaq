# betdaq

Betdaq API client to make SOAP requests to Betdaq betting API and parse responses.

See [http://betdaqtraders.com/](http://betdaqtraders.com/) for details of the API.

Includes an example API request (see main.go)

You will need to create a file `config.json` (see `config_example.json` for the format) with your account details.
Note that currently you have to have an account and register to use the API.

Requests implemented:
* GetOddsLadder

TODO:
<<<<<<< Updated upstream
* All of the others :-)
=======
* any other documentation for the structs?
* get rid of Call... prefix on API methods by changing all the struct names to not clash? Or move package?
* pull out acceptable values for parameters into enum type
* change package name to betdaq
* simplify method calls for no/few parameters? (eg GetOddsLadder())
>>>>>>> Stashed changes
