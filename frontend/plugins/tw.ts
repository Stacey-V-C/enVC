import { exec } from 'child_process'
import path from 'path'
import type { Plugin } from 'vite';

const runTailwindOnHmr = () => ({
  name: 'run-tailwind-on-hmr',
  // @ts-ignore
  handleHotUpdate({ file }) {
    if (file.endsWith('.tsx')) {
      const tailwindInputPath = path.resolve(__dirname, '../src/input.css')
      const tailwindOutputPath = path.resolve(__dirname, '../src/output.css')
      exec(`npx tailwindcss -i ${tailwindInputPath} -o ${tailwindOutputPath}`, (err, stdout, stderr) => {
        if (err) {
          console.log(err)
          return
        }
        console.log(stdout)
      }
      )
    }
  }
})

export default runTailwindOnHmr