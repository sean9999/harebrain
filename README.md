# Harebrain

Harebrain is a simple file-backed db with the following characteristics:

- every record is a file
- every record must know how to marshal and unmarshal itself
- every record has a `Hash()` function that produces a unique hash, which becomes the filename.

## Getting started

```go

import (
    "github.com/sean9999/harebrain"
)

type cat struct {
	Id    int
	Name  string
	Breed string
}

type catRecord = harebrain.JsonRecord[cat]

func main(){

    myCat := &cat{1,"Muffin","Calico"}

    db := harebrain.NewDatabase()
    db.Open("data")
    db.Table("cats").Insert(myCat)

}



```