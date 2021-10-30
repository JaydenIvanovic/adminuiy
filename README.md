# adminuiy

A simple tool to generate static html forms that can be hosted in an s3 bucket and submit to API endpoints you control.
Great for internal tooling where you don't want to build and maintain your own UI.
Instead, define the forms as json, and work on what matters.

## Generate form
`go run main.go ui.json ui.tmpl.html > ui.html`
