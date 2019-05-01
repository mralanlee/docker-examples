# Simple Node.js Application

You will not need to run `yarn` or `npm install`.

Note: Check out the Dockerfile and try to play around with it. I added some comments, but feel free to reach out in case you have any questions.

### Getting Started
1. Build your Docker Image.

`docker build -t <your-user-name>/simple-node-app .`

2. Run the following command:

`docker run -p 3000:8080 <your-user-name>/simple-node-app`

The anatomy of this command is you are telling Docker to `run` an image with from your local machines port (3000) to map it to 8080. If you don't want to occupy your port 3000, you can use anything you want like 45712.

3. Open Chrome and visit `http://localhost:<local-machine-port>` (i.e. `http://localhost:3000`)

4. To quit, hit `Crtl + C`.

5. Win, play around with the app or Dockerfile. You can make changes to the expose port, but make sure to change your `ENV` for PORT.
