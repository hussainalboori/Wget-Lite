package funcs

import (
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

var (
	SaveAs     = pflag.StringP("O", "O", "", "save the downloaded file with a different name")
	SaveDir    = pflag.StringP("P", "P", "", "directory to save the downloaded file")
	RateLimit  = pflag.StringP("rate-limit", "L", "0", "limit the download speed in bytes per second")
	InputFile  = pflag.StringP("i", "i", "", "specify a file with paths, and the program will download them async style :)")
	SilentMode = pflag.BoolP("B", "B", false, "Enables Silent mode")
	Mirror     = pflag.BoolP("mirror", "m", false, "Mirror a website's frontend by parsing html")
	Reject     = pflag.StringP("reject", "R", "", "Reject specific file types (comma-separated, e.g., .jpg,.png)")
	Exclude    = pflag.StringP("exclude", "X", "", "Exclude specific directories from mirroring (comma-separated)")
)

func parseRateLimit() (int64, error) {
	// Convert rate limit string to lowercase for case-insensitivity
	rateLimitStr := strings.ToLower(*RateLimit)

	if rateLimitStr == "" {
		return 0, nil
	}

	// Check if the rate limit string ends with 'k' or 'm'
	if strings.HasSuffix(rateLimitStr, "k") {
		// If 'k', convert to kilobytes
		value, err := strconv.ParseInt(strings.TrimSuffix(rateLimitStr, "k"), 10, 64)
		if err != nil {
			return 0, err
		}
		return value * 1024, nil
	} else if strings.HasSuffix(rateLimitStr, "m") {
		// If 'm', convert to megabytes
		value, err := strconv.ParseInt(strings.TrimSuffix(rateLimitStr, "m"), 10, 64)
		if err != nil {
			return 0, err
		}
		return value * 1024 * 1024, nil
	} else if strings.HasSuffix(rateLimitStr, "g") {
		// If 'g', convert to gigabytes
		value, err := strconv.ParseInt(strings.TrimSuffix(rateLimitStr, "g"), 10, 64)
		if err != nil {
			return 0, err
		}
		return value * 1024 * 1024 * 1024, nil
	}

	// If no 'k', 'm', or 'g', assume value is in bytes
	return strconv.ParseInt(rateLimitStr, 10, 64)
}
