import { defineConfig } from 'vite'
import angular from '@analogjs/vite-plugin-angular'

export default defineConfig({
  plugins: [angular()],
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
    port: 4200,
  },
  {{- end}}
})
