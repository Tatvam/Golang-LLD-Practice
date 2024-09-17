package models

type Player struct {
	name  string
	token *Token
}

func NewPlayer(name string, token *Token) *Player {
	return &Player{
		name:  name,
		token: token,
	}
}

func (f *Player) GetName() string {
	return f.name
}

func (f *Player) GetToken() *Token {
	return f.token
}
