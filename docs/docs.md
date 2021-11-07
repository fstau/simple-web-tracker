## Endpoints

```http
POST /v1/track

{
    "cts": 12345,
    "u": "anonymous",
    "s": "anonymous"
    "e": "pageview",
    "p": "/product",
    "q": "",
    "d": "",
}
```

```http
POST /v1/users

{
    "cts": 12345,
    "u": "anonymous",
    "ww": window.screen.width,
    "wh": window.screen.height,
    "waw": window.screen.availWidth,
    "wah": window.screen.availHeight,
    "o": window.screen.orientation.type,
}
```
