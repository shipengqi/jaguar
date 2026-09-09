import { Context, Next } from 'koa'

export async function errorMiddleware(ctx: Context, next: Next): Promise<void> {
  try {
    await next()
  } catch (err: any) {
    ctx.status = err.status || 500
    ctx.body = { error: err.message || 'Internal Server Error' }
    ctx.app.emit('error', err, ctx)
  }
}
