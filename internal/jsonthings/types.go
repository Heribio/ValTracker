package jsonthings

type Username struct {
    Name    string
    Tag     string
}

type TokenData struct {
    ValApiToken string
}

type WriteFavoriteParams struct {
	Player Username
}

type FavoriteData struct {
	Favorites []Username
}
