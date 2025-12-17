POST http://localhost:8080/cars
Content-Type: application/json

{
  "brand": "Toyota",
  "model": "Corolla",
  "year": 2020
}

HTTP/1.1 201
[Asserts]
jsonpath "$.id" exists
jsonpath "$.brand" == "Toyota"
jsonpath "$.model" == "Corolla"
jsonpath "$.year" == 2020