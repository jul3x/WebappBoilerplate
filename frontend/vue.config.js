const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  transpileDependencies: true,
    pages: {
        index: {
            entry: 'src/main.ts',
            template: 'public/index.html',
            filename: 'index.html',
            title: 'WebApp Template',
            chunks: ['chunk-vendors', 'chunk-common', 'index']
        },
    }
});
