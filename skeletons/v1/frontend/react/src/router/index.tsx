import { createBrowserRouter } from 'react-router-dom'
import { useAuthStore } from '@/store/auth.store'
import { Navigate } from 'react-router-dom'
import LoginPage from '@/views/auth/LoginPage'

function ProtectedRoute({ children }: { children: React.ReactNode }) {
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated)
  return isAuthenticated ? <>{children}</> : <Navigate to="/login" replace />
}

export const router = createBrowserRouter([
  { path: '/login', element: <LoginPage /> },
  {
    path: '/',
    element: <ProtectedRoute><div>Dashboard</div></ProtectedRoute>,
  },
])
