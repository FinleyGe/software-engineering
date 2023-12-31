import { ViteAliases } from "vite-aliases";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    ViteAliases({
      prefix: "@",
    }),
  ],
});
