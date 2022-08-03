import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    host: '127.0.0.1',
    port: '3000',
  },
  resolve: {
    alias: {
      "@": path.join(__dirname, "src"),
      "~": path.join(__dirname, "node_modules")
    }
  },
  css:{
    preprocessorOptions: {
       scss: {
         additionalData: `@import "@/styles/index.scss";`
       }
     }
  } 
})
