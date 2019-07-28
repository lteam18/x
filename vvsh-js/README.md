# vvsh

# How to start

Using npx

```bash
npx vvsh
```


```bash
npm install -g vvsh
vvsh
```

# Install independent binary

## Install in user mode, or install in global mode

Linux, MacOS, Windows Git-Bash

```bash
wget ""
```


```bash
curl ""
```

windows

```bat

```


## Webpack Config

```javascript

const path = require('path');

module.exports = {
  entry: './src/ghlx.ts',

  module: {
    rules: [{
        test: /.tsx?$/,
        use: 'ts-loader'
        // configFile: "./tsconfig.json"
    }]
  },
  resolve: {
    extensions: ['.tsx', '.ts', '.js']
  },
  output: {
    filename: 'ghlx.js',
    path: path.resolve(__dirname, 'webpack')
  },

  mode: "development",
//   mode: "production",
  target: "node",
  node: {
    __dirname: false,
  }
};

```
