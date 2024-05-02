import requests

API_ENDPOINT = 'http://pt_go:5000'

class AlbumService:

	@classmethod
	def get_all(cls):
		data = cls.__get_req(cls, '/api/albums')
		return data
	
	def __get_req(self, path):
		resp = requests.get(f'{API_ENDPOINT}{path}')
		if resp.status_code != 200:
			raise Exception(f"[{resp.status_code}] Failed to get data from path: {path}")
		return resp.json()