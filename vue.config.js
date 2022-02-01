/**
 * @type {import('@vue/cli-service').ProjectOptions}
 */

const path = require("path");
function resolve(dir) {
  return path.join(__dirname, dir);
}

module.exports = {
  publicPath: "/",
  outputDir: "build/www",
  assetsDir: "assets",
  indexPath: "index.html",
  pages: {
    index: {
      entry: "client/main.ts",
      template: "public/index.html",
      filename: "index.html",
      title: "My App",
    },
  },
  chainWebpack: (config) => {
    config.resolve.alias.set("@", resolve("client"));
  },
  lintOnSave: process.env.NODE_ENV !== "production",
  devServer: {
    port: 8081,
    https: true,
    proxy: "https://localhost:8443",
    overlay: {
      warnings: true,
      errors: true,
    },
  },
  productionSourceMap: true,
};
