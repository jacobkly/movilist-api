package search

/*

one endpoint with filters

returns season-level tv results

GET    /v1/search?q=batman                 # global multi-search (movies, tv seasons, actors)
GET    /v1/search?q=batman&type=movie     # movie-only search
GET    /v1/search?q=batman&type=tv        # tv season-only search
GET    /v1/search?q=ewan&type=actor       # actor-only search (low priority)

*/

func handler() {}
