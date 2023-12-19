# adminuiy

Proof of concept for a simple tool to generate static html forms that can be hosted in an s3 bucket and submit to API endpoints you control.
Great for internal tooling where you don't want to build and maintain your own UI.

## Generate form and run
```
go run main.go ui.json
open output/root.html
```

<img width="1438" alt="Screen Shot 2023-12-19 at 2 22 16 pm" src="https://github.com/JaydenIvanovic/adminuiy/assets/3058997/58c83124-5445-4a84-9639-fc2f98437eeb">
<img width="968" alt="Screen Shot 2023-12-19 at 2 23 39 pm" src="https://github.com/JaydenIvanovic/adminuiy/assets/3058997/624ccff8-38fb-4fff-9e2c-f79b10d32510">
