package hwprocess

import (
	"fmt"
	"log"

	"github.com/otiai10/gosseract"
)

func Process(p *string) string {
	log.Println(fmt.Sprint("Processing ", *p))
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(*p)

	text, err := client.Text()
	log.Println(fmt.Sprint("Image text: ", text, " error: ", err))
	return "Processed"
}
