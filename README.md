# FizzBuzz REST Server

The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by “fizz”, all multiples of 5 by “buzz”, and all multiples of 15 by “fizzbuzz”. The output would look like this :

```
"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,..."
```

The FizzBuzz REST Server exposes the REST API endpoint `GET /api/v1/fizzbuzz` that accepts five query parameters : two strings (say, `string1` and `string2`), and three integers (say, `int1`, `int2` and `limit`), and returns a JSON which contains a list of strings with numbers from 1 to limit, where :
- all multiples of `int1` are replaced by `string1`
- all multiples of `int2` are replaced by `string2`
- all multiples of `int1` and `int2` are replaced by `string1string2`

### Prerequisites

In order to build and run FizzBuzz, you need to install Go and Docker.

### Installing

#### Building FizzBuzz
```
make build
```

#### Running FizzBuzz in development environment
```
BIND_ADDRESS=localhost:8080 ./fizzbuzz
```

#### Building and running FizzBuzz in production environment using Docker
```
make docker
docker run -p 8080:8080 github.com/afourni/fizzbuzz
```

#### Using the API
```
curl 'http://localhost:8080/api/v1/fizzbuzz?int1=3&string1=fizz&int2=5&string2=buzz&limit=15'
```

## Running the tests

```
make test
```

## Running the benchmarks

```
make bench
```

## Authors

* **Anthony Fournier** - [afourni](https://github.com/afourni )

## License

This project is licensed under the "BEER-WARE licence" (Revision 42).
As long as you retain this notice you can do whatever you want with this stuff.
If we meet some day, and you think this stuff is worth it, you can buy me a beer in return.
