package cmd

import (
	"os"
	"io"
	"strings"
	"path/filepath"
	"github.com/lubouski/archiver/lib/vlc"

	"github.com/spf13/cobra"
)

var vlcUnpackCmd = &cobra.Command{
	Use: "vlc",
	Short: "Pack file using variable-length code",
	Run: unpack,
}

const unpackedExtension = "txt"

func unpack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleError(ErrEmptyPath)
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleError(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
        if err != nil {
                handleError(err)
        }

	// data -> Encode(data)
	// new line appears on created string
	packed := vlc.Decode(strings.TrimSuffix(string(data),"\n"))

	// save result to file
	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
        if err != nil {
                handleError(err)
        }
}

func unpackedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + unpackedExtension
}

func init() {
	unpackCmd.AddCommand(vlcUnpackCmd)
}
