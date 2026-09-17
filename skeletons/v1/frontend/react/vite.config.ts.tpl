import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path'

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: { '@': path.resolve(__dirname, './src') },
  },
  {{- if .Frontend.Embedded}}
  build: {
    outDir: '../internal/staticfs/web',
    emptyOutDir: true,
  },
  server: {
    proxy: { '/api': 'http://localhost:8080' },
  },
  {{- else}}
  build: {
    outDir: 'dist',
  },
  server: {
    port: 5173,
  },
  {{- end}}
})
