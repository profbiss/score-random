package score_random

import (
	"fmt"
)

func main()  {
	suggest := NewSuggest()
	suggest.Items = []Item{
		NewItem(1,10),
		NewItem(2,30),
		NewItem(3,60),
	}

	fmt.Println(suggest.GetItem())
}