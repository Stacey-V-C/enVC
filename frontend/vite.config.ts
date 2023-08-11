import { defineConfig } from 'vite';
import solidPlugin from 'vite-plugin-solid';
import runTailwindOnHmr from './plugins/tw';

export default defineConfig({
  plugins: [
    solidPlugin(), 
    // runTailwindOnHmr()
  ],
  build: {
    target: 'esnext',
  },
});
