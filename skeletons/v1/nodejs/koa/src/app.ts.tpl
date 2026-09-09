import Koa from 'koa'
import Router from '@koa/router'
import cors from '@koa/cors'
import bodyParser from 'koa-bodyparser'
import logger from 'koa-logger'
import { config } from './config'
import { userRouter } from './routes'

export const app = new Koa()

app.use(cors())
app.use(bodyParser())
app.use(logger())

const router = new Router({ prefix: '/api/v1' })
router.use('/users', userRouter.routes(), userRouter.allowedMethods())

app.use(router.routes())
app.use(router.allowedMethods())

app.on('error', (err: Error, ctx: Koa.Context) => {
  console.error('server error', err, ctx)
})
