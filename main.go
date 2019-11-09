package main

import (
	"fmt"
	//"net"
	//"net/http"
	//"net/rpc"
)


type Item struct{
	title string
	body string
}

var database []Item


func GetByName(title string) Item{
	var getitem Item

	for _, val:= range database{
		if val.title == title{
			getitem=val
		}
	}

	return getitem

}

func AddItem(item Item) Item{
	database = append(database, item)
	return item
}

func EditItem(title string, edit Item)Item{
	for ind, val :=range database{
		if val.title ==edit.title{
			database[ind] = edit
		}

	}
	return edit
}

func DeleteItem(i Item) Item {

	for ind, val := range database{
		if val.title == i.title && val.body == i.body{
			database = append(database[:ind], database[ind+1:]...)

		}
	}
	return i

}


func main(){

   fmt.Println("Initial Database : ", database)
   a := Item{"First", "ITEM 1"}
   b := Item{"Second", "ITEM 2"}

   AddItem(a)
   AddItem(b)

   fmt.Println("Database after appending elements :", database)

   c:= Item{"Third", "ITEM 3"}
   AddItem(c)

   EditItem("Third", Item{"THIRDSSS", "JFUYG"})

   fmt.Println("Database after editing : ", database)
   fmt.Println(GetByName("Third"))
   

}