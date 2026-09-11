import { apiClient } from '@/lib/axios'
import { useQuery, useMutation, useQueryClient } from '@tanstack/vue-query'

interface User {
  id: string
  name: string
  email: string
}

export const userKeys = {
  all: ['users'] as const,
  lists: () => [...userKeys.all, 'list'] as const,
  detail: (id: string) => [...userKeys.all, 'detail', id] as const,
}

export const useUsers = () =>
  useQuery({
    queryKey: userKeys.lists(),
    queryFn: () => apiClient.get<User[]>('/users').then((r) => r.data),
    staleTime: 5 * 60 * 1000,
  })

export const useUser = (id: string) =>
  useQuery({
    queryKey: userKeys.detail(id),
    queryFn: () => apiClient.get<User>(`/users/${id}`).then((r) => r.data),
    enabled: !!id,
  })

export const useUpdateUser = () => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: Partial<User> }) =>
      apiClient.patch(`/users/${id}`, data),
    onSuccess: (_, { id }) => {
      queryClient.invalidateQueries({ queryKey: userKeys.detail(id) })
    },
  })
}
