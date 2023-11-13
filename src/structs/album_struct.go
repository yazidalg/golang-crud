package album_struct

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var Albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "Al", Price: 55.12},
	{ID: "2", Title: "Big Brain", Artist: "Ahr", Price: 78.11},
	{ID: "3", Title: "Brain Trut", Artist: "Yunas", Price: 88.54},
}
