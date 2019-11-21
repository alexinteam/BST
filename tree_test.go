package main

func abcTree() *Tree {
	r := Node{
		Value: "1",
		Data:  "1",
		Left: &Node{
			Value: "2",
			Data:  "2",
			Left: &Node{
				Value: "3",
				Data:  "3",
				Right: &Node{
					Value: "4",
					Data:  "4",
				},
			},
			Right: &Node{
				Value: "5",
				Data:  "5",
				Left: &Node{
					Value: "6",
					Data:  "6",
					Right: &Node{
						Value: "7",
						Data:  "7",
					},
				},
			},
		},
		Right: &Node{
			Value: "8",
			Data:  "8",
		},
	}

	return &Tree{Root: &r}
}
