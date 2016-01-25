package main

//Item a representation of an Item and it's properties
type Item struct {
	Name   string
	Dmg    int
	Size   int
	Weight int
}

//Inventory holds an endless number of Items
type Inventory struct {
	Items []Item
}

//Stats contains a groupped information about stats of a character
type Stats struct {
	Strenght     int
	Agility      int
	Intelligence int
	Perception   int
	Luck         int
}

//Cast the cast of a player, like mage, rouge, warrior...
type Cast struct {
	name string
	id   int
}

//Race the race of the player, like elf, gnome, human, dwarf...
type Race struct {
	name string
	id   int
}

//Character is a player character
type Character struct {
	ID        string
	Inventory Inventory
	Name      string
	Stats     Stats
	Gold      int
	Level     int
	Race      int
	Cast      int
}
