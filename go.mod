module example.com/music-project-portal-APIs-GO

replace (
	example.com/music-project-portal-APIs-GO/notionApi => ./notionApi
	example.com/music-project-portal-APIs-GO/parser => ./parser
)

go 1.18

require (
	example.com/music-project-portal-APIs-GO/notionApi v0.0.0-00010101000000-000000000000
	example.com/music-project-portal-APIs-GO/parser v0.0.0-00010101000000-000000000000
)
