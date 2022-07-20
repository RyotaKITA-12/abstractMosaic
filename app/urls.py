from controllers import *

app.add_api_route('/', index)
app.add_api_route('/mosaic/', create_mosaic, methods=['POST'])
