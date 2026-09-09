{
  "name": "{{.App.NormalizedName}}",
  "version": "1.0.0",
  "description": "{{.App.Name}} - Koa REST API",
  "private": true,
  "main": "dist/server.js",
  "scripts": {
    "build": "tsc",
    "start": "node dist/server.js",
    "dev": "tsx watch src/server.ts",
    "format": "prettier --write \"src/**/*.ts\"",
    "lint": "oxlint src/",
    "test": "vitest run",
    "test:watch": "vitest",
    "test:cov": "vitest run --coverage"
  },
  "dependencies": {
    "koa": "^2.15.0",
    "@koa/router": "^12.0.0",
    "@koa/cors": "^5.0.0",
    "koa-bodyparser": "^4.4.1",
    "koa-logger": "^3.2.1",
    "dotenv": "^16.4.5"
  },
  "devDependencies": {
    "@types/koa": "^2.15.0",
    "@types/koa__router": "^12.0.4",
    "@types/koa__cors": "^5.0.0",
    "@types/koa-bodyparser": "^4.3.12",
    "@types/koa-logger": "^3.1.5",
    "@types/node": "^24.0.0",
    "@vitest/coverage-v8": "^4.0.0",
    "oxlint": "^1.0.0",
    "prettier": "^3.4.0",
    "tsx": "^4.0.0",
    "typescript": "^5.7.0",
    "vitest": "^4.0.0"
  }
}
