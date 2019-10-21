package model

import "time"

// Player represents a player
type Player struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Name      string     `gorm:"name" json:"name"`
	NameColor string     `gorm:"name_color" json:"name_color"`
	Tag       string     `gorm:"tag" json:"-"`
}

// PlayerData represents the state of a player at a given time
type PlayerData struct {
	ID                   uint       `gorm:"primary_key" json:"id"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
	DeletedAt            *time.Time `sql:"index" json:"-"`
	Player               Player     `json:"-"`
	PlayerID             uint       `json:"-"`
	TrophyCount          int        `json:"trophy_count"`
	ExpLevel             int        `json:"exp_level"`
	ExpPoints            int        `json:"exp_points"`
	ThreeV3Victories     int        `json:"3v3_victories"`
	SoloVictories        int        `json:"solo_victories"`
	DuoVictories         int        `json:"duo_victories"`
	BestRoboRumbleTime   int        `json:"best_robo_rumble_time"`
	BestTimeAsBigBrawler int        `json:"best_time_as_big_brawler"`
	TopBrawler           Brawler    `json:"top_brawler"`
	TopBrawlerID         uint       `json:"-"`
}

// Brawler represents a brawler
type Brawler struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	GameID    int        `gorm:"game_id" json:"game_id"`
	Name      string     `gorm:"name" json:"name"`
}
