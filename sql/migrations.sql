CREATE TABLE IF NOT EXISTS characters (
    id SERIAL,
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

CREATE TABLE IF NOT EXISTS duelPreparation (
    id SERIAL,
	selectingPlayer VARCHAR(255) NOT NULL,
    isReady         INT NOT NULL
);

CREATE TABLE IF NOT EXISTS duelPlayers(
    id SERIAL,
    preparationId INT NOT NULL,
    challenger VARCHAR(255) NOT NULL,
	challenged VARCHAR(255) NOT NULL,
    challengerChar VARCHAR(255),
	challengedChar VARCHAR(255)
);
