package cmd

import (
	"os"
	"io"
	"strings"
	"path/filepath"
	"errors"
	"github.com/lubouski/archiver/lib/vlc"

	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{
	Use: "vlc",
	Short: "Pack file using variable-length code",
	Run: pack,
}

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(_ *cobra.Command, args []string) {
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
	packed := vlc.Encode(strings.TrimSuffix(string(data),"\n"))

	// save result to file
	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
        if err != nil {
                handleError(err)
        }
}

func packedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension	
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
