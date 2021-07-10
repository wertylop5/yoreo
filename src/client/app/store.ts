import { configureStore } from '@reduxjs/toolkit'

import routinesReducer from '../features/routines/routinesSlice'

const store = configureStore({
	reducer: {
		routines: routinesReducer,
	},
})

export type RootState = ReturnType<typeof store.getState>

export type AppDispatch = typeof store.dispatch

