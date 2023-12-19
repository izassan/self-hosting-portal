import json
from flask import Flask, render_template

app = Flask(__name__, template_folder='templates')
json = json.load(open("services.json"))


@app.route('/')
def index():
    return render_template('index.html', services=json['services'])


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=9444, debug=True)
