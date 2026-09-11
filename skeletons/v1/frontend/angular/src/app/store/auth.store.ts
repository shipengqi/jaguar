import { signalStore, withState, withMethods, patchState } from '@ngrx/signals'

interface User {
  id: string
  name: string
  email: string
}

interface AuthState {
  user: User | null
  token: string | null
}

export const AuthStore = signalStore(
  { providedIn: 'root' },
  withState<AuthState>({ user: null, token: null }),
  withMethods((store) => ({
    setUser(user: User) {
      patchState(store, { user })
    },
    setToken(token: string) {
      patchState(store, { token })
    },
    logout() {
      patchState(store, { user: null, token: null })
    },
  }))
)
