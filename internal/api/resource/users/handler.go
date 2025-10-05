package users

/*

users api resouce is going to be done last (but before redis) as the db needs to be set up first

GET    /v1/users/{id}/watchlist           # list of movies + tv seasons
											(potentially add option with type query
											 param to seperate movies and tv)
POST   /v1/users/{id}/watchlist           # add movie or tv season entry
PATCH  /v1/users/{id}/watchlist/{media_id}  # edit movie or tv season entry
DELETE /v1/users/{id}/watchlist/{media_id}  # remove movie or tv season entry

*/

func handler() {}
