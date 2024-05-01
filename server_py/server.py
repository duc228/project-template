from flask import Flask, jsonify
from api.routes.todo_route import todo_route
from http import HTTPStatus

app = Flask(__name__)

app.register_blueprint(todo_route, url_prefix='/api/todo')

@app.route("/")
def hello_world():
    return "Hello, World!"

@app.route('/healthcheck', methods=['GET'])
def health_check():
    return jsonify({"status": "ok"}), HTTPStatus.OK

if __name__ == "__main__":
	app.run(host="0.0.0.0", port=5001, threaded=True, debug=True)