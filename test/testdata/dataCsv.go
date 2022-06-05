package test

import "github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"

var DataCharacters = [][]string{
	{"1", "Rick Sanchez", "Alive", "Male", "https://rickandmortyapi.com/api/character/avatar/1.jpeg", "https://rickandmortyapi.com/api/character/1", "2017-11-04T18:48:46.250Z"},
	{"2", "Morty Smith", "Alive", "Male", "https://rickandmortyapi.com/api/character/avatar/2.jpeg", "https://rickandmortyapi.com/api/character/2", "2017-11-04T18:50:21.651Z"},
	{"3", "Summer Smith", "Alive", "Female", "https://rickandmortyapi.com/api/character/avatar/3.jpeg", "https://rickandmortyapi.com/api/character/3", "2017-11-04T19:09:56.428Z"},
	{"4", "Beth Smith", "Alive", "Female", "https://rickandmortyapi.com/api/character/avatar/4.jpeg", "https://rickandmortyapi.com/api/character/4", "2017-11-04T19:22:43.665Z"},
	{"5", "Jerry Smith", "Alive", "Male", "https://rickandmortyapi.com/api/character/avatar/5.jpeg", "https://rickandmortyapi.com/api/character/5", "2017-11-04T19:26:56.301Z"},
}

var ExpectedResult = entity.ResponseBody{
	Results: []entity.Character{
		{ID: 1, Name: "Rick Sanchez", Status: "Alive", Gender: "Male", Image: "https://rickandmortyapi.com/api/character/avatar/1.jpeg", Url: "https://rickandmortyapi.com/api/character/1", Created: "2017-11-04T18:48:46.250Z"},
		{ID: 2, Name: "Morty Smith", Status: "Alive", Gender: "Male", Image: "https://rickandmortyapi.com/api/character/avatar/2.jpeg", Url: "https://rickandmortyapi.com/api/character/2", Created: "2017-11-04T18:50:21.651Z"},
		{ID: 3, Name: "Summer Smith", Status: "Alive", Gender: "Female", Image: "https://rickandmortyapi.com/api/character/avatar/3.jpeg", Url: "https://rickandmortyapi.com/api/character/3", Created: "2017-11-04T19:09:56.428Z"},
		{ID: 4, Name: "Beth Smith", Status: "Alive", Gender: "Female", Image: "https://rickandmortyapi.com/api/character/avatar/4.jpeg", Url: "https://rickandmortyapi.com/api/character/4", Created: "2017-11-04T19:22:43.665Z"},
		{ID: 5, Name: "Jerry Smith", Status: "Alive", Gender: "Male", Image: "https://rickandmortyapi.com/api/character/avatar/5.jpeg", Url: "https://rickandmortyapi.com/api/character/5", Created: "2017-11-04T19:26:56.301Z"},
	},
}
