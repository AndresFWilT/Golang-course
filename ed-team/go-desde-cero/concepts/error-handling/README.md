# Error Handling

Go works with errors and not with exceptions

It's important to control most of the errors that can occur during a call to a function, in that order of ideas you can use the `errors` library and the `fmt.Errorf()` to propagate the errors.To match errors you can use the `errors.Is(err1, err2)` and control them exceptions.
