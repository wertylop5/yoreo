import {
	createSlice,
	createEntityAdapter
} from '@reduxjs/toolkit'

import type {
	RootState
} from '../../app/store'

import type {
	Routine
} from '../../types/types'

enum RoutineFetchStatus {
	Idle = "IDLE",
	Fetching = "FETCHING",
	Success = "SUCCESS",
	Fail = "FAIL"
}

type RoutineFetchState = {
	fetchStatus: RoutineFetchStatus,
}

const initial: RoutineFetchState = {
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

export const {
	selectAll: selectAllRoutines,
	selectById: selectRoutineById,
	selectIds: selectRoutineIds
} = routinesAdapter.getSelectors((state: RootState) => state.routines)

