import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue({
    template: {
      compilerOptions: {
        isCustomElement: (tag) => ['menu-bar', 'logo', 'user-header', 'main-container', 'left-panel', 'contollr-container'].includes(tag)
      }
    }
  })],
  resolve: {
    extensions: ['.vue', '.js', '.jsx', '.json']
  }
})
