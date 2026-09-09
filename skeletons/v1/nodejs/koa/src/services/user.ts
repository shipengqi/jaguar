interface User {
  id: number
  name: string
  email: string
}

const users: User[] = []
let nextId = 1

export async function listUsers(): Promise<User[]> {
  return users
}

export async function addUser(data: { name: string; email: string }): Promise<User> {
  const user: User = { id: nextId++, ...data }
  users.push(user)
  return user
}
