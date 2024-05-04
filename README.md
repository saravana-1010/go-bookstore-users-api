# go-bookstore-users-api
Bookstore users api in Golang

Rest error structure examples:

{
    "message": "invalid json body",
    "code": 400,
    "error": "bad_request"
}

{
    "message": "user 123 not found",
    "code": 404,
    "error": "not_found"
}

{
    "message": "database is down",
    "code": 500,
    "error": "internal_server_error"
}
