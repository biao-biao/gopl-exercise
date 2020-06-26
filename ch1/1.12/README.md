go run main.go web

/?cycles=4&cycles=20

### why values for r.Form are lists of strings

Because in theory, each field could be present multiple times, [see here][1]

[1]: https://stackoverflow.com/questions/34839811/how-to-retrieve-form-data-as-array
