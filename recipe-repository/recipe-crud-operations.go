package recipe_repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-training/persistence-layer-for-users-and-recipies/db"
	"go-training/persistence-layer-for-users-and-recipies/entities"
	"log"
	"strings"
	"time"
)

type RecipeRepository struct {}

func(rr RecipeRepository) FindAll() ([]entities.Recipe, error){

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rows, err := db.DB.QueryContext(ctx, "SELECT * FROM recipes")
	if err != nil {
		return nil, fmt.Errorf("database query failed\n")
	}
	defer rows.Close()

	var recipes []entities.Recipe
	for rows.Next() {

		var recipe entities.Recipe
		var products, keywords string

		err := rows.Scan(&recipe.Id, &recipe.AuthorId, &recipe.RecipeTitle, &recipe.ShortDescr, &recipe.TimeToCook,
			&products, &recipe.RecipePictureURL, &recipe.DetailedDescr, &keywords, &recipe.Created, &recipe.Modified)
		if err != nil {
			return nil, fmt.Errorf("error scanning database record", err)
		}

		recipe.Products = strings.Split(products, " ")
		recipe.Keywords = strings.Split(keywords, " ")

		recipes = append(recipes, recipe)
	}
	return recipes, nil
}
func(rr RecipeRepository) FindAllPagedAndSorted(pageNumber int, pageSize int, sortingAttribute string, ascending bool) ([]entities.Recipe, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var query string
	offset := (pageNumber - 1) * pageSize
	if ascending {
		query = fmt.Sprintf("SELECT * FROM recipes ORDER BY %s LIMIT %d OFFSET %d ", sortingAttribute, pageSize, offset)
	} else {
		query = fmt.Sprintf("SELECT * FROM recipes ORDER BY %s DESC LIMIT %d OFFSET %d ", sortingAttribute, pageSize, offset)
	}

	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("database query failed\n")
	}
	defer rows.Close()

	var recipes []entities.Recipe
	for rows.Next() {

		var recipe entities.Recipe
		var products, keywords string

		err := rows.Scan(&recipe.Id, &recipe.AuthorId, &recipe.RecipeTitle, &recipe.ShortDescr, &recipe.TimeToCook, &products,
			&recipe.RecipePictureURL, &recipe.DetailedDescr, &keywords, &recipe.Created, &recipe.Modified)
		if err != nil {
			return nil, fmt.Errorf("error scanning database record\n", err)
		}

		recipe.Products = strings.Split(products, " ")
		recipe.Keywords = strings.Split(keywords, " ")
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}
func(rr RecipeRepository) FindByID(recipeId uint) (recipe entities.Recipe, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var products, keywords string

	err = db.DB.QueryRowContext(ctx, "SELECT * FROM recipes WHERE id = ?", recipeId).
		Scan(&recipe.Id, &recipe.AuthorId, &recipe.RecipeTitle, &recipe.ShortDescr, &recipe.TimeToCook, &products,
		&recipe.RecipePictureURL, &recipe.DetailedDescr, &keywords, &recipe.Created, &recipe.Modified)
	switch {
	case err == sql.ErrNoRows:
		return recipe, fmt.Errorf("recipe with id %d does not exist\n", recipeId)
	case err != nil:
		return recipe, fmt.Errorf("error quering the database\n")
	}

	recipe.Products = strings.Split(products, " ")
	recipe.Keywords = strings.Split(keywords, " ")

	return
}
func(rr RecipeRepository) FindAllByTitle(partOfTitle string) ([]entities.Recipe, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	query := fmt.Sprintf( "SELECT * FROM recipes WHERE recipe_title LIKE '%%%s%%'", partOfTitle)
	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying database", err)
	}
	defer rows.Close()

	var recipes []entities.Recipe
	for rows.Next() {

		var recipe entities.Recipe
		var products, keywords string

		scanErr :=rows.Scan(&recipe.Id, &recipe.AuthorId, &recipe.RecipeTitle, &recipe.ShortDescr, &recipe.TimeToCook, &products,
			&recipe.RecipePictureURL, &recipe.DetailedDescr, &keywords, &recipe.Created, &recipe.Modified)
		if scanErr != nil {
			return nil, fmt.Errorf("error scanning record")
		}

		recipe.Products = strings.Split(products, " ")
		recipe.Keywords = strings.Split(keywords, " ")

		recipes = append(recipes, recipe)

	}

	return recipes, nil
}
func(rr RecipeRepository) FindAllByProducts(products []string) ([]entities.Recipe, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	query := "SELECT * FROM recipes WHERE "
	for i, prod := range products {
		query += fmt.Sprintf( " products LIKE '%%%s%%'", prod)
		if i < len(products) - 1 {
			query += fmt.Sprintf( " AND")
		}
	}
	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying database")
	}
	defer rows.Close()

	var recipes []entities.Recipe
	for rows.Next() {

		var recipe entities.Recipe
		var products, keywords string

		scanErr :=rows.Scan(&recipe.Id, &recipe.AuthorId, &recipe.RecipeTitle, &recipe.ShortDescr, &recipe.TimeToCook, &products,
			&recipe.RecipePictureURL, &recipe.DetailedDescr, &keywords, &recipe.Created, &recipe.Modified)
		if scanErr != nil {
			return nil, fmt.Errorf("error scanning record")
		}

		recipe.Products = strings.Split(products, " ")
		recipe.Keywords = strings.Split(keywords, " ")

		recipes = append(recipes, recipe)

	}

	return recipes, nil

}
func(rr RecipeRepository) FindAllByTags(tags []string) ([]entities.Recipe, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	query := "SELECT * FROM recipes WHERE "
	for i, tag := range tags {
		query += fmt.Sprintf( " keywords LIKE '%%%s%%'", tag)
		if i < len(tags) - 1 {
			query += fmt.Sprintf( " OR")
		}
	}
	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying database", err)
	}
	defer rows.Close()

	var recipes []entities.Recipe
	for rows.Next() {

		var recipe entities.Recipe
		var products, keywords string

		scanErr :=rows.Scan(&recipe.Id, &recipe.AuthorId, &recipe.RecipeTitle, &recipe.ShortDescr, &recipe.TimeToCook, &products,
			&recipe.RecipePictureURL, &recipe.DetailedDescr, &keywords, &recipe.Created, &recipe.Modified)
		if scanErr != nil {
			return nil, fmt.Errorf("error scanning record")
		}

		recipe.Products = strings.Split(products, " ")
		recipe.Keywords = strings.Split(keywords, " ")

		recipes = append(recipes, recipe)

	}

	return recipes, nil

}
func(rr RecipeRepository) Create(recipe *entities.Recipe) error {

	products := convertSliceToString(recipe.Products)
	keywords := convertSliceToString(recipe.Keywords)

	res, err := db.DB.Exec("INSERT INTO recipes(author_id, recipe_title, short_description, time_to_cook, products, recipe_pictureURL, detailed_description," +
		"keywords, created, modified) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		3, recipe.RecipeTitle, recipe.ShortDescr, recipe.TimeToCook,
		products, recipe.RecipePictureURL, recipe.DetailedDescr, keywords, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("database query failed\n", err)
	}

	numRows, err :=res.RowsAffected()
	if err != nil || numRows != 1 {
		return fmt.Errorf("error creating recipe %s\n", recipe.RecipeTitle)
	}
	insId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting id of created recipe: %s\n", recipe.RecipeTitle)
	}
	recipe.Id = uint(insId)

	return nil

}
func(rr RecipeRepository) CreateBatch(recipes *[]entities.Recipe) error {


	tr, err := db.DB.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction")
	}

	for i := 0; i < len(*recipes); i++{

		products := convertSliceToString((*recipes)[i].Products)
		keywords := convertSliceToString((*recipes)[i].Keywords)

		res, err := tr.Exec("INSERT INTO recipes(author_id, recipe_title, short_description, time_to_cook, products," +
			" recipe_pictureURL, detailed_description, keywords, created, modified) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", (*recipes)[i].AuthorId,
			(*recipes)[i].RecipeTitle, (*recipes)[i].ShortDescr, (*recipes)[i].TimeToCook, products, (*recipes)[i].RecipePictureURL,
			(*recipes)[i].DetailedDescr, keywords, time.Now(), time.Now())

		if err != nil {
			if rollBackErr := tr.Rollback(); rollBackErr != nil {
				return fmt.Errorf("creation failed, unable to rollback")
			}
			return fmt.Errorf("creation failed", err)
		}

		rows, err := res.RowsAffected()
		if err != nil || rows != 1{
			return err
		}
		fmt.Printf("Created recipe: %s\n", (*recipes)[i].RecipeTitle)

		var recipe_id int64
		if recipe_id, err = res.LastInsertId(); err != nil {
			return fmt.Errorf("error generating id for recipe: %s\n", (*recipes)[i].RecipeTitle)
		}
		(*recipes)[i].Id = uint(recipe_id)
	}
		// COMMIT TRANSACTION
	if err := tr.Commit(); err != nil {
			return fmt.Errorf("error commiting transaction")
		}

	return nil
}
func(rr RecipeRepository) Update(recipe *entities.Recipe) error  {

	rows := db.DB.QueryRow("SELECT id FROM recipes WHERE id = ?", recipe.Id)

	var dtbsId int
	err := rows.Scan(&dtbsId)
	switch {
	case err == sql.ErrNoRows:
		return fmt.Errorf("no recipe with id %d exists", recipe.Id)
	case err != nil:
		return fmt.Errorf("database query failed\n")
	}

	products := convertSliceToString(recipe.Products)
	keywords := convertSliceToString(recipe.Keywords)

	res, err := db.DB.Exec("UPDATE recipes SET recipe_title=?, short_description=?, time_to_cook=?, products=?," +
		" recipe_pictureURL=?,detailed_description=?, keywords=?, modified=? WHERE recipe_title=?",
		recipe.RecipeTitle, recipe.ShortDescr, recipe.TimeToCook, products, recipe.RecipePictureURL,
		recipe.DetailedDescr, keywords, time.Now(), recipe.RecipeTitle)
	if err != nil {
		return fmt.Errorf("database query failed\n")
	}
	numRows, err := res.RowsAffected()
	if err != nil || numRows != 1 {
		return fmt.Errorf("recipe update failed: no rows affected\n")
	}

	return nil

}
func(rr RecipeRepository) DeleteByID(recipeID uint) (recipe entities.Recipe, err error) {
	recipe, err = rr.FindByID(recipeID)
	if err != nil {
		return recipe, fmt.Errorf("error deleting recipe -> %s", err)
	}

	res, err := db.DB.Exec("DELETE FROM recipes WHERE id = ?", recipeID)
	if err != nil {
		return recipe, fmt.Errorf("database query failed\n")
	}
	numRows, err := res.RowsAffected()
	if err != nil || numRows != 1 {
		return recipe, fmt.Errorf("error deleting recipe: zero affected rows\n")
	}

	return
}
func(rr RecipeRepository) Count() int {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var count int
	rows, err := db.DB.QueryContext(ctx, "SELECT COUNT(*)  FROM recipes")
	if err != nil {
		log.Println(err)
	}

	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Println(err)
		}
	}

	return count
}
func convertSliceToString(items []string) (appended string) {

	for _, item := range items {
		appended += fmt.Sprintf(item + " ")
	}
	return
}
