import * as React from "react";
import { useCreateUserMutation } from "../../services/user"

const CreateUserForm = () => {
	const [
		createUser,
		{
			data: data,
			error: err,
		}
	] = useCreateUserMutation()

	const [name, setName] = React.useState("")

	const onNameSubmit = (event: React.FormEvent) => {
		console.log(name)
		createUser({name})
		event.preventDefault()
	}

	if (err) console.error(err)

	if (data) console.log(data)

	return (
		<form method="POST" onSubmit={onNameSubmit}>
			<input
				type="text"
				name="userName"
				value={name}
				onChange={event => setName(event.target.value)}
			/>
			<input type="submit" name="submit" />
		</form>
	);
}

export default CreateUserForm
