import {Configuration} from '@nuxt/types'

const nuxtConfig: Configuration = {
  mode: 'spa',
  env: {},
  head: {
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      {
        hid: 'description',
        name: 'description',
        content: process.env.npm_package_description || '',
      }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
    ],
  },
  loading: { color: '#3B8070' },
  // css: ['~/assets/css/main.css'],
  plugins: ['@/plugins/compositionAPI'],
  build: {
    extend(config, { isClient }) {
      if (isClient) {
        config.devtool = '#source-map'
      }
    },
  },
  buildModules: [
    '@nuxt/typescript-build',
  ],
  modules: [
    'nuxt-buefy',
    '@nuxtjs/axios',
    '@nuxtjs/pwa',
    '@nuxtjs/toast',
    ['nuxt-sass-resources-loader', ['@/assets/blumatheme.scss']],
  ],
  axios: {},
  toast: {
    position: 'top-center',
    duration: 2000,
    fullWidth: true,
    iconPack : 'material',
    register: [
      {
        name: 'allSuccess',
        message: 'ファイルの送信が完了しました',
        options: {
          type: 'info',
        },
      },
      {
        name: 'badRequest',
        message: 'リクエストが不正です',
        options: {
          type: 'warn',
        },
      },
      {
        name: 'unknownError',
        message: 'エラーが発生しました',
        options: {
          type: 'error',
        },
      },
    ],
  },

  generate: {
    fallback: true,
  },
};

module.exports = nuxtConfig;
