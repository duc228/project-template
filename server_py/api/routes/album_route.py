from flask import Blueprint, Response
from http import HTTPStatus
import json
from services.album_service import AlbumService
from api.middlewares.auth_middleware import hello_middleware

album_route = Blueprint('album', __name__)

@album_route.route('/', methods=['GET'])
@hello_middleware
def get_todo():
	data = AlbumService.get_all()
	return Response(response=json.dumps({"code": 200, "data": data['data']}),status=HTTPStatus.OK,content_type='application/json')

