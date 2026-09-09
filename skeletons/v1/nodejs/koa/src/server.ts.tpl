import { app } from './app'
import { config } from './config'

app.listen(config.port, config.host, () => {
  console.log(`Server running at http://${config.host}:${config.port}`)
})
