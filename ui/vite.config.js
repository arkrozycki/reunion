import { defineConfig } from 'vite'
import Vue from '@vitejs/plugin-vue'
import Pages from 'vite-plugin-pages'
import ViteComponents from 'vite-plugin-components'
import { resolve } from 'path'
import pkg from './package.json'

process.env.VITE_APP_BUILD_EPOCH = new Date().getTime()
process.env.VITE_APP_VERSION = pkg.version

export default defineConfig({
  plugins: [
    Vue({
      include: [/\.vue$/],
    }),
    Pages({
      pagesDir: [
        {
          dir: 'src/pages',
          baseRoute: '',
        },
      ],
    }),
    ViteComponents({
      extensions: ['vue'],
      dirs: ['src/components', 'src/layouts']
    }),
  ],
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src'),
    },
  },
  server: {
    open: false,
    host: '0.0.0.0',
    port: 8088,
    hmr: {
      // protocol: 'ws',
      // clientPort: 8088,
    },
  },
})
