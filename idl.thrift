namespace go api

struct Request {
        1: string message
}

struct Response {
        1: string message
}

service Hello {
    Response Echo(1: Request req)
    string Hello(string req)
}