package test

import "github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"

var TestCharacters = [][]string{
	{"1", "Rick Sanchez", "Alive", "Male", "https://rickandmortyapi.com/api/character/avatar/1.jpeg", "https://rickandmortyapi.com/api/character/1", "2017-11-04T18:48:46.250Z"},
	{"2", "Morty Smith", "Alive", "Male", "https://rickandmortyapi.com/api/character/avatar/2.jpeg", "https://rickandmortyapi.com/api/character/2", "2017-11-04T18:50:21.651Z"},
	{"3", "Summer Smith", "Alive", "Female", "https://rickandmortyapi.com/api/character/avatar/3.jpeg", "https://rickandmortyapi.com/api/character/3", "2017-11-04T19:09:56.428Z"},
	{"4", "Beth Smith", "Alive", "Female", "https://rickandmortyapi.com/api/character/avatar/4.jpeg", "https://rickandmortyapi.com/api/character/4", "2017-11-04T19:22:43.665Z"},
	{"5", "Jerry Smith", "Alive", "Male", "https://rickandmortyapi.com/api/character/avatar/5.jpeg", "https://rickandmortyapi.com/api/character/5", "2017-11-04T19:26:56.301Z"},
}

var InfoData = entity.Info{
	Count: 5,
	Pages: 1,
	Next:  "",
	Prev:  "",
}
