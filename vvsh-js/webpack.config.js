
const path = require('path');

module.exports = {
  entry: './src/main.ts',

  module: {
    rules: [{
        test: /.tsx?$/,
        use: 'ts-loader',
        parser: {
            amd: false, // 禁用 AMD
            commonjs: true, // 禁用 CommonJS
            system: false, // 禁用 SystemJS
            harmony: false, // 禁用 ES2015 Harmony import/export
            requireInclude: false, // 禁用 require.include
            requireEnsure: false, // 禁用 require.ensure
            requireContext: false, // 禁用 require.context
            browserify: false, // 禁用特殊处理的 browserify bundle
            requireJs: false, // 禁用 requirejs.*
            node: true, // 禁用 __dirname, __filename, module, require.extensions, require.main 等。
          }
        // configFile: "./tsconfig.json"
    }
    ],
    
  },
  resolve: {
    extensions: ['.tsx', '.ts', '.js']
  },
  output: {
    filename: 'index.js',
    path: path.resolve(__dirname, 'webpack')
  },

  mode: "development",
//   mode: "production",
  target: "node",
  node: {
    __dirname: false,
  }
};
