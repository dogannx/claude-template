# Dev-oriented image (HMR). For production, build static assets and serve via nginx.
FROM node:lts-alpine
WORKDIR /app
RUN corepack enable
COPY package.json yarn.lock .yarnrc.yml ./
RUN yarn install --immutable
COPY . .
EXPOSE 5173
CMD ["yarn", "dev", "--host"]
