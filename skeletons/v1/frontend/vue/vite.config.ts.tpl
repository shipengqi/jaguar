import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: { '@': path.resolve(__dirname, './src') },
  },
  {{- if .Frontend.Embedded}}
  build: {
    outDir: '../staticfs/web',
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
