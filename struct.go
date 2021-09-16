package main

import "fmt"

type client struct {
	name     string
	nickname string
	smoke    bool
	actions  []string
}

type profile struct {
	client
	title string
}

type cat struct {
	name string
	age  int
}

func main() {

	client1 := client{
		name:     "João",
		nickname: "Pereira",
		smoke:    true,
		actions:  []string{"run", "walk"},
	}
	client2 := client{"Carlos", "Alberto", false, []string{"talk", "move"}}

	fmt.Println(client1)
	fmt.Println(client2)

	profile1 := profile{
		client: client{
			name:     "Miguel",
			nickname: "Jonas",
			smoke:    true,
		},
		title: "admin",
	}

	fmt.Println(profile1)

	fmt.Println(profile1.client.name)

	// Prop promoted because its the only one that have its name
	fmt.Println(profile1.name)

	// anonimous struct
	newStruct := struct {
		name string
	}{
		name: "joao",
	}

	fmt.Println(newStruct)

	for _, value := range client1.actions {

		fmt.Println(value)
	}

	clients := map[string]client{
		"Pereira": client1,
		"Carlos": {
			name:     "Miguel",
			nickname: "Jonas",
			smoke:    true,
		},
	}

	for _, value := range clients {

		fmt.Println(value)

		for _, value2 := range value.actions {

			fmt.Println(value2)
		}
	}

	var clients2 = make(map[string]client)

	clients2["Luana"] = client{
		nickname: "Jonas",
		smoke:    true,
	}

	fmt.Println(clients2)

	anonimousStruct := struct {
		actions map[int]string
		colors  []string
	}{
		actions: map[int]string{1: "run", 2: "walk"},
		colors:  []string{"blue", "green"},
	}

	fmt.Println(anonimousStruct)

	cat1 := cat{
		name: "Rodolfo",
		age:  11,
	}

	cat1.jump(22)
}

func (c cat) jump(time int) {

	fmt.Println(c.name, "has", c.age, "and jumped times", time)
}
