# Web Page Analyzer

## How to Run
Can either use `go run .` or build it using `go build` and then run the resulting `web-page-analyzer` artifact. If building the artifact, set the `GIN_MODE=release` environment variable to switch off the debug logs the Gin framework otherwise logs.

If needed, host and port can be configured as follows:
```shell
$ ./web-page-analyzer -host=127.0.0.1 -port=9090
```

By default the service runs on `localhost:8080`

## Assumptions/decisions Made
- Internal links are either links without the scheme and hostname or absolute URLs containing the same host name
- Inaccessible links are any response not in the 2xx class
- For simplicity, assumed that a form with one input field of type `password` is a login form. The down side is this will also pick registration forms which doesn't require the user to confirm the password as login forms.

## Further Improvements
- Could potentially improve error handling more. Especially to maybe differentiate between different errors and returning a more user friendly error than simply taking the error message provided by the error.
- Could probably refactor this to be better suited for unit testing. 
- Could maybe look at making the HEAD calls made to the links found in the page asynchronous so that it won't take so long to analyze a page.
