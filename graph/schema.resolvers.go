package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/achristie/gql-sample/graph/generated"
	"github.com/achristie/gql-sample/graph/model"
)

func (r *mutationResolver) UpsertCharacter(ctx context.Context, input model.CharacterInput) (*model.Character, error) {
	id := input.ID
	var character model.Character
	character.Name = input.Name
	character.CliqueType = input.CliqueType

	n := len(r.Resolver.CharacterStore)
	if n == 0 {
		r.Resolver.CharacterStore = make(map[string]model.Character)
	}

	if id != nil {
		cs, ok := r.Resolver.CharacterStore[*id]
		if !ok {
			return nil, fmt.Errorf("not found")
		}
		if input.IsHero != nil {
			character.IsHero = *input.IsHero
		} else {
			character.IsHero = cs.IsHero
		}
		r.Resolver.CharacterStore[*id] = character
	} else {
		nid := strconv.Itoa(n + 1)
		character.ID = nid
		if input.IsHero != nil {
			character.IsHero = *input.IsHero
		}
		r.Resolver.CharacterStore[nid] = character
	}

	return &character, nil
}

func (r *queryResolver) Character(ctx context.Context, id string) (*model.Character, error) {
	character, ok := r.Resolver.CharacterStore[id]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return &character, nil
}

func (r *queryResolver) Characters(ctx context.Context, cliqueType model.CliqueType) ([]*model.Character, error) {
	characters := make([]*model.Character, 0)
	for idx := range r.Resolver.CharacterStore {
		character := r.Resolver.CharacterStore[idx]
		if character.CliqueType == cliqueType {
			characters = append(characters, &character)
		}
	}
	return characters, nil
}

func (r *queryResolver) Outage(ctx context.Context, id string) (*model.WRDOutage, error) {
	// var o *model.WRDOutage
	client, err := NewClient(
		"https://api.platts.com/refinery-data/v1/outage-alerts?PageSize=2",
		"")

	if err != nil {
		fmt.Errorf("error creating client %s", err)
	}

	resp, err := client.Get()
	if err != nil {
		fmt.Errorf("Problem fetching data, %s", err)
	}

	// bytes, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Errorf("Error converting data to bytes, %s", err)
	// }

	var obj map[string]json.RawMessage
	err = json.NewDecoder(resp.Body).Decode(&obj)

	if err != nil {
		fmt.Errorf("Problem decoding json, %v", err)
	}

	results := obj["results"]

	model := &model.WRDOutage{}
	err = json.Unmarshal(results, &model)

	if err != nil {
		fmt.Errorf("Error unmarshaling raw json, %v", err)
	}

	fmt.Printf("%s, %s, %s, %s", results['outagejId'], results[''])

	// model := &model.WRDOutage{}
	// err = json.NewDecoder(resp.Body).Decode(model)

	// if err != nil {
	// 	fmt.Errorf("issue decoding json, %s", err)
	// }

	// fmt.Printf("data: %v", model)

	return model, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Pogues(ctx context.Context) ([]*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) Kooks(ctx context.Context) ([]*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}
