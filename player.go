package main

import "log"

// Item a representation of an Item and it's properties.
type Item struct {
	Name   string
	ID     int
	Dmg    int
	Weight int
	Armor  int
	Value  int
}

// Inventory holds an endless number of Items
type Inventory struct {
	Items    []Item
	Capacity int
}

// Stats contains a groupped information about stats of a character
type Stats struct {
	Strenght     int
	Agility      int
	Intelligence int
	Perception   int
	Luck         int
	Constitution int
}

// Body Represents a body of a Player which defines what he wears,
// Player will always automatically wear the best gear.
type Body struct {
	LRing   Item
	RRing   Item
	Armor   Item
	Head    Item
	Weapond Item
	Shield  Item
}

// Cast the cast of a player, like mage, rouge, warrior...
type Cast struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// Race the race of the player, like elf, gnome, human, dwarf...
type Race struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// Character is a player character.
type Character struct {
	ID          string
	Inventory   Inventory
	Body        Body
	Name        string
	Stats       Stats
	Hp          int
	MaxHp       int
	CurrentXp   int
	NextLevelXp int
	Gold        int
	Level       int
	Race        int
	Cast        int
}

// Rest will Replenish Health.
func (c *Character) Rest() {
	c.Hp = c.MaxHp
	log.Println("Player is fully rested.")
	mdb.Update(*c)
}

// SellItems will clear the inventory.
func (c *Character) SellItems() {
	for _, v := range c.Inventory.Items {
		c.Gold += v.Value
	}
	log.Println("Player has sold all items. Current gold is:", c.Gold)
	c.Inventory.Items = []Item{}
	mdb.Update(*c)
}

// Attack an enemy during an encounter.
func (c *Character) Attack(e Enemy) {
	// Fight until enemy is dead, or player is below 25%.
	log.Println("Attacking enemy:", e)
	playerHpLimit := int(float64(c.Hp) * 0.25)
	playerDamage := c.Body.Weapond.Dmg - e.Armor
	if playerDamage <= 0 {
		playerDamage = 1
	}
	enemyDamage := e.Damage - (c.Body.Armor.Armor + c.Body.Shield.Armor + c.Body.LRing.Armor + c.Body.RRing.Armor)
	if enemyDamage <= 0 {
		enemyDamage = 1
	}
	for c.Hp > playerHpLimit && e.Hp > 0 {
		e.Hp -= playerDamage
		c.Hp -= enemyDamage
		// log.Println("player hp:", c.Hp)
		// log.Println("enemy hp:", e.Hp)
	}
	if e.Hp <= 0 {
		log.Println("Player won!")
		c.CurrentXp += e.Xp
		//TODO: award items here as well.
		mdb.Update(*c)
		return
	}
	log.Println("Enemy won. Player has fled with hp:", c.Hp)
	mdb.Update(*c)
}
