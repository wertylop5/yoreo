const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");

module.exports = {
	entry: {
		root: './src/client/index.tsx',
		pageOne: './src/client/pageOne.tsx'
	},
	mode: "development",
	devtool: 'inline-source-map',
	module: {
		rules: [
			{
				test: /\.(ts|js)x?$/,
				exclude: /node_modules/,
				use: [
					{
						loader: 'babel-loader',
					},
					{
						loader: 'ts-loader',
					}
				],
			},
		],
	},
	resolve: {
		extensions: ['.tsx', '.ts', '.js'],
	},
	output: {
		path: path.resolve(__dirname, 'dist'),
		filename: '[name].bundle.js'
	},
	plugins: [
		new HtmlWebpackPlugin({
			template: 'src/client/index.html',
			chunks: ["root"],
			filename: path.resolve(__dirname, 'dist/index.html'),
		}),
		new HtmlWebpackPlugin({
			template: 'src/client/index.html',
			chunks: ["pageOne"],
			filename: path.resolve(__dirname, 'dist/pageOne.html'),
		})
	]
};

