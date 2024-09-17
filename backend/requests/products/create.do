let {
    base = "http://localhost:8080";
    shopId = "c3c18fa3-4b98-4b36-b314-3c0ec0fda3ac";
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
        "shop_id": "$shopId",
        "price_amount": 120,
        "price_currency": "PEN"
    }`;
}