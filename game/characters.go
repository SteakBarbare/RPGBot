package game

/*
WeaponSkill			Chance to hit or defend for a melee attack
BalisticSkill		Chance to hit for a ranged attack
Strength			Increase melee damage every 10 points & help carrying heavy burden
Endurance			Increase damage resistance every 10 points & help resisting some effects
Agility				Increase dodge chances and help resisting some effects
Willpower			Help using spells & resisting some effects
Fellowship			Used for interaction with NPCs
Hitpoints			Character hitpoints, if it reaches 0, the character may suffer minor to lethal injuries
*/

type PlayerChar struct {
	Name, Player                                                                               string
	WeaponSkill, BalisticSkill, Strength, Endurance, Agility, Willpower, Fellowship, Hitpoints int
}
