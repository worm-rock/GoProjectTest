package controller

import (
	//"clean/core"
	// "pokedex/domain"
	// "presenter"

	"gateway"
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func GetDebile() core.PaginationResult[domain.Pokemon] {
//		return core.NewPaginationResult([]domain.Pokemon{{Name: "pikatchu"}, {Name: "tortank"}}, 2, 1, 0)
//	}
//
// []domain.Pokemon{{Name: "pikatchu"}, {Name: "tortank"}
func TestGetPokemonIntegration(t *testing.T) {
	assert := assert.New(t)
	controller := NewControllerTestWithInit([]domain.Pokemon{{Name: "pikatchu"}})
	presenter := controller.GetMyPokemons()

	expected := "[{\"name\":\"pikatchu\"}]"
	actual, _ := presenter.Print()

	assert.Equal(expected, actual)
}

func TestAddPokemonIntegration(t *testing.T) {
	assert := assert.New(t)

	controller := NewControllerTest()
	presenter := controller.AddPokemons([]string{"pikatchu"})

	expected := "[{\"name\":\"pikatchu\"}]"
	actual, _ := presenter.Print()

	assert.Equal(expected, actual)
}

func TestAddAndGetPokemonIntegration(t *testing.T) {
	//Given
	assert := assert.New(t)
	controller := NewControllerTestWithInit([]domain.Pokemon{})
	//when
	controller.AddPokemons([]string{"pikatchu"})
	presenter := controller.GetMyPokemons()
	//Then
	expected := "[{\"name\":\"pikatchu\"}]"
	actual, _ := presenter.Print()
	assert.Equal(expected, actual)
}

func NewControllerTest() PokedexController[string] {
	repo := gateway.Repo[domain.Pokemon]{Context: []domain.Pokemon{}}
	return NewControllerJSonAndMemory(repo)
}
func NewControllerTestWithInit(buf []domain.Pokemon) PokedexController[string] {
	repo := gateway.Repo[domain.Pokemon]{Context: buf}
	return NewControllerJSonAndMemory(repo)
}
