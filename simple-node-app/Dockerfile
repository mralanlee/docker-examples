# FROM dictates the image you want to pull from. You can search and find images and their versions
# from dockerhub.com
FROM node:latest

# Your docker container starts from / (root path), specify a folder you want your app to be.
# This portion is especially important if you have multiple containers starting in a docker-compose
WORKDIR /app

# Copy your files from the current directory to your WORKDIR
# If you have a build folder or a dist folder, you would could do /dist
COPY . /app

# Add your environment variables here
# First argument is the key, second is the value
ENV NODE_ENV PRODUCTION
ENV PORT 8080

# Run is like executing a command in the command line, depending on the type of image, you could use cURL or WGET
# or other neat stuff, however since we do not have linux and we're only centralized to node, we can only have
# node based commands (and npm)
RUN npm install

# This tells Docker which port to keep open.
# Since our node application runs on 8080, we will expose 8080.
# This does not mean it will open your local machines port 8080.
EXPOSE 8080

# Command that you want to use to execute your container
CMD ["npm", "start"]
