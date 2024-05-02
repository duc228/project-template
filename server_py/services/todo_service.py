from flask import abort
import requests
from common.error.error import CustomException

API_ENDPOINT = 'https://jsonplaceholder.typicode.com'

class TodoService:

	@classmethod
	def get_all(cls):
		todo_data = cls.__get_req(cls, '/todos')
		return todo_data
	
	@classmethod
	def get_by_id(cls, todo_id):
		todo_data = cls.__get_req(cls, f'/tod1os/{todo_id}')
		return todo_data
	
	def __get_req(self, path):
		resp = requests.get(f'{API_ENDPOINT}{path}')
		if resp.status_code != 200:
			# raise CustomException("No such user!", status_code=404)
			abort(400, description="Resource not foundggg123")
		return resp.json()