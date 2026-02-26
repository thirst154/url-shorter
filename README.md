# URL Shorter

A URL shortener written in Go.

Deployed at [url.thirst.dev](https://url.thirst.dev)

## Plan

- GET /
- Returns a html page with a form to enter a URL

- POST /
- Takes a URL and returns a shortened URL
- e.g {"url": "https://www.google.com"} -> {"short": "http://URL/2m4i0f"}
- e.g. {"url": "https://www.google.com", suggest:"google"} -> {"short": "http://URL/google"}
- e.g. {"url": "https://www.google.com", suggest:"google"} -> 409 Conflict {"error": "URL already exists"}

- GET /:short
- Takes a short URL and returns the original URL
