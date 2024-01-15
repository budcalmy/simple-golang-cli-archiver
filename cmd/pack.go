package cmd

import (
	"archiver/lib/compression/types/table/haffman"
	"archiver/lib/compression/types/table/shannon_fano"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"archiver/lib/compression"
	"archiver/lib/compression/types"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   pack,
}

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(cmd *cobra.Command, args []string) {

	var encoder compression.Encoder

	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	filePath := args[0]

	archiveFilename := filePath + "." + packedExtension

	method := cmd.Flag("method").Value.String()

	switch method {
	case "sh_fan":
		encoder = vlc.New(shannonfano.NewGenerator())
	case "huffman":
		encoder = vlc.New(haffman.NewGenerator())
	default:
		cmd.PrintErr("unknown compression method\t")
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

	packed := encoder.Encode(string(data))

	err = os.WriteFile(archiveFilename, packed, 0644)
	if err != nil {
		handleErr(err)
	}

	fmt.Printf("Упаковка файла %s в архив %s\n", filePath, archiveFilename)
}

func init() {
	rootCmd.AddCommand(packCmd)

	packCmd.Flags().StringP("method", "m", "", "compression method: sh_fan, huffman")

	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
