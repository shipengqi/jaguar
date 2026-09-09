import Router from '@koa/router'
import { getUsers, createUser } from '../controllers/user'

export const userRouter = new Router()

userRouter.get('/', getUsers)
userRouter.post('/', createUser)
