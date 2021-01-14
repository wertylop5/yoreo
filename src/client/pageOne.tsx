import * as React from "react";
import * as ReactDOM from "react-dom";

type AppProps = {message: string}

const Hello = (props: AppProps) => {
	return (<>
		<p>{props.message}</p>
	</>);
}

ReactDOM.render(
	<Hello message="pageOne"/>,
	document.getElementById("root")
);

