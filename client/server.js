import fs from 'node:fs/promises'
import path from 'node:path'
import { fileURLToPath } from 'node:url'
import express from 'express'

const isProduction = process.env.NODE_ENV === 'production'
const port = process.env.VITE_PORT || 3000
const base = process.env.BASE || '/'

const __dirname = path.dirname(fileURLToPath(import.meta.url))
const resolve = (p) => path.resolve(__dirname, p)

async function createServer() {
  const app = express()
  let vite

  if (!isProduction) {
    vite = await (
      await import('vite')
    ).createServer({
      server: { middlewareMode: true },
      appType: 'custom',
      base,
    })
    app.use(vite.middlewares)
  } else {
    app.use('/assets', express.static(resolve('dist/client/assets'), {
      immutable: true, // Arquivos com hash podem ser cacheados para sempre
      maxAge: '1y'
    }))

    app.use(express.static(resolve('dist/client'), { index: false }))
  }

  app.use('/{*splat}', async (req, res) => {
    try {
      const url = req.originalUrl

      let template
      let render

      if (!isProduction) {
        template = await fs.readFile(resolve('index.html'), 'utf-8')
        template = await vite.transformIndexHtml(url, template)
        render = (await vite.ssrLoadModule('/src/entry-server.tsx')).render
      } else {
        template = await fs.readFile(resolve('dist/client/index.html'), 'utf-8')
        render = (await import('./dist/server/entry-server.js')).render
      }

      const appHtml = render(url) // Renderiza o componente React para HTML
      const html = template.replace(`<!--ssr-outlet-->`, appHtml) // Injeta o HTML no template

      res.status(200).set({ 'Content-Type': 'text/html' }).end(html)
    } catch (e) {
      if (vite) {
        vite.ssrFixStacktrace(e)
      }
      console.log(e.stack)
      res.status(500).end(e.stack)
    }
  })

  return { app }
}

createServer().then(({ app }) =>
  app.listen(port, () => {
    console.log(`http://localhost:${port}`)
  })
)
