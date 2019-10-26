package graph

import (
	"github.com/duncanleo/brawl-scraper/db"
	"github.com/duncanleo/brawl-scraper/model"
	"github.com/graphql-go/graphql"
)

func Schema() graphql.Schema {
	var brawlerType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Brawler",
			Fields: graphql.Fields{
				"id":      &graphql.Field{Type: graphql.Int},
				"game_id": &graphql.Field{Type: graphql.Int},
				"name":    &graphql.Field{Type: graphql.String},
			},
		},
	)

	var playerDataType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "PlayerData",
			Fields: graphql.Fields{
				"id":                       &graphql.Field{Type: graphql.Int},
				"created_at":               &graphql.Field{Type: graphql.DateTime},
				"trophy_count":             &graphql.Field{Type: graphql.Int},
				"exp_level":                &graphql.Field{Type: graphql.Int},
				"exp_points":               &graphql.Field{Type: graphql.Int},
				"tvt_victories":            &graphql.Field{Type: graphql.Int},
				"solo_victories":           &graphql.Field{Type: graphql.Int},
				"duo_victories":            &graphql.Field{Type: graphql.Int},
				"best_robo_rumble_time":    &graphql.Field{Type: graphql.Int},
				"best_time_as_big_brawler": &graphql.Field{Type: graphql.Int},
				"top_brawler": &graphql.Field{
					Type:        brawlerType,
					Description: "Get top brawler",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						playerData, ok := p.Source.(model.PlayerData)
						if ok {
							var brawler model.Brawler
							err := db.DB.
								Where(model.Brawler{ID: playerData.TopBrawlerID}).
								Find(&brawler).
								Error
							return brawler, err
						}
						return nil, nil
					},
				},
			},
		},
	)
	var playerType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Player",
			Fields: graphql.Fields{
				"id":         &graphql.Field{Type: graphql.Int},
				"created_at": &graphql.Field{Type: graphql.DateTime},
				"name":       &graphql.Field{Type: graphql.String},
				"name_color": &graphql.Field{Type: graphql.String},
				"tag":        &graphql.Field{Type: graphql.String},
				"datas": &graphql.Field{
					Type:        graphql.NewList(playerDataType),
					Description: "Get player datas by player id",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						player, ok := p.Source.(model.Player)
						if ok {
							var playerDatas []model.PlayerData
							err := db.DB.
								Where(model.PlayerData{PlayerID: player.ID}).
								Find(&playerDatas).
								Error
							return playerDatas, err
						}
						return nil, nil
					},
				},
			},
		},
	)

	var queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"player": &graphql.Field{
					Type:        playerType,
					Description: "get player by id",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						id, ok := p.Args["id"].(int)
						if ok {
							var player model.Player
							err := db.DB.Find(&player, id).Error
							return player, err
						}
						return nil, nil
					},
				},
				"players": &graphql.Field{
					Type:        graphql.NewList(playerType),
					Description: "Get players",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						var players []model.Player
						err := db.DB.Find(&players).Error
						return players, err
					},
				},
				"brawlers": &graphql.Field{
					Type:        graphql.NewList(brawlerType),
					Description: "Get brawlers",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						var brawlers []model.Brawler
						err := db.DB.Find(&brawlers).Error
						return brawlers, err
					},
				},
			},
		},
	)

	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)
	return schema
}
