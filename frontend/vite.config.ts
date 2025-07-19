import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import path from "path";
import tsconfigPaths from "vite-tsconfig-paths";
import fs from "fs";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), tsconfigPaths()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
  server: {
    host: "0.0.0.0",
    port: 3000,
    https: {
      key: fs.readFileSync(path.resolve(__dirname, "./localhost+2-key.pem")),
      cert: fs.readFileSync(path.resolve(__dirname, "./localhost+2.pem")),
    },
    proxy: {
      "/v1": {
        target: "https://backend:8080",
        changeOrigin: true,
        secure: false,
      },
    },
  },
});
