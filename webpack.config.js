const source = "./";
const target = "./static";
const webpack = require("webpack");
const path = require("path");
const ExtractTextPlugin = require("extract-text-webpack-plugin");
const { CheckerPlugin, TsConfigPathsPlugin } = require("awesome-typescript-loader");
const context = path.join(__dirname, source);

module.exports = {
   entry: {
      app: "./modules/person/app.tsx",
   },
   output: {
      path: path.join(__dirname, target),
      filename: "js/[name].js",
      publicPath: "/static/"
   },
   module: {
      rules: [
         {
            test: /\.tsx?$/,
            exclude: [/node_modules/, /test\.tsx?$/, /lib\/types\//],
            loader: ["awesome-typescript-loader"]
         },
         // Extract CSS and SASS imports into separate combined files.
         {
            test: /\.scss$/,
            exclude: /node_modules/,
            loader: ExtractTextPlugin.extract({
               fallback: "style-loader",
               // disable URL following (fonts and images)
               // https://github.com/webpack/css-loader
               use: ["css-loader?-url&minimize", "sass-loader"]
            }),
         },
         // Do not bundle image references but leave them as regular URLs.
         {
            test: /\.(jpg|jpeg|gif|png)$/,
            exclude: /node_modules/,
            use: "url-loader?limit=1024&name=img/[name].[ext]"
         }
      ]
   },
   resolve: {
      extensions: [".js", ".ts", ".tsx", ".json", ".sass"],
      plugins: [new TsConfigPathsPlugin()]
   },
   plugins: [
      new CheckerPlugin(),
      /**
       * Combine referenced CSS or SCSS into single file.
       *
       * https://github.com/webpack/extract-text-webpack-plugin
       * https://github.com/webpack/webpack/tree/master/examples/multiple-entry-points-commons-chunk-css-bundle
       */
      new ExtractTextPlugin({
         filename: "css/[name].css",
         allChunks: true
      }),
      new webpack.LoaderOptionsPlugin({
         minimize: true,
         debug: false,
         sassLoader: {
            includePaths: [
               path.resolve(__dirname, source + "/style")
            ]
         }
      }),
      // Create common chunk with all node_modules regardless of how often they"re
      // used in other chunks
      // https://webpack.github.io/docs/list-of-plugins.html#commonschunkplugin
      new webpack.optimize.CommonsChunkPlugin({
         name: "common",
         minChunks: module =>
            (typeof module.userRequest === "string") &&
            module.userRequest.includes("node_modules")
      }),
      // https://medium.com/webpack/webpack-3-official-release-15fd2dd8f07b
      // https://medium.com/webpack/webpack-freelancing-log-book-week-5-7-4764be3266f5
      // this plugin is preventing watcher from emitting correct bundles
      //new webpack.optimize.ModuleConcatenationPlugin(),
      new webpack.SourceMapDevToolPlugin({
         test: [/\.js$/],
         exclude: [/common\./],
         filename: "js/[name].js.map",
      }),
      new webpack.optimize.UglifyJsPlugin({
         compress: {
            warnings: false,
            screw_ie8: true
         },
         // mangled variables are trouble for source maps
         mangle: isProd,
         output: { comments: false },
         sourceMap: !isProd
      }),
      // production Node environment needed to trigger optimized React build
      new webpack.DefinePlugin({
         "process.env.NODE_ENV": JSON.stringify("production") // JSON.stringify(process.env.NODE_ENV)
      })
   ]
};