CREATE TABLE IF NOT EXISTS characters (
    charName VARCHAR(255) NOT NULL,
    player BIGINT NOT NULL,
    weaponSkill INT NOT NULL,
    balisticSkill INT NOT NULL,
    strength INT NOT NULL,
    endurance INT NOT NULL,
    agility INT NOT NULL,
    willpower INT NOT NULL,
    fellowship INT NOT NULL,
    hitpoints INT NOT NULL
);