let {
    base = "http://localhost:8080";
}

do {
    url = "$base/products";
    method = "POST";
    headers = {
        "Content-Type": "application/json"
    };
    body = `{
        "name": "test product",
        "description": "test product description",
        "price_amount": 120,
        "price_currency": "PEN"
    }`;
}