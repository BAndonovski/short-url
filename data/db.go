package data

import (
	"fmt"
	"time"

	"github.com/BAndonovski/short-url/data/model"
)

type SLDb map[string]*model.ShortLink

var db SLDb = SLDb{}

func Set(original, short string) error {
	if _, ok := db[short]; ok {
		return fmt.Errorf("%s already exists", short)
	}

	db[short] = &model.ShortLink{
		Short:     short,
		Original:  original,
		CreatedOn: time.Now(),
	}
	return nil
}

func Get(short string) (*model.ShortLink, error) {
	val, ok := db[short]
	if ok {
		val.LastVisit = time.Now()
		val.Visits += 1
		return val, nil
	}
	return nil, fmt.Errorf("%s not found", short)
}
