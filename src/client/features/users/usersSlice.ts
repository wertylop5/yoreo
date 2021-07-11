import {
	createSlice,
	createEntityAdapter
} from '@reduxjs/toolkit'

import type {
	RootState
} from '../../app/store'

enum UserFetchStatus {
	Idle = "IDLE",
	Fetching = "FETCHING",
	Success = "SUCCESS",
	Fail = "FAIL"
}

type UserFetchState = {
	fetchStatus: UserFetchStatus
}

type User = {
	id: number,
	name: string
}

const initial: UserFetchState = {
	fetchStatus: UserFetchStatus.Idle
}

const usersAdapter = createEntityAdapter<User>()

const initialState = usersAdapter.getInitialState(initial)

const usersSlice = createSlice({
	name: 'users',
	initialState,
	reducers: {}
})

export default usersSlice.reducer

export const {
	selectAll: selectAllUsers,
	selectById: selectUsersById,
	selectIds: selectUserIds
} = usersAdapter.getSelectors((state: RootState) => state.users)

