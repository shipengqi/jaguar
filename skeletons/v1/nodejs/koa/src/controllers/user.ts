import { Context } from 'koa'
import { listUsers, addUser } from '../services/user'

export async function getUsers(ctx: Context): Promise<void> {
  ctx.body = { users: await listUsers() }
}

export async function createUser(ctx: Context): Promise<void> {
  const body = ctx.request.body as { name: string; email: string }
  const user = await addUser(body)
  ctx.status = 201
  ctx.body = { user }
}
