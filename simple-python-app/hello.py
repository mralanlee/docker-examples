from flask import Flask
app = Flask(__name__)

@app.route("/")
def hello():
    return "<html><img src='https://media.giphy.com/media/QvGpESBc34emw8XdDG/giphy.gif'></html>"

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=5000)
