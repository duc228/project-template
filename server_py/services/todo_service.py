import requests

API_ENDPOINT = 'https://jsonplaceholder.typicode.com'

class TodoService:

	@classmethod
	def get_all(cls):
		todo_data = cls.__get_req(cls, '/todos')
		return todo_data
	
	@classmethod
	def get_by_id(cls, todo_id):
		todo_data = cls.__get_req(cls, f'/todos/{todo_id}')
		return todo_data
	
	def __get_req(self, path):
		resp = requests.get(f'{API_ENDPOINT}{path}')
		if resp.status_code != 200:
			raise Exception(f"[{resp.status_code}] Failed to get data from path: {path}")
		return resp.json()