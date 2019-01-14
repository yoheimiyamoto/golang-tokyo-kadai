package simpletree

import "testing"

func TestPrint(t *testing.T) {
	tre := &tree{
		FirstChild: &tree{
			Name: "one",
			FirstChild: &tree{
				Name: "two",
			},
			// NextSibling: &tree{
			// 	Name: "two",
			// },
		},
	}
	print(tre)
}
