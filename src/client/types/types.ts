export type User = {
	id: number,
	name: string
}

export type Routine = {
	id: number,
	name: string,
	owner: number,
	collaborators: any,
	data: string,
}

