package domain

import "clean/core"

type IAddPokemon interface {
	Add(AddPokemonQuery) ([]Pokemon, error)
}
type IGetPokedemon interface {
	Get(GetPokemonQuery) (core.PaginationResult[Pokemon], error)
}
type IGetPokedex interface {
	Get(GetPokemonQuery) (core.PaginationResult[Pokemon], error)
}
