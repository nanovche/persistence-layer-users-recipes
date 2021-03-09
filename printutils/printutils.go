package printutils

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"go-training/persistence-layer-for-users-and-recipies/entities"
	"os"
	"time"
)

func PrintUsers(entities []entities.User) {
	var tableRows []table.Row
	for _, u := range entities {
		row := table.Row{u.Id, u.UserName, u.LoginName, u.Password, u.Gender, u.UserRole, u.PictureURL, u.ShortDescr, u.Created.Format(time.ANSIC), u.Modified.Format(time.ANSIC)}
		tableRows = append(tableRows, row)
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "User name", "Login name", "Password", "Gender", "User role", "PictureURL", "Short description", "Created", "Modified"})
	t.AppendRows(tableRows)
	t.Render()
}

func PrintRecipes(entities []entities.Recipe) {
	var tableRows []table.Row
	for _, r := range entities {
		row := table.Row{r.Id, r.AuthorId, r.RecipeTitle, r.ShortDescr, r.TimeToCook, r.Products, r.RecipePictureURL, r.DetailedDescr, r.Keywords, r.Created.Format(time.ANSIC), r.Modified.Format(time.ANSIC)}
		tableRows = append(tableRows, row)
	}
	c := table.ColumnConfig{WidthMax: 70}
	t := table.NewWriter()
	t.SetColumnConfigs([]table.ColumnConfig{c})
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Author ID", "Title", "Short Description", "Cooking time(mins)", "Ingredients", "PictureURL", "Detailed Description", "Keywords", "Created", "Modified"})
	t.AppendRows(tableRows)
	t.Render()
}