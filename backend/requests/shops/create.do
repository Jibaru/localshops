let {
    base = "http://localhost:8080";
}

do {
    url = "$base/shops";
    method = "POST";
    headers = {
        "Content-Type": "application/json"
    };
    body = `{
        "name": "example",
        "description": "test",
        "latitude": -12.358355657471748,
        "longitude": -76.79657241488
    }`;
}