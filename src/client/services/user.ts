import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'

import type {
	User
} from '../types/types'

export const userApi = createApi({
	reducerPath: 'userApi',
	baseQuery: fetchBaseQuery({
		baseUrl: '/user/'
	}),
	endpoints: (builder) => ({
		getUserByName: builder.query<User, string>({
			query: (name) => `get/${name}`
		}),
		createUser: builder.mutation<User, Partial<User>>({
			query: ({ name }) => ({
				url: `create`,
				method: 'POST',
				body: {
					name
				}
			})
			// TODO: do any response processing if necessary
		})
	})
})

export const {
	useGetUserByNameQuery
} = userApi

