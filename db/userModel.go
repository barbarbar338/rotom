package db

type UserModel struct {
	ID                     string
	PokemonGOLevel         uint64
	PokemonGOUsername      string
	PokemonGOFriendCode    string
	PokemonUniteLevel      uint64
	PokemonUniteUsername   string
	PokemonUniteFriendCode string
}
