/*
Package env provides a similar interface for environment variables as flag/pflag package.

The reason for this is to make configuration management in a CLI application as easy as possible.
Environment variables are mostly used in server applications (eg. apps running in docker) though,
so this solution is mostly beneficial for them.

Use this package along with the flag package from the stdlib or the POSIX compliant pflag
(github.com/spf13/pflag) library. The general recommendation is to let flags take precedence
over environment variables (when it makes sense to accept a configuration via both methods).
To achieve this, make sure to parse flags after environment variables.

Here is an example for a regular application using both flags and environment variables:

	package main

	import (
		flag "github.com/spf13/pflag"
		"github.com/goph/env"
	)

	func main() {
		var debug bool

		// Flags
		listen := flag.String("listen", ":8080", "Listen on this port")
		flag.BoolVar(&debug, "debug", false, "Whether to serve extra debug information in logs")

		// Env vars
		env.BoolVar(&debug, "debug", false, "Whether to serve extra debug information in logs")

		err := env.Parse()
		if err != nil {
			panic(err)
		}

		err = flag.Parse()
		if err != nil {
			panic(err)
		}

		// ... further code
	}
 */
package env
