# Specify a Base Image
FROM node:20-alpine

# Set the working directory
WORKDIR /usr/src/app

# Install some dependencies
COPY ./ ./
RUN npm install

# Default command
CMD ["npm", "start"]
