package graph

import (
	"github.com/duncanleo/brawl-scraper/db"
	"github.com/duncanleo/brawl-scraper/model"
	"github.com/graphql-go/graphql"
)

func Schema() graphql.Schema {
	var playerType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Player",
			Fields: graphql.Fields{
				"id":         &graphql.Field{Type: graphql.Int},
				"name":       &graphql.Field{Type: graphql.String},
				"name_color": &graphql.Field{Type: graphql.String},
				"tag":        &graphql.Field{Type: graphql.String},
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
