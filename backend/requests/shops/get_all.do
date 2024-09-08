let {
    base = "http://localhost:8080";
}

do {
    url = "$base/shops";
    method = "GET";
    headers = {
        "Content-Type": "application/json"
    };
}