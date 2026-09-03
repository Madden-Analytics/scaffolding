import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  // compiler: true runs the React Compiler, which memoises components for you.
  plugins: [react({ compiler: true })],
})
