from flask import Flask, jsonify
from api.routes.todo_route import todo_route
from api.routes.album_route import album_route
from http import HTTPStatus
from api.middlewares.log_middleware import Middleware

app = Flask(__name__)

app.wsgi_app = Middleware(app.wsgi_app)

app.register_blueprint(todo_route, url_prefix='/api/todo')
app.register_blueprint(album_route, url_prefix='/api/album')

@app.route("/")
def hello_world():
    return "Hello, World!"

@app.route('/healthcheck', methods=['GET'])
def health_check():
    return jsonify({"status": "ok"}), HTTPStatus.OK

if __name__ == "__main__":
	app.run(host="0.0.0.0", port=5001, threaded=True, debug=True)