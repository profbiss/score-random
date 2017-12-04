# score-random
```Go
package main

import (
	"fmt"
  	"github.com/profbiss/score-random"
)

func main()  {
	suggest := score_random.NewSuggest()
	suggest.Items = []score_random.Item{
		score_random.NewItem(1,10),
		score_random.NewItem(2,30),
		score_random.NewItem(3,60),
	}

	fmt.Println(suggest.GetItem())
}
```

```
&{2 30 0.3 0}
```
