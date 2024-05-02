from flask import Flask, jsonify
from api.routes.todo_route import todo_route
from api.routes.album_route import album_route
from http import HTTPStatus
from api.middlewares.log_middleware import Middleware
from werkzeug.exceptions import HTTPException
import json
from common.error.error import CustomException

app = Flask(__name__)

# Middlware
app.wsgi_app = Middleware(app.wsgi_app)

app.register_blueprint(todo_route, url_prefix='/api/todo')
app.register_blueprint(album_route, url_prefix='/api/album')

@app.route("/")
def hello_world():
    return "Hello, World!"

@app.route('/healthcheck', methods=['GET'])
def health_check():
    return jsonify({"status": "ok"}), HTTPStatus.OK

# Error handling
# https://flask.palletsprojects.com/en/3.0.x/errorhandling/#generic-exception-handlers
@app.errorhandler(HTTPException)
def handle_exception(e):
    """Return JSON instead of HTML for HTTP errors."""
    # start with the correct headers and status code from the error
    response = e.get_response()
    # replace the body with JSON
    response.data = json.dumps({
        "code": e.code,
        "message": e.description,
    })
    response.content_type = "application/json"
    return response

# @app.errorhandler(CustomException)
# def invalid_api_usage(e):
#     return jsonify(e.to_dict()), e.status_code

if __name__ == "__main__":
	app.run(host="0.0.0.0", port=5001, threaded=True, debug=True)