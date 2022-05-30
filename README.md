Generic type wrappers for any map or slice


	type User struct {
		Name     string
		LastName string
	}
	//Declare type that contains map to merge
	type Data struct {
		Users Map[string, User]
	}
	data := Data{
		Users: map[string]User{
			"johnny76": {
				Name:     "Jonh",
				LastName: "Smith",
			},
			"sammy1789": {
				Name:     "Samuel",
				LastName: "Smith",
			},
		},
	}

	addUsers := map[string]User{
		"willy16": {
			Name:     "William",
			LastName: "Darrel",
		},
		"jenny12": {
			Name:     "Jennifer",
			LastName: "Gunning",
		},
	}

	//merge
	result := data.Users.Merge(addUsers)
	fmt.Println("merged result", result)
