package tv

/*

for tv, it will all be with seasons (matching anilist behavior). to represent entire shows, we do the first season.

GET    /v1/tv/trending           # trending tv shows (seasons)
GET    /v1/tv/upcoming           # upcoming tv shows (new seasons)
GET    /v1/tv/popular            # popular tv shows (all time -> popularity, top 100 -> average score) (first season shown as base)

# Series-level
GET    /v1/tv/{series_id}/recommendations  # similar/recommended series (last on priority)
	- have to figure something out as i will have first season represent an entire show. it's fine since every season in tmdb is just
	connected with series id and then season number so easy to find for any. the recommendation will just be generalized for the entire
	show not particular season.

# Season-level (for tracking/watchlist/ranking)
GET    /v1/tv/seasons/{season_id}     # full season details

*/

func handler() {}
