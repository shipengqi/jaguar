import { Injectable, signal } from '@angular/core'

interface User {
  id: string
  name: string
  email: string
}

@Injectable({ providedIn: 'root' })
export class AuthService {
  private _user = signal<User | null>(null)
  private _token = signal<string | null>(localStorage.getItem('token'))

  isAuthenticated = () => !!this._token()
  user = this._user.asReadonly()

  setUser(user: User) {
    this._user.set(user)
  }

  setToken(token: string) {
    this._token.set(token)
    localStorage.setItem('token', token)
  }

  logout() {
    this._user.set(null)
    this._token.set(null)
    localStorage.removeItem('token')
  }
}
