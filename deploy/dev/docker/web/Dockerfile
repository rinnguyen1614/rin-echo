# pull official base image
FROM node:18.6.0-alpine


ENV NODE_ENV development
RUN apk --no-cache add yarn 
ADD start.sh /
RUN chmod +x /start.sh
CMD ["/start.sh"]