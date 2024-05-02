from werkzeug.wrappers import Request, Response

class Middleware():
    def __init__(self, app):
        self.app = app
        self.username = 'TestUser'
        self.password = 'TestPass'

    def __call__(self, environ, start_response):
        request = Request(environ)
        print('Logging Middleware: ', request.url, request.method, request.headers, request.data)
       
        return self.app(environ, start_response)
