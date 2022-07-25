# pull official base image
FROM node:18.6.0-alpine


# add `/usr/src/web/node_modules/.bin` to $PATH
ENV PATH /usr/src/web/node_modules/.bin:$PATH

# set working directory
WORKDIR /usr/src/web

# install app dependencies
COPY ./src/web/package.json /usr/src/web/
COPY ./src/web/package-lock.json /usr/src/web/
# without --silent 
RUN npm install --force
RUN npm install react-scripts@5.0.1 -g --silent
# RUN mkdir -p /usr/src/web/node_modules/.cache && chmod -R 777 /usr/src/web/node_modules/.cache
# Copy in the rest of the application, still as root
COPY ./src/web/ /usr/src/web/

# start app
CMD ["npm", "start"]