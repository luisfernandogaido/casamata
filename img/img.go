package img

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

const (
	originalColor = "#0072ad"
)

type Pattern struct {
	Color  string
	Suffix string
}

var (
	patterns = []Pattern{
		{"#a65718", "hover"},
		{"#000000", "preto"},
		{"#ffffff", "branco"},
		{"#fff333", "contraste"},
	}
)

func Img() {
	var (
		fs  = flag.NewFlagSet("img", flag.ExitOnError)
		in  string
		out string
	)
	fs.StringVar(&in, "in", "./", "pasta de entrada")
	fs.StringVar(&out, "out", "./out", "pasta de saída")
	fs.Parse(os.Args[2:])
	if err := os.MkdirAll(out, 0644); err != nil {
		log.Fatal(err)
	}
	entries, err := os.ReadDir(in)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		fileNameIn := path.Join(in, entry.Name())
		if path.Ext(fileNameIn) != ".svg" {
			continue
		}
		b, err := os.ReadFile(fileNameIn)
		if err != nil {
			log.Fatal(err)
		}
		contentIn := string(b)
		for _, p := range patterns {
			ext := path.Ext(entry.Name())
			nameOut := strings.TrimSuffix(entry.Name(), ext) + "-" + p.Suffix + ext
			fileNameOut := path.Join(out, nameOut)
			fmt.Println(nameOut)
			contentOut := strings.ReplaceAll(contentIn, originalColor, p.Color)
			if err := os.WriteFile(fileNameOut, []byte(contentOut), 0644); err != nil {
				log.Fatal(err)
			}
		}
	}
}
