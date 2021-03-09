package recipesrecords

import (
	"go-training/persistence-layer-for-users-and-recipies/entities"
)

func GetRecipes() []entities.Recipe {

	return []entities.Recipe{
		{
			RecipeTitle:      "Taco Meat",
			ShortDescr:       "Make a big batch and use for a multiple of different meals: add to bean burritos for a more" +
				" hearty and tasty burrito, or add a can of beans (kidney, red, or pinto) and use for Frito® Pie," +
				" Navajo tacos, or taco salad.",
			TimeToCook:       15,
			Products:         []string{"ground beef", "onion", "salt", "ground cumin", "tomatoes"},
			RecipePictureURL: "https://imagesvc.meredithcorp.io/v3/mm/image?url=https%3A%2F%2Fimages.media-allrecipes.com%2Fuserphotos%2F5281437.jpg&w=596&h=596&c=sc&poi=face&q=85",
			DetailedDescr:    "Step 1\nHeat a large skillet over medium-high heat. Cook and stir beef in the hot skillet until" +
				" browned and crumbly, 5 to 7 minutes.\n\nStep 2\nSeason beef with onion powder, garlic salt, celery salt," +
				" and cumin. Pour tomato sauce over the beef, stir to coat, and simmer until thickened, slightly, about 5 minutes.",
			Keywords:         []string{"beef", "meat", "taco"},
		},
		{
			RecipeTitle:      "Old School Mac n' Cheese",
			ShortDescr:       "This is a completely unpretentious, down-home macaroni and cheese recipe just like my grandma and mom always made.",
			TimeToCook:       75,
			RecipePictureURL: "https://imagesvc.meredithcorp.io/v3/mm/image?url=https%3A%2F%2Fimages.media-allrecipes.com%2Fuserphotos%2F1059713.jpg&w=596&h=399&c=sc&poi=face&q=85",
			Products:         []string{"macaroni", "butter", "milk", "mustard", "onion", "pepper", "cheese", "salt"},
			DetailedDescr:    "Step 1\nPreheat oven to 375 degrees F (190 degrees C).\n\nStep 2\nBring a large pot of lightly" +
				" salted water to a boil. Cook elbow macaroni in the boiling water, stirring occasionally until cooked through" +
				" but firm to the bite, 8 minutes; drain.\n\nStep 3\nMelt butter in a large pot over medium-low heat. Slowly add flour to" +
				" butter, whisking constantly; cook until brown and the mixture no longer smells of flour, about 5 minutes. Pour 1 cup milk" +
				" into the flour mixture, whisking continually until fully incorporated, about 45 seconds; repeat twice. Add remaining 3" +
				" cups milk to the mixture, whisking to incorporate. Stir Worcestershire sauce, mustard powder, onion powder, and cayenne" +
				" pepper into the mixture; season with salt and black pepper.\n\nStep 4\nReduce heat to low. Cook sauce, whisking frequently, " +
				"until it begins to thicken, about 10 minutes. Add about half the package of shredded Cheddar cheese; stir continually until " +
				"the cheese melts completely. Repeat with remaining half package of Cheddar cheese and the American cheese, about 4 ounces" +
				" at a time. Once cheese is entirely incorporated, remove sauce from heat.\n\nStep 5\nStir drained macaroni into the" +
				" cheese sauce to coat. Divide macaroni between two 9x13-inch baking dishes.\n\nStep 6\nMix crushed potato chips, 1 " +
				"cup shredded Cheddar cheese, and Parmesan cheese in a bowl. Top the macaroni with the potato chip mixture evenly. " +
				"Spray the potato chip mixture with cooking spray.\n\nStep 7\nBake in preheated oven until the crust is golden brown and" +
				" the sauce is bubbling, 35 to 45 minutes.",
			Keywords:         []string{"cheese", "mac", "old school"},
		},
		{
			RecipeTitle:      "Basic Mashed Potatoes",
			ShortDescr:       "If you love good, old fashioned mashed potatoes this is the perfect recipe. Simple and delicious.\n\n",
			TimeToCook:       35,
			Products:         []string{"potatoes", "butter", "milk", "salt", "pepper"},
			RecipePictureURL: "https://imagesvc.meredithcorp.io/v3/mm/image?url=https%3A%2F%2Fimages.media-allrecipes.com%2Fuserphotos%2F5352583.jpg&w=596&h=596&c=sc&poi=face&q=85",
			DetailedDescr:    "Step 1\nBring a pot of salted water to a boil. Add potatoes and cook until tender but still" +
				" firm, about 15 minutes; drain.\n\nStep 2\nIn a small saucepan heat butter and milk over low heat until " +
				"butter is melted. Using a potato masher or electric beater, slowly blend milk mixture into potatoes " +
				"until smooth and creamy. Season with salt and pepper to taste.",
			Keywords:         []string{"potatoes", "milk", "mashed"},
		},
		{
			RecipeTitle:      "Baked Chicken Schnitzel",
			ShortDescr:       "Growing up, chicken schnitzel was a classic. I decided to make this dish oven-friendly using" +
				" less oil, and an easier cleanup. This dish tastes great with potato salad, or mashed potatoes and a nice" +
				" crisp salad.",
			TimeToCook:       30,
			Products:         []string{"chicken", "pepper", "paprika", "eggs", "lemon", "bread"},
			RecipePictureURL: "https://imagesvc.meredithcorp.io/v3/mm/image?url=https%3A%2F%2Fimages.media-allrecipes.com%2Fuserphotos%2F6316160.jpg&w=596&h=596&c=sc&poi=face&q=85",
			DetailedDescr:    "Step 1\nPreheat oven to 425 degrees F (220 degrees C). Line a large baking sheet with" +
				" aluminum foil and drizzle olive oil over foil. Place baking sheet in preheated oven.\n\nStep 2\nFlatten" +
				" chicken breasts so they are all about 1/4-inch thick. Season chicken with salt and pepper.\n\nStep 3\nMix" +
				" flour and paprika together on a large plate. Beat eggs with salt and pepper in a shallow bowl. Mix bread " +
				"crumbs and lemon zest together on a separate large plate. Dredge each chicken piece in flour mixture, " +
				"then egg, and then bread crumbs mixture and set aside in 1 layer on a clean plate. Repeat with remaining " +
				"chicken.\n\nStep 4\nRemove baking sheet from oven and arrange chicken in 1 layer on the sheet. Drizzle more " +
				"olive oil over each piece of coated chicken.\n\nStep 5\nBake in the preheated oven for 5 to 6 minutes. Flip" +
				" chicken and continue baking until no longer pink in the center and the breading is lightly browned, " +
				"5 to 6 minutes more. An instant-read thermometer inserted into the center should read at least 165 degrees " +
				"F (74 degrees C).",
			Keywords:         []string{"chicken breasts", "degrees", "mixture", "bread crumbs", "schnitzel"},
		},
		{
			RecipeTitle:      "Simple White Cake",
			ShortDescr:       "This cake was sent home from our children's school. It is the simplest, great tasting cake I've ever made. Great to make with the kids, especially for cupcakes.",
			TimeToCook:       50,
			Products:         []string{"sugar", "butter", "eggs", "vanilla", "baking powder", "milk"},
			RecipePictureURL: "https://imagesvc.meredithcorp.io/v3/mm/image?url=https%3A%2F%2Fimages.media-allrecipes.com%2Fuserphotos%2F6316160.jpg&w=596&h=596&c=sc&poi=face&q=85",
			DetailedDescr:    "Step 1 Preheat oven to 350 degrees F (175 degrees C). Grease and flour a 9x9 inch pan or line a muffin pan with paper liners. " +
				"Step 2 In a medium bowl, cream together the sugar and butter. Beat in the eggs, one at a time, then stir" +
				" in the vanilla. Combine flour and baking powder, add to the creamed mixture and mix well. Finally stir " +
				"in the milk until batter is smooth. Pour or spoon batter into the prepared pan. Step 3 Bake for 30 to 40 " +
				"minutes in the preheated oven. For cupcakes, bake 20 to 25 minutes. Cake is done when it springs back to" +
				" the touch.",
			Keywords:         []string{"sugar", "cake", "cupcakes", "bake"},
		},
		{
			RecipeTitle:      "Quick Crispy Parmesan Chicken Breasts",
			ShortDescr:       "These are delicious, easy, quick, and so versatile! Eat them plain, topped with your favorite spaghetti sauce, or sliced on a Caesar salad",
			TimeToCook:       35,
			Products:         []string{"chicken", "cheese", "butter", "wine", "bread", "salt", "onion"},
			RecipePictureURL: "https://imagesvc.meredithcorp.io/v3/mm/image?url=https%3A%2F%2Fimages.media-allrecipes.com%2Fuserphotos%2F4947131.jpg&w=596&h=596&c=sc&poi=face&q=85",
			DetailedDescr:    "Step 1 Preheat oven to 400 degrees F (200 degrees C). Line a baking sheet with aluminum foil and spray with" +
				" cooking spray. Step 2 Whisk bread crumbs, Parmesan cheese, paprika, salt, and black pepper together in a shallow bowl. Stir butter, white wine, mustard, and garlic" +
				" together in another bowl. Step 3 Dip each chicken breast half into melted butter mixture; press into bread" +
				" crumb mixture to evenly coat. Place breaded chicken in a single layer on the prepared baking sheet." +
				" Pat any leftover bread crumb mixture onto chicken breasts. Step 4 Bake chicken in the preheated oven" +
				" until no longer pink in the center and the juices run clear, about 20 minutes. An instant-read thermometer" +
				" inserted into the center should read at least 165 degrees F (74 degrees C).",
			Keywords:         []string{"crispy", "quick", "breasts", "chicken", "wine"},
		},
		{
			RecipeTitle:      "Stove Pot Roast With Mashed Potatoes",
			ShortDescr:       "Wonderful flavors make the meat the star of the dish by combining simple ingredients for a mouth-watering meal. Our southern family" +
				" has passed this recipe down for many years and enjoyed many memories dining on this delicious home-cooked classic.",
			TimeToCook:       390,
			Products:         []string{"beef", "milk", "onion", "salt", "carrots", "garlic", "pepper", "potatoes"},
			RecipePictureURL: "https://imagesvc.meredithcorp.io/v3/mm/image?url=https%3A%2F%2Fimages.media-allrecipes.com%2Fuserphotos%2F949315.jpg&w=596&h=596&c=sc&poi=face&q=85",
			DetailedDescr:    "Season chuck roast with salt and black pepper; sear in a large, deep skillet" +
				" or Dutch oven over medium heat until browned, about 10 minutes per side. Step 2 Pour beef broth and" +
				" water into the skillet with roast. Arrange onion wedges and garlic cloves around the meat. Spread carrots" +
				" atop roast and place sprig of rosemary atop carrots. Turn heat to medium-low and simmer until tender," +
				" about 6 hours. Step 3 Cover potatoes with water in a large pot and bring to a boil; reduce heat to low" +
				" and simmer until tender, about 30 minutes. Drain. Mash potatoes with butter and half the evaporated milk" +
				" until smooth; slowly mash remaining evaporated milk into potatoes to achieve the desired consistency." +
				" Season with salt. Step 4 Remove 1 or 2 cloves of garlic from skillet and mash cloves on top of the" +
				" roast; serve with mashed potatoes.",
			Keywords:         []string{"roast", "mashed", "potatoes", "pot"},
		},
		{
			RecipeTitle:      "Tatertot Casserole",
			ShortDescr:       "Quick and easy casserole everyone will love.",
			TimeToCook:       40,
			Products:         []string{"tatertots", "mushrooms", "onion", "salt", "cheese", "pepper"},
			RecipePictureURL: "https://imagesvc.meredithcorp.io/v3/mm/image?url=https%3A%2F%2Fimages.media-allrecipes.com%2Fuserphotos%2F913092.jpg&w=595&h=595&c=sc&poi=face&q=85",
			DetailedDescr:    "Instructions Checklist Step 1 Preheat oven to 350 degrees F (175 degrees C). Step 2 Cook and" +
				" stir ground beef in a large skillet over medium heat until no longer pink and completely browned, 7 to 10" +
				" minutes; season with salt and black pepper. Stir cream of mushroom soup into the cooked ground beef;" +
				" pour the mixture into a 9x13-inch baking dish. Layer tater tots evenly over the ground beef mixture;" +
				" top with Cheddar cheese. Step 3 Bake until tater tots are golden brown and hot, 30 to 45 minutes.",
			Keywords:         []string{"tatertos", "casserole"},
		},
		{
			RecipeTitle:      "Ultimate Tofu Breakfast Burrito Bowls",
			ShortDescr:       "Tofu scrambles up just like eggs, and with some clever spices," +
				 "even non-vegans will barely notice the difference. Try setting out toppings to let family or guests " +
				"assemble their own burrito bowls.",
			TimeToCook:       45,
			Products:         []string{"tomatoes", "onion", "salt", "avocado", "cilantro", "potatoes"},
			DetailedDescr: "Step 1 Preheat a large, heavy skillet over medium-high heat. Add 2 tablespoons oil. Break tofu apart over skillet into bite-size pieces, sprinkle with salt and pepper, then cook, stirring frequently with a thin metal spatula, until liquid cooks out and tofu browns, about 10 minutes. (If you notice liquid collecting in pan, increase heat to evaporate water.)" +
				" Be sure to get under the tofu when you stir, scraping the bottom of the pan where the good," +
				" crispy stuff is and keeping it from sticking.Step 2 Add onion and garlic powders, turmeric, juice, and remaining tablespoon" +
				" oil and toss to coat. Cook 5 minutes more. Step 3 Preheat a heavy-bottomed saucepan over medium-high heat. Add oil. Cook onion and jalapenos with a pinch of salt, stirring, until translucent, about 5 minutes, Add garlic and cook, stirring, until fragrant, about 30 seconds. Add tomatoes, cumin, and remaining salt, and cook, stirring," +
				" until tomatoes become saucy, about 5 minutes. Add cilantro and lemon juice. Let cilantro wilt in. Add beans and heat through," +
				" stirring occasionally, about 2 minutes. Taste for salt and seasoning. Step 4 Spoon some hash browns into each bowl, followed by a scoop" +
				" of beans and a scoop of scramble. Top with avocado, a squeeze of fresh lemon juice, and a sprinkle of cilantro. Serve with hot sauce.",
			Keywords:         []string{"tofu", "burrito", "avocado", "cilantro"},
		},
		{
			RecipeTitle:      "Rice Paper Fake Bacon",
			ShortDescr:       "Some fake bacons are hit-and-miss, but because bacon is the one thing most people miss when switching to a vegetarian or vegan way of life, it's expected a fitting alternative is created",
			TimeToCook:       21,
			Products:         []string{"maple", "miso paste", "paprika", "sesame oil", "pepper"},
			DetailedDescr:    "Step 1 Preheat the oven to 400 degrees F (200 degrees C).Step 2 Whisk sesame oil, soy sauce, maple syrup, miso, liquid smoke, paprika," +
			" and black pepper together in a shallow bowl. Step 3 Stack 2 pieces of rice paper and soak them in a bowl of cold water until" +
			" slightly soft and pliable, about 30 seconds.Cut the hydrated rice paper into strips using a pizza cutter. Dip each strip in the" +
			" soy mixture and place on a sheet of parchment paper. Carefully lay the parchment paper directly on the oven rack." +
			" Step 4 Bake in the preheated oven until strips are dry, 6 to 8 minutes.",
			Keywords:         []string{"fake", "bacon", "rice", "miso paste"},
		},

	}
}
