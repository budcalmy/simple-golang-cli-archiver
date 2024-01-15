package cmd

import (
	"archiver/lib/compression/types/table/haffman"
	"archiver/lib/compression/types/table/shannon_fano"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"archiver/lib/compression"
	"archiver/lib/compression/types"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file",
	Run:   unpack,
}

func unpack(cmd *cobra.Command, args []string) {

	var decoder compression.Decoder

	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	filePath := args[0]

	originalFilename := strings.TrimSuffix(filePath, filepath.Ext(filePath))

	method := cmd.Flag("method").Value.String()

	switch method {
	case "sh_fan":
		decoder = vlc.New(shannonfano.NewGenerator())
	case "huffman":
		decoder = vlc.New(haffman.NewGenerator())
	default:
		cmd.PrintErr("unknown method compression\t")
	}

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed := decoder.Decode(data)

	err = os.WriteFile(originalFilename, []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}

}

func init() {
	rootCmd.AddCommand(unpackCmd)

	unpackCmd.Flags().StringP("method", "m", "", "decompression method: sh_fan, huffman")

	if err := unpackCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
