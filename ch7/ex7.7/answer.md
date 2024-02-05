# Exercise 7.7: 

Explain why the help message contains °C when the default value of 20.0 does not.

## Explanation:

We use a `flag.CommandLine.Var` to call `FlagSet.Var`, which calls the `String` method on our `Celsius` type, returning `20°C`.