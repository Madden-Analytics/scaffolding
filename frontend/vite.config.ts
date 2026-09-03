import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    // Calls to /api/* are forwarded to the Go backend, so the frontend can talk
    // to it during development without any CORS configuration.
    proxy: {
      // Anchored and slash-terminated so that unrelated paths such as
      // /apiary are not swept into the proxy.
      '^/api/': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
  },
})
