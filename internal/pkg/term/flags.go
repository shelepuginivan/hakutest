package term

import "os"

// ResolveOutput is a helper function that accepts return of
// [*pflag.FlagSet.GetString](). If output is not specified (i.e. `err != nil`)
// or output is equal to dash (`-`), [os.Stdout] is returned. Otherwise
// ResolveOutput returns result of [os.Create].
func ResolveOutput(output string, err error) (*os.File, error) {
	if err != nil {
		return os.Stdout, nil
	}

	if output == "-" {
		return os.Stdout, nil
	}

	return os.Create(output)
}
