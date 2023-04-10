# primeChecker
A simple prime number checker written in Golang

Send request via curl:

Error request
curl --location 'http://localhost:1010/check-prime' \
header 'Content-Type: application/json' \
data '[1, 2, "nan"]'

Success request
curl --location 'http://localhost:1010/check-prime' \
header 'Content-Type: application/json' \
data '[1, 2, 3, 4]'

Success response
{
    false,
    true,
    true,
    false
}

Error response
{
    "error": "invalid input"
}
