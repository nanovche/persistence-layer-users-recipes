package entities

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id         uint      `header:"ID"`
	UserName   string    `header:"Username"`
	LoginName  string    `header:"LoginName"`
	Password   string    `header:"Password"`
	Gender     string    `header:"Gender"`
	UserRole   string    `header:"UserRole"`
	PictureURL string    `header:"PictureURL"`
	ShortDescr string    `header:"ShortDescr"`
	Active     bool      `header:"Active"`
	Created    time.Time `header:"Created"`
	Modified   time.Time `header:"Modified"`
}

type Recipe struct {
	Id          		uint      `header:"ID"`
	AuthorId    		uint      `header:"AuthorId"`
	RecipeTitle 		string    `header:"RecipeTitle"`
	ShortDescr   		string    `header:"ShortDescr"`
	TimeToCook     		int    	  `header:"TimeToCook"`
	Products   			[]string  `header:"Products"`
	RecipePictureURL	string    `header:"RecipePictureURL"`
	DetailedDescr 		string    `header:"DetailedDescr"`
	Keywords   			[]string  `header:"Keywords"`
	Created    			time.Time `header:"Created"`
	Modified   			time.Time `header:"Modified"`
}