package hwprocess

import (
	"fmt"
	"log"
)

func Process(p *string) string {
	log.Println(fmt.Sprint("Processing ", *p))
	return "Processed"
}
