package safetest

import (
	"fmt"

	"github.com/flan0910/SafeURLBot/modules"
	
	"github.com/google/safebrowsing"
)

func Safebrow(data []string) bool{
	key := modules.ConfigLoader().ApiKey
	sb, err := safebrowsing.NewSafeBrowser(safebrowsing.Config{
		APIKey: key,
	})
	if err != nil {
		fmt.Println(err)
	}

	ret, err := sb.LookupURLs(data)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range ret {
		if len(v) != 0 {
			return true
		}
	}
	return false
}