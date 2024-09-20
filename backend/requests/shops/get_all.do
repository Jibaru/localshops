let {
    base = "http://localhost:8080";
    correlationId = uuid();
}

do {
    url = "$base/shops";
    method = "GET";
    headers = {
        "Content-Type": "application/json",
        "X-Correlation-Id": correlationId
    };
}