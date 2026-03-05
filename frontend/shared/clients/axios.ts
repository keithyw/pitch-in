import axios, { AxiosInstance, AxiosError } from 'axios'
import { RefreshResponse } from '@pitch-in/shared/types'

export interface AxiosOptions {
	baseUrl?: string
	getToken: () => string | null
	getRefreshToken: () => string | null
	onRefreshSuccess: (res: RefreshResponse) => void
	onLogout: () => void
	refreshUrl: string
	authUrls: string[]
}

export const createAxiosClient = (opts: AxiosOptions): AxiosInstance => {
	const client = axios.create({ baseURL: opts.baseUrl })
	let isRefreshing = false
	let failedQueue: any[] = []

	const processQueue = (
		error: AxiosError | null,
		token: string | null = null,
	) => {
		failedQueue.forEach((p) => (error ? p.reject(error) : p.resolve(token)))
		failedQueue = []
	}

	client.interceptors.request.use((config) => {
		const token = opts.getToken()
		if (token) config.headers.Authorization = `Bearer ${token}`
		return config
	})

	client.interceptors.response.use(
		(res) => res,
		async (error: AxiosError) => {
			const req: any = error.config
			if (error.response?.status !== 401 || req._retry)
				return Promise.reject(error)
			// if the url is token url
			if (opts.authUrls.some((url) => req.url?.includes(url)))
				return Promise.reject(error)
			if (isRefreshing) {
				return new Promise((resolve, reject) => {
					failedQueue.push({ resolve, reject })
				})
			}

			req._retry = true
			isRefreshing = true

			try {
				const refresh = opts.getRefreshToken()
				const postRes = await axios.post(`${opts.baseUrl}${opts.refreshUrl}`, {
					refresh,
				})
				opts.onRefreshSuccess(postRes.data)
				processQueue(null, postRes.data.access) // not sure if this will be the structure in terms of the refresh url payload
				return client(req)
			} catch (err) {
				processQueue(err as AxiosError)
				opts.onLogout()
			} finally {
				isRefreshing = false
			}
		},
	)

	return client
}
