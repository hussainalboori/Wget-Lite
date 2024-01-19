package funcs

import "github.com/spf13/pflag"

var (
	SaveAs     = pflag.StringP("O", "O", "", "save the downloaded file with a different name")
	SaveDir    = pflag.StringP("P", "P", "", "directory to save the downloaded file")
	RateLimit  = pflag.Int64P("rate-limit", "R", 0, "limit the download speed in bytes per second")
	InputFile  = pflag.StringP("i", "i", "", "specify a file with paths, and the program will download them async style :)")
	SilentMode = pflag.BoolP("B", "B", false, "Enables Silent mode")
	// Mirror     = pflag.BoolP("mirror", "mirror", false, "Mirror a website's frontend by parsing html")
)
