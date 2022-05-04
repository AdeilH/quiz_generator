# Go Quiz Generator

## Generate a quiz through a json file and golang

### Version 0.1

1. Can Parse and Generate a Quiz
2. Json files can passed as flag
3. Final Score Page
4. Retry Option
5. End Quiz Option(Kills the process for now needs a better way)

### Improvements needed

1. Improve HTML and CSS
2. Look for scalibility
3. Improve handlers and routing
4. Refactor (priority: remove json responses that are not being used)
5. A json generator?
6. Render Code Snippets.
7. Make background image part of json

### How to run

1.No external dependencies so all you need to do is run

`go run *.go`

2.You can give flags for your json file and background image with -f and -i

`go run *.go -f quiz_file -i background_image`
