import { boot } from 'quasar/wrappers'
import VueGtag from 'vue-gtag'
// "async" is optional;
// more info on params: https://v2.quasar.dev/quasar-cli/boot-files
export default boot(({ app }) => {
  app.use(
    VueGtag,
    {
      appName: 'Web3Eye',
      config: {
        id: 'G-R045T863KM'
      },
    },
  )
})
