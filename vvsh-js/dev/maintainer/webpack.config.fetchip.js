
const path = require('path');

module.exports = {
  entry: './src/scripts/dingIPAddress.ts',

  module: {
    rules: [{
        test: /.tsx?$/,
        use: 'ts-loader'
        // configFile: "./tsconfig.json"
    }
    ]
  },
  resolve: {
    extensions: ['.tsx', '.ts', '.js']
  },
  output: {
    filename: 'dingIPAddress.js',
    path: path.resolve(__dirname, 'webpack')
  },

  mode: "development",
//   mode: "production",
  target: "node",
  node: {
    __dirname: false,
  }
};
