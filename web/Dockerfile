FROM node:22.14.0-bullseye AS build

WORKDIR /app

COPY ./fim-web /app/

RUN npm install -g npm@latest && \
  npm install && \
  npm run build

FROM node:22.14.0-slim AS production

WORKDIR /app

COPY --from=build /app /app

CMD ["npm", "start"]