package special

import (
	"strings"

	"github.com/jabbalaci/alap/lib"
)

func DubReplaceFileNames(source string, extra string) string {
	if len(extra) == 0 {
		return source
	}
	// else
	filename_only := lib.TrimExtension(extra)
	source = strings.ReplaceAll(source, "\"main\"", "\""+filename_only+"\"")
	source = strings.ReplaceAll(source, "\"main.d\"", "\""+filename_only+".d\"")
	return source
}

func JavaReplaceFileNames(source string, extra string) string {
	if len(extra) == 0 {
		return source
	}
	// else
	filename_only := lib.TrimExtension(extra)
	source = strings.ReplaceAll(source, "class Main", "class "+filename_only)
	return source
}
