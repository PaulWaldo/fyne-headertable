package headertable

import (
	"fmt"
	"time"
)

func DisplayAsString(i interface{}) string {
	return fmt.Sprintf("%s", i)
}

func DisplayAsCurrency(i interface{}) string {
	return fmt.Sprintf("%0.2f", i)
}

func DisplayAsTime(formatString string, i interface{}) string {
	t := i.(time.Time)
	return t.Format(formatString)
}

func DisplayAsISODate(i interface{}) string {
	return DisplayAsTime("2006-01-02", i)
}
