package merge

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExample(t *testing.T) {
	type User struct {
		Name     string
		LastName string
	}
	//Declare type that contains map to merge
	type Data struct {
		Users     Map[string, User] // same as map[string]User
		Blacklist Slice[string]     // same as []string
	}

	actual := Data{
		Blacklist: []string{"sammy1789"},
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
	addBlacklist := []string{"deleted_user_111"}
	//merge
	actual.Users.Merge(addUsers)
	actual.Blacklist.Merge(addBlacklist)

	fmt.Println("merged result", actual)

	//validate
	expected := Data{
		Blacklist: []string{"sammy1789", "deleted_user_111"},
		Users: map[string]User{
			"johnny76": {
				Name:     "Jonh",
				LastName: "Smith",
			},
			"sammy1789": {
				Name:     "Samuel",
				LastName: "Smith",
			},
			"willy16": {
				Name:     "William",
				LastName: "Darrel",
			},
			"jenny12": {
				Name:     "Jennifer",
				LastName: "Gunning",
			},
		},
	}
	deepEqual := reflect.DeepEqual(
		expected,
		actual,
	)
	if !deepEqual {
		t.Fatal("Expected:", expected, "Actual:", actual)
	}
}
