# GRAPHQL API TO GET THE SMALLEST PRIME ROMAN NUMBER

This is an example GraphQL API to receive a string that contain roman numerals and return the smallest prime roman numeral and the respective number.

## EXAMPLE

Mutation

```
mutation {
  search(input: { text: "AXIBIV" })
  {
    number
    value
  }
}
```

Return 

```
{
  "data": {
    "search": {
      "number": "XI",
      "value": 11
    }
  }
}
```

## HOW TO RUN

To start this api just run in terminal inside of the src folder:

> `go run cmd/server.go`

## DOCKER IMAGE
To build the docker image run the fallowing command in the terminal

>`docker build -t min_roman_numeral -f build/Dockerfile .`

To run the image, use the command bellow

>`docker run -d -p 8080:8080 min_roman_numeral`

## TESTS

A test file is included as `schema_resolver_test.go` it has a map containing the inputs and expected results.
