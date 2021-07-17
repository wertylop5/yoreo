import {
	createAsyncThunk,
	createEntityAdapter,
	createSlice,
	PayloadAction
} from '@reduxjs/toolkit'

import type {
	RootState
} from '../../app/store'

import type {
	User
} from '../../types/types'

enum UserFetchStatus {
	Idle = "IDLE",
	Fetching = "FETCHING",
	Success = "SUCCESS",
	Fail = "FAIL"
}

type UserFetchState = {
	fetchStatus: UserFetchStatus
}

type CreateUserPayload = {
	name: string,
}

const initial: UserFetchState = {
	fetchStatus: UserFetchStatus.Idle
}

const usersAdapter = createEntityAdapter<User>()

const initialState = usersAdapter.getInitialState(initial)

export const createUser = createAsyncThunk('users/createUser', async (name: string) => {
	const response = await fetch("/user/create", { method: "POST" })
})

const usersSlice = createSlice({
	name: 'users',
	initialState,
	reducers: {},
	extraReducers: (builder) => {
		builder
			.addCase(createUser.pending, (state, action) => {})
			.addCase(createUser.fulfilled, (state, action) => {})
			.addCase(createUser.rejected, (state, action) => {})
	}
})

export default usersSlice.reducer

export const {
	selectAll: selectAllUsers,
	selectById: selectUsersById,
	selectIds: selectUserIds
} = usersAdapter.getSelectors((state: RootState) => state.users)

