import { configureStore } from '@reduxjs/toolkit'

import routinesReducer from '../features/routines/routinesSlice'
import usersReducer from '../features/users/usersSlice'

const store = configureStore({
	reducer: {
		routines: routinesReducer,
		users: usersReducer,
	},
})

export type RootState = ReturnType<typeof store.getState>

export type AppDispatch = typeof store.dispatch

