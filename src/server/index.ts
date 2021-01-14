import express from "express";

const PORT = 3000;
const HOST = "localhost";

const app = express();

app.use(express.static("dist"))

app.listen(5000, "localhost", () => {
	console.log("listening");
});

