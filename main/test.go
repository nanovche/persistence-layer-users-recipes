package main

import (
	"fmt"
	"go-training/persistence-layer-for-users-and-recipies/entities"
	"go-training/persistence-layer-for-users-and-recipies/printutils"
	recipe_repository "go-training/persistence-layer-for-users-and-recipies/recipe-repository"
	"go-training/persistence-layer-for-users-and-recipies/testdata/recipesrecords"
	"go-training/persistence-layer-for-users-and-recipies/testdata/usersrecords"
	user_repository "go-training/persistence-layer-for-users-and-recipies/user-repository"
)

func TestUsers(userIds *[]uint) {

	users := usersrecords.GetUsers()
	ur := user_repository.UserRepository{}

	for _,u := range users {
		if err := ur.Create(&u); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Created user: %s\n", u.UserName)
		}
	}

	//FindAll
	if users, err := ur.FindAll(); err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("All users:")
		printutils.PrintUsers(users)
	}

	fmt.Printf("\n\n")
	//FindAllPagedAndSorted
	//-success
	pagedRes, err := ur.FindAllPagedAndSorted(3, 3, "userName", false)
	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println("Descending ordering by \"userName\", page 3, pagesize: 3")
		printutils.PrintUsers(pagedRes)
	}
	fmt.Printf("\n\n")
	pagedRes, err = ur.FindAllPagedAndSorted(4, 2, "created", true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ascending ordering by \"created\", page 4, pagesize: 2")
		printutils.PrintUsers(pagedRes)
	}

	fmt.Printf("\n\n")

	var user entities.User
	//FindById
	//-success
	if user, err = ur.FindByID(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User 2:")
		printutils.PrintUsers([]entities.User{user})
	}
	fmt.Printf("\n\n")
	//-error
	if user, err = ur.FindByID(34); err != nil {
		fmt.Println(err)
	} else {
		printutils.PrintUsers([]entities.User{user})
	}

	fmt.Printf("\n\n")

	//FindByUsername
	//-success
	if user, err = ur.FindByUsername("zonecooking"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User zonecooking:")
		printutils.PrintUsers([]entities.User{user})
	}
	//-error
	fmt.Printf("\n\n")
	if user, err = ur.FindByUsername("nonexistent"); err != nil {
		fmt.Println(err)
	} else {
		printutils.PrintUsers([]entities.User{user})
	}

	fmt.Printf("\n\n")

	//DeleteById
	//-success
	if user, err = ur.DeleteByID(4); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted user")
		printutils.PrintUsers([]entities.User{user})
	}
	fmt.Printf("\n\n")
	//-error
	if user, err = ur.DeleteByID(13); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted user")
		printutils.PrintUsers([]entities.User{user})
	}

	//add it back
	user = 	entities.User{
			Id: 		4,
			UserName:   "twinevidence",
			LoginName:  "theChef",
			Password:   "3zs`/Pq`P?",
			Gender:     "female",
			UserRole:   "user",
			PictureURL: "nema",
			ShortDescr: "hippie with east-cuisine interests",
			Active:     true,
	}
	if err := ur.Create(&user); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Created user: %s\n", user.UserName)
	}

	fmt.Printf("\n\n")

	//Update
	//-success
	user = entities.User{
			UserName:   "idiotworse",
			LoginName:  "terminator",
			Password:   ".[deD5_DJX",
			Gender:     "male",
			UserRole:   "admin",
			PictureURL: "nema",
			ShortDescr: "cooking guru",
			Active:     false,
	}
	if err = ur.Update(&user); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User updated successfully!")
	}
	//err
	user = entities.User{
		Id: 8,
		UserName:   "nosuchuser",
		LoginName:  "Ivelin",
		Password:   "99u>11Uvve]n",
		Gender:     "male",
		UserRole:   "user",
		PictureURL: "",
		ShortDescr: "the crazy cook",
		Active:     false,
	}
	if err = ur.Update(&user); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User updated successfully!")
	}

	fmt.Printf("\n\n")

	//Count
	fmt.Println("Total records: ", ur.Count())

	if users, err := ur.FindAll(); err != nil{
		fmt.Println(err)
	} else {
		for _, u := range users {
			*userIds = append(*userIds, u.Id)
		}
	}

}
func TestRecipes(userIds []uint) {
	rr := recipe_repository.RecipeRepository{}
	recipes := recipesrecords.GetRecipes()

	for i := 0; i < len(recipes); i++{
		recipes[i].AuthorId = userIds[i]
	}

	if err := rr.CreateBatch(&recipes); err != nil {
		fmt.Println(err)
	} else {
		printutils.PrintRecipes(recipes)
	}


	fmt.Printf("\n\nFind all\n")
	//	FindAll
	if recipes, err := rr.FindAll(); err != nil {
		fmt.Println(err)
	} else {
		printutils.PrintRecipes(recipes)
	}

	fmt.Printf("\n\nFind Paged and Sorted\n")
	//FindAllPagedAndSorted
	if recipes, err := rr.FindAllPagedAndSorted(4, 2 , "time_to_cook", true); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sorted by \"cooking time\": ascending")
		printutils.PrintRecipes(recipes)
	}

	if recipes, err := rr.FindAllPagedAndSorted(2, 4 , "recipe_title", false); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sorted by \"recipe_title\": descending")
		printutils.PrintRecipes(recipes)
	}

	//	FindById
	//	-success
	fmt.Printf("\n\nFind By Id\n")
	if recipe, err := rr.FindByID(6); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("-success")
		fmt.Printf("Recipe with id %d found\n\n", recipe.Id)
		printutils.PrintRecipes([]entities.Recipe{recipe})
	}
	//-error
	if recipe, err := rr.FindByID(1000); err != nil {
		fmt.Printf("\n-error\n")
		fmt.Println(err)
	} else {
		fmt.Printf("Recipe with id %d found\n", recipe.Id)
		printutils.PrintRecipes([]entities.Recipe{recipe})
	}


	fmt.Printf("\n\nFind By Title")
	//	FindByTitle
	if recipes, err := rr.FindAllByTitle("Mashed"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("-success")
		printutils.PrintRecipes(recipes)
	}

	fmt.Printf("\n\nFind By Products")
	//	FindByProducts
	if recipes, err := rr.FindAllByProducts([]string{"cheese", "onion"}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("-success")
		printutils.PrintRecipes(recipes)
	}

	fmt.Printf("\n\nFind By Tags\n")
	//	FindByTags
	if recipes, err := rr.FindAllByTags([]string{"cheese", "onion"}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("-success")
		printutils.PrintRecipes(recipes)
	}

	fmt.Printf("\n\nDelete by id")
	//	DeleteById
	var title string //for create
	if recipe, err := rr.DeleteByID(2); err != nil {
		fmt.Printf("\n-error\n")
		fmt.Println(err)
	} else {
		fmt.Printf("-success\n\n")
		fmt.Printf("Recipe with id %d deleted\n\n", recipe.Id)
		printutils.PrintRecipes([]entities.Recipe{recipe})
		title = recipe.RecipeTitle
	}
	if recipe, err := rr.DeleteByID(1000); err != nil {
		fmt.Printf("\n-error\n")
		fmt.Println(err)
	} else {
		fmt.Printf("Recipe with id %d deleted\n\n", recipe.Id)
		printutils.PrintRecipes([]entities.Recipe{recipe})
	}


	var recipeToCreate entities.Recipe
	recipes = recipesrecords.GetRecipes()
	for _, r := range recipes {
		if r.RecipeTitle == title {
			recipeToCreate = r
		}
	}

	fmt.Printf("\n\nCREATE")
	//	Create
	if err := rr.Create(&recipeToCreate); err != nil {
		fmt.Printf("\n-error\n")
		fmt.Println(err)
	} else {
		fmt.Printf("-success\n\n")
		fmt.Printf("Recipe with id %d created\n\n", recipeToCreate.Id)
	}

	fmt.Printf("\n\nUPDATE")
	//Update
	recipeToCreate.TimeToCook = 99
	if err := rr.Update(&recipeToCreate); err != nil {
		fmt.Printf("\n-error\n")
		fmt.Println(err)
	} else {
		fmt.Printf("-success\n\n")
		fmt.Printf("recipe \"%s\" successfully updated", recipeToCreate.RecipeTitle)
	}

	fmt.Printf("\n\nCOUNT\n")
	//Count
	fmt.Printf("Total number of recipes: %d", rr.Count())
}