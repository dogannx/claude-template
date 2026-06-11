# Dev-oriented image (HMR). For production, build static assets and serve via nginx.
FROM node:lts-alpine
WORKDIR /app
COPY package.json yarn.lock* ./
RUN yarn install 2>/dev/null || true
COPY . .
EXPOSE 5173
CMD ["yarn", "dev", "--host"]
