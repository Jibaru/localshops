let {
    base = "http://localhost:8080";
    correlationId = uuid();
}

do {
    url = "$base/shops";
    method = "POST";
    headers = {
        "Content-Type": "application/json",
        "X-Correlation-Id": correlationId
    };
    body = `{
        "name": "example",
        "description": "test",
        "latitude": -12.358355657471748,
        "longitude": -76.79657241488
    }`;
}