import {
	createSlice,
	createEntityAdapter
} from '@reduxjs/toolkit'

enum RoutineFetchStatus {
	Idle = "IDLE",
	Fetching = "FETCHING",
	Success = "SUCCESS",
	Fail = "FAIL"
}

type RoutineState = {
	fetchStatus: RoutineFetchStatus,
}

type Routine = {
	id: number,
	name: string,
	owner: number,
	collaborators: any,
	data: string,
}

const initial: RoutineState = {
	fetchStatus: RoutineFetchStatus.Idle
}

const routinesAdapter = createEntityAdapter<Routine>()

const initialState = routinesAdapter.getInitialState(initial)

const routinesSlice = createSlice({
	name: 'routines',
	initialState,
	reducers: {}
})

export default routinesSlice.reducer

