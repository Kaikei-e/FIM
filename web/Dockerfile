FROM node:22.12.0-bullseye AS build

WORKDIR /app

COPY ./fim-web /app/

RUN npm install &&\
  npm run build

FROM node:22.12.0-slim AS production

WORKDIR /app

COPY --from=build /app /app

CMD ["npm", "start"]