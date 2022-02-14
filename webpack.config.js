const path = require('path');
const CopyPlugin = require('copy-webpack-plugin');
const HtmlPlugin = require('html-webpack-plugin');
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

const env = process.env.NODE_ENV || 'development';

module.exports = {
    mode: env,
    devtool: env === 'development' ? 'source-map' : false,
    devServer: {
//        watchFiles: ['src/**/*.html']
    },
    entry: {
        'index': './src/index.ts'
    },
    output: {
        path: path.join(__dirname, 'dist'),
        filename: '[name]-[hash].js',
        clean: true
    },
    module: {
        rules: [
            {
                test: /\.ts$/,
                use: 'ts-loader'
            },
            {
                test: /\.(scss|css)$/,
                use: [MiniCssExtractPlugin.loader, 'css-loader', 'sass-loader'],
            }
        ]
    },
    resolve: {
        modules: [
            'node_modules',
        ],
        extensions: [
            '.ts',
            '.js'
        ]
    },
    plugins: [
        new MiniCssExtractPlugin({
            filename: 'style.css',
        }),
        new HtmlPlugin({
            template: 'src/index.html',
            filename: 'index.html'
        }),
        // new CopyPlugin({
        //     patterns: [
        //         {from: 'src/index.html', to: path.join(__dirname, 'dist')}
        //     ]
        // })
    ]
};