package cmd

import (
	"Archiver/lib/compression"
	"Archiver/lib/compression/vlc"
	"errors"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   pack,
}

const packedExtension = "vlc"

var EmptyPath = errors.New("path to file if not specified")

func pack(cmd *cobra.Command, args []string) {
	var encoder compression.Encoder

	if len(args) == 0 || args[0] == "" {
		handlerErr(EmptyPath)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		encoder = vlc.New()
	default:
		cmd.PrintErr("unknown method: " + method)
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handlerErr(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handlerErr(err)
	}

	packed := encoder.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), packed, 0644)
	if err != nil {
		handlerErr(err)
	}
}

func packedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	rootCmd.AddCommand(packCmd)

	packCmd.Flags().StringP("method", "m", "", "compression method: vlc")

	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
