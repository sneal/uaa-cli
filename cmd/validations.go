package cmd

import (
	"code.cloudfoundry.org/uaa-cli/utils"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"errors"
	"fmt"
)

func avalableFormats() []string {
	return []string{"jwt", "opaque"}
}

func availableFormatsStr() string {
	return "[" + strings.Join(avalableFormats(), ", ") + "]"
}

func validateTokenFormat(cmd *cobra.Command, tokenFormat string) {
	if !utils.Contains(avalableFormats(), tokenFormat) {
		log.Errorf(`The token format "%v" is unknown. Available formats: %v`, tokenFormat, availableFormatsStr())
		cmd.Usage()
		os.Exit(1)
	}
}

func validateTokenFormatError(tokenFormat string) error {
	if !utils.Contains(avalableFormats(), tokenFormat) {
		return errors.New(fmt.Sprintf(`The token format "%v" is unknown. Available formats: %v`, tokenFormat, availableFormatsStr()))
	}
	return nil
}
