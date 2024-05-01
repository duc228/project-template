from flask import Blueprint, jsonify, Response
from http import HTTPStatus
import json
from services.todo_service import TodoService

todo_route = Blueprint('todo', __name__)

@todo_route.route('/', methods=['GET'])
def get_todo():
	todo_data = TodoService.get_all()
	return Response(response=json.dumps(todo_data),status=HTTPStatus.OK,content_type='application/json')

@todo_route.route('/<int:todo_id>', methods=['GET'])
def get_todo_by_id(todo_id):
	todo_data = TodoService.get_by_id(todo_id)
	return jsonify({"data": todo_data}), HTTPStatus.OK