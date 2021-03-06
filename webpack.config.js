const path = require('path');
const CopyPlugin = require('copy-webpack-plugin');
const HtmlPlugin = require('html-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

const env = process.env.NODE_ENV || 'development';

module.exports = {
    mode: env,
    devtool: env === 'development' ? 'source-map' : false,
    devServer: {
        allowedHosts: ['.ngrok.io'],
    },
    entry: {
        index: './fe/index.ts',
        register: './fe/register.ts',
        game: './fe/game.ts',
        finish: './fe/finish.ts',
    },
    output: {
        path: path.join(__dirname, 'dist'),
        filename: 'js/[name]-[hash].js',
        clean: true,
    },
    module: {
        rules: [
            {
                test: /\.ts$/,
                use: 'ts-loader',
            },
            {
                test: /\.(scss|css)$/,
                use: [MiniCssExtractPlugin.loader, 'css-loader', 'sass-loader'],
            },
        ],
    },
    resolve: {
        modules: ['node_modules'],
        extensions: ['.ts', '.js'],
    },
    plugins: [
        new MiniCssExtractPlugin({
            filename: '[name]-[hash].css',
        }),
        new HtmlPlugin({
            template: 'fe/index.html',
            filename: 'index.html',
            chunks: ['index'],
        }),
        new HtmlPlugin({
            template: 'fe/register.html',
            filename: 'register/index.html',
            chunks: ['register'],
        }),
        new HtmlPlugin({
            template: 'fe/game.html',
            filename: 'game/index.html',
            chunks: ['game'],
        }),
        new HtmlPlugin({
            template: 'fe/finish.html',
            filename: 'finish/index.html',
            chunks: ['finish'],
        }),
        new CopyPlugin({
            patterns: [
                {from: 'fe/assets/ryugen.mp3', to: path.join(__dirname, 'dist/assets')},
                {from: 'fe/assets/turn.mp3', to: path.join(__dirname, 'dist/assets')},
                {from: 'fe/assets/start.mp3', to: path.join(__dirname, 'dist/assets')},
                {from: 'fe/assets/alarm.mp3', to: path.join(__dirname, 'dist/assets')},
                {from: 'fe/assets/favicon.ico', to: path.join(__dirname, 'dist')},
                {from: 'fe/assets/rule/index.html', to: path.join(__dirname, 'dist/rule')},
                {from: 'fe/assets/rule/style.css', to: path.join(__dirname, 'dist/rule')},
            ],
        }),
    ],
};
