# Neo Postman

As a QA Engineer, we usually do an API testing with Postman but there are some weakness of Postman

```
1. Can't record the result
2. Can't do an assertion result
```

So their is neo-postman which help you to do an API testing and make your job more easier

|                  | Postman                | Neo-Post                              |
|------------------|------------------------|---------------------------------------|
| API Testing      | v                      | v                                     |
| Assertion result | v (but need put logic) | v (just put the expected result body) |
| Record           | x                      | v                                     |

### Sample Request:

```json
{
  "method": "POST",
  "path": "/v1/test/endpoint",
  "description": "As a user i should bla bla bla",
  "request_header": {
    "Content-Type": "application/json",
    "Authorization": "Bearer jwtToken"
  },
  "request_body": {
    "email": "randysteven12@gmail.com",
    "password": "test_1234"
  },
  "expected_response_code": 201,
  "expected_response": {
  }
}
```

| Field                  | mandatory | Description                                                       |
|------------------------|-----------|-------------------------------------------------------------------|
| method                 | M         | Required the HTTP method [POST, GET, PUT, PATCH, DELETE]          |
| path                   | M         | The endpoint of API                                               |
| request_header         | M         | The request header for API                                        |
| request_body           | O         | Give the request body for the API actually this field is optional |
| expected_response_code | M         | Expected response code that user/QA need to check                 |
| expected_response      | O         | Expected response body that user/QA need to check                 |

### Sample Response:

```json
{
  "message": "success get response",
  "data": {
    "test_result": {
      "id": 12,
      "result_status": "expected"
    }
  }
}
```

```json
{
  "message": "success get response",
  "data": {
    "test_result": {
      "id": 11,
      "result_status": "unexpected",
      "expected_response_code": 201,
      "actual_response_code": 200
    }
  }
}
```