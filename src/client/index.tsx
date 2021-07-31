import * as React from "react";
import * as ReactDOM from "react-dom";
import { Provider } from "react-redux";

import { store } from "./app/store"
import CreateUserForm from "./features/users/CreateUserForm"

type AppProps = {message: string}

const Hello = (props: AppProps) => {
	return (<>
		<p>{props.message}</p>
		<CreateUserForm />
	</>);
}

ReactDOM.render(
	<Provider store={store}>
		<Hello message="hi"/>
	</Provider>,
	document.getElementById("root")
);

