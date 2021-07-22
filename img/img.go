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
	buider := strings.Builder{}
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
		ext := path.Ext(entry.Name())
		nameIn := strings.TrimSuffix(entry.Name(), ext)
		contentIn := string(b)
		if err := os.WriteFile(path.Join(out, entry.Name()), []byte(contentIn), 0644); err != nil {
			log.Fatal(err)
		}
		for _, p := range patterns {
			nameOut := strings.TrimSuffix(entry.Name(), ext) + "-" + p.Suffix + ext
			fileNameOut := path.Join(out, nameOut)
			fmt.Println(nameOut)
			contentOut := strings.ReplaceAll(contentIn, originalColor, p.Color)
			if err := os.WriteFile(fileNameOut, []byte(contentOut), 0644); err != nil {
				log.Fatal(err)
			}
		}

		selectors := `button.{{filename}}, a.button.{{filename}} {
    border-radius: 50%;
    padding: 0;
    background-image: url("../img/{{filename}}.svg");
    background-size: 1.5rem;
}

button.{{filename}}:hover, a.button.{{filename}}:hover {
    background-image: url("../img/{{filename}}-hover.svg");
}

button.{{filename}}:active, a.button.{{filename}}:active {
    border: 2px solid var(--cor-link-hover);
    box-shadow: none;
}

body.alto-contraste button.{{filename}}, body.alto-contraste a.button.{{filename}} {
    background-image: url("../img/{{filename}}-contraste.svg");
    background-color: transparent;
}

body.alto-contraste button.{{filename}}:hover, body.alto-contraste a.button.{{filename}}:hover,
body.alto-contraste button.{{filename}}:active, body.alto-contraste a.button.{{filename}}:active {
    border: 2px solid var(--cor-link-contraste);
}

`
		selectors = strings.ReplaceAll(selectors, "{{filename}}", nameIn)
		buider.WriteString(selectors)
	}
	cssFile := path.Join(out, "botoes.css")
	if err := os.WriteFile(cssFile, []byte(buider.String()), 0644); err != nil {
		log.Fatal(err)
	}
}
