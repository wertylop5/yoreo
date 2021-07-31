import { configureStore } from '@reduxjs/toolkit'

import { userApi } from '../services/user'
import routinesReducer from '../features/routines/routinesSlice'
import usersReducer from '../features/users/usersSlice'

export const store = configureStore({
	reducer: {
		routines: routinesReducer,
		users: usersReducer,
		[userApi.reducerPath]: userApi.reducer
	},
	middleware: (getDefaultMiddleware) =>
		getDefaultMiddleware().concat(userApi.middleware),
})

export type RootState = ReturnType<typeof store.getState>

export type AppDispatch = typeof store.dispatch

